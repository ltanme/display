package controller

import (
	"context"

	v1 "display/api/v1"
	"display/internal/model"
	"display/internal/service"
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
		Data:        "",
		Title:       "",
	})
	return
}
