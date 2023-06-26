package cnSystem

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cnSystem"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    cnSystemReq "github.com/flipped-aurora/gin-vue-admin/server/model/cnSystem/request"
)

type DeviceDataTemplateService struct {
}

// CreateDeviceDataTemplate 创建DeviceDataTemplate记录
// Author [piexlmax](https://github.com/piexlmax)
func (devicedatatemplateService *DeviceDataTemplateService) CreateDeviceDataTemplate(devicedatatemplate *cnSystem.DeviceDataTemplate) (err error) {
	err = global.GVA_DB.Create(devicedatatemplate).Error
	return err
}

// DeleteDeviceDataTemplate 删除DeviceDataTemplate记录
// Author [piexlmax](https://github.com/piexlmax)
func (devicedatatemplateService *DeviceDataTemplateService)DeleteDeviceDataTemplate(devicedatatemplate cnSystem.DeviceDataTemplate) (err error) {
	err = global.GVA_DB.Delete(&devicedatatemplate).Error
	return err
}

// DeleteDeviceDataTemplateByIds 批量删除DeviceDataTemplate记录
// Author [piexlmax](https://github.com/piexlmax)
func (devicedatatemplateService *DeviceDataTemplateService)DeleteDeviceDataTemplateByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]cnSystem.DeviceDataTemplate{},"id in ?",ids.Ids).Error
	return err
}

// UpdateDeviceDataTemplate 更新DeviceDataTemplate记录
// Author [piexlmax](https://github.com/piexlmax)
func (devicedatatemplateService *DeviceDataTemplateService)UpdateDeviceDataTemplate(devicedatatemplate cnSystem.DeviceDataTemplate) (err error) {
	err = global.GVA_DB.Save(&devicedatatemplate).Error
	return err
}

// GetDeviceDataTemplate 根据id获取DeviceDataTemplate记录
// Author [piexlmax](https://github.com/piexlmax)
func (devicedatatemplateService *DeviceDataTemplateService)GetDeviceDataTemplate(id uint) (devicedatatemplate cnSystem.DeviceDataTemplate, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&devicedatatemplate).Error
	return
}

// GetDeviceDataTemplateInfoList 分页获取DeviceDataTemplate记录
// Author [piexlmax](https://github.com/piexlmax)
func (devicedatatemplateService *DeviceDataTemplateService)GetDeviceDataTemplateInfoList(info cnSystemReq.DeviceDataTemplateSearch) (list []cnSystem.DeviceDataTemplate, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&cnSystem.DeviceDataTemplate{})
    var devicedatatemplates []cnSystem.DeviceDataTemplate
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&devicedatatemplates).Error
	return  devicedatatemplates, total, err
}
