package repository_cart

import (
	"errors"
	"gorm.io/gorm"
	"property-finder-go-bootcamp-homework/database/postgres"
	"property-finder-go-bootcamp-homework/internal/.config/messages"
	"property-finder-go-bootcamp-homework/internal/domain/cart"
)

type CartRepository struct {
}

func New() ICartRepository {
	return &CartRepository{}
}

func (r *CartRepository) GetCartsByUserID(userID uint) ([]cart.Cart, error) {
	db := postgres.ConnectDB()
	defer postgres.Disconnect(db)

	newCart := []cart.Cart{}
	response := db.Table("carts").Where("user_id = ?", userID).Where("is_completed", false).Find(&newCart)
	if response.Error != nil {
		if errors.Is(response.Error, gorm.ErrRecordNotFound) {
			return newCart, messages.NO_RECORD_FOUND
		}
		return nil, response.Error
	}
	return newCart, nil
}

func (r *CartRepository) Create(newCart cart.Cart) error {
	db := postgres.ConnectDB()
	defer postgres.Disconnect(db)

	response := db.Create(&newCart)
	if response.Error != nil {
		return messages.DATABASE_OPERATION_FAILED
	}

	return nil
}

func (r *CartRepository) CountByProductID(productID uint) (int64, error) {
	db := postgres.ConnectDB()
	defer postgres.Disconnect(db)

	var count int64 = 0
	response := db.Table("carts").Where("product_id = ?", productID).Count(&count)

	if response.Error != nil {
		if errors.Is(response.Error, gorm.ErrRecordNotFound) {
			return count, messages.NO_RECORD_FOUND
		}
		return 0, response.Error
	}

	return count, nil
}

func (r *CartRepository) Delete(userID, productID uint) error {
	db := postgres.ConnectDB()
	defer postgres.Disconnect(db)

	deletedCart := cart.Cart{}
	response := db.Table("carts").Where("user_id = ?", userID).Where("product_id = ?", productID).Where("is_completed", false).First(&deletedCart).Unscoped().Delete(&deletedCart)
	if response.Error != nil && response.RowsAffected != 0 {
		return messages.DATABASE_OPERATION_FAILED
	}
	if response.RowsAffected == 0 {
		return messages.NO_RECORD_FOUND
	}
	return nil
}

func (r *CartRepository) Complete(userID, orderID uint) error {
	db := postgres.ConnectDB()
	defer postgres.Disconnect(db)

	response := db.Table("carts").Where("user_id = ?", userID).Where("is_completed", false).Update("order_id", orderID).Update("is_completed", true)
	if response.Error != nil && response.RowsAffected != 0 {
		return messages.DATABASE_OPERATION_FAILED
	}
	if response.RowsAffected == 0 {
		return messages.NO_RECORD_FOUND
	}

	return nil
}
