package controller

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"

	v1 "display/api/v1"
	"display/internal/consts"
	"display/internal/model"
	"display/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

// 汉字发布，管理
var Characters = cCharacters{}

type cCharacters struct{}

type Dictionary struct {
	Character     string `json:"character"`
	Acjk          string `json:"acjk"`
	Set           string `json:"set"`
	Definition    string `json:"definition"`
	Pinyin        string `json:"pinyin"`
	Radical       string `json:"radical"`
	Decomposition string `json:"decomposition"`
}

// Index characters list
func (a *cCharacters) Index(ctx context.Context, req *v1.CharactersIndexReq) (res *v1.CharactersIndexRes, err error) {
	req.Type = consts.ContentTypeCharacters
	getListRes, err := service.Content().GetList(ctx, model.ContentGetListInput{
		Type:       req.Type,
		CategoryId: req.CategoryId,
		Page:       req.Page,
		Size:       req.Size,
		Sort:       req.Sort,
	})
	if err != nil {
		return nil, err
	}
	service.View().Render(ctx, model.View{
		ContentType: req.Type,
		Data:        getListRes,
		Title: service.View().GetTitle(ctx, &model.ViewGetTitleInput{
			ContentType: req.Type,
			CategoryId:  req.CategoryId,
		}),
	})
	return
}

func getSvg(ctx context.Context, ch string) string {
	svgstr := ""
	fontN := decUnicode(ch)

	filepath := g.Cfg().MustGet(ctx, "characters.svgsZhHans").String() + strconv.Itoa(fontN) + ".svg"

	if _, err := os.Stat(filepath); !os.IsNotExist(err) {
		// path/to/whatever exists
		data, err := ioutil.ReadFile(filepath)
		if err != nil {
			fmt.Println(err)
			return svgstr
		}
		svgstr = string(data)
	}
	return svgstr
}

//读取json文件
func readDictionarayZhans(ctx context.Context, c string) ([]string, int) {

	path := g.Cfg().MustGet(ctx, "characters.dictionaryZhHans").String()

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return []string{}, 0
	}
	s := string(data)
	r := regexp.MustCompile(`\{"character":"` + string(c) + `[^}]+\}`)
	match := r.FindString(s)
	sum := 0
	var pinyin []string
	if len(match) > 0 {
		l := match
		r2 := regexp.MustCompile(`"pinyin":\["([^]]+)"\]`)
		pinyinMatch := r2.FindStringSubmatch(l)
		if len(pinyinMatch) > 1 {
			pinyin = strings.Split(strings.ReplaceAll(pinyinMatch[1], "\"", ""), " ")
		} else {
			pinyin = []string{}
		}

		var dict Dictionary

		// 将 JSONL 字符串转换为一个 io.Reader 对象
		reader := bufio.NewReader(strings.NewReader(l))
		// 创建一个 JSON 解码器
		decoder := json.NewDecoder(reader)
		decoder.Decode(&dict)
		re := regexp.MustCompile(`\d+(\.\d+)?`)
		matches := re.FindAllString(dict.Acjk, -1)
		//计算字体笔划
		for _, str := range matches {
			if n, err := strconv.Atoi(str); err == nil {
				sum += n
			}
		}
		// fmt.Printf("总笔划：%d", sum)
	}
	return pinyin, sum
}

func formatPrononciation(a []string) string {
	s := ""
	first := true
	for _, b := range a {
		if !first {
			s += ", "
		} else {
			first = false
		}
		s += regexp.MustCompile(`\([^\)]*\)`).ReplaceAllString(b, "")
	}
	return s
}

func decUnicode(c string) int {
	r := 0
	r1 := int(c[0])
	if len(c) == 1 {
		return r1
	}
	r2 := int(c[1])
	if len(c) == 2 {
		return ((r1 & 31) << 6) + (r2 & 63)
	}
	r3 := int(c[2])
	if len(c) == 3 {
		return ((r1 & 15) << 12) + ((r2 & 63) << 6) + (r3 & 63)
	}
	r4 := int(c[3])
	if len(c) == 4 {
		return ((r1 & 7) << 18) + ((r2 & 63) << 12) + ((r3 & 63) << 6) + (r4 & 63)
	}
	return r
}

func getMax(s []int) int {
	max := s[0]
	for _, n := range s {
		if n > max {
			max = n
		}
	}
	return max
}

//Show  .cCharaters show
func (a *cCharacters) Show(ctx context.Context, req *v1.CharactersShowReq) (res *v1.CharactersShowRes, err error) {

	// g.Log().Info(ctx, consts.ContentTypeCharacters)
	// g.Log().Info(ctx, "----------------------")
	// g.Log().Info(ctx, getSvg("我"))
	// g.Log().Info(ctx, readDictionarayZhans("许"))
	// g.Log().Info(ctx, formatPrononciation(readDictionarayZhans("许")))
	// g.Log().Info(ctx, "----------------------")

	data, err := service.Content().GetList(ctx, model.ContentGetListInput{
		Type:       consts.ContentTypeCharacters,
		CategoryId: 15,
		Page:       req.Page,
		Size:       18,
		Sort:       0,
	})

	if err != nil {
		fmt.Print(data)
	}

	cop := make([]model.CharactersOutput, 0)
	resp := make(map[string]interface{})

	pinyin := []string{}
	num := 0
	maxMum := []int{}

	for _, item := range data.List {

		pinyin, num = readDictionarayZhans(ctx, item.Content.Title)
		maxMum = append(maxMum, num)
		cop = append(cop, model.CharactersOutput{
			Svg:    getSvg(ctx, item.Content.Title),
			Pinyin: formatPrononciation(pinyin),
			Num:    num,
		})
	}

	resp["Cop"] = cop
	resp["MaxNum"] = getMax(maxMum) + 3

	service.View().Render(ctx, model.View{
		ContentType: consts.ContentTypeCharacters,
		Data:        resp,
	})
	return
}

// Detail .cCharaters details
func (a *cCharacters) Detail(ctx context.Context, req *v1.CharactersDetailReq) (res *v1.CharactersDetailRes, err error) {
	getDetailRes, err := service.Content().GetDetail(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if getDetailRes == nil {
		service.View().Render404(ctx)
		return nil, nil
	}
	if err = service.Content().AddViewCount(ctx, req.Id, 1); err != nil {
		return res, err
	}
	service.View().Render(ctx, model.View{
		ContentType: consts.ContentTypeCharacters,
		Data:        getDetailRes,
		Title: service.View().GetTitle(ctx, &model.ViewGetTitleInput{
			ContentType: getDetailRes.Content.Type,
			CategoryId:  getDetailRes.Content.CategoryId,
			CurrentName: getDetailRes.Content.Title,
		}),
		BreadCrumb: service.View().GetBreadCrumb(ctx, &model.ViewGetBreadCrumbInput{
			ContentId:   getDetailRes.Content.Id,
			ContentType: getDetailRes.Content.Type,
			CategoryId:  getDetailRes.Content.CategoryId,
		}),
	})
	return
}
