package main

import (
    "time"
)

type Karyawan struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    Nama      string    `gorm:"type:varchar(100)" json:"nama"`
    Divisi    string    `gorm:"type:text" json:"divisi"`
    CreatedAt time.Time `json:"created_at"`
}

type Buku struct {
    ID          uint   `gorm:"primaryKey" json:"id"`
    Judul       string `gorm:"type:varchar(200)" json:"judul"`
    Pengarang   string `gorm:"type:varchar(200)" json:"pengarang"`
    KodeBuku    string `gorm:"type:varchar(50);uniqueIndex" json:"kode_buku"`
    JumlahStock int    `json:"jumlah_stock"`
}

type Pinjaman struct {
    ID               uint       `gorm:"primaryKey" json:"id"`
    IdAnggota        uint       `json:"id_anggota"`
    IdBuku           uint       `json:"id_buku"`
    TglPinjam        time.Time  `json:"tgl_pinjam"`
    TglKembalikan    *time.Time `json:"tgl_kembalikan"`
    StatusPeminjaman bool       `json:"status_peminjaman"` // true = returned
}
