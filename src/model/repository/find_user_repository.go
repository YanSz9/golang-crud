package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/YanSz9/golang-crud/src/configuration/logger"
	"github.com/YanSz9/golang-crud/src/configuration/rest_err"
	"github.com/YanSz9/golang-crud/src/model"
	"github.com/YanSz9/golang-crud/src/model/repository/entity"
	"github.com/YanSz9/golang-crud/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (ur *userRepository) FindUserByEmail(
	email string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByEmail repository",
		zap.String("journey", "findUserByEmail"))

	collection_name := os.Getenv(MONGODB_USER_DB)
	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "email", Value: email}}
	err := collection.FindOne(
		context.Background(),
		filter,
	).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf(
				"User not found with this email: %s", email)
			logger.Error(errorMessage,
				err,
				zap.String("journey", "findUserByEmail"))
			return nil, rest_err.NewNotFoundErorr(errorMessage)
		}
		errorMessage := "Error trying to find user by email"
		logger.Error(errorMessage,
			err,
			zap.String("journey", "findUserByEmail"))
		return nil, rest_err.NewInternalServerError(errorMessage)
	}
	logger.Info("findUserByEmail repository executed sucessfully",
		zap.String("journey", "findUserByEmail"),
		zap.String("email", email),
		zap.String("userId", userEntity.ID.Hex()))

	return converter.ConvertEntityToDomain(*userEntity), nil
}

func (ur *userRepository) FindUserByID(
	id string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByID repository",
		zap.String("journey", "findUserByID"))

	collection_name := os.Getenv(MONGODB_USER_DB)
	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "_id", Value: id}}
	err := collection.FindOne(
		context.Background(),
		filter,
	).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf(
				"User not found with this id: %s", id)
			logger.Error(errorMessage,
				err,
				zap.String("journey", "findUserByID"))
			return nil, rest_err.NewNotFoundErorr(errorMessage)
		}
		errorMessage := "Error trying to find user by id"
		logger.Error(errorMessage,
			err,
			zap.String("journey", "findUserById"))
		return nil, rest_err.NewInternalServerError(errorMessage)
	}
	logger.Info("findUserByEmail repository executed sucessfully",
		zap.String("journey", "findUserByEmail"),
		zap.String("userId", userEntity.ID.Hex()))

	return converter.ConvertEntityToDomain(*userEntity), nil
}
