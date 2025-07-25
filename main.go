package main

import (
	"log"

	api "github.com/cybercoder/restbill/pkg/api"
	database "github.com/cybercoder/restbill/pkg/database"
	"github.com/cybercoder/restbill/pkg/database/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("The .env file not loaded")
	}
	db := database.Init()
	db.AutoMigrate(
		&models.Currency{},
		&models.Category{},
		&models.Product{},
		&models.Addon{},
		&models.ProductPrice{},
		&models.AddonPrice{},
		&models.Invoice{},
		&models.InvoiceItem{},
		&models.Cart{},
		&models.CartItem{},
		&models.CartItemAddons{},
	)
	router := gin.Default()
	api.SetupRoutes(router)
	router.Run("0.0.0.0:8000")
}
