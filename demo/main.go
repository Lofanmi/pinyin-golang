package main

import (
	"fmt"

	"pinyin-golang/pinyin"
)

func main() {
	dict := pinyin.NewDict()

	// ----
	// 简单用法
	// ----

	// Redis shì yí gè Key-Value cún chǔ xì tǒng.
	str := `Redis 是一个 Key-Value 存储系统。`
	fmt.Println(dict.Sentence(str).Unicode())

	s := ""

	// wǒ hé shí néng bào fù
	s = dict.Convert(`我，何時能暴富？`, " ").Unicode()
	fmt.Println(s)
	// wǒ, hé shí néng bào fù?
	s = dict.Sentence(`我，何時能暴富？`).Unicode()
	fmt.Println(s)

	// ----
	// 转换接口: Dict.Convert
	// ----

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

	// ----
	// 句子接口: Dict.Sentence
	// ----

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

	// ----
	// 转换人名: Dict.Name
	// ----

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

	// ----
	// 转换拼音简写: Dict.Abbr
	// ----

	// 转换简体中文和繁体中文, 输出为带 连字符- 分隔的拼音字符串首字符
	// m-q-w-x-h-c-s-n
	s = dict.Abbr(`万俟沃喜欢吃酸奶`, "-")
	fmt.Println(s)

	// ----
	// 转换为字符串 slice: ToSlice
	// ----
	// wo3 he2 shi2 neng2 bao4 fu4
	s = dict.Convert(`我，何時能暴富？`, " ").ASCII()
	fmt.Println(s)

	// [wo3 he2 shi2 neng2 bao4 fu4]
	fmt.Printf("%v\n", pinyin.ToSlice(s))

	// ----
	// 新增功能: 只转换中文和繁体字符
	// ----
	fmt.Println("\n---- 只转换中文和繁体字符的新功能 ----")

	// 使用 ConvertOnlyChinese
	fmt.Println("使用 ConvertOnlyChinese:")
	// Ni3-hao3-我-de-peng2-you3!
	s = dict.ConvertOnlyChinese(`Nihao我的朋友!`, "-").ASCII()
	fmt.Println(s)

	// Nǐ-hǎo-wǒ-de-péng-yǒu!
	s = dict.ConvertOnlyChinese(`Nihao我的朋友!`, "-").Unicode()
	fmt.Println(s)

	// Ni-hao-wo-de-peng-you!
	s = dict.ConvertOnlyChinese(`Nihao我的朋友!`, "-").None()
	fmt.Println(s)

	// 使用 SentenceOnlyChinese
	fmt.Println("\n使用 SentenceOnlyChinese:")
	// Ni hao wo3 de peng2 you3!
	s = dict.SentenceOnlyChinese(`Nihao我的朋友!`).ASCII()
	fmt.Println(s)

	// Ni hao wǒ de péng yǒu!
	s = dict.SentenceOnlyChinese(`Nihao我的朋友!`).Unicode()
	fmt.Println(s)

	// Ni hao wo de peng you!
	s = dict.SentenceOnlyChinese(`Nihao我的朋友!`).None()
	fmt.Println(s)

	// 更复杂的例子
	fmt.Println("\n更复杂的例子:")
	// Hello, wo3 shi4 yi1 ge4 example! 123 numbers bu4 hui4 bei4 bian4 huan4.
	s = dict.SentenceOnlyChinese(`Hello, 我是一个 example! 123 numbers不会被变换.`).ASCII()
	fmt.Println(s)

	// Hello, wǒ shì yī gè example! 123 numbers bù huì bèi biàn huàn.
	s = dict.SentenceOnlyChinese(`Hello, 我是一个 example! 123 numbers不会被变换.`).Unicode()
	fmt.Println(s)

	// Hello, wo shi yi ge example! 123 numbers bu hui bei bian huan.
	s = dict.SentenceOnlyChinese(`Hello, 我是一个 example! 123 numbers不会被变换.`).None()
	fmt.Println(s)
}
