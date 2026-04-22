package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/guizo792/mini-go-api/api"
	"github.com/guizo792/mini-go-api/internal/tools"
	log "github.com/sirupsen/logrus"
)

type OrderHandler struct {
	DB tools.DatabaseInterface
}

func (h *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	var params = api.OrderParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	err = decoder.Decode(&params, r.URL.Query())

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var orderDetails *tools.OrderDetails
	orderDetails, err = h.DB.GetUserOrder(params.Username)
	if err != nil || orderDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var response = api.OrderResponse{
		Code:     http.StatusOK,
		OrderId:  (*orderDetails).OrderId,
		Product:  (*orderDetails).Product,
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
