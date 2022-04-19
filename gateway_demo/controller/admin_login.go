package controller

import (
	"github.com/e421083458/gin_scaffold/dto"
	"github.com/gin-gonic/gin"
)

type AdminLoginController struct {
}

func AdminLoginRegist(group *gin.RouterGroup) {
	// adminLogin:=&AdminLoginController{}
	// group.POST()
}

func (adminlogin *AdminLoginController) AdminLogin(c *gin.Context) {
	params := &dto.AdminLoginInput{}

}
