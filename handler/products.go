package handler

import (
	"net/http"
	"product-dagang/helper"
	"product-dagang/products"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	service products.Service
}

func NewProductHandler(service products.Service) *productHandler {
	return &productHandler{service}
}

func (h *productHandler) GetAllProducts(c *gin.Context) {
	products, err := h.service.GetAllProducts()
	if err != nil {
		response := helper.APIResponse("Error to get products", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to get products", http.StatusOK, "success", products)
	c.JSON(http.StatusOK, response)
}

func (h *productHandler) ProductPagination(c *gin.Context) {
	pageStr := c.Param("page")
	page, _ := strconv.Atoi(pageStr)
	offset := (page - 1) * 10

	products, err := h.service.ProductPagination(offset)
	if err != nil {
		response := helper.APIResponse("Error to get products", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to get products", http.StatusOK, "success", products)
	c.JSON(http.StatusOK, response)
}

func (h *productHandler) GetProductByID(c *gin.Context) {
	var input products.GetProductByInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get product by ID", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := h.service.GetProductByID(input)
	if err != nil {
		response := helper.APIResponse("Failed to call service: GetProductByID()", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to get product by ID", http.StatusOK, "success", product)
	c.JSON(http.StatusOK, response)
}

func (h *productHandler) FilterByBrand(c *gin.Context) {
	brand := c.Param("brand")

	product, err := h.service.FilterByBrand(brand)
	if err != nil {
		response := helper.APIResponse("Failed to call service: FilterByBrand()", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to get product by Brand", http.StatusOK, "success", product)
	c.JSON(http.StatusOK, response)
}

func (h *productHandler) FilterByType(c *gin.Context) {
	productType := c.Param("type")

	product, err := h.service.FilterByType(productType)
	if err != nil {
		response := helper.APIResponse("Failed to call service: FilterByType()", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to get product by Type", http.StatusOK, "success", product)
	c.JSON(http.StatusOK, response)
}

func (h *productHandler) FilterByColor(c *gin.Context) {
	color := c.Param("color")

	product, err := h.service.FilterByColor(color)
	if err != nil {
		response := helper.APIResponse("Failed to call service: FilterByColor()", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to get product by Color", http.StatusOK, "success", product)
	c.JSON(http.StatusOK, response)
}

func (h *productHandler) FilterBySize(c *gin.Context) {
	sizeStr := c.Param("size")
	size, _ := strconv.Atoi(sizeStr)

	product, err := h.service.FilterBySize(size)
	if err != nil {
		response := helper.APIResponse("Failed to call service: FilterBySize()", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to get product by Size", http.StatusOK, "success", product)
	c.JSON(http.StatusOK, response)
}

func (h *productHandler) FilterByAvailability(c *gin.Context) {
	status := c.Param("status")

	product, err := h.service.FilterByAvailability(status)
	if err != nil {
		response := helper.APIResponse("Failed to call service: FilterByAvailability()", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to get product by Available Status", http.StatusOK, "success", product)
	c.JSON(http.StatusOK, response)
}

func (h *productHandler) CreateProduct(c *gin.Context) {
	var input products.ProductInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create product", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newProduct, err := h.service.CreateProduct(input)
	if err != nil {
		response := helper.APIResponse("Failed to call service: CreateProduct()", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create new product", http.StatusCreated, "success", newProduct)
	c.JSON(http.StatusCreated, response)
}

func (h *productHandler) UpdateProduct(c *gin.Context) {
	var inputID products.GetProductByInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to bind ID from uri", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData products.ProductInput

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to bind data from JSON request", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatedProduct, err := h.service.UpdateProduct(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed to call service: UpdateProduct()", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to update product", http.StatusOK, "success", updatedProduct)
	c.JSON(http.StatusOK, response)
}

func (h *productHandler) DeleteProduct(c *gin.Context) {
	var input products.GetProductByInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get product by ID", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = h.service.DeleteProduct(input)
	if err != nil {
		response := helper.APIResponse("Failed to call service: DeleteProduct()", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to delete product", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
