package entity

// Customer ..
type Customer struct {
	// {"custom_name": ""}
	CustomerID     int    `json:"customer_id"`
	Name           string `json:"full_name" binding:"min=3,max=100"`
	Address        string `json:"address" binding:"required"`
	Number         string `json:"number" binding:"required,number"`
	PassportNumber string `json:"passport_number" binding:"required"`
}
