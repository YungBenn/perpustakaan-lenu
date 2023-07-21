package routes

import (
	"perpustakaan-lenu/internal/buku"
	"perpustakaan-lenu/internal/mahasiswa"
	"perpustakaan-lenu/internal/peminjaman"
	"perpustakaan-lenu/internal/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupMahasiswaRoutes(app *fiber.App, db *pgxpool.Pool) {
	repo := repository.NewRepository(db)
	handler := mahasiswa.NewHandler(repo, db)

	router := app.Group("/api/v1/mahasiswa")

	router.Post("/", handler.Create)
	router.Get("/", handler.List)
	router.Put("/:id", handler.Update)
	router.Delete("/:id", handler.Delete)
}

func SetupBukuRoutes(app *fiber.App, db *pgxpool.Pool) {
	repo := repository.NewRepository(db)
	handler := buku.NewHandler(repo, db)

	router := app.Group("/api/v1/buku")

	router.Post("/", handler.Create)
	router.Get("/", handler.List)
	router.Put("/:id", handler.Update)
	router.Delete("/:id", handler.Delete)
}

func SetupPeminjamanRoutes(app *fiber.App, db *pgxpool.Pool) {
	repo := repository.NewRepository(db)
	handler := peminjaman.NewHandler(repo, db)

	router := app.Group("/api/v1/peminjaman")

	router.Post("/", handler.Create)
	router.Get("/", handler.List)
	router.Put("/:id", handler.Update)
	router.Delete("/:id", handler.Delete)
}