package repo

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"
	"yolllo-manager/models"
	"yolllo-manager/pkg/yolsdk"

	"github.com/jackc/pgx/v4"
)

func (r *Repository) CreateNewWallet() (walletAddress string, err error) {
	// init transaction
	ctx := context.Background()
	tx, err := r.Conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {

		return
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	var wallet models.Wallet
	wallet.CreatedAt = time.Now().Unix()
	err = tx.QueryRow(ctx, `
		INSERT INTO	wallets (
			created_at
		)
		VALUES ($1)
		RETURNING wallet_index`,
		wallet.CreatedAt).Scan(&wallet.WalletIndex)
	if err != nil {
		fmt.Println(err)

		return
	}

	wallet.WalletAddress, err = yolsdk.GetYolAddress(r.Config.Mnemonic, wallet.WalletIndex)
	if err != nil {

		return
	}

	// inc wallet counter
	_, err = tx.Exec(ctx, `
		UPDATE	wallets
		SET		wallet_address = $1
		WHERE	wallet_index = $2`,
		wallet.WalletAddress,
		wallet.WalletIndex,
	)
	if err != nil {
		fmt.Println(err)

		return
	}

	return wallet.WalletAddress, nil
}

func (r *Repository) GetWalletIndexByWalletAddress(walletAddress string) (walletIndex int64, err error) {
	var wallet models.Wallet
	rows, err := r.Conn.Query(context.Background(), `
		SELECT	*
		FROM	wallets
		WHERE	wallet_address 	= $1`,
		walletAddress)
	if err != nil {
		log.Println(err)

		return
	}
	for rows.Next() {
		err = rows.Scan(
			&wallet.WalletIndex,
			&wallet.WalletAddress,
			&wallet.CreatedAt,
		)
		if err != nil {
			log.Println(err)

			return
		}
	}

	if wallet.WalletIndex == 0 && wallet.WalletAddress == "" {
		err = errors.New("unknown address")

		return
	}

	return wallet.WalletIndex, nil
}
