package main

import (
	"fmt"

	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/config"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/repository"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/usecase"
)

func main() {
	c, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
	}

	dbConn, _ := config.NewDbConnection(c)
	db := dbConn.Conn()

	ur := repository.NewUserRepository(db)
	uuc := usecase.NewAuthenticationUseCase(ur)

	user, err := uuc.Login("admdin", "P@ssw0rd")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Selamat datang:", user.UserName)
}
