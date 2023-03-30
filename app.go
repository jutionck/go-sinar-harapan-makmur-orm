package main

import (
	"fmt"

	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/config"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/model"
)

func main() {
	c, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
	}

	dbConn, _ := config.NewDbConnection(c)
	// db := dbConn.Conn()
	dbConn.Migrate(
		&model.Brand{},
		&model.Vehicle{},
		&model.UserCredential{},
		&model.Customer{},
		&model.Employee{},
		&model.Transaction{},
	)
}
