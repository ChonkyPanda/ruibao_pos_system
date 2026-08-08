package db

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name  string
	Barcode string
	Price float32
	Stock int
	Category string
}

func GetAllProducts() ([]Product, error) {
	var products []Product
	result := DB.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

func CreateProduct(name, barcode, category string, price float32, stock int) (*Product, error) {
	product := Product {
		Name: name,
		Barcode: barcode,
		Price: price,
		Stock: stock,
		Category: category,
	}

	result := DB.Create(&product)
	if result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}

func DeleteProduct(barcode string) error {
	var product Product
	result := DB.Where("barcode = ?", barcode).First(&product)
	if result.Error != nil {
		return result.Error
	}

	return DB.Delete(&product).Error
}

func GetProductByBarcode(barcode string) (*Product, error) {
	var product Product
	result := DB.Where("barcode = ?", barcode).First(&product)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func UpdateStock(barcode string, newStock int) error {
	var product Product
	result := DB.Where("barcode = ?", barcode).First(&product)
	if result.Error != nil {
		return result.Error
	}

	product.Stock = newStock
	result = DB.Save(&product)
	return result.Error
}