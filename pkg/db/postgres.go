package db

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type dbConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
	Port     int    `json:"port"`
	SSLMode  string `json:"sslmode"`
}

func loadConfig(path string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Ошибка при открытии config.json: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := dbConfig{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalf("Ошибка при парсинге config.json: %v", err)
	}

	str := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", config.Host, config.User, config.Password, config.DBName, config.Port, config.SSLMode)
	return str
}

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(postgres.Open(loadConfig("./config.json")), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных: ", err)
	}

	DB = db
}
