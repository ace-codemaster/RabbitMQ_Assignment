package dbhelper

import (
	"EasternEnterprise/helpers/logginghelper"
	"EasternEnterprise/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var HotelDB *gorm.DB

func init() {
	// logginghelper.Info("Creating Database Connection")
	// dsn := "test:test@tcp(10.15.20.233:3306)/Test?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open("mysql", "test:test@tcp(10.15.20.233:3306)/Test?charset=utf8mb4&parseTime=True&loc=Local")
	// if err != nil {
	// 	fmt.Println("Error:>>>>>", err)
	// 	panic("Cant initialise db")
	// }
	// dbhelper.HotelDB = db

	// create a database connection

	logginghelper.Info("Creating Database Connection")

	//dbhelper.InitiateDB()

	dsn := "test:test@tcp(10.15.20.233:3306)/Test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logginghelper.Error("Error:>>>>>", err)
		panic("Cant initialise db")
	}
	HotelDB = db

	HotelDB.AutoMigrate(&model.Hotel{})
	HotelDB.AutoMigrate(&model.Room{})
	HotelDB.AutoMigrate(&model.RatePlan{})
	logginghelper.Info("Database Connection is created")

}
