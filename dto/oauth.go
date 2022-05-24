package dto

import (
	"github.com/gin-gonic/gin"
	"go_gateway/public"
)

/*tag:输出用json,输入用form,这里只用到form*/

type TokensInput struct {
	UserName string `json:"grant_type" form:"grant_type" comment:"授权类型" example:"client_credentials" validate:"required"`
	Scope    string `json:"scope" form:"scope" comment:"权限范围" example:"read_write" validate:"required"`
}

func (a *TokensInput) BindValidParam(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, a)
}

type TokensOutput struct {
	AccessToken string `json:"access_token" form:"access_token" `
	ExpiresIn   int    `json:"expires_in" form:"expires_in" `
	TokenType   string `json:"token_type" form:"token_type" `
	Scope       string `json:"scope" form:"scope" `
}
