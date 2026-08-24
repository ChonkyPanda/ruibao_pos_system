package repository

import (
	"ruibao_pos_system/database"
	
	"ruibao_pos_system/models"
)

func GetAllOrders() ([]models.Order, error) {
	var orders []models.Order
	err := database.DB.Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, err
}

func GetOrderByID(id uint) (*models.Order, error) {
	var order models.Order
		err := database.DB.First(&order, id).Error
	if err != nil {
		return nil, err
	}
	return &order, err
}

func GetOrdersByUserID(user_id uint) ([]models.Order, error) {
	var orders []models.Order
	err := database.DB.Find(&orders, user_id).Error
	if err != nil {
		return nil, err
	}
	return orders, err
}

func CreatOrder(user_id uint, transaction_id uint, items []models.OrderItem) (*models.Order, error) {
	order := &models.Order{
		UserID: user_id,
		TransactionID: transaction_id,
		Items: items,
	}

	err := database.DB.Create(order).Error
	if err != nil {
		return nil, err
	}
	return order, nil
}

func UpdateOrder(order *models.Order) error {
	return database.DB.Model(&models.Order{}).Where("id = ?", order.ID).Updates(order).Error
}

func DeleteOrder(id uint) error {
	return database.DB.Delete(&models.Order{}, id).Error
}