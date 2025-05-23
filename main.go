package main

import (
	"ukk-smkn2/database"
	"ukk-smkn2/docs"
	"ukk-smkn2/routes"

	"os"
	e "ukk-smkn2/entities"

	_ "ukk-smkn2/docs" // for Swagger docs

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

//	@title			API UKK SMKN2 SURABAYA DOCUMENTATION
//	@version		1.0
//	@description	Documentation for API UKK SMKN2 Surabaya
//	@host			0.0.0.0:8080
//	@BasePath		/

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	if os.Getenv("ENV") != "production" {
		docs.SwaggerInfo.Host = "0.0.0.0:8080"
	} else {
		docs.SwaggerInfo.Host = "ukk-production-4fd8.up.railway.app" // change later
	}

	app := fiber.New(fiber.Config{
		EnablePrintRoutes: true,
	})
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))
	database.Connect()
	db := database.DB
	// db.AutoMigrate(&e.User{}) //migrate later
	// db.AutoMigrate(&e.User{}, &e.otherEntities{})
	db.AutoMigrate(&e.Admin{})
	db.AutoMigrate(&e.Siswa{})
	db.AutoMigrate(&e.KategoriKarya{})
	db.AutoMigrate(&e.KaryaUKK{})
	db.AutoMigrate(&e.FileUpload{})
	db.AutoMigrate(&e.Review{})
	routes.SetupRoutes(app, db)
	app.Static("/swagger", "./swagger")
	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Listen(":8080")
}
