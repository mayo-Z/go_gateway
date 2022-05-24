package dto

import (
	"github.com/gin-gonic/gin"
	"go_gateway/public"
	"time"
)

type AdminSessionInfo struct {
	ID        int       `json:"id" `
	UserName  string    `json:"username" `
	LoginTime time.Time `json:"login_time" `
}

/*tag:输出用json,输入用form,这里只用到form*/

type AdminLoginInput struct {
	UserName string `json:"username" form:"username" comment:"管理员用户名" example:"admin" validate:"required,valid_username"`
	Password string `json:"password" form:"password" comment:"密码" example:"123456" validate:"required"`
}

func (a *AdminLoginInput) BindValidParam(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, a)
}

type AdminLoginOutput struct {
	Token string `json:"token" form:"token" comment:"token" example:"token" validate:""`
}
