package services

import (
	"Loan/entity"
	"database/sql"
	"fmt"
	"time"

	// _ ....
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "admin"
	password = "xoji"
	dbname   = "loan"
)

//LoanServices ..
type LoanServices interface {
	AddCostumer(entity.Customer) int
	AddLoan(entity.Loan) int
	RepayLoan(entity.RePayLoan) int
	GetAllLoans() []entity.Loan
	GetAllCustomers() []entity.Customer
}

type loanServices struct {
	loans     []entity.Loan
	customers []entity.Customer
}

//New LoanServices ..
func New() LoanServices {
	return &loanServices{}
}

func (service *loanServices) AddCostumer(customer entity.Customer) int {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	fmt.Print(err)
	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	db.SetConnMaxLifetime(time.Second * 5)
	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(100)
	defer db.Close()
	fmt.Print(customer)
	// sql := "INSERT INTO `customer`(full_name, address, number) values($1, $2, $3) RETURNING customer_id"
	// id := 0
	// err = db.QueryRow(sql, customer.Name, customer.Address, customer.Number).Scan(&id)
	address := ""
	sqlselect := "select address from customer where customer_id = 1"
	err = db.QueryRow(sqlselect).Scan(&address)
	fmt.Println(address)
	if err != nil {
		panic(err.Error())
	}

	return 1
}
func (service *loanServices) AddLoan(entity.Loan) int {

	return 1
}
func (service *loanServices) RepayLoan(entity.RePayLoan) int {

	return 1
}
func (service *loanServices) GetAllLoans() []entity.Loan {
	return service.loans
}
func (service *loanServices) GetAllCustomers() []entity.Customer {
	return service.customers
}
