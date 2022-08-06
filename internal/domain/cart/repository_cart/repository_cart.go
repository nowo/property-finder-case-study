package repository_cart

import (
	"fmt"
	"property-finder-go-bootcamp-homework/database/postgres"
	"property-finder-go-bootcamp-homework/internal/.config/general"
	"property-finder-go-bootcamp-homework/internal/.config/messages"
	"property-finder-go-bootcamp-homework/internal/domain/cart"
	"strconv"
)

type CartRepository struct {
}

func New() ICartRepository {
	return &CartRepository{}
}

func (r *CartRepository) GetCartInfoByUserID(userID uint) ([]cart.Cart, error) {
	db := postgres.ConnectDB()
	defer postgres.Disconnect(db)

	newCart := []cart.Cart{}
	response := db.Table("carts").Where("user_id = ?", userID).Find(&newCart)
	fmt.Println()
	if response.Error != nil {
		return nil, response.Error
	}
	fmt.Println("new Cart: ", newCart)
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
		return 0, messages.DATABASE_OPERATION_FAILED
	}

	return count, nil
}

func (r *CartRepository) IsAmountExceedByMonth(userID uint) (bool, error) {
	db := postgres.ConnectDB()
	defer postgres.Disconnect(db)

	//get this month cart amount by product
	var cartAmount int64 = 0

	subQuery := db.Select("product_id").Where("user_id = ?", userID).Where("created_at >= ?", "now() - INTERVAL '1 month'").Table("carts")
	response := db.Table("products").Where("id IN ?", subQuery).Select("SUM(price)").Scan(&cartAmount)
	if response.Error != nil {
		return false, response.Error
	}
	givenMonthInteger, err := strconv.Atoi(general.GIVEN_AMOUNT)
	if err != nil {
		return false, err
	}
	return int64(givenMonthInteger) < cartAmount, nil

}

func (r *CartRepository) Delete(userID, productID uint) error {
	db := postgres.ConnectDB()
	defer postgres.Disconnect(db)

	deletedCart := cart.Cart{}
	response := db.Table("carts").Where("user_id = ?", userID).Where("product_id = ?", productID).First(&deletedCart).Unscoped().Delete(&deletedCart)
	if response.Error != nil {
		return messages.DATABASE_OPERATION_FAILED
	}

	return nil
}
