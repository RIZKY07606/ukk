package main

import (
	"log"
	"os"

	"ukk-smkn2/database"
	"ukk-smkn2/docs"
	"ukk-smkn2/routes"
	e "ukk-smkn2/entities"

	_ "ukk-smkn2/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

func main() {
	host := os.Getenv("SWAGGER_HOST")
	if host == "" {
		host = "127.0.0.1:8080"
	}
	docs.SwaggerInfo.Host = host

	app := fiber.New(fiber.Config{
		EnablePrintRoutes: true,
	})
	app.Use(logger.New())

	db, err := database.Connect()
	if err != nil {
		log.Fatalf("❌ Database connection failed: %v", err)
	}

	if err := db.AutoMigrate(
		&e.Admin{},
		&e.Siswa{},
		&e.KategoriKarya{},
		&e.KaryaUKK{},
		&e.FileUpload{},
		&e.Review{},
	); err != nil {
		log.Fatalf("❌ AutoMigrate failed: %v", err)
	}
}
