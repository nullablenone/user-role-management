# User and Role Management API (Clean Architecture)

## Pengantar

Project ini awalnya dibuat dengan tujuan untuk membatasi akses user berdasarkan role, menggunakan JWT sebagai token untuk otorisasi dan pengecekan setiap route yang diproteksi.

Kemudian, project ini dikembangkan lebih lanjut dengan menerapkan **Clean Architecture** agar struktur kodenya lebih rapi dan terstandar. Selanjutnya, saya melengkapi project ini dengan dokumentasi **Swagger** untuk mempermudah testing API, dan terakhir mengimplementasikan **Redis** untuk caching guna meningkatkan performa.

---

## Konsep & Arsitektur

Proyek ini secara ketat mengadopsi prinsip-prinsip **Clean Architecture** untuk memisahkan *concerns* dan memastikan setiap lapisan memiliki tanggung jawab yang jelas. Arsitektur ini membuat kode lebih mudah diuji, dipelihara, dan dikembangkan seiring waktu.

Struktur lapisan pada proyek ini adalah sebagai berikut:

1.  **Domain Layer**: Inti dari aplikasi. Lapisan ini berisi *entitas* (model) dan aturan bisnis (`service`) yang tidak bergantung pada detail teknis apa pun (seperti *database* atau *framework*).
    * `internal/domain/{user,role,auth}/model.go`
    * `internal/domain/{user,role,auth}/service.go`
    * `internal/domain/{user,role}/repository.go` (Interface)

2.  **Infrastructure Layer**: Berisi semua detail teknis dan implementasi dari *interface* yang didefinisikan di *domain layer*. Ini mencakup koneksi *database* (PostgreSQL), *caching* (Redis), dan komponen eksternal lainnya.
    * `internal/infrastucture/repository/`
    * `internal/infrastucture/cache/`
    * `config/`
    * `utils/`

3.  **Presentation Layer**: Bertanggung jawab untuk menangani interaksi dengan dunia luar. Dalam proyek ini, lapisan ini diimplementasikan sebagai API *endpoint* menggunakan *framework* Gin.
    * `internal/domain/{user,role,auth}/handler.go`
    * `routes/routes.go`

Pemisahan ini memastikan bahwa logika bisnis inti (Domain) tetap murni dan tidak tercampur dengan detail implementasi teknis.

---

## Fitur Utama

-   **Otentikasi & Otorisasi Berbasis JWT**: Sistem login yang aman menghasilkan token JWT untuk akses ke *endpoint* yang dilindungi.
-   **Manajemen Pengguna (Admin)**: Operasi CRUD (Create, Read, Update, Delete) penuh untuk mengelola data pengguna.
-   **Manajemen Peran (Admin)**: Operasi CRUD untuk mengelola peran dan hak akses (`user` & `admin`).
-   **Role-Based Access Control (RBAC)**: *Middleware* untuk membatasi akses ke *endpoint* tertentu hanya untuk peran yang diizinkan.
-   **Caching Layer dengan Redis**: Implementasi *cache* pada *repository* pengguna (menggunakan *Decorator Pattern*) untuk mengurangi beban *database* dan mempercepat response time. Termasuk strategi *cache invalidation* yang cerdas.
-   **Robust Error Handling**: Menggunakan *sentinel errors* untuk membedakan kesalahan bisnis (cth: *data not found*) dan kesalahan teknis, menghasilkan respons HTTP yang lebih akurat.
-   **Dokumentasi API (Swagger)**: Dokumentasi API yang digenerasi secara otomatis dan interaktif.
-   **Password Hashing**: Menggunakan `bcrypt` untuk mengamankan *password* pengguna.
-   **Konfigurasi Terpusat**: Pengelolaan konfigurasi melalui file `.env`.

---


## Dokumentasi & Endpoint API

Dokumentasi API lengkap tersedia melalui Swagger. Setelah menjalankan aplikasi, akses URL berikut:

### Ringkasan Endpoint

| Method | Endpoint | Deskripsi | Akses |
| :--- | :--- | :--- | :--- |
| `POST` | `/register` | Registrasi pengguna baru | Publik |
| `POST` | `/login` | Login dan dapatkan token JWT | Publik |
| `GET` | `/user/profile` | Lihat profil pengguna saat ini | Pengguna (Login) |
| `GET` | `/admin/users` | Dapatkan semua pengguna | Admin |
| `POST` | `/admin/users` | Buat pengguna baru | Admin |
| `GET` | `/admin/users/{id}` | Dapatkan pengguna berdasarkan ID | Admin |
| `PUT` | `/admin/users/{id}` | Perbarui pengguna berdasarkan ID | Admin |
| `DELETE`| `/admin/users/{id}`| Hapus pengguna berdasarkan ID | Admin |
| `GET` | `/admin/roles` | Dapatkan semua peran | Admin |
| `POST` | `/admin/roles` | Buat peran baru | Admin |
| `GET` | `/admin/roles/{id}` | Dapatkan peran berdasarkan ID | Admin |
| `PUT` | `/admin/roles/{id}` | Perbarui peran berdasarkan ID | Admin |
| `DELETE`| `/admin/roles/{id}`| Hapus peran berdasarkan ID | Admin |

---

## Tumpukan Teknologi

* **Bahasa**: Golang
* **Framework**: Gin Gonic
* **Database**: PostgreSQL
* **Caching**: Redis
* **ORM**: GORM
* **Dokumentasi**: Swaggo
* **Lainnya**: `godotenv`, `jwt-go`, `bcrypt`
---

## Struktur Proyek

Berikut adalah struktur folder yang telah dirancang untuk mendukung Clean Architecture.

```
user-role-management/
├── config/
│   ├── cache.go                        # Setup koneksi Redis
│   ├── database.go                     # Setup koneksi database
│   └── env.go                          # Memuat variabel dari file .env
├── docs/
│   ├── docs.go                         # File utama yang digenerasi oleh Swaggo
│   ├── swagger.json                    # Spek OpenAPI dalam format JSON
│   └── swagger.yaml                    # Spek OpenAPI dalam format YAML
├── internal/
│   ├── domain/                         # Lapisan Domain (Inti Aplikasi)
│   │   ├── auth/
│   │   │   ├── dto.go                  # Objek transfer data untuk registrasi & login
│   │   │   ├── handler.go              # Handler untuk endpoint otentikasi
│   │   │   └── service.go              # Logika bisnis untuk otentikasi
│   │   ├── role/
│   │   │   ├── dto.go                  # DTO untuk operasi Role
│   │   │   ├── handler.go              # Handler untuk endpoint Role
│   │   │   ├── model.go                # Model domain untuk Role
│   │   │   ├── repository.go           # Interface (kontrak) untuk repository Role
│   │   │   └── service.go              # Logika bisnis untuk manajemen Role
│   │   └── user/
│   │       ├── dto.go                  # DTO untuk operasi User
│   │       ├── handler.go              # Handler untuk endpoint User
│   │       ├── model.go                # Model domain untuk User
│   │       ├── repository.go           # Interface (kontrak) untuk repository User
│   │       └── service.go              # Logika bisnis untuk manajemen User
│   ├── errors/
│   │   └── errors.go                   # Definisi custom error aplikasi
│   └── infrastucture/                  # Lapisan Infrastruktur (Detail Teknis)
│       ├── cache/
│       │   └── user_cache.go           # Implementasi cache repository untuk user
│       └── repository/
│           ├── db_models.go            # Model GORM untuk tabel 'users' & 'roles'
│           ├── role_repository.go      # Implementasi repository untuk peran
│           └── user_repository.go      # Implementasi repository untuk pengguna
├── middlewares/
│   ├── access.go                       # Middleware untuk kontrol akses berbasis peran (RBAC)
│   └── jwt.go                          # Middleware untuk validasi token JWT
├── routes/
│   └── routes.go                       # Definisi semua route API
├── seeders/
│   ├── role_seeder.go                  # Seeder untuk mengisi data peran default
│   └── user.seeder.go                  # Seeder untuk mengisi data pengguna default
├── utils/
│   ├── jwt.go                          # Utilitas pembuatan & validasi JWT
│   ├── password.go                     # Utilitas hashing & perbandingan password
│   └── response.go                     # Utilitas untuk format respons API standar
|
├── .gitignore                          # Daftar file yang diabaikan oleh Git
├── go.mod                              # Deklarasi modul Go
├── go.sum                              # File checksum modul Go
└── main.go                             # App entry point
```

*(Struktur ini didasarkan pada file yang diunggah)*

---

## Instalasi & Konfigurasi

Untuk menjalankan proyek ini secara lokal, ikuti langkah-langkah berikut:

1.  **Kloning Repositori**
    ```sh
    git clone https://github.com/nullablenone/user-role-management.git
    cd user-role-management
    ```

2.  **Konfigurasi Environment**
    Salin dari `.env.example` atau buat file `.env` baru di root proyek.
    ```env
    # Database
    DB_HOST=localhost
    DB_USER=postgres
    DB_PASS=password_anda
    DB_NAME=nama_database
    DB_PORT=5432
    DB_SSLMODE=disable

    # JWT
    SecretKey=kunci_rahasia_jwt_anda

    # Redis
    REDIS_ADDR=localhost:6379
    REDIS_PASSWORD=
    REDIS_DB=0
    ```
    *Pastikan variabel di atas diisi sesuai dengan konfigurasi lokal Anda.*

3.  **Instalasi Dependensi**
    ```sh
    go mod tidy
    ```

4.  **Jalankan Aplikasi**
    ```sh
    go run main.go
    ```
    Server akan berjalan di `http://localhost:8080`.

---

## Cara Menggunakan API

1.  **Dapatkan Token**: Lakukan `POST` request ke `/login` dengan email dan password `admin@gmail.com` / `admin@gmail.com` (data dari *seeder*).
2.  **Gunakan Token**: Salin token dari respons. Untuk mengakses *endpoint* yang dilindungi, tambahkan *header* `Authorization` dengan format `Bearer <token_anda>`.
3.  **Akses Swagger**: Buka `http://localhost:8080/swagger/index.html` dan klik tombol "Authorize" untuk memasukkan token Anda dan mencoba *endpoint* lainnya.




