# Durianpay Fullstack Take-Home

Repository ini berisi:
- `backend/` (Golang + SQLite)
- `frontend/` (Vue + Vite)

## Prasyarat
- Go 1.21+
- Node.js 20+
- Make

## Struktur
```
.
├─ backend/
├─ frontend/
└─ openapi.yaml
```

## Cara Install

### 1) Buka Terminal dan masuk ke folder kerja
Contoh buat folder kerja di Desktop:
```bash
cd ~/Desktop
mkdir proyek
cd proyek
```

### 2) Clone project
```bash
git clone https://github.com/rezhasanjaya/durianpay-fullstack.git
cd fullstack-boilerplate-master
```

### 3) Jalankan Backend (tanpa Docker)
```bash
cd backend
cp env.sample .env
make dep
make gen-secret
make openapi-gen
make run
```
Backend: `http://localhost:8080`

### 4) Jalankan Frontend (tanpa Docker)
Buka terminal baru:
```bash
cd ~/Desktop/proyek/fullstack-boilerplate-master/frontend
cp .env.example .env
npm install
npm run dev
```
Frontend: `http://localhost:3001`  
Login: `http://localhost:3001/login`

## Install via Docker (branch `docker`)
Silakan untuk fetch dockerfile pada branch `docker`:

Buka Terminal
```bash
git checkout -b docker
# jika branch sudah ada:
git checkout docker
```
Untuk Menjalankannya Masukan Perintah:
```bash
docker compose up --build
```
Setelah `docker compose up --build` berjalan, tekan tombol `v` di terminal (Docker Desktop shortcut) untuk membuka web.

**Akses setelah Docker berjalan:**
- Frontend: `http://localhost:3001`
- Login: `http://localhost:3001/login`
- Backend: `http://localhost:8080`

**Akun demo:**
- `cs@test.com` / `password`
- `operation@test.com` / `password`

## Testing Backend
```bash
cd backend
go test ./internal/module/payment/repository -v
```

## API
Lihat `openapi.yaml` di root.
