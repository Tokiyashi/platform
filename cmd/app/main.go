package main

import (
	"log"
	"os"
	_ "platform/docs" // важный импорт для swagger
	"platform/internal/server"

	"github.com/joho/godotenv"
)

// @title           Swagger ДЛЯ ОБРАЗОВАТЕЛЬНОЙ ПЛАТФОРМЫ (УМНИ)
// @version         1.0
// @description     Лучший в мире бэкендер, стример, блогер, фронтендер, тиктокер, лайкер и будущий актёр (Никита Стасыч Р.) забабахал бэкенд и документацию для тг-блогера, живой легенды известной в кругах фронтов по имени jormZ

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Enter your Bearer token in the format: Bearer <token>
// @scheme bearer

// @host      localhost:8080
// @BasePath  /
func main() {
	if os.Getenv("GO_ENV") != "production" {
		godotenv.Load()
	}
	if err := godotenv.Load(); err != nil {
		log.Println("Ошибка при загрузке .env файла")
		// return
	}

	s, err := server.New()

	if err != nil {
		log.Print(err)
		return
	}
	log.Println(("Server is running on port 8080"))

	s.Run()
}
