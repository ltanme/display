package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CharactersIndexReq struct {
	g.Meta `path:"/characters" method:"get" tags:"汉字" summary:"展示Characters列表页面"`
	ContentGetListCommonReq
}
type CharactersIndexRes struct {
	ContentGetListCommonRes
}

type CharactersShowReq struct {
	g.Meta `path:"/characters/show" method:"get" tags:"显示汉字笔顺" summary:"展示Characters动画页面" `
	CommonPaginationReq
}

type CharactersShowRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}

type CharactersDetailReq struct {
	g.Meta `path:"/characters/{Id}" method:"get" tags:"汉字" summary:"展示Characters详情页面" `
	Id     uint `in:"path" v:"min:1#请选择查看的内容" dc:"内容id"`
}

type CharactersDetailRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}
