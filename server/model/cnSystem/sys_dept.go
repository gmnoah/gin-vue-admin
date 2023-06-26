// 自动生成模板SysDept
package cnSystem

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
	
)

// SysDept 结构体
type SysDept struct {
      global.GVA_MODEL
      DeptId  *int `json:"deptId" form:"deptId" gorm:"column:dept_id;comment:;size:10;"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:部门名称;size:50;"`
      Sort  *int `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`
      CreateTime  *time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;"`
      UpdateTime  *time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:修改时间;"`
      DelFlag  string `json:"delFlag" form:"delFlag" gorm:"column:del_flag;comment:是否删除  -1：已删除  0：正常;"`
      ParentId  *int `json:"parentId" form:"parentId" gorm:"column:parent_id;comment:;size:10;"`
      TenantId  *int `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:;size:10;"`
      IsTenant  string `json:"isTenant" form:"isTenant" gorm:"column:is_tenant;comment:0: 部门 1：租户 （公司）;"`
      LeaderId  string `json:"leaderId" form:"leaderId" gorm:"column:leader_id;comment:;size:64;"`
      Lon  *float64 `json:"lon" form:"lon" gorm:"column:lon;comment:;size:22;"`
      Lat  *float64 `json:"lat" form:"lat" gorm:"column:lat;comment:;size:22;"`
      Province  string `json:"province" form:"province" gorm:"column:province;comment:;size:32;"`
      City  string `json:"city" form:"city" gorm:"column:city;comment:;size:64;"`
      District  string `json:"district" form:"district" gorm:"column:district;comment:;size:128;"`
      Street  string `json:"street" form:"street" gorm:"column:street;comment:;size:128;"`
      Attr1  string `json:"attr1" form:"attr1" gorm:"column:attr1;comment:;size:128;"`
      Attr2  string `json:"attr2" form:"attr2" gorm:"column:attr2;comment:省市区区分;size:128;"`
      Attr3  string `json:"attr3" form:"attr3" gorm:"column:attr3;comment:公司类型（五金 电子 三小 印刷 家居 其它）;size:128;"`
      Attr4  string `json:"attr4" form:"attr4" gorm:"column:attr4;comment:单位类型;size:128;"`
      Title  string `json:"title" form:"title" gorm:"column:title;comment:;size:128;"`
      LogoImage  string `json:"logoImage" form:"logoImage" gorm:"column:logo_image;comment:;size:256;"`
      BackImage  string `json:"backImage" form:"backImage" gorm:"column:back_image;comment:;size:256;"`
}


// TableName SysDept 表名
func (SysDept) TableName() string {
  return "sys_dept"
}

