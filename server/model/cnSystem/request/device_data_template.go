package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/cnSystem"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type DeviceDataTemplateSearch struct{
    cnSystem.DeviceDataTemplate
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
