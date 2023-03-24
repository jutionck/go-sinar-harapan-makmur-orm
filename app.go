package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Accessing Database")
	// Koneksi DB
	host := "localhost"
	user := "postgres"
	password := "P@ssw0rd"
	database := "db_sinar_harapan_makmur_v2"
	port := "5432"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, password, database, port)
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db, err := conn.DB()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
}
