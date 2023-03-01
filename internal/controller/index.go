package controller

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	v1 "display/api/v1"
	"display/internal/model"
	"display/internal/service"

	"github.com/gogf/gf/encoding/gjson"
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

	service.View().Render(ctx, model.View{
		ContentType: "",
		Data:        "",
		Title:       "",
	})
	return
}

func (a *cIndex) Clock(ctx context.Context, req *v1.ClockReq) (res *v1.ClockRes, err error) {

	service.View().Render(ctx, model.View{
		ContentType: "",
		Data:        GetCurrentDate(),
		Title:       "",
	})
	return
}

//返回当前天时间
func GetCurrentDate() string {
	// 初始化全局变量
	var timestamp string
	resp, err := http.Get("http://api.m.taobao.com/rest/api3.do?api=mtop.common.getTimestamp")
	if err != nil {
		fmt.Println(err)
		return "1677675248035"
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "1677675248035"
	}

	// 解析json获取t字段
	j, err := gjson.DecodeToJson(string(body))

	if err != nil {
		fmt.Println(err)
		return "1677675248035"
	}
	timestamp = j.GetString("data.t")
	return timestamp
}
