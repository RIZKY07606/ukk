package common

import (
	"ukk-smkn2/features/admin/create"
	getprofile "ukk-smkn2/features/admin/getProfile"
	"ukk-smkn2/infrastructure/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(api fiber.Router, db *gorm.DB) {
	// if all routes should be authenticated
	// group := api.Group("/routes", middleware.JWTProtected())

	group := api.Group("/admin")
	group.Post("/register", create.Register(db))

	// if just one of them
	group.Get("/profile", middleware.JWTProtected(), getprofile.Profile(db))
}
