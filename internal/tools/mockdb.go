package tools

import (
	"time"
	log "github.com/sirupsen/logrus"
)

type mockDb struct{}

var mockLoginDetails = map[string]LoginDetails {
	"michael": {
		AuthToken: "123abc",
		Username: "michael",
	},
	"andy": {
		AuthToken: "321bkd",
		Username: "andy",
	},
	"pam": {
		AuthToken: "pam432",
		Username: "pampam",
	},
}

var mockOrderDetails = map[string]OrderDetails {
	"michael": {
		OrderId: "111",
		Product: "Boss Mug",
		Quantity: 1,
	},
	"andy": {
		OrderId: "321",
		Product: "Pizza",
		Quantity: 12,
	},
	"pam": {
		OrderId: "532",
		Product: "Printer",
		Quantity: 2,
	},
}

func (d *mockDb) GetUserLoginDetails(username string) (*LoginDetails, error) {
	// Simulate DB call
	time.Sleep(time.Second * 1)

	data, ok := mockLoginDetails[username]
	if !ok {
		return nil, nil
	}

	log.WithFields(log.Fields{
		"data": data,
	}).Info("Retrieved User Data from DB")

	return &data, nil
}

func (d *mockDb) GetUserOrder(username string) (*OrderDetails, error) {
	// Simulate DB call
	time.Sleep(time.Second * 1)

	data, ok := mockOrderDetails[username]

	if !ok {
		return nil, nil
	}

	log.WithFields(log.Fields{
		"data": data,
	}).Info("Retrieved Order Data from DB")

	return &data, nil
}

func (d *mockDb) SetupDatabase() error {
	return nil
}
