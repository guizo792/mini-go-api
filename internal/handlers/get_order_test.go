package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/guizo792/mini-go-api/internal/tools"
)

type stubDB struct{}

func (s *stubDB) SetupDatabase() error { return nil }

func (s *stubDB) GetUserLoginDetails(username string) (*tools.LoginDetails, error) {
	return nil, nil
}

func (s *stubDB) GetUserOrder(username string) (*tools.OrderDetails, error) {
	if username != "michael" {
		return nil, nil
	}

	return &tools.OrderDetails{
		OrderId:  "111",
		Product:  "Boss Mug",
		Quantity: 1,
	}, nil
}

func TestGetOrder_OK(t *testing.T) {
	h := OrderHandler{DB: &stubDB{}}

	req := httptest.NewRequest(http.MethodGet, "/user/orders?username=michael", nil)
	rr := httptest.NewRecorder()

	h.GetOrder(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}

	var got tools.OrderDetails
	if err := json.NewDecoder(rr.Body).Decode(&got); err != nil {
		t.Fatalf("decode failed: %v", err)
	}

	if got.OrderId != "111" {
		t.Fatalf("expected order 111, got %s", got.OrderId)
	}
}
