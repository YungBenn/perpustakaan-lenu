CREATE TABLE mahasiswa (
    id BIGSERIAL PRIMARY KEY,
    nama TEXT NOT NULL,
    nim TEXT NOT NULL,
    jurusan TEXT NOT NULL
);

CREATE TABLE buku (
    id BIGSERIAL PRIMARY KEY,
    judul TEXT NOT NULL,
    penulis TEXT NOT NULL,
    kuantitas INT NOT NULL,
    tempat_penyimpanan TEXT NOT NULL  
);

CREATE TABLE peminjaman (
    id BIGSERIAL PRIMARY KEY,
    mahasiswa_id BIGINT REFERENCES mahasiswa(id) ON DELETE CASCADE,
    buku_id BIGINT REFERENCES buku(id) ON DELETE CASCADE,
    tanggal_peminjaman DATE NOT NULL,
    tanggal_batas_pengembalian DATE NOT NULL,
    tanggal_pengembalian DATE 
);