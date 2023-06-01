# trans
Golang翻译库，用于项目国际化语言处理，基于配置文件翻译。

## 配置表格式

|  zh | en  | es  |  fr | ja  |
| ------------ | ------------ | ------------ | ------------ | ------------ |
| 你好  |  hello | Hola!  |  Bonjour |  こんにちは |

- 说明
  - 第一行为语言代码，可参照[语言代码表](http://www.lingoes.net/zh/translator/langcode.htm)配置
  - 每一列配置对应语言代码的文本
  - 可参照[例子](./example/trans.csv)进行配置

## 代码使用

```go
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
```

- 执行输出结果

```shell
原文本: 你好, 中文翻译文本：你好
原文本: 你好, 英文翻译文本：hello
原文本: 你好, 西班牙语翻译文本：Hola!
原文本: 你好, 法语翻译文本：Bonjour
原文本: 你好, 日语翻译文本：こんにちは
```

## 使用说明

### Config

```go
FilePath: 配置文件的路径
SourceLan: 原文本的语言代码，例如：需要将中文翻译为其他语言，那么 SourceLan 对应值得为 zh，如果是将英语翻译为其他语言，那么值为 en
```

### 翻译方法

```go
trans.Translate(trans.EN, text)

第一个参数：需要为什么语言
第二个参数：需要翻译的文本
```