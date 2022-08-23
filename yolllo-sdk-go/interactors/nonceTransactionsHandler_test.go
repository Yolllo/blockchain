package interactors

import (
	"context"
	"errors"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/ElrondNetwork/elrond-sdk-erdgo/core"
	"github.com/ElrondNetwork/elrond-sdk-erdgo/data"
	"github.com/ElrondNetwork/elrond-sdk-erdgo/testsCommon"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewNonceTransactionHandler(t *testing.T) {
	t.Parallel()

	nth, err := NewNonceTransactionHandler(nil, time.Minute, false)
	require.Nil(t, nth)
	assert.Equal(t, ErrNilProxy, err)

	nth, err = NewNonceTransactionHandler(&testsCommon.ProxyStub{}, time.Second-time.Nanosecond, false)
	require.Nil(t, nth)
	assert.True(t, errors.Is(err, ErrInvalidValue))
	assert.True(t, strings.Contains(err.Error(), "for intervalToResend in NewNonceTransactionHandler"))

	nth, err = NewNonceTransactionHandler(&testsCommon.ProxyStub{}, time.Minute, false)
	require.NotNil(t, nth)
	require.Nil(t, err)

	require.Nil(t, nth.Close())
}

func TestNonceTransactionsHandler_GetNonce(t *testing.T) {
	t.Parallel()

	testAddress, _ := data.NewAddressFromBech32String("erd1zptg3eu7uw0qvzhnu009lwxupcn6ntjxptj5gaxt8curhxjqr9tsqpsnht")
	currentNonce := uint64(664)

	numCalls := 0
	proxy := &testsCommon.ProxyStub{
		GetAccountCalled: func(address core.AddressHandler) (*data.Account, error) {
			if address.AddressAsBech32String() != testAddress.AddressAsBech32String() {
				return nil, errors.New("unexpected address")
			}

			numCalls++

			return &data.Account{
				Nonce: currentNonce,
			}, nil
		},
	}

	nth, _ := NewNonceTransactionHandler(proxy, time.Minute, false)
	nonce, err := nth.GetNonce(context.Background(), nil)
	assert.Equal(t, ErrNilAddress, err)
	assert.Equal(t, uint64(0), nonce)

	nonce, err = nth.GetNonce(context.Background(), testAddress)
	assert.Nil(t, err)
	assert.Equal(t, currentNonce, nonce)

	nonce, err = nth.GetNonce(context.Background(), testAddress)
	assert.Nil(t, err)
	assert.Equal(t, currentNonce+1, nonce)

	assert.Equal(t, 2, numCalls)

	require.Nil(t, nth.Close())
}

func TestNonceTransactionsHandler_SendMultipleTransactionsResendingEliminatingOne(t *testing.T) {
	t.Parallel()

	testAddress, _ := data.NewAddressFromBech32String("erd1zptg3eu7uw0qvzhnu009lwxupcn6ntjxptj5gaxt8curhxjqr9tsqpsnht")
	currentNonce := uint64(664)

	mutSentTransactions := sync.Mutex{}
	numCalls := 0
	sentTransactions := make(map[int][]*data.Transaction)
	proxy := &testsCommon.ProxyStub{
		GetAccountCalled: func(address core.AddressHandler) (*data.Account, error) {
			if address.AddressAsBech32String() != testAddress.AddressAsBech32String() {
				return nil, errors.New("unexpected address")
			}

			return &data.Account{
				Nonce: atomic.LoadUint64(&currentNonce),
			}, nil
		},
		SendTransactionsCalled: func(txs []*data.Transaction) ([]string, error) {
			mutSentTransactions.Lock()
			defer mutSentTransactions.Unlock()

			sentTransactions[numCalls] = txs
			numCalls++
			hashes := make([]string, len(txs))

			return hashes, nil
		},
		SendTransactionCalled: func(tx *data.Transaction) (string, error) {
			mutSentTransactions.Lock()
			defer mutSentTransactions.Unlock()

			sentTransactions[numCalls] = []*data.Transaction{tx}
			numCalls++

			return "", nil
		},
	}

	numTxs := 5
	nth, _ := NewNonceTransactionHandler(proxy, time.Second*2, false)
	txs := createMockTransactions(testAddress, numTxs, atomic.LoadUint64(&currentNonce))
	for i := 0; i < numTxs; i++ {
		_, err := nth.SendTransaction(context.TODO(), txs[i])
		require.Nil(t, err)
	}

	time.Sleep(time.Second * 3)
	_ = nth.Close()

	mutSentTransactions.Lock()
	defer mutSentTransactions.Unlock()

	numSentTransaction := 5
	numSentTransactions := 1
	assert.Equal(t, numSentTransaction+numSentTransactions, len(sentTransactions))
	for i := 0; i < numSentTransaction; i++ {
		assert.Equal(t, 1, len(sentTransactions[i]))
	}
	assert.Equal(t, numTxs-1, len(sentTransactions[numSentTransaction])) // resend
}

func TestNonceTransactionsHandler_SendMultipleTransactionsResendingEliminatingAll(t *testing.T) {
	t.Parallel()

	testAddress, _ := data.NewAddressFromBech32String("erd1zptg3eu7uw0qvzhnu009lwxupcn6ntjxptj5gaxt8curhxjqr9tsqpsnht")
	currentNonce := uint64(664)

	mutSentTransactions := sync.Mutex{}
	numCalls := 0
	sentTransactions := make(map[int][]*data.Transaction)
	proxy := &testsCommon.ProxyStub{
		GetAccountCalled: func(address core.AddressHandler) (*data.Account, error) {
			if address.AddressAsBech32String() != testAddress.AddressAsBech32String() {
				return nil, errors.New("unexpected address")
			}

			return &data.Account{
				Nonce: atomic.LoadUint64(&currentNonce),
			}, nil
		},
		SendTransactionCalled: func(tx *data.Transaction) (string, error) {
			mutSentTransactions.Lock()
			defer mutSentTransactions.Unlock()

			sentTransactions[numCalls] = []*data.Transaction{tx}
			numCalls++

			return "", nil
		},
	}

	numTxs := 5
	nth, _ := NewNonceTransactionHandler(proxy, time.Second*2, false)
	txs := createMockTransactions(testAddress, numTxs, atomic.LoadUint64(&currentNonce))
	for i := 0; i < numTxs; i++ {
		_, err := nth.SendTransaction(context.Background(), txs[i])
		require.Nil(t, err)
	}

	atomic.AddUint64(&currentNonce, uint64(numTxs))
	time.Sleep(time.Second * 3)
	_ = nth.Close()

	mutSentTransactions.Lock()
	defer mutSentTransactions.Unlock()

	//no resend operation was made because all transactions were executed (nonce was incremented)
	assert.Equal(t, 5, len(sentTransactions))
	assert.Equal(t, 1, len(sentTransactions[0]))
}

func TestNonceTransactionsHandler_SendTransactionResendingEliminatingAll(t *testing.T) {
	t.Parallel()

	testAddress, _ := data.NewAddressFromBech32String("erd1zptg3eu7uw0qvzhnu009lwxupcn6ntjxptj5gaxt8curhxjqr9tsqpsnht")
	currentNonce := uint64(664)

	mutSentTransactions := sync.Mutex{}
	numCalls := 0
	sentTransactions := make(map[int][]*data.Transaction)
	proxy := &testsCommon.ProxyStub{
		GetAccountCalled: func(address core.AddressHandler) (*data.Account, error) {
			if address.AddressAsBech32String() != testAddress.AddressAsBech32String() {
				return nil, errors.New("unexpected address")
			}

			return &data.Account{
				Nonce: atomic.LoadUint64(&currentNonce),
			}, nil
		},
		SendTransactionCalled: func(tx *data.Transaction) (string, error) {
			mutSentTransactions.Lock()
			defer mutSentTransactions.Unlock()

			sentTransactions[numCalls] = []*data.Transaction{tx}
			numCalls++

			return "", nil
		},
	}

	numTxs := 1
	nth, _ := NewNonceTransactionHandler(proxy, time.Second*2, false)
	txs := createMockTransactions(testAddress, numTxs, atomic.LoadUint64(&currentNonce))

	hash, err := nth.SendTransaction(context.Background(), txs[0])
	require.Nil(t, err)
	require.Equal(t, "", hash)

	atomic.AddUint64(&currentNonce, uint64(numTxs))
	time.Sleep(time.Second * 3)
	_ = nth.Close()

	mutSentTransactions.Lock()
	defer mutSentTransactions.Unlock()

	//no resend operation was made because all transactions were executed (nonce was incremented)
	assert.Equal(t, 1, len(sentTransactions))
	assert.Equal(t, numTxs, len(sentTransactions[0]))
}

func TestNonceTransactionsHandler_SendTransactionErrors(t *testing.T) {
	t.Parallel()

	testAddress, _ := data.NewAddressFromBech32String("erd1zptg3eu7uw0qvzhnu009lwxupcn6ntjxptj5gaxt8curhxjqr9tsqpsnht")
	currentNonce := uint64(664)

	var errSent error
	proxy := &testsCommon.ProxyStub{
		GetAccountCalled: func(address core.AddressHandler) (*data.Account, error) {
			if address.AddressAsBech32String() != testAddress.AddressAsBech32String() {
				return nil, errors.New("unexpected address")
			}

			return &data.Account{
				Nonce: atomic.LoadUint64(&currentNonce),
			}, nil
		},
		SendTransactionCalled: func(tx *data.Transaction) (string, error) {
			return "", errSent
		},
	}

	numTxs := 1
	nth, _ := NewNonceTransactionHandler(proxy, time.Second*2, false)
	txs := createMockTransactions(testAddress, numTxs, atomic.LoadUint64(&currentNonce))

	hash, err := nth.SendTransaction(context.Background(), nil)
	require.Equal(t, ErrNilTransaction, err)
	require.Equal(t, "", hash)

	errSent = errors.New("expected error")

	hash, err = nth.SendTransaction(context.Background(), txs[0])
	require.True(t, errors.Is(err, errSent))
	require.Equal(t, "", hash)
}

func createMockTransactions(addr core.AddressHandler, numTxs int, startNonce uint64) []*data.Transaction {
	txs := make([]*data.Transaction, 0, numTxs)
	for i := 0; i < numTxs; i++ {
		tx := &data.Transaction{
			Nonce:     startNonce,
			Value:     "1",
			RcvAddr:   addr.AddressAsBech32String(),
			SndAddr:   addr.AddressAsBech32String(),
			GasPrice:  100000,
			GasLimit:  50000,
			Data:      nil,
			Signature: "sig",
			ChainID:   "3",
			Version:   1,
		}

		txs = append(txs, tx)
		startNonce++
	}

	return txs
}

func TestNonceTransactionsHandler_SendTransactionsWithGetNonce(t *testing.T) {
	t.Parallel()

	testAddress, _ := data.NewAddressFromBech32String("erd1zptg3eu7uw0qvzhnu009lwxupcn6ntjxptj5gaxt8curhxjqr9tsqpsnht")
	currentNonce := uint64(664)

	mutSentTransactions := sync.Mutex{}
	numCalls := 0
	sentTransactions := make(map[int][]*data.Transaction)
	proxy := &testsCommon.ProxyStub{
		GetAccountCalled: func(address core.AddressHandler) (*data.Account, error) {
			if address.AddressAsBech32String() != testAddress.AddressAsBech32String() {
				return nil, errors.New("unexpected address")
			}

			return &data.Account{
				Nonce: atomic.LoadUint64(&currentNonce),
			}, nil
		},
		SendTransactionCalled: func(tx *data.Transaction) (string, error) {
			mutSentTransactions.Lock()
			defer mutSentTransactions.Unlock()

			sentTransactions[numCalls] = []*data.Transaction{tx}
			numCalls++

			return "", nil
		},
	}

	numTxs := 5
	nth, _ := NewNonceTransactionHandler(proxy, time.Second*2, false)
	txs := createMockTransactionsWithGetNonce(t, testAddress, 5, nth)
	for i := 0; i < numTxs; i++ {
		_, err := nth.SendTransaction(context.Background(), txs[i])
		require.Nil(t, err)
	}

	atomic.AddUint64(&currentNonce, uint64(numTxs))
	time.Sleep(time.Second * 3)
	_ = nth.Close()

	mutSentTransactions.Lock()
	defer mutSentTransactions.Unlock()

	//no resend operation was made because all transactions were executed (nonce was incremented)
	assert.Equal(t, numTxs, len(sentTransactions))
	assert.Equal(t, 1, len(sentTransactions[0]))
}

func TestNonceTransactionsHandler_SendDuplicateTransactions(t *testing.T) {
	testAddress, _ := data.NewAddressFromBech32String("erd1zptg3eu7uw0qvzhnu009lwxupcn6ntjxptj5gaxt8curhxjqr9tsqpsnht")
	currentNonce := uint64(664)

	numCalls := 0
	proxy := &testsCommon.ProxyStub{
		GetAccountCalled: func(address core.AddressHandler) (*data.Account, error) {
			if address.AddressAsBech32String() != testAddress.AddressAsBech32String() {
				return nil, errors.New("unexpected address")
			}

			return &data.Account{
				Nonce: atomic.LoadUint64(&currentNonce),
			}, nil
		},
		SendTransactionCalled: func(tx *data.Transaction) (string, error) {
			require.LessOrEqual(t, numCalls, 1)
			currentNonce++
			return "", nil
		},
	}

	nth, _ := NewNonceTransactionHandler(proxy, time.Second*60, true)

	nonce, err := nth.GetNonce(context.Background(), testAddress)
	require.Nil(t, err)
	tx := &data.Transaction{
		Nonce:     nonce,
		Value:     "1",
		RcvAddr:   testAddress.AddressAsBech32String(),
		SndAddr:   testAddress.AddressAsBech32String(),
		GasPrice:  100000,
		GasLimit:  50000,
		Data:      nil,
		Signature: "sig",
		ChainID:   "3",
		Version:   1,
	}
	_, err = nth.SendTransaction(context.Background(), tx)
	require.Nil(t, err)
	acc := nth.getOrCreateAddressNonceHandler(testAddress)
	t.Run("after sending first tx, nonce shall increase", func(t *testing.T) {
		require.Equal(t, acc.computedNonce+1, currentNonce)
	})
	t.Run("sending the same tx, NonceTransactionHandler shall return ErrTxAlreadySent "+
		"and computedNonce shall not increase", func(t *testing.T) {
		nonce, err = nth.GetNonce(context.Background(), testAddress)
		_, err = nth.SendTransaction(context.Background(), tx)
		require.Equal(t, err, ErrTxAlreadySent)
		require.Equal(t, nonce, currentNonce)
		require.Equal(t, acc.computedNonce+1, currentNonce)
	})

}

func createMockTransactionsWithGetNonce(
	tb testing.TB,
	addr core.AddressHandler,
	numTxs int,
	nth *nonceTransactionsHandler,
) []*data.Transaction {
	txs := make([]*data.Transaction, 0, numTxs)
	for i := 0; i < numTxs; i++ {
		nonce, err := nth.GetNonce(context.Background(), addr)
		require.Nil(tb, err)

		tx := &data.Transaction{
			Nonce:     nonce,
			Value:     "1",
			RcvAddr:   addr.AddressAsBech32String(),
			SndAddr:   addr.AddressAsBech32String(),
			GasPrice:  100000,
			GasLimit:  50000,
			Data:      nil,
			Signature: "sig",
			ChainID:   "3",
			Version:   1,
		}

		txs = append(txs, tx)
	}

	return txs
}

func TestNonceTransactionsHandler_ForceNonceReFetch(t *testing.T) {
	t.Parallel()

	testAddress, _ := data.NewAddressFromBech32String("erd1zptg3eu7uw0qvzhnu009lwxupcn6ntjxptj5gaxt8curhxjqr9tsqpsnht")
	currentNonce := uint64(664)

	proxy := &testsCommon.ProxyStub{
		GetAccountCalled: func(address core.AddressHandler) (*data.Account, error) {
			if address.AddressAsBech32String() != testAddress.AddressAsBech32String() {
				return nil, errors.New("unexpected address")
			}

			return &data.Account{
				Nonce: atomic.LoadUint64(&currentNonce),
			}, nil
		},
	}

	nth, _ := NewNonceTransactionHandler(proxy, time.Minute, false)
	_, _ = nth.GetNonce(context.Background(), testAddress)
	_, _ = nth.GetNonce(context.Background(), testAddress)
	newNonce, err := nth.GetNonce(context.Background(), testAddress)
	require.Nil(t, err)
	assert.Equal(t, atomic.LoadUint64(&currentNonce)+2, newNonce)

	err = nth.ForceNonceReFetch(nil)
	assert.Equal(t, ErrNilAddress, err)

	err = nth.ForceNonceReFetch(testAddress)
	assert.Nil(t, err)

	newNonce, err = nth.GetNonce(context.Background(), testAddress)
	assert.Equal(t, nil, err)
	assert.Equal(t, atomic.LoadUint64(&currentNonce), newNonce)
}
