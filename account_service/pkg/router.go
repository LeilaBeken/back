package pkg

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Routes() {
	r := gin.Default()

	r.GET("/accounts", GetAccounts)
	r.GET("/accounts/:id", GetAccount)
	r.POST("/accounts", CreateAccount)
	r.PUT("/accounts/:id", AuthMiddleware(), UpdateAccount)
	r.DELETE("/accounts/:id", DeleteAccount)
	r.POST("/login", Login)

	fmt.Println("Account service is running on :8080")
	r.Run(":8080")
}
