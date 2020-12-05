package middlewares

import "github.com/gin-gonic/gin"

//BasicAuth ...
func BasicAuth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"joker": "P@ssw0rd",
	})
}
