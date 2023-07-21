package main

import (
	"context"
	"log"
	"os"
	"perpustakaan-lenu/internal/routes"
	"perpustakaan-lenu/pkg/database"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func buildServer() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := database.Connect(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		return err
	}

	app := fiber.New()

	routes.SetupMahasiswaRoutes(app, db)
	routes.SetupBukuRoutes(app, db)
	routes.SetupPeminjamanRoutes(app, db)

	port := os.Getenv("PORT")
	app.Listen(":" + port)

	return nil
}

func main() {
	if err := buildServer(); err != nil {
		panic(err)
	}
}
