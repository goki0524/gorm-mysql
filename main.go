package main

import (
	"fmt"
	"os"

	"github.com/goki0524/gorm-mysql/helpers/utils"
	"github.com/goki0524/gorm-mysql/server"
)

func init() {
	utils.EnvLoad()
}

func main() {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	fmt.Println("serving on " + host + ":" + port)
	server := server.Router{
		Port: port,
	}
	server.Start()

}
