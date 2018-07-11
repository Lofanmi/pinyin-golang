# pinyin-golang

[![Build Status](https://travis-ci.org/Lofanmi/pinyin-golang.svg)](https://travis-ci.org/Lofanmi/pinyin-golang)
[![codecov](https://codecov.io/gh/Lofanmi/pinyin-golang/branch/master/graph/badge.svg)](https://codecov.io/gh/Lofanmi/pinyin-golang)
[![Go Report Card](https://goreportcard.com/badge/github.com/Lofanmi/pinyin-golang)](https://goreportcard.com/report/github.com/Lofanmi/pinyin-golang)

采用 [CC-CEDICT](https://cc-cedict.org/wiki/start) 词典的中文转拼音类库, 提供更为准确的中文转拼音解决方案.

使用 `Go` 编写, 觉得好用的给 `Star` 呗~

觉得**好用**, 而不是觉得**有用**. 如果不好用, 欢迎向我提 issue, 我会抽空不断改进它!

真希望有人打钱打钱打钱给我啊哈哈哈哈!!!

Go 1.7+测试通过, 1.6及以下应该也可以, 不过单元测试跑不了.

# 优势

采用 cc-cedict 词库, 并在其基础上针对语义优化词库顺序, 尽可能提高转换准确率.

1. 使用反引号分隔的新格式, 解析起来更容易;
2. 按照中文字符长度排序, 字符长的优先被替换为拼音;
3. 部分常用字[大: da4/dai4]是多音字, 调换前后顺序能更大概率匹配到正确的读音;
4. 同音字前后次序, 取决于英文释义的丰富程度. 一般来说, 越常用的字(词)意义越广泛, 转换为拼音时应给予更高的优先级;
5. 原词典的韵母 v 使用 u: 表达, 本词典采用 v 的形式; 如绿色 [lu:4 se4] => [lv4 se4];
6. 原词典的轻声, 使用5声调, 本词典略去轻声调, 不加数字5; 如打量 [da3 liang5] => [da3 liang];
7. 接口简洁而有丰富规范:
	- 支持转换时指定拼音与拼音之间的分隔号;
	- 支持不带声调的输出;
	- 支持带声调的 ASCII [zhong1 wen2] 和 Unicode [zhōng wén] 两种形式的输出;
	- 支持姓氏转换;
	- 支持拼音首字母缩写形式;
	- 支持句子和段落转换;
8. 词典格式:

原格式:
> Traditional Simplified [pin1 yin1] /English equivalent 1/equivalent 2/

现格式:
> Traditional'Simplified'pin1 yin1'English equivalent 1/equivalent 2

# 如何安装

```bash
go get -u -v github.com/Lofanmi/pinyin-golang
```

# 用法

## 转换接口: Dict.Convert

输入中文字符串, 指定拼音与拼音之间的分隔号, 返回特定格式的拼音字符串.

```go
import (
	"github.com/Lofanmi/pinyin-golang/pinyin"
)

func test() {
	filename := "/path/to/cedict.lofanmi"
	dict, err := pinyin.NewDict(filename)
	if err != nil {
		panic(err)
	}

	s = `我，何时能暴富？`

	// 输入繁体中文, 输出为带 空格 分隔的拼音字符串
	// ASCII 格式显示
	// wo3 he2 neng2 bao4 fu4
	s = dict.Convert(s, " ", Traditional).ASCII()
	
	// 输入简体中文, 输出为带 连字符- 分隔的拼音字符串
	// Unicode 格式显示
	// wǒ-hé-shí-néng-bào-fù
	s = dict.Convert(s, "-", Simplified).Unicode()
	
	// 转换简体中文和繁体中文, 转换为带 斜杆/ 分隔的拼音字符串
	// 不显示声调
	// wo/he/shi/neng/bao/fu
	s = dict.Convert(s, "/", All).None()
}
```

## 句子接口: Dict.Sentence

输入中文字符串, 保留标点符号, 并转换中文标点为英文标点, 返回特定格式的拼音字符串.

```go
func test() {
	s = `我，何时能暴富？`

	// 输入繁体中文, 输出为带 空格 分隔的拼音字符串
	// ASCII 格式显示
	// wo3, he2 shi2 neng2 bao4 fu4?
	s = dict.Sentence(s, Traditional).ASCII()
	
	// 输入简体中文, 输出为带 连字符- 分隔的拼音字符串
	// Unicode 格式显示
	// wǒ, hé shí néng bào fù?
	s = dict.Sentence(s, Simplified).Unicode()
	
	// 转换简体中文和繁体中文, 转换为带 斜杆/ 分隔的拼音字符串
	// 不显示声调
	// wo, he shi neng bao fu?
	s = dict.Sentence(s, All).None()
}
```

## 转换人名: Dict.Name

输入姓氏人名, 返回特定格式的拼音字符串.

```go
func test() {
	s = `万俟沃喜欢吃酸奶`

	// 输入繁体中文, 输出为带 空格 分隔的人名拼音字符串
	// ASCII 格式显示
	// mo4 qi2 wo4 xi3 huan1 chi1 suan1 nai3
	s = dict.Name(s, " ", Traditional).ASCII()
	
	// 输入简体中文, 输出为带 连字符- 分隔的人名拼音字符串
	// Unicode 格式显示
	// mò-qí-wò-xǐ-huan-chī-suān-nǎi
	s = dict.Name(s, "-", Simplified).Unicode()
	
	// 转换简体中文和繁体中文, 转换为带 斜杆/ 分隔的人名拼音字符串
	// 不显示声调
	// mo/qi/wo/xi/huan/chi/suan/nai
	s = dict.Name(s, "/", All).None()
}
```

## 转换拼音简写: Dict.Abbr

输入中文字符串, 指定拼音与拼音之间的分隔号, 返回特定格式的拼音字符串的简写.

```go
func test() {
	s = `万俟沃喜欢吃酸奶`

	// 转换简体中文和繁体中文, 输出为带 连字符- 分隔的拼音字符串首字符
	// m-q-w-x-h-c-s-n
	s = dict.Abbr(s, "-", All)
}
```

## 转换为字符串slice: ToSlice

有时候可能需要对转换的结果做进一步处理, 可以使用 `ToSlice` 接口:

```go
func test() {
	s = `我，何时能暴富？`

	// wo3 he2 neng2 bao4 fu4
	s = dict.Convert(s, " ", Traditional).ASCII()

	// [wo3 he2 neng2 bao4 fu4]
	fmt.Printf("%v", ToSlice(s))
}
```

# Contribution

欢迎提意见及完善词库

# License

MIT
