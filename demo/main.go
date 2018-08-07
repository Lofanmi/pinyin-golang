package main

import (
	"fmt"

	"github.com/Lofanmi/pinyin-golang/pinyin"
)

func main() {
	filename := "E:/go/src/github.com/Lofanmi/pinyin-golang/pinyin/cedict.lofanmi"
	dict, err := pinyin.NewDict(filename)
	if err != nil {
		panic(err)
	}

	s := ""

	// ----
	// 转换接口: Dict.Convert
	// ----
	// 输入繁体中文, 输出为带 空格 分隔的拼音字符串
	// ASCII 格式显示
	// wo3 he2 shi2 neng2 bao4 fu4
	s = dict.Convert(`我，何時能暴富？`, " ", pinyin.Traditional).ASCII()
	fmt.Println(s)

	// 输入简体中文, 输出为带 连字符- 分隔的拼音字符串
	// Unicode 格式显示
	// wǒ-hé-shí-néng-bào-fù
	s = dict.Convert(`我，何时能暴富？`, "-", pinyin.Simplified).Unicode()
	fmt.Println(s)

	// 转换简体中文和繁体中文, 转换为带 斜杆/ 分隔的拼音字符串
	// 不显示声调
	// wo/he/shi/neng/bao/fu
	s = dict.Convert(`我，何时能暴富？`, "/", pinyin.All).None()
	fmt.Println(s)

	// ----
	// 句子接口: Dict.Sentence
	// ----
	// 输入繁体中文, 输出为带 空格 分隔的拼音字符串
	// ASCII 格式显示
	// wo3, he2 shi2 neng2 bao4 fu4?
	s = dict.Sentence(`我，何時能暴富？`, pinyin.Traditional).ASCII()
	fmt.Println(s)

	// 输入简体中文, 输出为带 空格 分隔的拼音字符串
	// Unicode 格式显示
	// wǒ, hé shí néng bào fù?
	s = dict.Sentence(`我，何时能暴富？`, pinyin.Simplified).Unicode()
	fmt.Println(s)

	// 转换简体中文和繁体中文, 转换为带 空格 分隔的拼音字符串
	// 不显示声调
	// wo, he shi neng bao fu?
	s = dict.Sentence(`我，何时能暴富？`, pinyin.All).None()
	fmt.Println(s)

	// ----
	// 转换人名: Dict.Name
	// ----
	// 输入繁体中文, 输出为带 空格 分隔的人名拼音字符串
	// ASCII 格式显示
	// mo4 qi2 wo4 xi3 huan1 chi1 suan1 nai3
	s = dict.Name(`万俟沃喜欢吃酸奶`, " ", pinyin.Traditional).ASCII()
	fmt.Println(s)

	// 输入简体中文, 输出为带 连字符- 分隔的人名拼音字符串
	// Unicode 格式显示
	// mò-qí-wò-xǐ-huan-chī-suān-nǎi
	s = dict.Name(`万俟沃喜欢吃酸奶`, "-", pinyin.Simplified).Unicode()
	fmt.Println(s)

	// 转换简体中文和繁体中文, 转换为带 斜杆/ 分隔的人名拼音字符串
	// 不显示声调
	// mo/qi/wo/xi/huan/chi/suan/nai
	s = dict.Name(`万俟沃喜欢吃酸奶`, "/", pinyin.All).None()
	fmt.Println(s)

	// ----
	// 转换拼音简写: Dict.Abbr
	// ----
	// 转换简体中文和繁体中文, 输出为带 连字符- 分隔的拼音字符串首字符
	// m-q-w-x-h-c-s-n
	s = dict.Abbr(`万俟沃喜欢吃酸奶`, "-", pinyin.All)
	fmt.Println(s)

	// ----
	// 转换为字符串 slice: ToSlice
	// ----
	// wo3 he2 shi2 neng2 bao4 fu4
	s = dict.Convert(`我，何時能暴富？`, " ", pinyin.Traditional).ASCII()
	fmt.Println(s)

	// [wo3 he2 shi2 neng2 bao4 fu4]
	fmt.Printf("%v", pinyin.ToSlice(s))
}
