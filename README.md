# URL Shortener

Aplikasi URL Shortener adalah layanan pemendek URL yang dibangun menggunakan Go (Golang). Aplikasi ini memungkinkan pengguna untuk membuat versi pendek dari URL panjang, mengelola URL yang telah dipendekan, dan menggunakan autentikasi Google OAuth untuk keamanan.

## Fitur Utama

- Pemendekkan URL dengan custom alias
- Autentikasi menggunakan Google OAuth
- Manajemen URL (membuat, membaca, menghapus)
- Redirect otomatis ke URL asli
- Pelacakan URL berdasarkan pengguna
- Analitik penggunaan URL
- RESTful API

## Teknologi yang Digunakan

- **Go** - Bahasa pemrograman utama
- **Gin** - Web framework
- **GORM** - ORM (Object Relational Mapping)
- **SQLite** - Database
- **JWT** - Token autentikasi
- **Google OAuth2** - Autentikasi pengguna

## Persyaratan Sistem

- Go 1.24.0 atau yang lebih baru
- SQLite

## Konfigurasi

Buat file `.env` di root proyek dengan konfigurasi berikut:

```env
BASE_URL=http://localhost:8080
GOOGLE_CLIENT_ID=your_google_client_id
GOOGLE_CLIENT_SECRET=your_google_client_secret
JWT_SECRET=your_jwt_secret
```

## Cara Menjalankan

1. Clone repositori
```bash
git clone https://github.com/Auxesia23/url_shortener.git
```

2. Masuk ke direktori proyek
```bash
cd url_shortener
```

3. Install dependensi
```bash
go mod download
```

4. Jalankan aplikasi
```bash
go run cmd/api/main.go
```

Aplikasi akan berjalan di `http://localhost:8080`

## Struktur Proyek

```
.
├── cmd/
│   └── api/
│       ├── api.go
│       └── main.go
├── internal/
│   ├── auth/
│   │   ├── google.go
│   │   └── jwt.go
│   ├── db/
│   │   └── sqlite.go
│   ├── handlers/
│   │   ├── redirect_handler.go
│   │   ├── url_handler.go
│   │   └── user_handler.go
│   ├── mapper/
│   │   ├── analytic_mapper.go
│   │   └── url_mapperr.go
│   ├── middlewares/
│   │   └── auth_middleware.go
│   ├── models/
│   │   ├── analytic.go
│   │   ├── url.go
│   │   └── user.go
│   ├── repositories/
│   │   ├── analytic_repository.go
│   │   ├── url_repository.go
│   │   └── user_repository.go
│   ├── services/
│   │   ├── analytic_service.go
│   │   ├── redirect_service.go
│   │   ├── url_service.go
│   │   └── user_service.go
│   └── utils/
│       ├── auth_validator.go
│       └── hashPassword.go
└── go.mod
```

## API Endpoints

### URL Endpoints

- `POST /api/urls` - Membuat URL pendek baru
- `GET /api/urls` - Mendapatkan daftar URL pengguna
- `DELETE /api/urls/:shortened` - Menghapus URL pendek
- `GET /:shortened` - Redirect ke URL asli
- `GET /api/urls/:shortened/analytics` - Mendapatkan statistik penggunaan URL

### Autentikasi Endpoints

- `GET /auth/google/login` - Memulai proses login Google
- `GET /auth/google/callback` - Callback URL untuk autentikasi Google

## Model Data

### URL Model
```go
type Url struct {
    Original  string
    Shortened string
    UserEmail string
    CreatedAt time.Time
}
```

### User Model
```go
type User struct {
    Email   string
    Name    string
    Picture string
}
```

### Analytic Model
```go
type Analytic struct {
    ShortenedUrl    string
    VisitedAt       time.Time
    UserAgent       string
    IPAddress       string
}
```

## Kontribusi

Kontribusi selalu diterima! Silakan buat pull request untuk perbaikan atau penambahan fitur.
