package routes

import (
	"tugas/domain/service"
	"github.com/gofiber/fiber/v2"
	"tugas/domain/repository"
	."tugas/domain/middleware"
)

func Alumni(app *fiber.App, userRepo repository.UserRepository) {
    app.Get("/alumni", JWTAuth(userRepo), RequireRole("admin", "user"), service.GetAllAlumniService)
    app.Get("/alumni/:nim", JWTAuth(userRepo), RequireRole("admin", "user"), service.CheckAlumniService)
    app.Post("/alumni", JWTAuth(userRepo), RequireRole("admin"), service.CreateAlumniService)
    app.Put("/alumni/:nim", JWTAuth(userRepo), RequireRole("admin"), service.UpdateAlumniService)
    app.Delete("/alumni/:nim", JWTAuth(userRepo), RequireRole("admin"), service.DeleteAlumniService)
}
