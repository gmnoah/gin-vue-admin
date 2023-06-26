package cnSystem

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ClientDeviceRouter struct {
}

// InitClientDeviceRouter 初始化 ClientDevice 路由信息
func (s *ClientDeviceRouter) InitClientDeviceRouter(Router *gin.RouterGroup) {
	clientDeviceRouter := Router.Group("clientDevice").Use(middleware.OperationRecord())
	clientDeviceRouterWithoutRecord := Router.Group("clientDevice")
	var clientDeviceApi = v1.ApiGroupApp.CnSystemApiGroup.ClientDeviceApi
	{
		clientDeviceRouter.POST("createClientDevice", clientDeviceApi.CreateClientDevice)   // 新建ClientDevice
		clientDeviceRouter.DELETE("deleteClientDevice", clientDeviceApi.DeleteClientDevice) // 删除ClientDevice
		clientDeviceRouter.DELETE("deleteClientDeviceByIds", clientDeviceApi.DeleteClientDeviceByIds) // 批量删除ClientDevice
		clientDeviceRouter.PUT("updateClientDevice", clientDeviceApi.UpdateClientDevice)    // 更新ClientDevice
	}
	{
		clientDeviceRouterWithoutRecord.GET("findClientDevice", clientDeviceApi.FindClientDevice)        // 根据ID获取ClientDevice
		clientDeviceRouterWithoutRecord.GET("getClientDeviceList", clientDeviceApi.GetClientDeviceList)  // 获取ClientDevice列表
	}
}
