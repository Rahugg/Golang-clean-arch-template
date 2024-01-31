package service

import (
	"github.com/Rahugg/Golang-clean-arch-template/internal/model"
	"github.com/gin-gonic/gin"
)

type UserService interface {
	Create(ctx *gin.Context, payload *model.UserInput) error
	Get(ctx *gin.Context, uuid string) (*model.User, error)
}
