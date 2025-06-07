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
- **PostgreSQL** - Database
- **JWT** - Token autentikasi
- **Google OAuth2** - Autentikasi pengguna
- **Docker** & **Docker Compose** - Deployment

## Persyaratan Sistem

- Go 1.24.0 atau yang lebih baru (jika menjalankan tanpa Docker)
- Docker & Docker Compose (jika ingin menjalankan dengan container)
- PostgreSQL

## Konfigurasi

Buat file `.env` di root proyek dengan konfigurasi berikut:

```env
GOOGLE_CLIENT_ID=your_google_client_id
GOOGLE_CLIENT_SECRET=your_google_client_secret
GOOGLE_REDIRECT_URI=http://localhost:8080/v1/auth/google/callback
SECRET_KEY=your_jwt_secret
BASE_URL=http://localhost:8080
IPINFO_TOKEN=your_ipinfo_token

DATABASE_URL=postgres://postgres:postgres@postgres:5432/postgres?sslmode=disable
POSTGRES_DB=postgres
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
```

> **Catatan:**  
> - Pastikan `DATABASE_URL` sesuai dengan konfigurasi pada `compose.yaml`.
> - Untuk development, Anda bisa menggunakan nilai default seperti di atas.

## Cara Menjalankan

### Dengan Docker Compose

1. **Clone repositori**
    ```bash
    git clone https://github.com/Auxesia23/url_shortener.git
    cd url_shortener
    ```

2. **Salin file contoh environment**
    ```bash
    cp .env.example .env
    # Edit .env sesuai kebutuhan
    ```

3. **Jalankan aplikasi**
    ```bash
    docker compose up --build
    ```

Aplikasi akan berjalan di `http://localhost:8080`

### Tanpa Docker

1. **Install dependensi**
    ```bash
    go mod download
    ```

2. **Jalankan PostgreSQL**  
   Pastikan database PostgreSQL sudah berjalan dan environment variable sudah diatur.

3. **Jalankan aplikasi**
    ```bash
    go run cmd/api/main.go
    ```

## Struktur Proyek

```
.
├── cmd/
│   └── api/
│       ├── api.go
│       └── main.go
├── internal/
│   ├── auth/
│   ├── db/
│   ├── handlers/
│   ├── mapper/
│   ├── middlewares/
│   ├── models/
│   ├── repositories/
│   ├── services/
│   └── utils/
├── Dockerfile
├── compose.yaml
├── go.mod
├── go.sum
├── .env.example
└── README.md
```

## API Endpoints

### URL Endpoints

- `POST /v1/urls` - Membuat URL pendek baru
- `GET /v1/urls` - Mendapatkan daftar URL pengguna
- `DELETE /v1/urls/:shortened` - Menghapus URL pendek
- `GET /:shortened` - Redirect ke URL asli
- `GET /v1/urls/:shortened` - Mendapatkan detail URL & analitik

### Autentikasi Endpoints

- `GET /v1/auth/google` - Memulai proses login Google
- `GET /v1/auth/google/callback` - Callback URL untuk autentikasi Google

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
    ShortenedUrl string
    IpAddress    string
    Country      string
    UserAgent    string
}
```

## Kontribusi

Kontribusi selalu diterima! Silakan buat pull request untuk perbaikan atau penambahan fitur.
