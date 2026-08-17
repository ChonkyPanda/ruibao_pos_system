package service

import (
	"ruibao_pos_system/models"
	"ruibao_pos_system/repository"
)

func GetAllProducts() ([]models.Product, error) {
	return repository.GetAllProducts()
}

func GetProductByID(id uint) (*models.Product, error) {
	return repository.GetProductByID(id)
}

func GetProductByBarcode(barcode string) (*models.Product, error) {
	return repository.GetProductByBarcode(barcode)
}

func CreateProduct(name, barcode string, price float64, stock int, category string) (*models.Product, error) {
	return repository.CreateProduct(name, barcode, price, stock, category)
}

func UpdateProduct(p *models.Product) error {
	return repository.UpdateProduct(p)
}

func DeleteProduct(id uint) error {
	return repository.DeleteProduct(id)
}