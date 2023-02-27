package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type SettingIndexReq struct {
	g.Meta `path:"/setting/index" method:"get" tags:"设置中心" summary:"设置首页"`
}
type SettingIndexRes struct {
}

type SettingSaveReq struct {
	g.Meta `path:"/setting/save" method:"post" tags:"设置中心保存" summary:"restfull"`
	K      string `json:"k" dc:"设置key"`
	V      string `json:"v" dc:"设置value"`
}
type SettingSaveRes struct {
}
