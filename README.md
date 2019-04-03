# pinyin-golang

[![Build Status](https://travis-ci.org/Lofanmi/pinyin-golang.svg)](https://travis-ci.org/Lofanmi/pinyin-golang)
[![codecov](https://codecov.io/gh/Lofanmi/pinyin-golang/branch/master/graph/badge.svg)](https://codecov.io/gh/Lofanmi/pinyin-golang)
[![Go Report Card](https://goreportcard.com/badge/github.com/Lofanmi/pinyin-golang)](https://goreportcard.com/report/github.com/Lofanmi/pinyin-golang)

`Go 语言` 的中文转拼音类库, 提供更为准确的中文转拼音解决方案.

拼音词库编译到二进制可执行文件中, 部署方便. `Go 1.7+` 单元测试通过.

`好用` ? 右上角 `Star` ! 欢迎 `Issue` 和 `Pull Request` , 我会不断改进它!

> 注: 词库来源于 [安正超](https://overtrue.me/) 的 PHP 开源项目: [overtrue/pinyin](https://github.com/overtrue/pinyin)

# 安装

```bash
go get -u -v github.com/Lofanmi/pinyin-golang/pinyin
```

# 用法

## 引入

```go
import (
	"github.com/Lofanmi/pinyin-golang/pinyin"
)
```

## 转换接口: Dict.Convert

输入中文字符串, 指定拼音与拼音之间的分隔号, 返回特定格式的拼音字符串.

```go
// 输入繁体中文, 输出为带 空格 分隔的拼音字符串
// ASCII 格式显示
// wo3 he2 shi2 neng2 bao4 fu4
s = dict.Convert(`我，何時能暴富？`, " ").ASCII()
fmt.Println(s)

// 输入简体中文, 输出为带 连字符- 分隔的拼音字符串
// Unicode 格式显示
// wǒ-hé-shí-néng-bào-fù
s = dict.Convert(`我，何时能暴富？`, "-").Unicode()
fmt.Println(s)

// 转换简体中文和繁体中文, 转换为带 斜杆/ 分隔的拼音字符串
// 不显示声调
// wo/he/shi/neng/bao/fu
s = dict.Convert(`我，何时能暴富？`, "/").None()
fmt.Println(s)
```

## 句子接口: Dict.Sentence

输入中文字符串, 保留标点符号, 并转换中文标点为英文标点, 返回特定格式的拼音字符串.

```go
// 输入繁体中文, 输出为带 空格 分隔的拼音字符串
// ASCII 格式显示
// wo3, he2 shi2 neng2 bao4 fu4?
s = dict.Sentence(`我，何時能暴富？`).ASCII()
fmt.Println(s)

// 输入简体中文, 输出为带 空格 分隔的拼音字符串
// Unicode 格式显示
// wǒ, hé shí néng bào fù?
s = dict.Sentence(`我，何时能暴富？`).Unicode()
fmt.Println(s)

// 转换简体中文和繁体中文, 转换为带 空格 分隔的拼音字符串
// 不显示声调
// wo, he shi neng bao fu?
s = dict.Sentence(`我，何时能暴富？`).None()
fmt.Println(s)
```

## 转换人名: Dict.Name

输入姓氏人名, 返回特定格式的拼音字符串.

```go
// 输入繁体中文, 输出为带 空格 分隔的人名拼音字符串
// ASCII 格式显示
// mo4 qi2 wo4 xi3 huan1 chi1 suan1 nai3
s = dict.Name(`万俟沃喜欢吃酸奶`, " ").ASCII()
fmt.Println(s)

// 输入简体中文, 输出为带 连字符- 分隔的人名拼音字符串
// Unicode 格式显示
// mò-qí-wò-xǐ-huan-chī-suān-nǎi
s = dict.Name(`万俟沃喜欢吃酸奶`, "-").Unicode()
fmt.Println(s)

// 转换简体中文和繁体中文, 转换为带 斜杆/ 分隔的人名拼音字符串
// 不显示声调
// mo/qi/wo/xi/huan/chi/suan/nai
s = dict.Name(`万俟沃喜欢吃酸奶`, "/").None()
fmt.Println(s)
```

## 转换拼音简写: Dict.Abbr

输入中文字符串, 指定拼音与拼音之间的分隔号, 返回特定格式的拼音字符串的简写.

```go
// 转换简体中文和繁体中文, 输出为带 连字符- 分隔的拼音字符串首字符
// m-q-w-x-h-c-s-n
s = dict.Abbr(`万俟沃喜欢吃酸奶`, "-")
fmt.Println(s)
```

## 转换为字符串 slice: ToSlice

有时候可能需要对转换的结果做进一步处理, 可以使用 `ToSlice` 接口:

```go
// wo3 he2 shi2 neng2 bao4 fu4
s = dict.Convert(`我，何時能暴富？`, " ").ASCII()
fmt.Println(s)

// [wo3 he2 shi2 neng2 bao4 fu4]
fmt.Printf("%v", pinyin.ToSlice(s))
```

# Contribution

欢迎提意见及完善词库

# License

MIT
