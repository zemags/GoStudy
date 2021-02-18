package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/zemags/GoStudy/repeat_study/modern_solutions/first_web_app/app/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type planetMass struct {
	gorm.Model
	id       int
	planetId int
	radius   int
}

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	conf := config.NewConfig()
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", conf.DBParams.User, conf.DBParams.Password, conf.DBParams.DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect db")
	}

	// create table and migrate
	db.AutoMigrate(&Comet{})

	// insert val
	for _, comet := range comets {
		db.Create(&comet)
	}

	// read
	var comet1 Comet
	db.First(&comet1, 2) // finc comet with integer primary key 2
	fmt.Println(comet1)
	db.First(&comet1, "comet_short_name = ?", "Lovejoy")
	fmt.Println(comet1.CometShortName)

	// update
	db.Model(&comet1).Where("id = ?", 2).Update("comet_short_name", "Lovejoy2")

	// delelte
	db.Delete(&comets, []int{3, 4, 5, 6, 7})
}
