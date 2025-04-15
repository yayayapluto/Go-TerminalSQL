Pastikan sudah ada Go Modules
```bash
go mod init .
```

Lalu install GORM dan driver database-nya:
```bash
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```

## 2. Menjalankan `main.go`
Jalankan file `main.go` pake command:
```bash
go run main.go
```
