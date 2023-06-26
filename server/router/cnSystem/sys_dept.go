package cnSystem

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SysDeptRouter struct {
}

// InitSysDeptRouter 初始化 SysDept 路由信息
func (s *SysDeptRouter) InitSysDeptRouter(Router *gin.RouterGroup) {
	sysDeptRouter := Router.Group("sysDept").Use(middleware.OperationRecord())
	sysDeptRouterWithoutRecord := Router.Group("sysDept")
	var sysDeptApi = v1.ApiGroupApp.CnSystemApiGroup.SysDeptApi
	{
		sysDeptRouter.POST("createSysDept", sysDeptApi.CreateSysDept)   // 新建SysDept
		sysDeptRouter.DELETE("deleteSysDept", sysDeptApi.DeleteSysDept) // 删除SysDept
		sysDeptRouter.DELETE("deleteSysDeptByIds", sysDeptApi.DeleteSysDeptByIds) // 批量删除SysDept
		sysDeptRouter.PUT("updateSysDept", sysDeptApi.UpdateSysDept)    // 更新SysDept
	}
	{
		sysDeptRouterWithoutRecord.GET("findSysDept", sysDeptApi.FindSysDept)        // 根据ID获取SysDept
		sysDeptRouterWithoutRecord.GET("getSysDeptList", sysDeptApi.GetSysDeptList)  // 获取SysDept列表
	}
}
