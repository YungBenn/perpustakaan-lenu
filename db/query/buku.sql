-- name: ListBuku :many
SELECT * FROM buku;

-- name: CreateBuku :one
INSERT INTO buku (judul, penulis, kuantitas, tempat_penyimpanan)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateBuku :one
UPDATE buku
SET judul = $2, penulis = $3, kuantitas = $4, tempat_penyimpanan = $5
WHERE id = $1
RETURNING *;

-- name: DeleteBuku :exec
DELETE FROM buku
WHERE id = $1;
