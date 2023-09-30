package pkg

import (
	md "account_service/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetAccounts(c *gin.Context) {
	db, err := GetDB()
	if err != nil {
		panic(err)
	}
	var accounts []md.Account
	db.Find(&accounts)
	c.JSON(http.StatusOK, accounts)
}

func GetAccount(c *gin.Context) {
	db, err := GetDB()
	if err != nil {
		panic(err)
	}
	id := c.Param("id")
	var account md.Account
	if err := db.First(&account, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		return
	}
	c.JSON(http.StatusOK, account)
}

func CreateAccount(c *gin.Context) {
	db, err := GetDB()
	if err != nil {
		panic(err)
	}
	var account md.Account
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	account.Password, err = hashPassword(account.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	db.Create(&account)
	c.JSON(http.StatusOK, account)
}

func UpdateAccount(c *gin.Context) {
	db, err := GetDB()
	if err != nil {
		panic(err)
	}
	id := c.Param("id")
	var account md.Account
	if err := db.First(&account, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		return
	}

	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Save(&account)
	c.JSON(http.StatusOK, account)
}

func DeleteAccount(c *gin.Context) {
	db, err := GetDB()
	if err != nil {
		panic(err)
	}
	id := c.Param("id")
	var account md.Account
	if err := db.First(&account, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		return
	}
	db.Delete(&account)
	c.JSON(http.StatusNoContent, gin.H{})
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
