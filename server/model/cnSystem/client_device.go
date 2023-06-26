// 自动生成模板ClientDevice
package cnSystem

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// ClientDevice 结构体
type ClientDevice struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:名称;size:64;"`
      Code  string `json:"code" form:"code" gorm:"column:code;comment:产品编码;size:64;"`
      Series  string `json:"series" form:"series" gorm:"column:series;comment:产品系列 （集装箱式储能、移动应急储能、用户保障储能、梯次技术）;size:64;"`
      Phone  string `json:"phone" form:"phone" gorm:"column:phone;comment:;size:64;"`
      GatewayId  string `json:"gatewayId" form:"gatewayId" gorm:"column:gateway_id;comment:;size:64;"`
      DeptId  string `json:"deptId" form:"deptId" gorm:"column:dept_id;comment:;size:64;"`
      Lon  *float64 `json:"lon" form:"lon" gorm:"column:lon;comment:;size:22;"`
      Lat  *float64 `json:"lat" form:"lat" gorm:"column:lat;comment:;size:22;"`
}


// TableName ClientDevice 表名
func (ClientDevice) TableName() string {
  return "client_device"
}

