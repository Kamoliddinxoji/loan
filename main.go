package main

import (
	"Loan/controller"
	"Loan/middlewares"
	"Loan/services"
	"database/sql"
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
			c.JSON(200, id)
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
			c.JSON(200, id)
		}

	})
	server.GET("/testCon", func(c *gin.Context) {
		db, err := sql.Open("mysql", "admin:xoji@tcp(127.0.0.1:5432)/loan")
		if err != nil {
			panic(err.Error())

		}

		add := ""
		db.QueryRow("select address from customer where customer_id = 1;").Scan(&add)
		defer db.Close()
		c.JSON(200, add)
	})

	server.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
