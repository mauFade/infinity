package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/mauFade/infinity/internal/config"
)

func main() {
	config.ConnectDB()

	app := fiber.New()

	SetupRoutes(app)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", os.Getenv("APP_PORT"))))
}
