package initializer

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	// .env fileをloadし、環境変数として扱えるようにする
	err := godotenv.Load()

	if err != nil {
		log.Fatalln("Error loading .env file")
	}
}
