package repository

import (
	"ruibao_pos_system/database"
	
	"ruibao_pos_system/models"
)

func GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	err := database.DB.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func GetProductByID(id uint) (*models.Product, error) {
	var p models.Product
	err := database.DB.First(&p, id).Error
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func GetProductByBarcode(barcode string) (*models.Product, error) {
	var p models.Product
	err := database.DB.Where("barcode = ?", barcode).First(&p).Error
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func CreateProduct(name, barcode string, price float64, stock int, category string) (*models.Product, error) {
	p := &models.Product{
		Name:     name,	
		Barcode:  barcode,
		Price:    price,
		Stock:    stock,
		Category: category,
	}
	err := database.DB.Create(p).Error
	if err != nil {
		return nil, err
	}
	return p, nil
}

func UpdateProduct(p *models.Product) error {
	return database.DB.Model(&models.Product{}).Where("id = ?", p.ID).Updates(p).Error
}

func DeleteProduct(id uint) error {
	return database.DB.Delete(&models.Product{}, id).Error
}
