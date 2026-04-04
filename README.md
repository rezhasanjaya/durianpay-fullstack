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

## Menjalankan Backend
```bash
cd backend
cp env.sample .env
make dep
make gen-secret
make openapi-gen
make run
```
Backend berjalan di `http://localhost:8080`.

## Menjalankan Frontend
```bash
cd frontend
npm install
npm run dev
```
Frontend berjalan di `http://localhost:3001`.

## Testing Backend (contoh)
```bash
cd backend
go test ./internal/module/payment/repository -v
```

## API
Lihat `openapi.yaml` di root.
