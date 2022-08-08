# Property Finder Basket Service

[![Property Finder](https://avatars.githubusercontent.com/u/7037387?s=200&v=4)](https://www.propertyfinder.ae/)

This repository is an example of simple basket service with jwt authentication
## Table OF Contents
* [General info](#general-info)
* [Technologies](#technologies)
* [Clone the project](#clone)
* [Setup](#setup)
* [Test](#test)

## Clone the project

```
$ git clone https://github.com/nowo/property-finder-case-study
$ cd property-finder-case-study
```
## General info
This project is a simple basket service app with jwt. It is produced for property finder case study. There is some product and users can add this product to their carts. User can see their cart, delete product from their cart or complete the order. There is some discount algorithms when completing the orders.
- Every fourth order whose total is more than given amount may have discount
depending on products. Products whose VAT is %1 don’t have any discount
but products whose VAT is %8 and %18 have discount of %10 and %15
respectively.
- If there are more than 3 items of the same product, then fourth and
subsequent ones would have %8 off.
- If the customer made purchase which is more than given amount in a month
then all subsequent purchases should have %10 off.
## Technologies
Project is created with:
* Go version: 1.18
* logrus
* gorm
* postgresql
* gomock
* fiber
* goconvey
* docker




## [Setup](#setup)
##### To run the project you must add environment file to root of the project
##### Environment Variables:
These environment variables are accepted:
- POSTGRES_HOST
- POSTGRES_PORT
- POSTGRES_USER
- POSTGRES_PASS
- POSTGRES_DB
- POSTGRES_SSL
- POSTGRES_HOST_AUTH_METHOD
- JWT_SECRET_KEY
- PORT
- GIVEN_AMOUNT
-
#### There is no product by default. Because of that you must include product by yourself.  To insert product add this command to <b>connection.go<b> file inside to Migration function
  ```
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
        DB.AutoMigrate(&user.User{}, &product.Product{}, &cart.Cart{}, &order.Order{})
```
To run project with docker, type this command. This command will create server and postqresql images
 ```
 $  docker compose up
 ```

 ## [test](#test)
 #### To run tests type these commands sequentially
 ```
 $ make install-dependencies
 $ make make-mocks
 $ make run-test
 ```



