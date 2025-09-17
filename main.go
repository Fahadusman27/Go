package main

import (
	"log"

	. "tugas/domain/config"
	"tugas/domain/repository"
	"tugas/domain/routes"

	"fmt"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func LoggerMiddleware(c *fiber.Ctx) error {
	fmt.Println("Request:", c.Method(), c.Path())
	return c.Next()
}

func main() {
    LoadEnv()
    db := ConnectDB()
    if err := db.Ping(); err != nil {
        log.Fatal("Koneksi database gagal: ", err)
    }
	
	userRepo := repository.NewUserRepository(db)

    app := routes.NewApp(db)
	routes.AuthRoutes(app, userRepo)
	routes.Alumni(app, userRepo)
	routes.PekerjaanAlumni(app, userRepo)
    port := "3000"

    log.Fatal(app.Listen(":" + port))
}

