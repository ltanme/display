package main

import (
	"fmt"
	"html"
	"regexp"
)

func unichr(u int8) string {
	htmlEntity := fmt.Sprintf("&#%d;", u)
	return html.UnescapeString(htmlEntity)
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

func main() {

	// filepath := "E:/aa/go-demo/display/resource/public/svgsZhHant/19968.svg"

	// if _, err := os.Stat(filepath); !os.IsNotExist(err) {
	// 	// path/to/whatever exists
	// 	data, err := ioutil.ReadFile(filepath)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	fmt.Println(string(data))
	// }

	str := `"acjk":"每⿱𠂉2母.5"`
	re := regexp.MustCompile(`"acjk":\d+(\.\d+)?`)
	matches := re.FindAllString(str, -1)
	fmt.Println(matches)

}
