package controller

import (
	"github.com/e421083458/gateway_demo/dao"
	"github.com/e421083458/gateway_demo/dto"
	"github.com/e421083458/gateway_demo/middleware"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
)

type ServiceController struct {
}

func ServiceRegist(group *gin.RouterGroup) {
	service := &ServiceController{}
	group.GET("/service_list", service.ServiceList)
}

// ListPage godoc
// @Summary 服务列表
// @Description 服务列表
// @Tags 服务管理
// @ID /service/service_list
// @Accept  json
// @Produce  json
// @Param info query string false  "关键词"
// @Param page_size query string true  "每页个数"
// @Param page_no query string true  "当前页数"
// @Success 200 {object} middleware.Response{data=dto.ServiceListOutput} "success"
// @Router /service/service_list[get]
func (service *ServiceController) ServiceList(c *gin.Context) {

	params := &dto.ServiceListInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}

	serviceInfo := &dao.ServiceInfo{}
	list, total, err := serviceInfo.PageList(c, tx, params)
	if err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	outList := []dto.ServiceListItemOutput{}
	for _, listItem := range list {

		outItem := dto.ServiceListItemOutput{
			ID:          int64(listItem.ID),
			ServicName:  listItem.ServiceName,
			ServiceDesc: listItem.ServiceDesc,
		}
		outList = append(outList, outItem)
	}

	out := &dto.ServiceListOutput{
		Total: total,
		List:  outList,
	}

	//执行成功的输出
	middleware.ResponseSuccess(c, out)
}
