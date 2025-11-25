# 📚 Tugure Library System  
Backend (Go) + Frontend (React) + PostgreSQL

Sistem perpustakaan sederhana yang memungkinkan:
- CRUD Karyawan
- CRUD Buku
- Transaksi Peminjaman (CRU)
- Daftar peminjam yang belum mengembalikan buku
- Detail pinjaman per anggota
- History peminjaman per buku

---

## 🚀 1. Backend (Go)

### **1.1. Konfigurasi Environment**
Edit file:

```
backend/.env
```

Set database credentials Anda:

```
DB_HOST=localhost
DB_USER=postgres
DB_PASS=yourpassword
DB_NAME=library_db
DB_PORT=5432
PORT=8080
```

---

### **1.2. Jalankan Backend**

```
cd backend
go mod tidy
go run *.go
```

Backend berjalan pada:

```
http://localhost:8080
```

Semua endpoint berada di bawah prefix `/api`.

---

## 🌐 2. Frontend (React + Tailwind)

### **2.1. Install dependencies**

```
cd frontend
npm install
```

### **2.2. Jalankan frontend**

```
npm start
```

Frontend berjalan di:

```
http://localhost:3000
```

Frontend menggunakan environment variable:

```
REACT_APP_API_URL=http://localhost:8080
```

---

## 🗄 3. Database (PostgreSQL)

### **3.1. Import Schema**

```
db/schema.sql
```

Import dengan psql:

```
psql -U postgres -d library_db -f db/schema.sql
```

### **3.2. Backup**

```
db/backup.sql
```

---

## 📌 Endpoint Penting

### Daftar Peminjam Belum Mengembalikan
```
GET /api/pinjaman/overdue/anggota
```

### Detail Pinjaman Anggota
```
GET /api/pinjaman/anggota/:id
```

Semua endpoint ada di:

```
backend/routes.go
```

---

## ✔ Final Deliverables

- Backend Go Project (`/backend`)
- Frontend React (`/frontend`)
- Database schema & backup (`/db/schema.sql`, `db/backup.sql`)
- README.md (file ini)
