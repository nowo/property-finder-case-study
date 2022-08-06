package application

import (
	"fmt"
	"property-finder-go-bootcamp-homework/database/postgres"
	"property-finder-go-bootcamp-homework/internal/api/router"
)

func Start() {
	fmt.Println("Application started!!!")
	postgres.Migration(postgres.ConnectDB())
	router.Router()
	//err := service_cart.New().AddToCart(1, 1)
	//if err != nil {
	//	fmt.Println("erdal12")
	//	fmt.Println(err)
	//}
}
