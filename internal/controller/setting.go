package controller

import (
	"context"

	v1 "display/api/v1"
	"display/internal/dao"
	"display/internal/model"
	"display/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

// setting 管理
var Setting = cSetting{}

type cSetting struct{}

func (a *cSetting) Index(ctx context.Context, req *v1.SettingIndexReq) (res *v1.SettingIndexRes, err error) {
	HomepageUrl, _ := service.Setting().GetVar(ctx, "HomepageUrl")
	ExamDate, _ := service.Setting().GetVar(ctx, "ExamDate")
	Playlist, _ := service.Setting().GetVar(ctx, "Playlist")
	m := g.Map{"HomepageUrl": HomepageUrl, "ExamDate": ExamDate, "Playlist": Playlist}
	service.View().Render(ctx, model.View{
		ContentType: "",
		Data:        m,
	})
	return
}

func (a *cSetting) Save(ctx context.Context, req *v1.SettingSaveReq) (res *v1.SettingSaveRes, err error) {

	if req.K == "HomepageUrl" {
		v, _ := service.Setting().GetVar(ctx, "HomepageUrl")
		if !v.IsNil() {
			//更新
			dao.Setting.Ctx(ctx).Data(req).Where("k", req.K).Update()
		} else {
			//插入
			dao.Setting.Ctx(ctx).Data(req).OmitEmpty().Insert()
		}
	}

	if req.K == "ExamDate" {
		v, _ := service.Setting().GetVar(ctx, "ExamDate")
		if !v.IsNil() {
			//更新
			dao.Setting.Ctx(ctx).Data(req).Where("k", req.K).Update()
		} else {
			//插入
			dao.Setting.Ctx(ctx).Data(req).OmitEmpty().Insert()
		}
	}

	if req.K == "Playlist" {
		v, _ := service.Setting().GetVar(ctx, "Playlist")
		if !v.IsNil() {
			//更新
			dao.Setting.Ctx(ctx).Data(req).Where("k", req.K).Update()
		} else {
			//插入
			dao.Setting.Ctx(ctx).Data(req).OmitEmpty().Insert()
		}
	}

	return
}
