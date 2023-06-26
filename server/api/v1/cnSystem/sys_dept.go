package cnSystem

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/cnSystem"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    cnSystemReq "github.com/flipped-aurora/gin-vue-admin/server/model/cnSystem/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type SysDeptApi struct {
}

var sysDeptService = service.ServiceGroupApp.CnSystemServiceGroup.SysDeptService


// CreateSysDept 创建SysDept
// @Tags SysDept
// @Summary 创建SysDept
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cnSystem.SysDept true "创建SysDept"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysDept/createSysDept [post]
func (sysDeptApi *SysDeptApi) CreateSysDept(c *gin.Context) {
	var sysDept cnSystem.SysDept
	err := c.ShouldBindJSON(&sysDept)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := sysDeptService.CreateSysDept(&sysDept); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSysDept 删除SysDept
// @Tags SysDept
// @Summary 删除SysDept
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cnSystem.SysDept true "删除SysDept"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysDept/deleteSysDept [delete]
func (sysDeptApi *SysDeptApi) DeleteSysDept(c *gin.Context) {
	var sysDept cnSystem.SysDept
	err := c.ShouldBindJSON(&sysDept)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := sysDeptService.DeleteSysDept(sysDept); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSysDeptByIds 批量删除SysDept
// @Tags SysDept
// @Summary 批量删除SysDept
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SysDept"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /sysDept/deleteSysDeptByIds [delete]
func (sysDeptApi *SysDeptApi) DeleteSysDeptByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := sysDeptService.DeleteSysDeptByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSysDept 更新SysDept
// @Tags SysDept
// @Summary 更新SysDept
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cnSystem.SysDept true "更新SysDept"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysDept/updateSysDept [put]
func (sysDeptApi *SysDeptApi) UpdateSysDept(c *gin.Context) {
	var sysDept cnSystem.SysDept
	err := c.ShouldBindJSON(&sysDept)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := sysDeptService.UpdateSysDept(sysDept); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSysDept 用id查询SysDept
// @Tags SysDept
// @Summary 用id查询SysDept
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cnSystem.SysDept true "用id查询SysDept"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysDept/findSysDept [get]
func (sysDeptApi *SysDeptApi) FindSysDept(c *gin.Context) {
	var sysDept cnSystem.SysDept
	err := c.ShouldBindQuery(&sysDept)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if resysDept, err := sysDeptService.GetSysDept(sysDept.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resysDept": resysDept}, c)
	}
}

// GetSysDeptList 分页获取SysDept列表
// @Tags SysDept
// @Summary 分页获取SysDept列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cnSystemReq.SysDeptSearch true "分页获取SysDept列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysDept/getSysDeptList [get]
func (sysDeptApi *SysDeptApi) GetSysDeptList(c *gin.Context) {
	var pageInfo cnSystemReq.SysDeptSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := sysDeptService.GetSysDeptInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
