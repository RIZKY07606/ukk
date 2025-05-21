package getall

import (
	"ukk-smkn2/entities"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetAllSiswa godoc
//
//	@Summary		Get all siswa
//	@Description	Get list of all siswa
//	@Tags			siswa
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	GetAllSiswaResponseWrapper
//	@Failure		500	{object}	map[string]string
//	@Router			/api/siswa [get]
func GetAllSiswa(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var siswaList []entities.Siswa
		if err := db.Find(&siswaList).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch siswa"})
		}

		var responseData []GetAllSiswaResponse
		for _, s := range siswaList {
			responseData = append(responseData, GetAllSiswaResponse{
				SiswaID:  s.ID.String(),
				NIS:      s.NIS,
				Nama:     s.Nama,
				Kelas:    s.Kelas,
				Jurusan:  s.Jurusan,
				Angkatan: s.Angkatan,
			})
		}

		return c.JSON(GetAllSiswaResponseWrapper{
			Code:    200,
			Message: "Success",
			Data:    responseData,
		})
	}
}
