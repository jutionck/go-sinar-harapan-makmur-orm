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

	er := repository.NewEmployeeRepository(db)
	euc := usecase.NewEmployeeUseCase(er)

	// managerID := "31725863-8375-47ed-8ac5-b82afef34028"
	// employee := &model.Employee{
	// 	FirstName:   "Fadli",
	// 	LastName:    "Rahman",
	// 	Address:     "Cianjur",
	// 	Email:       "fadli.rahman@enigmacamp.com",
	// 	PhoneNumber: "082180221161",
	// 	Bod:         time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC),
	// 	Position:    "Sales Pemasaran",
	// 	Salary:      10000000,
	// 	ManagerID:   &managerID,
	// 	UserCredential: model.UserCredential{
	// 		UserName: "fadli.rahmand@enigmacamp.com",
	// 		Password: "P@ssw0rd",
	// 		IsActive: false,
	// 	},
	// }

	// employees, _ := euc.FindAll()
	// for _, employee := range employees {
	// 	fmt.Println("Karyawan:", employee.FirstName, employee.LastName)
	// 	if employee.Manager != nil {
	// 		fmt.Println("Atasan:", employee.Manager.FirstName, employee.Manager.LastName)
	// 	}
	// 	fmt.Println()
	// }

	employees, _ := euc.FindAllEmployeeByManager("31725863-8375-47ed-8ac5-b82afef34028")
	for _, employee := range employees {
		fmt.Println("Karyawan:", employee.FirstName, employee.LastName)
		fmt.Println()
	}

}
