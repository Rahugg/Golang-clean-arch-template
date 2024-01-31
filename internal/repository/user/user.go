package user

import (
	"github.com/Rahugg/Golang-clean-arch-template/internal/model"
	"github.com/gin-gonic/gin"
)

func (r *repository) Get(ctx *gin.Context, id string) (*model.User, error) {
	return nil, nil
}

func(r *repository)Create(ctx *gin.Context, payload *model.UserInput) error {
	return nil
}