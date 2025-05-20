package main

import (
	"ukk-smkn2/database"
	"ukk-smkn2/docs"
	"ukk-smkn2/routes"

	"os"
	e "ukk-smkn2/entities"

	_ "ukk-smkn2/docs" // for Swagger docs

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

//	@title			API UKK SMKN2 SURABAYA DOCUMENTATION
//	@version		1.0
//	@description		Documentation for API UKK SMKN2 Surabaya
//	@host			ukk-3738-91aq7hdc.leapcell.dev
//	@BasePath		/

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if os.Getenv("ENV") != "production" {
		docs.SwaggerInfo.Host = "localhost:" + port
	} else {
		docs.SwaggerInfo.Host = "ukk-3738-91aq7hdc.leapcell.dev" // change later
	}

	app := fiber.New(fiber.Config{
		EnablePrintRoutes: true,
	})
	app.Use(logger.New())
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
	app.Get("/swagger/*", swagger.HandlerDefault)

	err := app.Listen("0.0.0.0:" + port)
	if err != nil {
		panic(err)
	}

}
