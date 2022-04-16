package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	Name    string
	Surname string
}

func main() {
	db, err := gorm.Open(postgres.Open("postgres://user:pwd@localhost:5432/gorm-db?sslmode=disable"))
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&Person{})
	if err != nil {
		log.Fatal(err)
	}

	// db.Create(&Person{Name: "John", Surname: "Smith"})
	// db.Create(&Person{Name: "Fred"})

	var count int64
	db.Model(&Person{}).Where("name = ?", "John").Count(&count)
	log.Println(count)
}
