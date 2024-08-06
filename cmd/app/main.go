package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/mauFade/infinity/internal/config"
	"github.com/mauFade/infinity/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := config.GetDSN()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	var GlobalDatabase *gorm.DB
	GlobalDatabase = db

	db.AutoMigrate(&models.User{}, &models.Address{})

	app := fiber.New()

	SetupRoutes(app)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", os.Getenv("APP_PORT"))))
}
