package main

import (
	"fmt"
	"os"

	"github.com/goki0524/gorm-mysql/config"
	"github.com/goki0524/gorm-mysql/helper/utils"
	"github.com/goki0524/gorm-mysql/server"
	"gorm.io/gorm"
)

func init() {
	utils.EnvLoad()
}

func main() {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	// Mysql Setup
	mysqlConfig := config.NewMysqlConfig()
	db := mysqlConfig.Connection()
	injectMysqlRepository(db)

	fmt.Println("serving on " + host + ":" + port)
	server := server.Router{
		Port: port,
	}
	server.Start()

}

func injectMysqlRepository(db *gorm.DB) {

}
