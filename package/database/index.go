package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Hero struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Race         string
	Lifepoint    int
	Armor        int
	Force        int
	Vitesse      int
	Constitution int
	Magie        int
	Controle     int
	Inteligence  int
	Sagesse      int
	Charisme     int
}

func init() {

	var user string = "root"
	var password string = "example"
	var host string = "localhost"
	var port int = 3302
	var dbname string = "test02"

	option := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, dbname)

	db, err := gorm.Open(mysql.Open(option), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	// Auto Migrate
	db.AutoMigrate(&Hero{})

	// hero := Hero{0, "goku", "sayan", 0, 0}

	var hero Hero

	db.First(&hero)

	db.Model(&hero).Updates(Hero{Force: 15, Vitesse: 13, Constitution: 14, Controle: 10, Inteligence: 8, Sagesse: 11, Charisme: 10, Magie: 7, Armor: 1, Lifepoint: 20}).Where("race='sayan'")

	// db.Create(&hero)

}
