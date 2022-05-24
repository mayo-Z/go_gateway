package dto

//DTO (Data Transfer Object)数据传输对象

import (
	"github.com/gin-gonic/gin"
	"go_gateway/public"
	"time"
)

type AdminInfoOutput struct {
	ID           int       `json:"id" `
	Name         string    `json:"name" `
	LoginTime    time.Time `json:"login_time" `
	Avatar       string    `json:"avatar" `
	Introduction string    `json:"introduction" `
	Roles        []string  `json:"roles" `
}

type ChangePwdInput struct {
	Password string `json:"password" form:"password" comment:"密码" example:"123456" validate:"required"`
}

func (a *ChangePwdInput) BindValidParam(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, a)
}
