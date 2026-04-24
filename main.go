package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"

	"absensi_karyawan/auth"
	"absensi_karyawan/database"
)

func main() {
	app := fiber.New()

	// middleware global - allow all origins
	app.Use(cors.New())

	// koneksi DB
	db := database.ConnectDB()

	// inject dependency
	repo := &auth.Repository{DB: db}
	service := &auth.Service{Repo: repo}
	handler := &auth.Handler{Service: service}

	// rate limit khusus login
	loginlimiter := limiter.New(limiter.Config{
		Max:        15,
		Expiration: 15 * time.Minute,
	})

	// routes auth
	app.Post("/login", loginlimiter, handler.Login)
	app.Post("/refresh", handler.Refresh)
	app.Post("/logout", handler.Logout)
	app.Post("/admin/create-user",
		auth.AuthMiddleware,
		auth.AdminOnly,
		handler.CreateUser,
	)

	// protected route contoh
	app.Get("/profile", auth.AuthMiddleware, func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"user_id": c.Locals("user_id"),
			"role":    c.Locals("role"),
		})
	})

	log.Println("Server running on http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}
