package main

import (
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
)

// Karyawan CRUD
func CreateKaryawan(c *gin.Context) {
    var k Karyawan
    if err := c.ShouldBindJSON(&k); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := DB.Create(&k).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, k)
}

func ListKaryawan(c *gin.Context) {
    var list []Karyawan
    DB.Find(&list)
    c.JSON(http.StatusOK, list)
}

func GetKaryawan(c *gin.Context) {
    var k Karyawan
    id := c.Param("id")
    if err := DB.First(&k, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "not found"}); return
    }
    c.JSON(http.StatusOK, k)
}

func UpdateKaryawan(c *gin.Context) {
    var k Karyawan
    id := c.Param("id")
    if err := DB.First(&k, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "not found"}); return
    }
    var payload Karyawan
    if err := c.ShouldBindJSON(&payload); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return
    }
    k.Nama = payload.Nama
    k.Divisi = payload.Divisi
    DB.Save(&k)
    c.JSON(http.StatusOK, k)
}

func DeleteKaryawan(c *gin.Context) {
    id := c.Param("id")
    DB.Delete(&Karyawan{}, id)
    c.Status(http.StatusNoContent)
}

// Buku CRUD
func CreateBuku(c *gin.Context) {
    var b Buku
    if err := c.ShouldBindJSON(&b); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return
    }
    if err := DB.Create(&b).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return
    }
    c.JSON(http.StatusCreated, b)
}

func ListBuku(c *gin.Context) {
    var list []Buku
    DB.Find(&list)
    c.JSON(http.StatusOK, list)
}

func GetBuku(c *gin.Context) {
    var b Buku
    id := c.Param("id")
    if err := DB.First(&b, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "not found"}); return
    }
    c.JSON(http.StatusOK, b)
}

func UpdateBuku(c *gin.Context) {
    var b Buku
    id := c.Param("id")
    if err := DB.First(&b, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "not found"}); return
    }
    var payload Buku
    if err := c.ShouldBindJSON(&payload); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return
    }
    b.Judul = payload.Judul
    b.Pengarang = payload.Pengarang
    b.KodeBuku = payload.KodeBuku
    b.JumlahStock = payload.JumlahStock
    DB.Save(&b)
    c.JSON(http.StatusOK, b)
}

func DeleteBuku(c *gin.Context) {
    id := c.Param("id")
    DB.Delete(&Buku{}, id)
    c.Status(http.StatusNoContent)
}

// Pinjaman (Create, Return/Update, List)
func CreatePinjaman(c *gin.Context) {
    var p Pinjaman
    if err := c.ShouldBindJSON(&p); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return
    }
    // check if anggota exists
    var k Karyawan
    if err := DB.First(&k, p.IdAnggota).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "anggota not found"}); return
    }
    // check if buku exists
    var b Buku
    if err := DB.First(&b, p.IdBuku).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "buku not found"}); return
    }
    // check active loans count for anggota
    var count int64
    DB.Model(&Pinjaman{}).Where("id_anggota = ? AND status_peminjaman = false", p.IdAnggota).Count(&count)
    if count >= 5 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "limit 5 buku aktif"}); return
    }
    p.TglPinjam = time.Now()
    p.StatusPeminjaman = false
    if err := DB.Create(&p).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return
    }
    c.JSON(http.StatusCreated, p)
}

func ReturnPinjaman(c *gin.Context) {
    id := c.Param("id")
    var p Pinjaman
    if err := DB.First(&p, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "not found"}); return
    }
    t := time.Now()
    p.TglKembalikan = &t
    p.StatusPeminjaman = true
    DB.Save(&p)
    c.JSON(http.StatusOK, p)
}

func ListPinjamanByAnggota(c *gin.Context) {
    id := c.Param("id")
    var list []Pinjaman
    DB.Where("id_anggota = ?", id).Find(&list)
    c.JSON(http.StatusOK, list)
}

// History per buku
func HistoryByBuku(c *gin.Context) {
    id := c.Param("id")
    var list []Pinjaman
    DB.Where("id_buku = ?", id).Order("tgl_pinjam desc").Find(&list)
    c.JSON(http.StatusOK, list)
}

// List overdue loans (detailed)
func ListOverdue(c *gin.Context) {
    cutoff := time.Now().AddDate(0, 0, -14)
    var list []Pinjaman
    DB.Where("status_peminjaman = false AND tgl_pinjam < ?", cutoff).Find(&list)
    c.JSON(http.StatusOK, list)
}

// List distinct anggota who have overdue books
func ListAnggotaOverdue(c *gin.Context) {
    cutoff := time.Now().AddDate(0, 0, -14)
    type Row struct {
        IdAnggota uint   `json:"id_anggota"`
        Nama      string `json:"nama"`
        Divisi    string `json:"divisi"`
    }
    var rows []Row
    DB.Raw(`SELECT DISTINCT k.id as id_anggota, k.nama, k.divisi FROM tbl_pinjaman p JOIN tbl_karyawan k ON p.id_anggota = k.id WHERE p.status_peminjaman = false AND p.tgl_pinjam < ?`, cutoff).Scan(&rows)
    c.JSON(http.StatusOK, rows)
}
