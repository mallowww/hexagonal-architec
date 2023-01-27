package main

import (
	"bank/repository"
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

	_ = customerRepository

	customer, err := customerRepository.GetById(2000)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)


	// customers, err := customerRepository.GetAll()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(customers)
}
