package products

import "time"

type Product struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Brand string `json:"brand"`
	Type string `json:"type"`
	Color string `json:"color"`
	Size int `json:"size"`
	Quantity int `json:"quantity"`
	Status string `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}