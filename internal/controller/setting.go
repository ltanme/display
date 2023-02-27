package controller

import (
	"context"
	"fmt"

	v1 "display/api/v1"
	"display/internal/dao"
	"display/internal/model"
	"display/internal/model/entity"
	"display/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// setting 管理
var Setting = cSetting{}

type cSetting struct{}

func (a *cSetting) Index(ctx context.Context, req *v1.SettingIndexReq) (res *v1.SettingIndexRes, err error) {

	service.View().Render(ctx, model.View{
		ContentType: "",
		Data:        "getListRes",
	})
	return
}

func (a *cSetting) Save(ctx context.Context, req *v1.SettingSaveReq) (res *v1.SettingSaveRes, err error) {
	// g.Log().Info(ctx, req)
	err2 := service.Setting().Set(ctx, req.K, string(req.V))
	fmt.Println(err2)
	// aa := service.Setting().Set(ctx, string(req.K), string(req.V))
	// fmt.Println(aa)

	// service.Setting().GetVar("homepageUrl")

	r, err := dao.Setting.SettingDao.DB().Save(ctx, "gf_setting", gdb.Map{
		"k": "1111",
		"v": "3333",
	})

	fmt.Println(r)

	dao.Setting.Ctx(ctx).Data(entity.Setting{
		K: "HomepageUrl",
		V: "33333333333333333333",
	}).Save()

	g.Log().Info(ctx, req.K, req.V)
	// v, err := service.Setting().GetVar(ctx, "abcd")

	g.Log().Info(ctx, "查到v值", 0)
	return
}
