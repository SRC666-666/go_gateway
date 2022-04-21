package dto

import (
	"time"

	"github.com/e421083458/gateway_demo/public"
	"github.com/gin-gonic/gin"
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
	Password string `json:"password" form:"password" comment:"密码" example:"123456" validate:"required"` //管理员密码
}

//参数校验
func (param *ChangePwdInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}
