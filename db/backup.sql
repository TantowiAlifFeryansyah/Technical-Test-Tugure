-- Schema for library_db (PostgreSQL)
CREATE TABLE IF NOT EXISTS tbl_karyawan (
  id SERIAL PRIMARY KEY,
  nama VARCHAR(100),
  divisi TEXT,
  created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE IF NOT EXISTS tbl_buku (
  id SERIAL PRIMARY KEY,
  judul VARCHAR(200),
  pengarang VARCHAR(200),
  kode_buku VARCHAR(50) UNIQUE,
  jumlah_stock INT DEFAULT 1
);

CREATE TABLE IF NOT EXISTS tbl_pinjaman (
  id SERIAL PRIMARY KEY,
  id_anggota INT REFERENCES tbl_karyawan(id),
  id_buku INT REFERENCES tbl_buku(id),
  tgl_pinjam TIMESTAMP,
  tgl_kembalikan TIMESTAMP NULL,
  status_peminjaman BOOLEAN DEFAULT false
);

-- sample data
INSERT INTO tbl_karyawan (nama, divisi) VALUES ('Andi', 'IT'), ('Budi', 'HR'), ('Sari', 'Finance');
INSERT INTO tbl_buku (judul, pengarang, kode_buku, jumlah_stock) VALUES
  ('Belajar Go', 'Penulis A', 'BK-001', 2),
  ('Pemrograman Web', 'Penulis B', 'BK-002', 1),
  ('Basis Data', 'Penulis C', 'BK-003', 1);
