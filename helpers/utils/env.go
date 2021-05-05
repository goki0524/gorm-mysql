package utils

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

// EnvLoad : .envファイルをロードする.
func EnvLoad() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file. err:" + fmt.Sprint(err))
	}
}

// EnvLoadForTest : Test用の.envファイルをロードする.
func EnvLoadForTest() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("error loading .env file. err:" + fmt.Sprint(err))
	}
}
