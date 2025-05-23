# UKK SMKN2

Requirements :
- [PostgreSQL] CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
- [Go] 1.20 up

RESTful API :
- [Go](https://golang.org/)
- [Fiber](https://github.com/gofiber/fiber)
- [GORM](https://gorm.io/)
- JWT Authentication
- Swagger Docs (via swaggo)

---

## 🔧 How to Clone

### 1. Clone
```bash
git clone https://github.com/RIZKY07606/ukk.git
cd ukk
```

### 2. Create .env file
```bash
DATABASE_URL='host=yourhost user=userName password=pass dbname=dbName port=5432 sslmode=require'
JWT_SECRET='your_jwt_secret'
PORT=8080
ENV='development'
```

### 3. Install Dependency
```bash
go mod tidy
```

### 4. Generate Swagger
```bash
go install github.com/swaggo/swag/cmd/swag@latest
swag init
```

### 5. Setup Database
```bash
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
```

### 6. Run
```bash
go run main.go
```