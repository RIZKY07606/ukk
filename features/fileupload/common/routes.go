package common

import (
	"ukk-smkn2/features/fileupload/create"
	"ukk-smkn2/features/fileupload/delete"
	"ukk-smkn2/features/fileupload/getall"
	"ukk-smkn2/features/fileupload/getbyid"
	"ukk-smkn2/features/fileupload/update"
	"ukk-smkn2/infrastructure/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(api fiber.Router, db *gorm.DB) {
	groupFileUpload := api.Group("/fileupload", middleware.JWTProtected())
	groupFileUpload.Post("/", create.CreateFileUpload(db))
	groupFileUpload.Get("/", getall.GetAllFileUpload(db))
	groupFileUpload.Get("/:id", getbyid.GetFileUploadByID(db))
	groupFileUpload.Put("/:id", update.UpdateFileUpload(db))
	groupFileUpload.Delete("/:id", delete.DeleteFileUpload(db))

}
