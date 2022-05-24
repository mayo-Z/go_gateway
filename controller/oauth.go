package controller

import (
	"encoding/base64"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
	"go_gateway/dao"
	"go_gateway/dto"
	"go_gateway/middleware"
	"go_gateway/public"
	"strings"
	"time"
)

//jwt的token生成接口
//取出app_id secret
//生成app_list(类似service.LoadOnce)
//匹配app_id
//基于jwt生成token
//生成output

type OAuthController struct{}

func OAuthRegister(group *gin.RouterGroup) {
	oauth := &OAuthController{}
	group.POST("/tokens", oauth.Tokens)
}

// Tokens godoc
// @Summary 获取TOKEN
// @Description 获取TOKEN
// @Tags OAUTH
// @ID /oauth/tokens
// @Accept  json
// @Produce  json
// @Param body body dto.TokensInput true "body"
// @Success 200 {object} middleware.Response{data=dto.TokensOutput} "success"
// @Router /oauth/tokens [post]
func (oauth *OAuthController) Tokens(ctx *gin.Context) {
	params := &dto.TokensInput{}
	if err := params.BindValidParam(ctx); err != nil {
		middleware.ResponseError(ctx, 2000, err)
		return
	}
	splits := strings.Split(ctx.GetHeader("Authorization"), " ")
	if len(splits) != 2 {
		middleware.ResponseError(ctx, 2001, errors.New("用户名或密码格式错误"))
		return
	}
	appSecret, err := base64.StdEncoding.DecodeString(splits[1])
	if err != nil {
		middleware.ResponseError(ctx, 2002, err)
		return
	}
	//fmt.Println("appSecret ",string(appSecret))

	parts := strings.Split(string(appSecret), ":")
	if len(parts) != 2 {
		middleware.ResponseError(ctx, 2003, errors.New("用户名或密码格式错误"))
		return
	}

	appList := dao.AppManagerHandler.GetAppList()
	for _, appInfo := range appList {
		if appInfo.AppID == parts[0] && appInfo.Secret == parts[1] {
			claims := jwt.StandardClaims{
				ExpiresAt: time.Now().Add(public.JwtExpires * time.Second).In(lib.TimeLocation).Unix(),
				Issuer:    appInfo.AppID,
			}
			token, err := public.JwtEncode(claims)
			if err != nil {
				middleware.ResponseError(ctx, 2004, err)
				return
			}
			output := &dto.TokensOutput{
				AccessToken: token,
				ExpiresIn:   public.JwtExpires,
				TokenType:   "Bearer",
				Scope:       "read_write",
			}
			middleware.ResponseSuccess(ctx, output)
			return
		}
	}
	middleware.ResponseError(ctx, 2005, errors.New("未匹配正确的app信息"))
}
