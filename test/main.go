package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gogf/gf/encoding/gjson"
)

func main() {
	// 初始化全局变量
	var timestamp string

	resp, err := http.Get("http://api.m.taobao.com/rest/api3.do?api=mtop.common.getTimestamp")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 解析json获取t字段
	j, err := gjson.DecodeToJson(string(body))

	// 将t字段转换为时间格式
	tm, err := strconv.ParseInt(j.GetString("data.t"), 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	timestamp = time.Unix(0, tm*int64(time.Millisecond)).Format("2006-01-02 15:04:05")

	fmt.Println(timestamp)

	// 阻塞主线程
	select {}
}
