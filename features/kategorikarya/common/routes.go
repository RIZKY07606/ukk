package common

import (
	"ukk-smkn2/features/kategorikarya/create"
	"ukk-smkn2/features/kategorikarya/delete"
	"ukk-smkn2/features/kategorikarya/getall"
	"ukk-smkn2/features/kategorikarya/getbyid"
	"ukk-smkn2/features/kategorikarya/update"
	"ukk-smkn2/infrastructure/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(api fiber.Router, db *gorm.DB) {
	groupKategori := api.Group("/kategori-karya", middleware.JWTProtected())
	groupKategori.Post("/", create.CreateKategoriKarya(db))
	groupKategori.Get("/", getall.GetAllKategoriKarya(db))
	groupKategori.Get("/:id", getbyid.GetKategoriKaryaByID(db))
	groupKategori.Put("/:id", update.UpdateKategoriKarya(db))
	groupKategori.Delete("/:id", delete.DeleteKategoriKarya(db))

}
