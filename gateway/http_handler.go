package main

import (
	"net/http"

	"github.com/brunodev09/industrial-oms-go/common"
	pb "github.com/brunodev09/industrial-oms-go/common/api"
)

type handler struct {
	client pb.OrderServiceClient
}

func NewHandler(client pb.OrderServiceClient) *handler {
	return &handler{client}
}

func (h *handler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	cid := r.PathValue("customerId")
	var items []*pb.ItemsWithQuantity
	if err := common.UnmarshalJSON(r, &items); err != nil {
		common.WriteErrorJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	h.client.CreateOrder(r.Context(), &pb.CreateOrderRequest{
		CustomerId: cid,
		Items:      items,
	})
}

func (h *handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/customers/{customerId}/orders", h.HandleCreateOrder)
}
