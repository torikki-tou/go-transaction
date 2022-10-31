package repo

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/torikki-tou/go-transaction/common"
)

type UserRepository interface {
	ChangeBalance(clientId string, delta int) (int, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (c *userRepository) ChangeBalance(clientId string, delta int) (int, error) {
	ctx := context.Background()
	tx, err := c.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})
	if err != nil {
		return 0, err
	}

	var balance int
	rows, err := tx.QueryContext(ctx, fmt.Sprintf("SELECT balance FROM clients WHERE id = '%s'", clientId))
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}
	defer func(rows *sql.Rows) { _ = rows.Close() }(rows)
	for rows.Next() {
		err := rows.Scan(&balance)
		if err != nil {
			return 0, err
		}
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
