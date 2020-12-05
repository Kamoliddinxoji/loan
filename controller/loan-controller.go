package controller

import (
	"Loan/entity"
	"Loan/services"

	"github.com/gin-gonic/gin"
)

//LoanController ....
type LoanController interface {
	AddCustomer(ctx *gin.Context) (int, error)
	AddLoan(ctx *gin.Context) (int, error)
	RePayLoan(ctx *gin.Context) (int, error)
	GetAllLoans() []entity.Loan
	GetAllCustomers() []entity.Customer
}
type controller struct {
	services services.LoanServices
}

//New LoanController ...
func New(services services.LoanServices) LoanController {
	return &controller{
		services: services,
	}
}

//AddCustomer ...
func (c *controller) AddCustomer(ctx *gin.Context) (int, error) {
	var customer entity.Customer
	err := ctx.ShouldBindJSON(&customer)
	if err != nil {
		return 0, err
	}
	id := c.services.AddCostumer(customer)
	return id, nil
}

func (c *controller) AddLoan(ctx *gin.Context) (int, error) {
	var loan entity.Loan
	err := ctx.ShouldBindJSON(&loan)
	if err != nil {
		return 0, err
	}
	id := c.services.AddLoan(loan)
	return id, nil
}
func (c *controller) RePayLoan(ctx *gin.Context) (int, error) {
	var repayloan entity.RePayLoan
	err := ctx.ShouldBindJSON(&repayloan)
	if err != nil {
		return -1, err
	}
	id := c.services.RepayLoan(repayloan)
	return id, nil
}
func (c *controller) GetAllLoans() []entity.Loan {
	return c.services.GetAllLoans()
}
func (c *controller) GetAllCustomers() []entity.Customer {
	return c.services.GetAllCustomers()
}
