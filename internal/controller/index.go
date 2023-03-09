package controller

import (
	"context"
	"encoding/json"
	"strconv"
	"strings"
	"time"

	v1 "display/api/v1"
	"display/internal/model"
	"display/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
)

// 首页接口
var Index = cIndex{}

type cIndex struct{}

func (a *cIndex) Index(ctx context.Context, req *v1.IndexReq) (res *v1.IndexRes, err error) {
	if getListRes, err := service.Content().GetList(ctx, model.ContentGetListInput{
		Type:       req.Type,
		CategoryId: req.CategoryId,
		Page:       req.Page,
		Size:       req.Size,
		Sort:       req.Sort,
	}); err != nil {
		return nil, err
	} else {
		service.View().Render(ctx, model.View{
			ContentType: req.Type,
			Data:        getListRes,
			Title:       "首页",
		})
	}
	return
}

func (a *cIndex) Display(ctx context.Context, req *v1.DisplayReq) (res *v1.DisplayRes, err error) {
	v, _ := service.Setting().GetVar(ctx, "HomepageUrl")
	s := g.NewVar(v).String()
	s2 := []string{"http://localhost" + g.Cfg().MustGet(ctx, "server.address").String() + "/clock"}
	s3 := make(map[string][]string)
	if v != nil && gstr.Contains(s, ",") {
		s2 = strings.Split(s, ",")
	} else if v != nil {
		s2 = []string{s}
	}
	s3["HomepageUrl"] = s2
	service.View().Render(ctx, model.View{
		ContentType: "",
		Data:        s3,
		Title:       "",
	})
	return
}

func (a *cIndex) Clock(ctx context.Context, req *v1.ClockReq) (res *v1.ClockRes, err error) {

	v, _ := service.Setting().GetVar(ctx, "CurrentDate")

	var date2 *v1.CurrentDate

	json.Unmarshal([]byte(g.NewVar(v).String()), &date2)

	service.View().Render(ctx, model.View{
		ContentType: "",
		Data:        date2,
		Title:       "",
	})
	return
}

//视频播放页
func (a *cIndex) Video(ctx context.Context, req *v1.VideoReq) (res *v1.VideoRes, err error) {

	CurrentDate, _ := service.Setting().GetVar(ctx, "CurrentDate")
	var date2 *v1.CurrentDate

	json.Unmarshal([]byte(g.NewVar(CurrentDate).String()), &date2)

	examDateStr, _ := service.Setting().GetVar(ctx, "ExamDate")
	Playlist, _ := service.Setting().GetVar(ctx, "Playlist")

	examDate, _ := time.Parse("2006-1-2", g.NewVar(examDateStr).String()) // 将考试日期字符串解析为 time.Time 类型
	timeInt64, _ := strconv.ParseInt(date2.Timestampms, 10, 64)
	currentTime := time.Unix(0, timeInt64*int64(time.Millisecond)) // 将当前时间戳转换为 time.Time 类型

	DaysLeft := int(examDate.Sub(currentTime).Hours() / 24) // 计算距离考试还有几天

	s := g.NewVar(Playlist).String()
	var pd []v1.PlaylistData
	if !Playlist.IsNil() && gstr.Contains(s, ",") {
		s2 := strings.Split(s, ",")
		for _, v := range s2 {
			cpd := v1.PlaylistData{}
			if gstr.Contains(v, "|") {
				s3 := strings.Split(v, "|")
				cpd.Url = s3[0]
				cpd.Title = s3[1]
			}
			pd = append(pd, cpd)
		}

	} else if !Playlist.IsNil() {
		if gstr.Contains(s, "|") {
			cpd := v1.PlaylistData{}
			s3 := strings.Split(s, "|")
			cpd.Url = s3[0]
			cpd.Title = s3[1]
			pd = append(pd, cpd)
		}
	}

	m := g.Map{
		"CurrentDate": date2,
		"Playlist":    pd,
		"DaysLeft":    DaysLeft,
	}

	service.View().Render(ctx, model.View{
		ContentType: "",
		Data:        m,
		Title:       "",
	})
	return
}
