package main

import (
	"fmt"
	"log"

	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/config"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/model"
)

func main() {
	fmt.Println("Migration")

	c, err := config.NewConfig()
	if err != nil {
		fmt.Println(err, "", 102)
		return
	}

	db, err := config.NewDbConnection(c)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Migrate(&model.Vehicle{})
	if err != nil {
		log.Fatal(err)
	}
}
