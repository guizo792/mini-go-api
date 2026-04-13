package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/guizo792/mini-go-api/api"
	"github.com/guizo792/mini-go-api/internal/tools"
	log "github.com/sirupsen/logrus"
	"github.com/gorilla/schema"
)

func GetOrder(w http.ResponseWriter, r *http.Request) {
	var params = api.OrderParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	err = decoder.Decode(&params, r.URL.Query())

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	var orderDetails *tools.OrderDetails
	orderDetails = (*database).GetUserOrder(params.Username)
	if orderDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var response = api.OrderResponse{
		Code: http.StatusOK,
		OrderId: (*orderDetails).OrderId,
		Product: (*orderDetails).Product,
		Quantity: (*orderDetails).Quantity,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

}
