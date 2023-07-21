package mahasiswa

import (
	"net/http"
	"perpustakaan-lenu/internal/repository"
	"perpustakaan-lenu/pkg/sqlc"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Handler struct {
	repo repository.Repository
	db   *pgxpool.Pool
}

func NewHandler(repo repository.Repository, db *pgxpool.Pool) *Handler {
	return &Handler{repo, db}
}

func (h *Handler) Create(c *fiber.Ctx) error {
	var req sqlc.CreateMahasiswaParams
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	ctx := c.Context()
	mahasiswa, err := h.repo.CreateMahasiswa(ctx, req)
	if err != nil {
		return err
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"status":  http.StatusCreated,
		"message": "Create new mahasiswa success",
		"data":    mahasiswa,
	})
}

func (h *Handler) List(c *fiber.Ctx) error {
	ctx := c.Context()
	mahasiswa, err := h.repo.ListMahasiswa(ctx)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "List all mahasiswa success",
		"data":    mahasiswa,
	})
}

func (h *Handler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	var req sqlc.UpdateMahasiswaParams
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	ctx := c.Context()
	req.ID = int64(id)
	mahasiswa, err := h.repo.UpdateMahasiswa(ctx, req)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Update mahasiswa success",
		"data":    mahasiswa,
	})
}

func (h *Handler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	ctx := c.Context()
	err = h.repo.DeleteMahasiswa(ctx, int64(id))
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Delete mahasiswa success",
	})
}
