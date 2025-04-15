package service

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type connService struct {
	DB *gorm.DB
}

func NewConnService() *connService {
	return &connService{}
}

func DefaultConnection(c *connService) {
	dsn := "root:@tcp(127.0.0.1:3306)/?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		panic("❌ Failed to connect database")
	}

	c.DB = db
}
