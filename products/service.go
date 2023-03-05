package products

type Service interface {
	GetAllProducts() ([]Product, error)
	GetProductByID(input GetProductByInput) (Product, error)
	FilterByBrand(brand string) ([]Product, error)
	FilterByType(productType string) ([]Product, error)
	FilterByColor(color string) ([]Product, error)
	FilterBySize(size int) ([]Product, error)
	FilterByAvailability(status string) ([]Product, error)
	ProductPagination(offset int) ([]Product, error)
	CreateProduct(input ProductInput) (Product, error)
	UpdateProduct(inputID GetProductByInput, inputData ProductInput) (Product, error)
	DeleteProduct(inputID GetProductByInput) (error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllProducts() ([]Product, error) {
	products, err := s.repository.FindAll()
	if err != nil {
		return []Product{}, err
	}

	return products, nil
}

func (s *service) ProductPagination(offset int) ([]Product, error) {
	products, err := s.repository.Pagination(offset)
	if err != nil {
		return []Product{}, err
	}

	return products, nil
}

func (s *service) GetProductByID(input GetProductByInput) (Product, error) {
	product, err := s.repository.FindByID(input.ID)
	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func (s *service) FilterByBrand(brand string) ([]Product, error) {
	products, err := s.repository.FindByBrand(brand)
	if err != nil {
		return []Product{}, err
	}

	return products, nil
}

func (s *service) FilterByType(productType string) ([]Product, error) {
	products, err := s.repository.FindByType(productType)
	if err != nil {
		return []Product{}, err
	}

	return products, nil
}

func (s *service) FilterByColor(color string) ([]Product, error) {
	products, err := s.repository.FindByColor(color)
	if err != nil {
		return []Product{}, err
	}

	return products, nil
}

func (s *service) FilterBySize(size int) ([]Product, error) {
	products, err := s.repository.FindBySize(size)
	if err != nil {
		return []Product{}, err
	}

	return products, nil
}

func (s *service) FilterByAvailability(status string) ([]Product, error) {
	products, err := s.repository.FindByAvailability(status)
	if err != nil {
		return []Product{}, err
	}

	return products, nil
}

func (s *service) CreateProduct(input ProductInput) (Product, error) {
	product := Product{
		Name: input.Name,
		Brand: input.Brand,
		Type: input.Type,
		Color: input.Color,
		Size: input.Size,
		Quantity: input.Quantity,
	}

	if input.Quantity != 0 {
		product.Status = "instock"
	} else {
		product.Status = "outstock"
	}

	newProduct, err := s.repository.Create(product)
	if err != nil {
		return Product{}, err
	}

	return newProduct, nil
}

func (s *service) UpdateProduct(inputID GetProductByInput, inputData ProductInput) (Product, error) {
	product, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return Product{}, err
	}

	product.Name = inputData.Name
	product.Brand = inputData.Brand
	product.Type = inputData.Type
	product.Color = inputData.Color
	product.Size = inputData.Size
	product.Quantity = inputData.Quantity

	if inputData.Quantity != 0 {
		product.Status = "instock"
	} else {
		product.Status = "outstock"
	}

	updatedProduct, err := s.repository.Update(product)
	if err != nil {
		return Product{}, err
	}

	return updatedProduct, nil
}

func (s *service) DeleteProduct(input GetProductByInput) (error) {
	product, err := s.repository.FindByID(input.ID)
	if err != nil {
		return err
	}

	err = s.repository.Delete(product.ID)
	if err != nil {
		return err
	}
	
	return nil
}