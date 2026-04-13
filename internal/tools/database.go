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
	GetUserLoginDetails(username string) *LoginDetails
	GetUserOrder(username string) *OrderDetails
	SetupDatabase() error
}

func NewDatabase() (*DatabaseInterface, error) {
	var database DatabaseInterface = &mockDb{}

	var err error = database.SetupDatabase()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &database, nil
}
