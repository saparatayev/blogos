package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	PORT      = 0
	DB_DRIVER = ""
	DBURL     = ""
)

func Load() {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	PORT, err = strconv.Atoi(os.Getenv("API_PORT"))

	if err != nil {
		// log.Println(err)
		PORT = 9000
	}

	DB_DRIVER = os.Getenv("DB_DRIVER")

	DBURL = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc-local", os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))
}
