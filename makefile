deps:
	go get -u golang.org/x/crypto/bcrypt
	go get -u github.com/joho/godotenv
	go get -u github.com/golang-jwt/jwt/v4
	go get -u gorm.io/gorm
	go get -u gorm.io/driver/postgres
	go get github.com/labstack/echo/v4

.PHONY: all
all: deps