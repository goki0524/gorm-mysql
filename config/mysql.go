package config

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql struct {
	host     string
	port     string
	user     string
	password string
	database string
	isDegug  bool
}

func (m *Mysql) setMysqlEnv() {
	m.host = os.Getenv("DB_HOST")
	m.port = os.Getenv("DB_PORT")
	m.user = os.Getenv("DB_USER")
	m.password = os.Getenv("DB_PASSWORD")
	m.database = os.Getenv("DB_DATABASE")

	m.isDegug = false
	if os.Getenv("ENV") == envLocal && os.Getenv("DB_DEBUG") == "true" ||
		os.Getenv("DB_DEBUG") == "True" ||
		os.Getenv("DB_DEBUG") == "TRUE" {
		m.isDegug = true
	}
}

func (m *Mysql) Connection() *gorm.DB {
	m.setMysqlEnv()

	dsn := m.user + ":" + m.password + "@tcp(" + m.host + ":" + m.port + ")/" + m.database + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	if m.isDegug {
		return db.Debug()
	}
	return db
}
