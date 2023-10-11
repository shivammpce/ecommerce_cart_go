package main

import (
	"github.com/ecommerce_cart/database"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.Connect()

	r.Run(":8080")
}
