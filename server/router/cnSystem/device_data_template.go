package cnSystem

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DeviceDataTemplateRouter struct {
}

// InitDeviceDataTemplateRouter 初始化 DeviceDataTemplate 路由信息
func (s *DeviceDataTemplateRouter) InitDeviceDataTemplateRouter(Router *gin.RouterGroup) {
	devicedatatemplateRouter := Router.Group("devicedatatemplate").Use(middleware.OperationRecord())
	devicedatatemplateRouterWithoutRecord := Router.Group("devicedatatemplate")
	var devicedatatemplateApi = v1.ApiGroupApp.CnSystemApiGroup.DeviceDataTemplateApi
	{
		devicedatatemplateRouter.POST("createDeviceDataTemplate", devicedatatemplateApi.CreateDeviceDataTemplate)   // 新建DeviceDataTemplate
		devicedatatemplateRouter.DELETE("deleteDeviceDataTemplate", devicedatatemplateApi.DeleteDeviceDataTemplate) // 删除DeviceDataTemplate
		devicedatatemplateRouter.DELETE("deleteDeviceDataTemplateByIds", devicedatatemplateApi.DeleteDeviceDataTemplateByIds) // 批量删除DeviceDataTemplate
		devicedatatemplateRouter.PUT("updateDeviceDataTemplate", devicedatatemplateApi.UpdateDeviceDataTemplate)    // 更新DeviceDataTemplate
	}
	{
		devicedatatemplateRouterWithoutRecord.GET("findDeviceDataTemplate", devicedatatemplateApi.FindDeviceDataTemplate)        // 根据ID获取DeviceDataTemplate
		devicedatatemplateRouterWithoutRecord.GET("getDeviceDataTemplateList", devicedatatemplateApi.GetDeviceDataTemplateList)  // 获取DeviceDataTemplate列表
	}
}
