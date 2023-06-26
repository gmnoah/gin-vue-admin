package cnSystem

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cnSystem"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    cnSystemReq "github.com/flipped-aurora/gin-vue-admin/server/model/cnSystem/request"
)

type ClientDeviceService struct {
}

// CreateClientDevice 创建ClientDevice记录
// Author [piexlmax](https://github.com/piexlmax)
func (clientDeviceService *ClientDeviceService) CreateClientDevice(clientDevice *cnSystem.ClientDevice) (err error) {
	err = global.GVA_DB.Create(clientDevice).Error
	return err
}

// DeleteClientDevice 删除ClientDevice记录
// Author [piexlmax](https://github.com/piexlmax)
func (clientDeviceService *ClientDeviceService)DeleteClientDevice(clientDevice cnSystem.ClientDevice) (err error) {
	err = global.GVA_DB.Delete(&clientDevice).Error
	return err
}

// DeleteClientDeviceByIds 批量删除ClientDevice记录
// Author [piexlmax](https://github.com/piexlmax)
func (clientDeviceService *ClientDeviceService)DeleteClientDeviceByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]cnSystem.ClientDevice{},"id in ?",ids.Ids).Error
	return err
}

// UpdateClientDevice 更新ClientDevice记录
// Author [piexlmax](https://github.com/piexlmax)
func (clientDeviceService *ClientDeviceService)UpdateClientDevice(clientDevice cnSystem.ClientDevice) (err error) {
	err = global.GVA_DB.Save(&clientDevice).Error
	return err
}

// GetClientDevice 根据id获取ClientDevice记录
// Author [piexlmax](https://github.com/piexlmax)
func (clientDeviceService *ClientDeviceService)GetClientDevice(id uint) (clientDevice cnSystem.ClientDevice, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&clientDevice).Error
	return
}

// GetClientDeviceInfoList 分页获取ClientDevice记录
// Author [piexlmax](https://github.com/piexlmax)
func (clientDeviceService *ClientDeviceService)GetClientDeviceInfoList(info cnSystemReq.ClientDeviceSearch) (list []cnSystem.ClientDevice, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&cnSystem.ClientDevice{})
    var clientDevices []cnSystem.ClientDevice
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&clientDevices).Error
	return  clientDevices, total, err
}
