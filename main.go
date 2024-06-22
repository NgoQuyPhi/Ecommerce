package main

import (
	"PJ/E-Commerce_2/database"
	"PJ/E-Commerce_2/routers"
)

func main() {
	dsn := "root:NgocBich1609@@@tcp(localhost:3306)/ecommerce2?charset=utf8mb4&parseTime=True&loc=Local"
	database.Connect(dsn)
	r := routers.InitRouter()

	r.Run(":5050")

}
