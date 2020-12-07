package main

import (
	"Loan/controller"
	"Loan/middlewares"
	"Loan/services"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	loanService    services.LoanServices     = services.New()
	loanController controller.LoanController = controller.New(loanService)
)

func setupOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	server := gin.New()
	setupOutput()
	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), gindump.Dump())

	server.GET("/loans", func(c *gin.Context) {
		c.JSON(200, loanController.GetAllLoans())
	})

	server.POST("/loans", func(c *gin.Context) {
		id, err := loanController.AddLoan(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(200, gin.H{
				"id": id,
			})
		}
	})

	server.GET("/customers", func(c *gin.Context) {
		c.JSON(200, loanController.GetAllCustomers())
	})
	server.POST("/customers", func(c *gin.Context) {
		fmt.Println("post qator")
		id, err := loanController.AddCustomer(c)
		if err != nil {
			fmt.Println("errr qator")
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			fmt.Println("succes qator")
			c.JSON(200, gin.H{
				"id": id,
			})
		}

	})

	server.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
