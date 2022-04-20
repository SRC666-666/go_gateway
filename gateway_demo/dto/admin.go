package dto

import (
	"errors"
	"time"

	"github.com/e421083458/gateway_demo/dto"
	"github.com/e421083458/gateway_demo/public"
	"github.com/e421083458/gorm"
	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"
)

type Admin struct {
	Id        int       `json:"id" gorm:"primary_key" description:"自增主键"`
	UserName  string    `json:"user_name" gorm:"column:user_name" description:"管理员用户名"`
	Salt      string    `json:"salt" gorm:"column:salt" description:"盐"`
	Password  string    `json:"password" gorm:"column:password" description:"密码"`
	UpdatedAt time.Time `json:"update_at" gorm:"column:update_at" description:"更新时间"`
	CreatedAt time.Time `json:"create_at" gorm:"column:create_at" description:"创建时间"`
	Isdelete  int       `json:"is_delete" gorm:"column:is_delete" description:"删除时间"`
}

func (t *Admin) TableName() string {
	return "area"
}

// func (t *Admin) Find(c *gin.Context, tx *gorm.DB, id string) (*Admin, error) {
// 	area := &Admin{}
// 	err := tx.WithContext(c).Where("id = ?", id).Find(area).Error
// 	// err := tx.SetCtx(public.GetGinTraceContext(c)).Where("id = ?", id).Find(area).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return area, nil
// }

func (t *Admin) Find(c *gin.Context, tx *gorm.DB, search *Admin) (*Admin, error) {
	area := &Admin{}
	// err := tx.WithContext(c).Where("id = ?", id).Find(area).Error
	err := tx.SetCtx(public.GetGinTraceContext(c)).Where(search).Find(area).Error
	if err != nil {
		return nil, err
	}
	return area, nil
}

func (t *Admin) LoginCheck(c *gin.Context, tx *gorm.DB, param *dto.AdminLoginInput) (*Admin, error) {
	adminInfo, err := t.Find(c, tx, (&Admin{UserName: param.UserName, Isdelete: 0}))
	if err != nil {
		return nil, errors.New("用户信息不存在")
	}
	// param.Password
	// adminInfo.Salt
	saltPassword := public.GenSaltPassword(adminInfo.Salt, param.Password)
	if adminInfo.Password != saltPassword {
		return nil, errors.New("密码错误，请重新输入")
	}
	return adminInfo, nil
}
