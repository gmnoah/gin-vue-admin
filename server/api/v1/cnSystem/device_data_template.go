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
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type DeviceDataTemplateApi struct {
}

var devicedatatemplateService = service.ServiceGroupApp.CnSystemServiceGroup.DeviceDataTemplateService


// CreateDeviceDataTemplate 创建DeviceDataTemplate
// @Tags DeviceDataTemplate
// @Summary 创建DeviceDataTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cnSystem.DeviceDataTemplate true "创建DeviceDataTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /devicedatatemplate/createDeviceDataTemplate [post]
func (devicedatatemplateApi *DeviceDataTemplateApi) CreateDeviceDataTemplate(c *gin.Context) {
	var devicedatatemplate cnSystem.DeviceDataTemplate
	err := c.ShouldBindJSON(&devicedatatemplate)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "TemplateId":{utils.NotEmpty()},
        "TemplateName":{utils.NotEmpty()},
        "DataName":{utils.NotEmpty()},
        "DataKey":{utils.NotEmpty()},
    }
	if err := utils.Verify(devicedatatemplate, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := devicedatatemplateService.CreateDeviceDataTemplate(&devicedatatemplate); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteDeviceDataTemplate 删除DeviceDataTemplate
// @Tags DeviceDataTemplate
// @Summary 删除DeviceDataTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cnSystem.DeviceDataTemplate true "删除DeviceDataTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /devicedatatemplate/deleteDeviceDataTemplate [delete]
func (devicedatatemplateApi *DeviceDataTemplateApi) DeleteDeviceDataTemplate(c *gin.Context) {
	var devicedatatemplate cnSystem.DeviceDataTemplate
	err := c.ShouldBindJSON(&devicedatatemplate)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := devicedatatemplateService.DeleteDeviceDataTemplate(devicedatatemplate); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteDeviceDataTemplateByIds 批量删除DeviceDataTemplate
// @Tags DeviceDataTemplate
// @Summary 批量删除DeviceDataTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DeviceDataTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /devicedatatemplate/deleteDeviceDataTemplateByIds [delete]
func (devicedatatemplateApi *DeviceDataTemplateApi) DeleteDeviceDataTemplateByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := devicedatatemplateService.DeleteDeviceDataTemplateByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateDeviceDataTemplate 更新DeviceDataTemplate
// @Tags DeviceDataTemplate
// @Summary 更新DeviceDataTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body cnSystem.DeviceDataTemplate true "更新DeviceDataTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /devicedatatemplate/updateDeviceDataTemplate [put]
func (devicedatatemplateApi *DeviceDataTemplateApi) UpdateDeviceDataTemplate(c *gin.Context) {
	var devicedatatemplate cnSystem.DeviceDataTemplate
	err := c.ShouldBindJSON(&devicedatatemplate)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "TemplateId":{utils.NotEmpty()},
          "TemplateName":{utils.NotEmpty()},
          "DataName":{utils.NotEmpty()},
          "DataKey":{utils.NotEmpty()},
      }
    if err := utils.Verify(devicedatatemplate, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := devicedatatemplateService.UpdateDeviceDataTemplate(devicedatatemplate); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindDeviceDataTemplate 用id查询DeviceDataTemplate
// @Tags DeviceDataTemplate
// @Summary 用id查询DeviceDataTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cnSystem.DeviceDataTemplate true "用id查询DeviceDataTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /devicedatatemplate/findDeviceDataTemplate [get]
func (devicedatatemplateApi *DeviceDataTemplateApi) FindDeviceDataTemplate(c *gin.Context) {
	var devicedatatemplate cnSystem.DeviceDataTemplate
	err := c.ShouldBindQuery(&devicedatatemplate)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if redevicedatatemplate, err := devicedatatemplateService.GetDeviceDataTemplate(devicedatatemplate.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redevicedatatemplate": redevicedatatemplate}, c)
	}
}

// GetDeviceDataTemplateList 分页获取DeviceDataTemplate列表
// @Tags DeviceDataTemplate
// @Summary 分页获取DeviceDataTemplate列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query cnSystemReq.DeviceDataTemplateSearch true "分页获取DeviceDataTemplate列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /devicedatatemplate/getDeviceDataTemplateList [get]
func (devicedatatemplateApi *DeviceDataTemplateApi) GetDeviceDataTemplateList(c *gin.Context) {
	var pageInfo cnSystemReq.DeviceDataTemplateSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := devicedatatemplateService.GetDeviceDataTemplateInfoList(pageInfo); err != nil {
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
