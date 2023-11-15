package controller

import (
	"github.com/YanSz9/golang-crud/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	err := rest_err.NewBadResquestError("Error on call route")
	c.JSON(err.Code, err)
}
