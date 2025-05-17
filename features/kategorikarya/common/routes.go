package common

import (
	"ukk-smkn2/features/siswa/create"
	"ukk-smkn2/features/siswa/delete"
	"ukk-smkn2/features/siswa/getall"
	"ukk-smkn2/features/siswa/getbyid"
	"ukk-smkn2/features/siswa/update"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(api fiber.Router, db *gorm.DB) {
	groupKategori := api.Group("/kategori-karya")
	groupKategori.Post("/", create.Handler(db))
	groupKategori.Get("/", getall.Handler(db))
	groupKategori.Get("/:id", getbyid.Handler(db))
	groupKategori.Put("/:id", update.Handler(db))
	groupKategori.Delete("/:id", delete.Handler(db))

}
