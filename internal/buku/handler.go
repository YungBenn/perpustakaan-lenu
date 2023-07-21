package buku

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
	var req sqlc.CreateBukuParams
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	ctx := c.Context()
	buku, err := h.repo.CreateBuku(ctx, req)
	if err != nil {
		return err
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"status":  http.StatusCreated,
		"message": "Create new buku success",
		"data":    buku,
	})
}

func (h *Handler) List(c *fiber.Ctx) error {
	ctx := c.Context()
	buku, err := h.repo.ListBuku(ctx)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "List all buku success",
		"data":    buku,
	})
}

func (h *Handler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	var req sqlc.UpdateBukuParams
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	ctx := c.Context()
	req.ID = int64(id)
	buku, err := h.repo.UpdateBuku(ctx, req)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Update buku success",
		"data":    buku,
	})
}

func (h *Handler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	ctx := c.Context()
	err = h.repo.DeleteBuku(ctx, int64(id))
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Delete buku success",
	})
}
