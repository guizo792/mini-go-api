package tools

import (
	log "github.com/sirupsen/logrus"
)

// Database collections
type LoginDetails struct {
	AuthToken string
	Username string
}

type OrderDetails struct {
	OrderId string
	Product string
	Quantity int64
}

type DatabaseInterface interface {
	GetUserLoginDetails(username string) (*LoginDetails, error)
	GetUserOrder(username string) (*OrderDetails, error)
	SetupDatabase() error
}

func NewDatabase(useMock bool) (DatabaseInterface, error) {
	var db DatabaseInterface
	if useMock {
		db = &mockDb{}
	} else {
		db = &postgresDb{}
	}

	var err error = db.SetupDatabase()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return db, nil
}
