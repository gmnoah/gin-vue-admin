package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/cnSystem"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type SysDeptSearch struct{
    cnSystem.SysDept
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
