# Durianpay Fullstack Take-Home

Repository ini berisi:
- `backend/` (Golang + SQLite)
- `frontend/` (Vue + Vite)

## Cara Install (Langkah Ringkas & Jelas)

### 1) Buka Terminal dan masuk ke folder kerja
Contoh buat folder kerja di Desktop:
```bash
cd ~/Desktop
mkdir proyek
cd proyek
```

### 2) Clone project
```bash
git clone <URL_REPO_KAMU>
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

## Install via Docker (branch `docker`)
Kamu sudah siapkan branch `docker`. Langkahnya:
```bash
git checkout docker
```
Lalu jalankan sesuai file docker di branch tersebut. Umumnya:
```bash
docker compose up --build
```
Jika nama file berbeda, ikuti README di branch `docker`.

## Testing Backend (contoh)
```bash
cd backend
go test ./internal/module/payment/repository -v
```

## API
Lihat `openapi.yaml` di root.
