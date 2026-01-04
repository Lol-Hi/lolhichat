package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"errors"
	"gorm.io/gorm"
	"backend/internal/dataaccess"
	"backend/internal/helpers"
)

type SignUp struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
// TODO: unify with the login struct as a general Auth struct

const HASH_COST = 14

func HandleSignUp (c *gin.Context) {
	var payload SignUp
	
	reqErr := c.ShouldBindJSON(&payload)
	if reqErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": reqErr.Error()})
		return
	}
	
	_, dbErr1 := dataaccess.GetUser(payload.Username)
	if dbErr1 == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
		return
	}
	if !errors.Is(dbErr1, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": dbErr1.Error()})
	}	

	passHash, hashErr := helpers.HashPassword(payload.Password, HASH_COST)
	if hashErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": hashErr.Error()})
		return
	}
	dbErr2 := dataaccess.CreateUser(payload.Username, passHash)
	if dbErr2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": dbErr2.Error()})
		return
	}

	
	c.JSON(http.StatusOK, nil)
}
