package cnSystem

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cnSystem"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    cnSystemReq "github.com/flipped-aurora/gin-vue-admin/server/model/cnSystem/request"
)

type SysDeptService struct {
}

// CreateSysDept 创建SysDept记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysDeptService *SysDeptService) CreateSysDept(sysDept *cnSystem.SysDept) (err error) {
	err = global.GVA_DB.Create(sysDept).Error
	return err
}

// DeleteSysDept 删除SysDept记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysDeptService *SysDeptService)DeleteSysDept(sysDept cnSystem.SysDept) (err error) {
	err = global.GVA_DB.Delete(&sysDept).Error
	return err
}

// DeleteSysDeptByIds 批量删除SysDept记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysDeptService *SysDeptService)DeleteSysDeptByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]cnSystem.SysDept{},"id in ?",ids.Ids).Error
	return err
}

// UpdateSysDept 更新SysDept记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysDeptService *SysDeptService)UpdateSysDept(sysDept cnSystem.SysDept) (err error) {
	err = global.GVA_DB.Save(&sysDept).Error
	return err
}

// GetSysDept 根据id获取SysDept记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysDeptService *SysDeptService)GetSysDept(id uint) (sysDept cnSystem.SysDept, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&sysDept).Error
	return
}

// GetSysDeptInfoList 分页获取SysDept记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysDeptService *SysDeptService)GetSysDeptInfoList(info cnSystemReq.SysDeptSearch) (list []cnSystem.SysDept, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&cnSystem.SysDept{})
    var sysDepts []cnSystem.SysDept
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&sysDepts).Error
	return  sysDepts, total, err
}
