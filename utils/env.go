package utils

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// LoadEnvFile load the .env file
func LoadEnvFile() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	val := os.Getenv("PORT")
	fmt.Println(val)

}

// GetENVasString get a tag from .env as string (default value)
func GetENVasString(val string) string {
	return os.Getenv(val)
}

// GetENVasINT convert the string from .env tag to INT
func GetENVasINT(val string) int {

	str := os.Getenv(val)
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
	}

	return num
}
