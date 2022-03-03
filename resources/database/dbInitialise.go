package database

import (
	"fmt"
	"log"
	"os"
	"rahulProj/student/beans"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {

	godotenv.Load()

	// dsn := "root:pass@tcp(127.0.0.1:3306)/projstudent?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	// fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("unable to connect to database")
	}

	DB = db

	student := beans.Student{
		Name:        "rahul",
		YearOfStudy: 1,
		Department:  "CSE",
		BloodGroup:  "A+",
		PhoneNo:     1234567891,
		Email:       "rahul@gmail.com",
		Location:    "Alwal",
	}

	db.AutoMigrate(&beans.Student{})
	if err := db.Create(&student).Error; err != nil {
		fmt.Println("db connection error2 ")
		log.Fatalln(err)
	}

}
