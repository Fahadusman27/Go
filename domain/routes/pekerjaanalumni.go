package routes

import (
	"tugas/domain/service"
	"github.com/gofiber/fiber/v2"
	"tugas/domain/repository"
	."tugas/domain/middleware"
)

func PekerjaanAlumni(app *fiber.App, userRepo repository.UserRepository) {
    app.Get("/pekerjaan", JWTAuth(userRepo), RequireRole("admin", "user"), service.GetAllPerkajaanAlumniService)
    app.Get("/pekerjaan/:id", JWTAuth(userRepo), RequireRole("admin", "user"), service.CheckPerkajaanAlumniService)
    app.Get("/pekerjaan/alumni/:id_alumni", JWTAuth(userRepo), RequireRole("admin"), service.CheckPerkajaanAlumniService)
    app.Post("/pekerjaan", JWTAuth(userRepo), RequireRole("admin"), service.CreatePerkajaanAlumniService)
    app.Put("/pekerjaan/:id_alumni", JWTAuth(userRepo), RequireRole("admin"), service.UpdatePerkajaanAlumniService)
    app.Delete("/pekerjaan/:id_alumni", JWTAuth(userRepo), RequireRole("admin"), service.DeletePerkajaanAlumniService)
}
