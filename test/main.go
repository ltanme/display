package main

import (
	"fmt"
	"time"
)

func main() {
	examDateStr := "2023-5-25"
	examDate, _ := time.Parse("2006-1-2", examDateStr) // 将考试日期字符串解析为 time.Time 类型

	currentTime := time.Unix(0, 1678348565301*int64(time.Millisecond)) // 将当前时间戳转换为 time.Time 类型

	daysLeft := int(examDate.Sub(currentTime).Hours() / 24) // 计算距离考试还有几天

	fmt.Printf("距离考试还有 %d 天\n", daysLeft)
}
