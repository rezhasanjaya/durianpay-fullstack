Backend (Golang + SQLite)

## Prasyarat
- Go 1.21+
- Make

## Setup & Run
```bash
cp env.sample .env
make dep
make gen-secret
make openapi-gen
make run
```

Server berjalan di `http://localhost:8080`.

## Endpoint Utama
- `POST /dashboard/v1/auth/login`
- `GET /dashboard/v1/payments?status=...&id=...&sort=...`
- `GET /dashboard/v1/widget`

## Testing (contoh)
```bash
go test ./internal/module/payment/repository -v
```
