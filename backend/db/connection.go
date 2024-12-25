package db

import (
	"fmt"
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"backend/models"
)

var DB *gorm.DB

func Init() {
	
}