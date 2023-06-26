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

type ClientDeviceApi struct {
}

var clientDeviceService = service.ServiceGroupApp.CnSystemServiceGroup.ClientDeviceService


// CreateClientDevice 创建ClientDevice
// @Tags ClientDevice
// @Summary 创建ClientDevice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cnSystem.ClientDevice true "创建ClientDevice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /clientDevice/createClientDevice [post]
func (clientDeviceApi *ClientDeviceApi) CreateClientDevice(c *gin.Context) {
	var clientDevice cnSystem.ClientDevice
	err := c.ShouldBindJSON(&clientDevice)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := clientDeviceService.CreateClientDevice(&clientDevice); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteClientDevice 删除ClientDevice
// @Tags ClientDevice
// @Summary 删除ClientDevice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cnSystem.ClientDevice true "删除ClientDevice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /clientDevice/deleteClientDevice [delete]
func (clientDeviceApi *ClientDeviceApi) DeleteClientDevice(c *gin.Context) {
	var clientDevice cnSystem.ClientDevice
	err := c.ShouldBindJSON(&clientDevice)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := clientDeviceService.DeleteClientDevice(clientDevice); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteClientDeviceByIds 批量删除ClientDevice
// @Tags ClientDevice
// @Summary 批量删除ClientDevice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ClientDevice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /clientDevice/deleteClientDeviceByIds [delete]
func (clientDeviceApi *ClientDeviceApi) DeleteClientDeviceByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := clientDeviceService.DeleteClientDeviceByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateClientDevice 更新ClientDevice
// @Tags ClientDevice
// @Summary 更新ClientDevice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cnSystem.ClientDevice true "更新ClientDevice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /clientDevice/updateClientDevice [put]
func (clientDeviceApi *ClientDeviceApi) UpdateClientDevice(c *gin.Context) {
	var clientDevice cnSystem.ClientDevice
	err := c.ShouldBindJSON(&clientDevice)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := clientDeviceService.UpdateClientDevice(clientDevice); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindClientDevice 用id查询ClientDevice
// @Tags ClientDevice
// @Summary 用id查询ClientDevice
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cnSystem.ClientDevice true "用id查询ClientDevice"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /clientDevice/findClientDevice [get]
func (clientDeviceApi *ClientDeviceApi) FindClientDevice(c *gin.Context) {
	var clientDevice cnSystem.ClientDevice
	err := c.ShouldBindQuery(&clientDevice)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reclientDevice, err := clientDeviceService.GetClientDevice(clientDevice.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reclientDevice": reclientDevice}, c)
	}
}

// GetClientDeviceList 分页获取ClientDevice列表
// @Tags ClientDevice
// @Summary 分页获取ClientDevice列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cnSystemReq.ClientDeviceSearch true "分页获取ClientDevice列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /clientDevice/getClientDeviceList [get]
func (clientDeviceApi *ClientDeviceApi) GetClientDeviceList(c *gin.Context) {
	var pageInfo cnSystemReq.ClientDeviceSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := clientDeviceService.GetClientDeviceInfoList(pageInfo); err != nil {
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
