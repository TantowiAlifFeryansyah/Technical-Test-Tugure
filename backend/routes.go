package main

import "github.com/gin-gonic/gin"

func setupRoutes(r *gin.Engine) {
    api := r.Group("/api")
    {
        // karyawan
        api.POST("/karyawan", CreateKaryawan)
        api.GET("/karyawan", ListKaryawan)
        api.GET("/karyawan/:id", GetKaryawan)
        api.PUT("/karyawan/:id", UpdateKaryawan)
        api.DELETE("/karyawan/:id", DeleteKaryawan)

        // buku
        api.POST("/buku", CreateBuku)
        api.GET("/buku", ListBuku)
        api.GET("/buku/:id", GetBuku)
        api.PUT("/buku/:id", UpdateBuku)
        api.DELETE("/buku/:id", DeleteBuku)

        // pinjaman
        api.POST("/pinjaman", CreatePinjaman)
        api.PUT("/pinjaman/return/:id", ReturnPinjaman)
        api.GET("/pinjaman/anggota/:id", ListPinjamanByAnggota)
        api.GET("/pinjaman/overdue", ListOverdue)
        api.GET("/pinjaman/overdue/anggota", ListAnggotaOverdue)
        api.GET("/pinjaman/history/buku/:id", HistoryByBuku)
    }
}
