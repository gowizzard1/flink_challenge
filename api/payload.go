package api

import (
	"flink_chalenge/model"
)

type LocationPayload struct {
	OrderId string           `json:"order_id"`
	History []model.Location `json:"history"`
}
