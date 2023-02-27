package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type HelloReq struct {
	g.Meta `path:"/hello" method:"get" tags:"hello" summary:"hello"`
}
type HelloRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}
type HelloDetailReq struct {
	g.Meta `path:"/hello/abc" method:"get" tags:"hello" summary:"hello"`
}
type HelloDetailRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>"`
}
