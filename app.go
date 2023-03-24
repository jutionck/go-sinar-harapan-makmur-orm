package main

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Declare Model
// Conventions
// -> Jika di struct fieldny ProductionYear maka otomatis di db production_year
// -> Untuk nama tabel akan di convert ke dalam plural, vehicle -> vehicles
type Vehicle struct {
	Brand          string
	Model          string
	ProductionYear int
	Color          string
	IsAutomatic    bool
	Stock          int
	SalePrice      int
	Status         string // enum: "Baru" & "Bekas"
	BaseModel
}

type BaseModel struct {
	ID        string `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (v *Vehicle) IsValidStatus() bool {
	return v.Status == "Baru" || v.Status == "Bekas"
}

func main() {
	fmt.Println("Getting Started")

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

	// Migrate
	// Akan terbuat tabel otomatis
	// Debug() -> Mengaktifkan log agar terlihat di balik layar
	conn.Debug().AutoMigrate(&Vehicle{})

}
