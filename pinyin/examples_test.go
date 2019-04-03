package pinyin

import "fmt"

func Example_simple() {
	var dict = NewDict()
	// 简单用法示例
	fmt.Println(dict.Sentence(`Redis 是一个 Key-Value 存储系统。`).Unicode())
	fmt.Println(dict.Convert(`我，何時能暴富？`, " ").Unicode())
	fmt.Println(dict.Sentence(`我，何時能暴富？`).Unicode())

	// Output:
	// Redis shì yí gè Key-Value cún chǔ xì tǒng.
	// wǒ hé shí néng bào fù
	// wǒ, hé shí néng bào fù?
}

func Example_convert() {
	var dict = NewDict()
	//转换接口示例

	// 输入繁体中文, 输出为带 空格 分隔的拼音字符串
	// ASCII 格式显示
	fmt.Println(dict.Convert(`我，何時能暴富？`, " ")) //or ⬇️
	fmt.Println(dict.Convert(`我，何時能暴富？`, " ").ASCII())

	// 输入简体中文, 输出为带 连字符- 分隔的拼音字符串
	// Unicode 格式显示
	fmt.Println(dict.Convert(`我，何时能暴富？`, "-").Unicode())

	// 转换简体中文和繁体中文, 转换为带 斜杆/ 分隔的拼音字符串
	// 不显示声调
	fmt.Println(dict.Convert(`我，何时能暴富？`, "/").None())

	// Output:
	// wo3 he2 shi2 neng2 bao4 fu4
	// wo3 he2 shi2 neng2 bao4 fu4
	// wǒ-hé-shí-néng-bào-fù
	// wo/he/shi/neng/bao/fu
}

func Example_sentence() {
	dict := NewDict()
	//句子接口示例

	// 输入繁体中文, 输出为带 空格 分隔的拼音字符串
	// ASCII 格式显示
	fmt.Println(dict.Sentence(`我，何時能暴富？`))

	// 输入简体中文, 输出为带 空格 分隔的拼音字符串
	// Unicode 格式显示
	fmt.Println(dict.Sentence(`我，何时能暴富？`).Unicode())

	// 转换简体中文和繁体中文, 转换为带 空格 分隔的拼音字符串
	// 不显示声调
	fmt.Println(dict.Sentence(`我，何时能暴富？`).None())

	// Output:
	// wo3, he2 shi2 neng2 bao4 fu4?
	// wǒ, hé shí néng bào fù?
	// wo, he shi neng bao fu?
}

func Example_name() {
	dict := NewDict()
	//转换人名接口

	// 输入繁体中文, 输出为带 空格 分隔的人名拼音字符串
	// ASCII 格式显示
	fmt.Println(dict.Name(`万俟沃喜欢吃酸奶`, " "))

	// 输入简体中文, 输出为带 连字符- 分隔的人名拼音字符串
	// Unicode 格式显示
	fmt.Println(dict.Name(`万俟沃喜欢吃酸奶`, "-").Unicode())

	// 转换简体中文和繁体中文, 转换为带 斜杆/ 分隔的人名拼音字符串
	// 不显示声调
	fmt.Println(dict.Name(`万俟沃喜欢吃酸奶`, "/").None())

	// Output:
	// mo4 qi2 wo4 xi3 huan1 chi1 suan1 nai3
	// mò-qí-wò-xǐ-huān-chī-suān-nǎi
	// mo/qi/wo/xi/huan/chi/suan/nai
}

func Example_abbr1() {
	dict := NewDict()
	// 转换拼音简写
	// 转换简体中文和繁体中文, 输出为带 连字符- 分隔的拼音字符串首字符
	fmt.Println(dict.Abbr(`万俟沃喜欢吃酸奶`, "-"))

	// Output: m-q-w-x-h-c-s-n
}

func Example_abbr2() {
	dict := NewDict()
	// 转换为字符串
	fmt.Println(dict.Convert(`我，何時能暴富？`, " "))
	fmt.Printf("%v", dict.Convert(`我，何時能暴富？`, " "))

	// Output:
	// wo3 he2 shi2 neng2 bao4 fu4
	// wo3 he2 shi2 neng2 bao4 fu4
}
