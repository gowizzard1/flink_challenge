package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"flink_chalenge/model"
	"flink_chalenge/service"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

type LocationHandler struct {
	service *service.Location
}

func NewLocationHandler(locationService *service.Location) *LocationHandler {
	return &LocationHandler{
		service: locationService,
	}
}

func (h *LocationHandler) AddLocation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderId, ok := vars["order_id"]
	if !ok {
		err := fmt.Errorf("error during rrid parsing")
		log.Err(err).Send()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	payload := new(model.Location)
	err := json.NewDecoder(r.Body).Decode(payload)
	if err != nil {
		log.Err(err).Send()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = payload.Validate()
	if err != nil {
		log.Err(err).Send()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.service.AppendLocation(orderId, *payload)
	respondWithJSON(w, http.StatusOK, payload)
}

func (h *LocationHandler) GetLocation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderId, ok := vars["order_id"]
	if !ok {
		err := fmt.Errorf("error during rrid parsing")
		log.Err(err).Send()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var max int
	maxStr := r.URL.Query().Get("max")
	if len(maxStr) > 0 {
		var err error
		max, err = strconv.Atoi(maxStr)
		if err != nil {
			log.Err(err).Send()
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
	response, err := h.service.GetLocation(orderId, max)
	if err != nil {
		if err != nil {
			log.Err(err).Send()
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
	}
	respondWithJSON(w, http.StatusOK, LocationPayload{
		OrderId: orderId,
		History: response,
	})
}

func (h *LocationHandler) DeleteLocation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderId, ok := vars["order_id"]
	if !ok {
		err := fmt.Errorf("error during rrid parsing")
		log.Err(err).Send()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := h.service.DeleteLocation(orderId)
	if err != nil {
		log.Err(err).Send()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// respondWithJSON write json response format
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		log.Err(err).Send()
	}
}
