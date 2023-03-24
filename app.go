package main

import (
	"fmt"

	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/config"
)

func main() {
	fmt.Println("Accessing Database")
	// Get Connection DB From Config

	c, err := config.NewConfig()
	if err != nil {
		fmt.Println(err, "", 102)
		return
	}

	config.NewDbConnection(c)

}
