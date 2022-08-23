package interactors

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/ElrondNetwork/elrond-go-core/core/check"
	"github.com/ElrondNetwork/elrond-sdk-erdgo/core"
	"github.com/ElrondNetwork/elrond-sdk-erdgo/data"
)

const minimumIntervalToResend = time.Second

// nonceTransactionsHandler is the handler used for an unlimited number of addresses.
// It basically contains a map of addressNonceHandler, creating new entries on the first
// access of a provided address. This struct delegates all the operations on the right
// instance of addressNonceHandler. It also starts a go routine that will periodically
// try to resend "stuck transactions" and to clean the inner state. The recommended resend
// interval is 1 minute. The Close method should be called whenever the current instance of
// nonceTransactionsHandler should be terminated and collected by the GC.
// This struct is concurrent safe.
type nonceTransactionsHandler struct {
	proxy              Proxy
	mutHandlers        sync.Mutex
	handlers           map[string]*addressNonceHandler
	checkForDuplicates bool
	cancelFunc         func()
	intervalToResend   time.Duration
}

// NewNonceTransactionHandler will create a new instance of the nonceTransactionsHandler. It requires a Proxy implementation
// and an interval at which the transactions sent are rechecked and eventually, resent.
// checkForDuplicates set as true will prevent sending a transaction with the same receiver, value and data.
func NewNonceTransactionHandler(proxy Proxy, intervalToResend time.Duration, checkForDuplicates bool) (*nonceTransactionsHandler, error) {
	if check.IfNil(proxy) {
		return nil, ErrNilProxy
	}
	if intervalToResend < minimumIntervalToResend {
		return nil, fmt.Errorf("%w for intervalToResend in NewNonceTransactionHandler", ErrInvalidValue)
	}

	nth := &nonceTransactionsHandler{
		proxy:              proxy,
		handlers:           make(map[string]*addressNonceHandler),
		intervalToResend:   intervalToResend,
		checkForDuplicates: checkForDuplicates,
	}

	ctx, cancelFunc := context.WithCancel(context.Background())
	nth.cancelFunc = cancelFunc
	go nth.resendTransactionsLoop(ctx, intervalToResend)

	return nth, nil
}

// GetNonce will return the nonce for the provided address
func (nth *nonceTransactionsHandler) GetNonce(ctx context.Context, address core.AddressHandler) (uint64, error) {
	if check.IfNil(address) {
		return 0, ErrNilAddress
	}

	anh := nth.getOrCreateAddressNonceHandler(address)

	return anh.getNonceUpdatingCurrent(ctx)
}

func (nth *nonceTransactionsHandler) getOrCreateAddressNonceHandler(address core.AddressHandler) *addressNonceHandler {
	nth.mutHandlers.Lock()
	addressAsString := string(address.AddressBytes())
	anh, found := nth.handlers[addressAsString]
	if !found {
		anh = newAddressNonceHandler(nth.proxy, address)
		nth.handlers[addressAsString] = anh
	}
	nth.mutHandlers.Unlock()

	return anh
}

// SendTransaction will store and send the provided transaction
func (nth *nonceTransactionsHandler) SendTransaction(ctx context.Context, tx *data.Transaction) (string, error) {
	if tx == nil {
		return "", ErrNilTransaction
	}

	addrAsBech32 := tx.SndAddr
	addressHandler, err := data.NewAddressFromBech32String(addrAsBech32)
	if err != nil {
		return "", fmt.Errorf("%w while creating address handler for string %s", err, addrAsBech32)
	}

	anh := nth.getOrCreateAddressNonceHandler(addressHandler)
	if nth.checkForDuplicates && anh.isTxAlreadySent(tx) {
		// TODO: add gas comparation logic EN-11887
		anh.decrementComputedNonce()
		return "", ErrTxAlreadySent
	}
	sentHash, err := anh.sendTransaction(ctx, tx)
	if err != nil {
		return "", fmt.Errorf("%w while sending transaction for address %s", err, addrAsBech32)
	}

	return sentHash, nil
}

func (nth *nonceTransactionsHandler) resendTransactionsLoop(ctx context.Context, intervalToResend time.Duration) {
	timer := time.NewTimer(intervalToResend)
	defer timer.Stop()

	for {
		timer.Reset(intervalToResend)

		select {
		case <-timer.C:
			nth.resendTransactions(ctx)
		case <-ctx.Done():
			log.Debug("finishing nonceTransactionsHandler.resendTransactionsLoop...")
			return
		}
	}
}

func (nth *nonceTransactionsHandler) resendTransactions(ctx context.Context) {
	nth.mutHandlers.Lock()
	defer nth.mutHandlers.Unlock()

	for _, anh := range nth.handlers {
		select {
		case <-ctx.Done():
			log.Debug("finishing nonceTransactionsHandler.resendTransactions...")
			return
		default:
		}

		resendCtx, cancel := context.WithTimeout(ctx, nth.intervalToResend)
		err := anh.reSendTransactionsIfRequired(resendCtx)
		log.LogIfError(err)
		cancel()
	}
}

// ForceNonceReFetch will mark the addressNonceHandler to re-fetch its nonce from the blockchain account.
// This should be only used in a fallback plan, when some transactions are completely lost (or due to a bug, not even sent in first time)
func (nth *nonceTransactionsHandler) ForceNonceReFetch(address core.AddressHandler) error {
	if check.IfNil(address) {
		return ErrNilAddress
	}

	anh := nth.getOrCreateAddressNonceHandler(address)
	anh.markReFetchNonce()

	return nil
}

// Close finishes the transactions resend go routine
func (nth *nonceTransactionsHandler) Close() error {
	nth.cancelFunc()

	return nil
}

// IsInterfaceNil returns true if there is no value under the interface
func (nth *nonceTransactionsHandler) IsInterfaceNil() bool {
	return nth == nil
}
