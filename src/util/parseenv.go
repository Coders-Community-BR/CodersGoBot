package util

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(name string) string {
	err := godotenv.Load()	
	if err != nil {
		log.Fatal(err)
	}

	env := os.Getenv(name)

	return env
}
