package create

import (
	"regexp"
	"strings"
	"time"

	e "ukk-smkn2/entities"
	r "ukk-smkn2/infrastructure/repositories/admin"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

// Register godoc
//
//	@Summary
//	@Description	Register a new admin
//	@Tags			admin
//	@Accept			json
//	@Produce		json
//	@Param			request	body		CreateAdminRequest	true	"Create admin request body"
//	@Success		200		{object}	CreateAdminResponseWrapper
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/api/admin/register [post]
func Register(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req CreateAdminRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
		}

		if strings.TrimSpace(req.Nama) == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Name is required"})
		}

		if !emailRegex.MatchString(req.Email) {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid email format"})
		}

		if len(req.Password) < 6 {
			return c.Status(400).JSON(fiber.Map{"error": "Password must be at least 6 characters"})
		}

		hashedPassword, _ := r.HashPassword(req.Password)
		admin := e.Admin{
			ID:        uuid.New(),
			Nama:      req.Nama,
			Email:     req.Email,
			Password:  hashedPassword,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		if err := r.CreateAdmin(db, &admin); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Email already exist"})
		}
		return c.JSON(e.SuccessResponse(&CreateAdminResponse{Nama: admin.Nama, Email: admin.Email}))
	}
}
