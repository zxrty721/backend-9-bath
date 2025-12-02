package config

import (
	"backend/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	godotenv.Load()

	dsn := os.Getenv("DATABASE_URL")
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto Migrate: สร้างตารางให้อัตโนมัติ ไม่ต้องเข้า phpMyAdmin ไปสร้างเอง
	DB.AutoMigrate(&models.User{}, &models.Product{}, &models.ProductLog{})
	fmt.Println("Database Connected & Migrated!")
}
