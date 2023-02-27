package controller

import (
	"context"
	v1 "display/api/v1"
	"display/internal/model"
	"display/internal/service"
)

var Hello = cHello{}

type cHello struct{}

func (a *cHello) Index(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {
	service.View().Render(ctx, model.View{
		ContentType: "hello content",
		Data:        "this data",
		Title:       service.View().GetTitle(ctx, &model.ViewGetTitleInput{}),
	})
	return
}

func (a *cHello) Abc(ctx context.Context, req *v1.HelloDetailReq) (res *v1.HelloDetailRes, err error) {

	service.View().Render(ctx, model.View{
		ContentType: "hello content",
		Data:        "this data",
		Title:       service.View().GetTitle(ctx, &model.ViewGetTitleInput{}),
	})
	return
}
