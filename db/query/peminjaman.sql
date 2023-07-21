-- name: ListPeminjaman :many
SELECT * FROM peminjaman;

-- name: CreatePeminjaman :one
INSERT INTO peminjaman (mahasiswa_id, buku_id, tanggal_peminjaman, tanggal_batas_pengembalian, tanggal_pengembalian)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdatePeminjaman :one
UPDATE peminjaman
SET mahasiswa_id = $2, buku_id = $3, tanggal_peminjaman = $4, tanggal_batas_pengembalian = $5, tanggal_pengembalian = $6
WHERE id = $1
RETURNING *;

-- name: DeletePeminjaman :exec
DELETE FROM peminjaman
WHERE id = $1;
