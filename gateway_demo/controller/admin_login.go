package controller

import (
	"github.com/e421083458/gateway_demo/dto"
	"github.com/e421083458/gateway_demo/middleware"
	"github.com/gin-gonic/gin"
)

type AdminLoginController struct {
}

func AdminLoginRegist(group *gin.RouterGroup) {
	adminLogin := &AdminLoginController{}
	group.POST("/login", adminLogin.AdminLogin)
}

// ListPage godoc
// @Summary 管理员登录
// @Description 管理员登录
// @Tags 管理员接口
// @ID /admin_login/login
// @Accept  json
// @Produce  json
// @Param body body dto.AdminLoginInput true "body"
// @Success 200 {object} middleware.Response{data=dto.AdminLoginOutput} "success"
// @Router /admin_login/login [post]
func (adminlogin *AdminLoginController) AdminLogin(c *gin.Context) {
	params := &dto.AdminLoginInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 1001, err)
		return
	}
	//1.通过params.UserName取得管理员信息 admininfo
	//2.admininfo.salt+params.Password sha256=>saltPassword
	//3.saltPassword==admininfo.password
	out := &dto.AdminLoginOutput{Token: params.UserName}
	//执行成功的输出
	middleware.ResponseSuccess(c, out)
}
