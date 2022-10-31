package repo

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/torikki-tou/go-transaction/common"
)

type ClientRepository interface {
	ChangeBalance(clientId string, delta int) (int, error)
}

type clientRepository struct {
	db *sql.DB
}

func NewClientRepository(db *sql.DB) ClientRepository {
	return &clientRepository{
		db: db,
	}
}

func (c *clientRepository) ChangeBalance(clientId string, delta int) (int, error) {
	ctx := context.Background()
	tx, err := c.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})
	if err != nil {
		return 0, err
	}

	var balance int
	err = tx.QueryRowContext(
		ctx,
		fmt.Sprintf("SELECT balance FROM clients WHERE id = '%s'", clientId),
	).Scan(&balance)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	changedBalance := balance + delta
	if changedBalance < 0 {
		return 0, &common.LowBalanceError{}
	}

	_, err = tx.ExecContext(ctx, fmt.Sprintf(
		"UPDATE clients SET balance = %d WHERE id = '%s'", changedBalance, clientId,
	))
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return changedBalance, nil
}
