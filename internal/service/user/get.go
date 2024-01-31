package user

import (
	"github.com/Rahugg/Golang-clean-arch-template/internal/model"
	"github.com/gin-gonic/gin"
)

func (s *service) Get(ctx *gin.Context, id string) (*model.User, error) {
	return nil, nil
}