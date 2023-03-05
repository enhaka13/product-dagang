package products

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Product, error)
	FindByID(ID int) (Product, error)
	FindByBrand(brand string) ([]Product, error)
	FindByType(productType string) ([]Product, error)
	FindByColor(color string) ([]Product, error)
	FindBySize(size int) ([]Product, error)
	FindByAvailability(status string) ([]Product, error)
	Pagination(offset int) ([]Product, error)
	Create(product Product) (Product, error)
	Update(product Product) (Product, error)
	Delete(ID int) (error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Product, error) {
	var products []Product

	if err := r.db.Find(&products).Error; err != nil {
		return []Product{}, err
	}

	return products, nil
}

func (r *repository) Pagination(offset int) ([]Product, error) {
	var products []Product

	if err := r.db.Limit(10).Offset(offset).Find(&products).Error; err != nil {
		return []Product{}, err
	}

	return products, nil
}

func (r *repository) FindByID(ID int) (Product, error) {
	var product Product

	if err := r.db.Where("id = ?", ID).Find(&product).Error; err != nil {
		return Product{}, err
	}

	return product, nil
}

func (r *repository) FindByBrand(brand string) ([]Product, error) {
	var products []Product

	if err := r.db.Where("brand = ?", brand).Find(&products).Error; err != nil {
		return []Product{}, err
	}

	return products, nil
}

func (r *repository) FindByType(productType string) ([]Product, error) {
	var products []Product

	if err := r.db.Where("type = ?", productType).Find(&products).Error; err != nil {
		return []Product{}, err
	}

	return products, nil
}

func (r *repository) FindByColor(color string) ([]Product, error) {
	var products []Product

	if err := r.db.Where("color = ?", color).Find(&products).Error; err != nil {
		return []Product{}, err
	}

	return products, nil
}

func (r *repository) FindBySize(size int) ([]Product, error) {
	var products []Product

	if err := r.db.Where("size = ?", size).Find(&products).Error; err != nil {
		return []Product{}, err
	}

	return products, nil
}

func (r *repository) FindByAvailability(status string) ([]Product, error) {
	var products []Product

	if err := r.db.Where("status = ?", status).Find(&products).Error; err != nil {
		return []Product{}, err
	}

	return products, nil
}

func (r *repository) Create(product Product) (Product, error) {
	if err := r.db.Create(&product).Error; err != nil {
		return Product{}, err
	}

	return product, nil
}

func (r *repository) Update(product Product) (Product, error) {
	if err := r.db.Save(&product).Error; err != nil {
		return Product{}, err
	}

	return product, nil
}

func (r *repository) Delete(ID int) (error) {
	var product Product
	if err := r.db.Where("id = ?", ID).Delete(&product).Error; err != nil {
		return err
	}

	return nil
}