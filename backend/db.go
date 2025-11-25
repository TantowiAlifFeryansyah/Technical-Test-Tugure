package main

import (
    "fmt"
    "log"
    "os"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func initDB() {
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        getenv("DB_HOST", "localhost"),
        getenv("DB_USER", "postgres"),
        getenv("DB_PASS", "postgres"),
        getenv("DB_NAME", "library_db"),
        getenv("DB_PORT", "5432"),
    )
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect database: %v", err)
    }
    DB = db
    // Auto migrate
    DB.AutoMigrate(&Karyawan{}, &Buku{}, &Pinjaman{})
}

func getenv(k, def string) string {
    v := os.Getenv(k)
    if v == "" {
        return def
    }
    return v
}
