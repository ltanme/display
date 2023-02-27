package model

// CharactersOutput 输出字体结构表
type CharactersOutput struct {
	Svg    string `json:"svg" description:"字体svg"`
	Pinyin string `json:"pinyin" description:"拼音"`
	Num    int    `json:"num" description:"笔画数"`
}

// CharactersInput 输入字体结构表
type CharactersInput struct {
	Svg    string `json:"svg" description:"字体svg"`
	Pinyin string `json:"pinyin" description:"拼音"`
}
