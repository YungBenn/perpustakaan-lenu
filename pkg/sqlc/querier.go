// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0

package sqlc

import (
	"context"
)

type Querier interface {
	CreateBuku(ctx context.Context, arg CreateBukuParams) (Buku, error)
	CreateMahasiswa(ctx context.Context, arg CreateMahasiswaParams) (Mahasiswa, error)
	CreatePeminjaman(ctx context.Context, arg CreatePeminjamanParams) (Peminjaman, error)
	DeleteBuku(ctx context.Context, id int64) error
	DeleteMahasiswa(ctx context.Context, id int64) error
	DeletePeminjaman(ctx context.Context, id int64) error
	ListBuku(ctx context.Context) ([]Buku, error)
	ListMahasiswa(ctx context.Context) ([]Mahasiswa, error)
	ListPeminjaman(ctx context.Context) ([]Peminjaman, error)
	UpdateBuku(ctx context.Context, arg UpdateBukuParams) (Buku, error)
	UpdateMahasiswa(ctx context.Context, arg UpdateMahasiswaParams) (Mahasiswa, error)
	UpdatePeminjaman(ctx context.Context, arg UpdatePeminjamanParams) (Peminjaman, error)
}

var _ Querier = (*Queries)(nil)
