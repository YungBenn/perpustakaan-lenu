package peminjaman

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
	var req sqlc.CreatePeminjamanParams
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	ctx := c.Context()
	peminjaman, err := h.repo.CreatePeminjaman(ctx, req)
	if err != nil {
		return err
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"status":  http.StatusCreated,
		"message": "Create new peminjaman data success",
		"data":    peminjaman,
	})
}

func (h *Handler) List(c *fiber.Ctx) error {
	ctx := c.Context()
	peminjaman, err := h.repo.ListPeminjaman(ctx)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "List all peminjaman data success",
		"data":    peminjaman,
	})
}

func (h *Handler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	var req sqlc.UpdatePeminjamanParams
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	ctx := c.Context()
	req.ID = int64(id)
	peminjaman, err := h.repo.UpdatePeminjaman(ctx, req)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Update peminjaman data success",
		"data":    peminjaman,
	})
}

func (h *Handler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	ctx := c.Context()
	err = h.repo.DeletePeminjaman(ctx, int64(id))
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Delete peminjaman data success",
	})
}
