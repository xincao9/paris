package model

type Item struct {
	Id     int       `json:"id"`
	Vector []float32 `json:"vector"`
}