package products

type GetProductByInput struct {
	ID int `uri:"id" binding:"required"`
}

type ProductInput struct {
	Name string `json:"name"`
	Brand string `json:"brand"`
	Type string `json:"type"`
	Color string `json:"color"`
	Size int `json:"size"`
	Quantity int `json:"quantity"`
}