package controller

import (
	"encoding/json"
	"fmt"

	"github.com/e421083458/gateway_demo/dao"
	"github.com/e421083458/gateway_demo/dto"
	"github.com/e421083458/gateway_demo/middleware"
	"github.com/e421083458/gateway_demo/public"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

type AdminController struct {
}

func AdminRegist(group *gin.RouterGroup) {
	adminLogin := &AdminController{}
	group.GET("/admin_info", adminLogin.AdminInfo)
	group.POST("/change_pwd", adminLogin.AdminInfo)
}

// AdminInfo godoc
// @Summary 管理员信息
// @Description 管理员信息
// @Tags 管理员接口
// @ID /admin/admin_info
// @Accept  json
// @Produce  json
// @Param body body dto.AdminInfoOutput true "body"
// @Success 200 {object} middleware.Response{data=dto.AdminInfoOutput} "success"
// @Router /admin/admin_info [get]
func (adminlogin *AdminController) AdminInfo(c *gin.Context) {
	//拿到信息
	sess := sessions.Default(c)
	sessInfo := sess.Get(public.AdminSessionInfoKey)
	// sessInfoStr:=sessInfo.(string)	//转成字符串,因为json.Unmarshal([]byte(sessInfo), adminSessionInfo)中的[]byte(sessInfo)sessInfo要是string
	adminSessionInfo := &dto.AdminSessionInfo{}
	if err := json.Unmarshal([]byte(fmt.Sprint(sessInfo)), adminSessionInfo); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	//组装信息
	//1.读取sessionKey对应的json转化为结构体
	//2.取出数据，然后封装输出结构体
	out := &dto.AdminInfoOutput{
		ID:           adminSessionInfo.ID,
		Name:         adminSessionInfo.UserName,
		LoginTime:    adminSessionInfo.LoginTime,
		Avatar:       "",
		Introduction: "",
		Roles:        []string{},
	}
	//执行成功的输出
	middleware.ResponseSuccess(c, out)
}

// ChangPwd godoc
// @Summary 修改密码
// @Description 修改密码
// @Tags 管理员接口
// @ID /admin/change_pwd
// @Accept  json
// @Produce  json
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /admin/admin_info [post]
func (adminlogin *AdminController) ChangePwd(c *gin.Context) {

	params := &dto.ChangePwdInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	//1.session读取用户信息到结构体sessInfo
	//2.sessInfo.ID 读取数据库信息 adminInfo
	//3.params.password+adminInfo.salt sha256 saltPassword
	//4.saltPassword==> adminInfo save执行数据库保存

	//拿到信息
	sess := sessions.Default(c)
	sessInfo := sess.Get(public.AdminSessionInfoKey)
	// sessInfoStr:=sessInfo.(string)	//转成字符串,因为json.Unmarshal([]byte(sessInfo), adminSessionInfo)中的[]byte(sessInfo)sessInfo要是string
	adminSessionInfo := &dto.AdminSessionInfo{}
	if err := json.Unmarshal([]byte(fmt.Sprint(sessInfo)), adminSessionInfo); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	//2.sessInfo.ID 读取数据库信息 adminInfo
	//从数据库中读取 adminInfo
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	adminInfo := &dao.Admin{}
	adminInfo, err = adminInfo.Find(c, tx, (&dao.Admin{
		UserName: adminInfo.UserName,
	}))
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}

	//3.生成加盐密码params.password+adminInfo.salt sha256 saltPassword
	saltPassword := public.GenSaltPassword(adminInfo.Salt, params.Password)
	adminInfo.Password = saltPassword
	//数据保存
	if err := adminInfo.Save(c, tx); err != nil {
		middleware.ResponseError(c, 2003, err)
		return
	}

	//执行成功的输出
	middleware.ResponseSuccess(c, "")
}
