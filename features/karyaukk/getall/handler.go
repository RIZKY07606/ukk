package getall

import (
	"ukk-smkn2/entities"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetAllKaryaUKK godoc
//
//	@Summary		Get all karya UKK
//	@Description	Get all karya UKK, filter optional by 'judul'
//	@Tags			karya_ukk
//	@Accept			json
//	@Produce		json
//	@Param			judul	query		string	false	"Filter by judul"
//	@Success		200		{object}	GetAllKaryaUKKResponseWrapper
//	@Failure		500		{object}	map[string]string
//	@Router			/api/karya-ukk [get]
func GetAllKaryaUKK(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var list []entities.KaryaUKK
		query := db.Preload("Siswa").Preload("Kategori").Preload("Admin")

		if judul := c.Query("judul"); judul != "" {
			query = query.Where("judul ILIKE ?", "%"+judul+"%")
		}

		if err := query.Find(&list).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data karya UKK"})
		}

		var resp []GetAllKaryaUKKResponse
		for _, k := range list {
			resp = append(resp, GetAllKaryaUKKResponse{
				KaryaID:    k.ID.String(),
				Judul:      k.Judul,
				Deskripsi:  k.Deskripsi,
				Thumbnail:  k.Thumbnail,
				Link:       k.Link,
				SiswaID:    k.SiswaID.String(),
				KategoriID: k.KategoriID.String(),
				AdminID:    k.AdminID.String(),
			})
		}

		return c.JSON(GetAllKaryaUKKResponseWrapper{
			Code:    200,
			Message: "Success",
			Data:    resp,
		})
	}
}
