package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type IndexReq struct {
	g.Meta `path:"/" method:"get" tags:"首页" summary:"首页"`
	ContentGetListCommonReq
}
type IndexRes struct {
	ContentGetListCommonRes
}

type DisplayReq struct {
	g.Meta `path:"/display" method:"get" tags:"播放显示" summary:"主页显示"`
}
type DisplayRes struct {
}

type ClockReq struct {
	g.Meta `path:"/clock" method:"get" tags:"倒计时" summary:"倒计时"`
}
type ClockRes struct {
}

type VideoReq struct {
	g.Meta `path:"/video" method:"get" tags:"视频播放" summary:"视频播放"`
}
type VideoRes struct {
}

type CurrentDate struct {
	Day         int
	Hour        int
	LunarDay    int
	LunarMonth  int
	LunarYear   int
	Minute      int
	Month       int
	Timestampms string
	Weekday     string
	Year        int
}

type PlaylistData struct {
	Url   string
	Title string
}
