-- name: ListMahasiswa :many
SELECT * FROM mahasiswa;

-- name: CreateMahasiswa :one
INSERT INTO mahasiswa (nama, nim, jurusan)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateMahasiswa :one
UPDATE mahasiswa
SET nama = $2, nim = $3, jurusan = $4
WHERE id = $1
RETURNING *;

-- name: DeleteMahasiswa :exec
DELETE FROM mahasiswa
WHERE id = $1;
