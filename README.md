#basket-service


To insert product add this command to connection.go file inside to Migration function before

Db.AutoMigrate(&Product{})
    db.Model(&Product{}).AddForeignKey("basket_id", "baskets(id)", "CASCADE", "CASCADE")
/*
newProduct := product.Product{
Model: gorm.Model{},
ProductInfo: entity_product.ProductInfo{
Name:        "product1",
Price:       100,
Vat:         8,
Quantity:    15,
Description: "erdal",
},
}
DB.Create(&newProduct)
fmt.Println("Migration completed")
*/

