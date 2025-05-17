package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	auth "ukk-smkn2/features/auth/common"
	fileupload "ukk-smkn2/features/fileupload/common"
	karyaukk "ukk-smkn2/features/karyaukk/common"
	kategorikarya "ukk-smkn2/features/kategorikarya/common"
	review "ukk-smkn2/features/review/common"
	role "ukk-smkn2/features/role/common"
	siswa "ukk-smkn2/features/siswa/common"
	user "ukk-smkn2/features/user/common"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	api := app.Group("/api")

	auth.RegisterRoutes(api, db)
	user.RegisterRoutes(api, db)
	role.RegisterRoutes(api, db)
	siswa.RegisterRoutes(api, db)
	review.RegisterRoutes(api, db)
	kategorikarya.RegisterRoutes(api, db)
	karyaukk.RegisterRoutes(api, db)
	fileupload.RegisterRoutes(api, db)
}
