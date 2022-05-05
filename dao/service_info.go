package dao

import (
	"time"

	"github.com/e421083458/gateway_demo/dto"
	"github.com/e421083458/gateway_demo/public"
	"github.com/e421083458/gorm"
	"github.com/gin-gonic/gin"
)

type ServiceInfo struct {
	ID          int       `json:"id" gorm:"primary_key" description:"自增主键"`
	LoadType    int       `json:"load_type" gorm:"column:load_type" description:"负载类型 0=http 1=tcp 2=grpc"`
	ServiceName string    `json:"service_name" gorm:"column:service_name" description:"服务名称"`
	ServiceDesc string    `json:"service_desc" gorm:"column:service_desc" description:"服务描述"`
	UpdatedAt   time.Time `json:"update_at" gorm:"column:update_at" description:"更新时间"`
	CreatedAt   time.Time `json:"create_at" gorm:"column:create_at" description:"创建时间"`
	Isdelete    int       `json:"is_delete" gorm:"column:is_delete" description:"是否已删除;0:否;1:是"`
}

func (t *ServiceInfo) TableName() string {
	return "gateway_service_info"
}

func (t *ServiceInfo) Find(c *gin.Context, tx *gorm.DB, search *ServiceInfo) (*ServiceInfo, error) {
	out := &ServiceInfo{}
	// err := tx.WithContext(c).Where("id = ?", id).Find(area).Error
	err := tx.SetCtx(public.GetGinTraceContext(c)).Where(search).Find(out).Error
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (t *ServiceInfo) Save(c *gin.Context, tx *gorm.DB) error {

	return tx.SetCtx(public.GetGinTraceContext(c)).Save(t).Error
}

func (t *ServiceInfo) PageList(c *gin.Context, tx *gorm.DB, param *dto.ServiceListInput) ([]ServiceInfo, int64, error) {
	total := 0
	list := []ServiceInfo{}
	offset := (param.PageNo - 1) * param.PageSize

	query := tx.SetCtx(public.GetGinTraceContext(c))
	query = query.Table(t.TableName()).Where("is_delete=0")
	if param.Info != "" {
		query = query.Where("(service_name like ? or service_desc like ?)", "%"+param.Info+"%", "%"+param.Info+"%")
	}
	if err := query.Limit(param.PageSize).Offset(offset).Order("id desc").Find(&list).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}

	query.Limit(param.PageSize).Offset(offset).Count(&total)

	return list, int64(total), nil

}

func (t *ServiceInfo) ServiceDetail(c *gin.Context, tx *gorm.DB, search *ServiceInfo) (*ServiceDeatil, error) {
	httpRule := &HttpRule{ServiceID: int64(search.ID)}
	httpRule, err := httpRule.Find(c, tx, httpRule)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	tcpRule := &TcpRule{ServiceID: int64(search.ID)}
	tcpRule, err = tcpRule.Find(c, tx, tcpRule)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	grpcRule := &GrpcRule{ServiceID: int64(search.ID)}
	grpcRule, err = grpcRule.Find(c, tx, grpcRule)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	accessControl := &AccessControl{ServiceID: int64(search.ID)}
	accessControl, err = accessControl.Find(c, tx, accessControl)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	loadBanlance := &LoadBalance{ServiceID: int64(search.ID)}
	loadBanlance, err = loadBanlance.Find(c, tx, loadBanlance)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	detail := &ServiceDeatil{
		Info:          search,
		HTTPRule:      httpRule,
		TCPRule:       tcpRule,
		GRPCRule:      grpcRule,
		LoadBalance:   loadBanlance,
		AccessControl: accessControl,
	}
	return detail, nil
}
