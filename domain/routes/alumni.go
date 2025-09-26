package routes

import (
	. "tugas/domain/middleware"
	"tugas/domain/model"
	"tugas/domain/service"

	"github.com/gofiber/fiber/v2"
)

func Alumni(app *fiber.App, userRepo *model.UserRepository) {
    app.Get("/alumni", JWTAuth(userRepo), RequireRole("admin", "user"), service.GetAllAlumniService)
    app.Get("/alumni/:nim", JWTAuth(userRepo), RequireRole("admin", "user"), service.CheckAlumniService)
    app.Post("/alumni", JWTAuth(userRepo), RequireRole("admin"), service.CreateAlumniService)
    app.Put("/alumni/:nim", JWTAuth(userRepo), RequireRole("admin"), service.UpdateAlumniService)
    app.Delete("/alumni/:nim", JWTAuth(userRepo), RequireRole("admin"), service.DeleteAlumniService)
}
