// 自动生成模板DeviceDataTemplate
package cnSystem

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// DeviceDataTemplate 结构体
type DeviceDataTemplate struct {
      global.GVA_MODEL
      TemplateId  *int `json:"templateid" form:"templateid" gorm:"column:template_id;comment:名称;size:4;"`
      TemplateName  string `json:"templatename" form:"templatename" gorm:"column:template_name;comment:产品编码;size:64;"`
      DataName  string `json:"dataname" form:"dataname" gorm:"column:data_name;comment:产品系列 （集装箱式储能、移动应急储能、用户保障储能、梯次技术）;size:64;"`
      DataKey  string `json:"datakey" form:"datakey" gorm:"column:data_key;comment:;size:64;"`
      DataUnit  string `json:"dataunit" form:"dataunit" gorm:"column:data_unit;comment:;size:64;"`
}


// TableName DeviceDataTemplate 表名
func (DeviceDataTemplate) TableName() string {
  return "device_data_template"
}

