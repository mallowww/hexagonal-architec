package main

import (
	"bank/repository"
	"bank/service"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	url := os.Getenv("DATABASE_URL")
	db, err := sqlx.Open("mysql", url)
	if err != nil {
		panic(err)
	}

	customerRepository := repository.NewCustomerRepositoryDB(db)
	customerService := service.NewCustomerService(customerRepository)

	customers, err := customerService.GetCustomer(2000)
	if err != nil {
		panic(err)
	}
	fmt.Println(customers)
}
