package api

import (
	"net/http"

	"github.com/gorilla/mux"

	"flink_chalenge/service"
)

func NewRouteHandler(locationService *service.Location) http.Handler {
	reportHandler := NewLocationHandler(locationService)
	router := mux.NewRouter()
	router.HandleFunc("/location/{order_id}/now", reportHandler.AddLocation).Methods(http.MethodPost)
	router.HandleFunc("/location/{order_id}", reportHandler.GetLocation).Methods(http.MethodGet)
	router.HandleFunc("/location/{order_id}", reportHandler.DeleteLocation).Methods(http.MethodDelete)

	return router
}
