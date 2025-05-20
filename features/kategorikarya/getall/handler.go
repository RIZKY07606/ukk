package getall

import (
	"ukk-smkn2/entities"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetAllKategoriKarya godoc
//
//	@Summary		Get all kategori karya
//	@Description	Get all kategori karya, optional filter by name (query param 'nama')
//	@Tags			kategori_karya
//	@Accept			json
//	@Produce		json
//	@Param			nama	query		string	false	"Filter by name"
//	@Success		200		{object}	GetAllKategoriKaryaResponseWrapper
//	@Failure		500		{object}	map[string]string
//	@Router			/api/kategori-karya [get]
func GetAllKategoriKarya(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var kategoris []entities.KategoriKarya
		namaFilter := c.Query("nama")

		query := db.Model(&entities.KategoriKarya{})
		if namaFilter != "" {
			query = query.Where("nama ILIKE ?", "%"+namaFilter+"%") // ILIKE for case-insensitive on Postgres; gunakan LIKE kalau MySQL
		}

		if err := query.Find(&kategoris).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch kategori karya"})
		}

		var data []GetAllKategoriKaryaResponse
		for _, k := range kategoris {
			data = append(data, GetAllKategoriKaryaResponse{
				KategoriID: k.ID.String(),
				Nama:       k.Nama,
			})
		}

		return c.JSON(GetAllKategoriKaryaResponseWrapper{
			Code:    200,
			Message: "Success",
			Data:    data,
		})
	}
}
