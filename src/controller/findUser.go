package controller

import (
	"net/http"
	"net/mail"

	"github.com/YanSz9/golang-crud/src/configuration/logger"
	"github.com/YanSz9/golang-crud/src/configuration/rest_err"
	"github.com/YanSz9/golang-crud/src/view"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) FindUserById(c *gin.Context) {
	logger.Info("Init findUserById controller",
		zap.String("journey", "findUserById"),
	)

	userId := c.Param("userId")

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error("Error trying to validate userId",
			err,
			zap.String("journey", "findUserById"),
		)
		errorMessage := rest_err.NewBadResquestError(
			"UserId is not a valid id",
		)
		c.JSON(errorMessage.Code, errorMessage)
		return
	}
	userDomain, err := uc.service.FindUserByIDServices(userId)
	if err != nil {
		logger.Error("Error trying to call findUserByIDServices",
			err,
			zap.String("journey", "findUserById"),
		)
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		userDomain,
	))
	logger.Info("FindUserById controller executed sucessfully",
		zap.String("journey", "findUserById"),
	)
}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	logger.Info("Init findUserByEmail controller",
		zap.String("journey", "findUserByEmail"),
	)

	userEmail := c.Param("userEmail")

	if _, err := mail.ParseAddress(userEmail); err != nil {
		logger.Error("Error trying to validate userEmail",
			err,
			zap.String("journey", "findUserByEmail"),
		)
		errorMessage := rest_err.NewBadResquestError(
			"UserEmail is not valid email",
		)
		c.JSON(errorMessage.Code, errorMessage)
		return
	}
	userDomain, err := uc.service.FindUserByEmailServices(userEmail)
	if err != nil {
		logger.Error("Error trying to call findUserByEmailServices",
			err,
			zap.String("journey", "findUserByEmail"),
		)
		c.JSON(err.Code, err)
		return
	}
	logger.Info("FindUserByEmail controller executed sucessfully",
		zap.String("journey", "findUserByEmail"),
	)
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		userDomain,
	))
}
