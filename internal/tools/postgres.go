package tools

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	log "github.com/sirupsen/logrus"
)

type postgresDb struct {
	pool *pgxpool.Pool
}

func (d *postgresDb) SetupDatabase() error {
	connString := os.Getenv("DATABASE_URL")
	if connString == "" {
		return fmt.Errorf("DATABASE_URL is not set in env vars")
	}

	pool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		return err
	}

	d.pool = pool

	return nil
}

func (d *postgresDb) GetUserLoginDetails(username string) (*LoginDetails, error) {
	var data LoginDetails
	err := d.pool.QueryRow(
		context.Background(),
		`SELECT
			auth_token, username
		 FROM users
		 	WHERE username = $1`,
		username,
	).Scan(&data.AuthToken, &data.Username)

	if err != nil {
		log.WithError(err).WithField("username", username).Error("Failed to get user login details")
		return nil, err
	}

	return &data, nil
}

func (d *postgresDb) GetUserOrder(username string) (*OrderDetails , error) {
	var data OrderDetails
	err := d.pool.QueryRow(
		context.Background(),
		`SELECT
			order_id, product, quantity
		FROM
			orders
		WHERE
			username = $1`,
		username,
	).Scan(&data.OrderId, &data.Product, &data.Quantity)

	if err != nil {
		log.WithError(err).WithField("username", username).Error("Failed to get user order")
		return nil, nil
	}

	return &data, nil
}
