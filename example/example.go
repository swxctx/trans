package main

import (
	"fmt"

	"github.com/swxctx/trans"
)

func main() {
	// 加载配置表
	if err := trans.Reload(&trans.Config{
		FilePath:  "./trans.csv",
		SourceLan: "zh",
	}); err != nil {
		fmt.Printf("Reload err-> %v", err)
		return
	}

	text := "你好"
	// 翻译中文
	fmt.Printf("原文本: %s, 中文翻译文本：%s\n", text, trans.Translate(trans.ZH, text))
	// 翻译英文
	fmt.Printf("原文本: %s, 英文翻译文本：%s\n", text, trans.Translate(trans.EN, text))
	// 翻译西班牙语
	fmt.Printf("原文本: %s, 西班牙语翻译文本：%s\n", text, trans.Translate(trans.ES, text))
	// 翻译法语
	fmt.Printf("原文本: %s, 法语翻译文本：%s\n", text, trans.Translate(trans.FR, text))
	// 翻译日语
	fmt.Printf("原文本: %s, 日语翻译文本：%s\n", text, trans.Translate(trans.JA, text))
}
