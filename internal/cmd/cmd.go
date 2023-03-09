package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"display/internal/controller"
	"display/internal/dao"
	"display/internal/service"

	"github.com/6tail/lunar-go/calendar"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/os/gcron"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/goai"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gmode"

	"display/internal/consts"
	"display/utility/response"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start focus server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			var (
				s   = g.Server()
				oai = s.GetOpenApi()
			)

			// OpenApi自定义信息
			oai.Info.Title = `API Reference`
			oai.Config.CommonResponse = response.JsonRes{}
			oai.Config.CommonResponseDataField = `Data`

			// 静态目录设置
			uploadPath := g.Cfg().MustGet(ctx, "upload.path").String()
			if uploadPath == "" {
				g.Log().Fatal(ctx, "文件上传配置路径不能为空")
			}
			s.AddStaticPath("/upload", uploadPath)

			g.View().SetConfigWithMap(g.Map{
				"paths": g.Slice{"/resource/template"},
			})

			gcron.AddSingleton("0 0/5 * * * *", func() {
				ms := GetCurrentDate()
				s2 := saveCurrentDate(ms)
				v, _ := service.Setting().GetVar(ctx, "CurrentDate")
				if !v.IsNil() {
					//更新
					dao.Setting.Ctx(ctx).Data(g.Map{
						"v": s2,
					}).Where("k", "CurrentDate").Update()
				} else {
					//插入
					dao.Setting.Ctx(ctx).Data(g.Map{
						"k": "CurrentDate",
						"v": s2,
					}).OmitEmpty().Insert()
				}
			})

			// 注册模板函数
			g.View().BindFunc("split", func(str, sep string) []string {
				return strings.Split(str, sep)
			})

			// HOOK, 开发阶段禁止浏览器缓存,方便调试
			if gmode.IsDevelop() {
				s.BindHookHandler("/*", ghttp.HookBeforeServe, func(r *ghttp.Request) {
					r.Response.Header().Set("Cache-Control", "no-store")
				})
			}

			// 前台系统自定义错误页面
			s.BindStatusHandler(401, func(r *ghttp.Request) {
				if !gstr.HasPrefix(r.URL.Path, "/admin") {
					service.View().Render401(r.Context())
				}
			})
			s.BindStatusHandler(403, func(r *ghttp.Request) {
				if !gstr.HasPrefix(r.URL.Path, "/admin") {
					service.View().Render403(r.Context())
				}
			})
			s.BindStatusHandler(404, func(r *ghttp.Request) {
				if !gstr.HasPrefix(r.URL.Path, "/admin") {
					service.View().Render404(r.Context())
				}
			})
			s.BindStatusHandler(500, func(r *ghttp.Request) {
				if !gstr.HasPrefix(r.URL.Path, "/admin") {
					service.View().Render500(r.Context())
				}
			})

			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
				)

				group.Bind(
					controller.Index,      // 首页
					controller.Login,      // 登录
					controller.Register,   // 注册
					controller.Category,   // 栏目
					controller.Topic,      // 主题
					controller.Article,    // 文章
					controller.Characters, //汉字
					controller.Search,     // 搜索
					controller.Captcha,    // 验证码
					controller.User,       // 用户
				)

				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(
						service.Middleware().Ctx,
						service.Middleware().ResponseHandler,
					)
					group.Bind(
						controller.Hello, //hello diy 页面
					)
				})

				// 权限控制路由
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(service.Middleware().Auth)
					group.Bind(
						controller.Profile, // 个人
						controller.Setting, // 数字内容设置
						controller.Content, // 内容
						controller.File,    // 文件
					)
				})
			})
			// 自定义丰富文档
			enhanceOpenAPIDoc(s)
			// 启动Http Server
			s.Run()
			return
		},
	}
)

func enhanceOpenAPIDoc(s *ghttp.Server) {
	openapi := s.GetOpenApi()
	openapi.Config.CommonResponse = ghttp.DefaultHandlerResponse{}
	openapi.Config.CommonResponseDataField = `Data`

	// API description.
	openapi.Info.Title = `Focus Project`
	openapi.Info.Description = ``

	// Sort the tags in custom sequence.
	openapi.Tags = &goai.Tags{
		{Name: consts.OpenAPITagNameIndex},
		{Name: consts.OpenAPITagNameLogin},
		{Name: consts.OpenAPITagNameRegister},
		{Name: consts.OpenAPITagNameArticle},
		{Name: consts.OpenAPITagNameTopic},
		{Name: consts.OpenAPITagNameContent},
		{Name: consts.OpenAPITagNameSearch},
		{Name: consts.OpenAPITagNameInteract},
		{Name: consts.OpenAPITagNameCategory},
		{Name: consts.OpenAPITagNameProfile},
		{Name: consts.OpenAPITagNameUser},
		{Name: consts.OpenAPITagNameMess},
	}
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

func saveCurrentDate(timestamp string) string {

	ms, _ := strconv.ParseInt(timestamp, 10, 64)

	// 将毫秒数转换为UTC时间
	t := time.Unix(ms/1000, (ms%1000)*int64(time.Millisecond))

	// 获取年、月、日、时、分
	year, month, day := t.Date()
	hour, min, _ := t.Clock()

	// 获取农历日期
	lunarDate := calendar.NewLunarFromYmd(year, int(month), day)
	lunarYear, lunarMonth, lunarDay := lunarDate.GetYear(), lunarDate.GetMonth(), lunarDate.GetDay()

	// 获取星期
	weekday := t.Weekday().String()

	// 将结果转换为JSON格式并输出
	result := map[string]interface{}{
		"year":        year,
		"month":       int(month),
		"day":         day,
		"hour":        hour,
		"minute":      min,
		"lunarYear":   lunarYear,
		"lunarMonth":  lunarMonth,
		"lunarDay":    lunarDay,
		"weekday":     weekday,
		"timestampms": timestamp,
	}

	jsonResult, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}
	return string(jsonResult)
}
