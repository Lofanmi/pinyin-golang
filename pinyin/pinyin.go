package pinyin

import (
	"regexp"
	"strings"
)

var (
	// punctuations 标点符号
	punctuations = []string{
		// 逗号
		"，", ",",
		// 句号
		"。", ".",
		// 感叹号
		"！", "!",
		// 问号
		"？", "?",
		// 冒号
		"：", ":",
		// 分号
		"；", ";",
		// 左/右单引号
		"‘", " '", "’", " '",
		// 左/右双引号
		"“", ` "`, "”", ` "`,
		// 左/右直角引号
		"「", " [", "」", " ]",
		"『", " [", "』", " ]",
		// 左/右括号
		"（", " (", "）", " )",
		"〔", " [", "〕", " ]",
		"【", " [", "】", " ]",
		"{", " {", "}", " }",
		// 省略号
		"……", "...",
		// 破折号
		"——", "-",
		// 连接号
		"—", "-",
		// 左/右斜杆
		"/", " /", "\\", " \\",
		// 波浪线
		"～", "~",
		// 书名号
		"《", " <", "》", " >",
		"〈", " <", "〉", " >",
		// 间隔号
		"·", " ·",
		// 顿号
		"、", ",",
	}
	// finals 韵母表
	finals = []string{
		// a
		"a1", "ā", "a2", "á", "a3", "ǎ", "a4", "à",
		// o
		"o1", "ō", "o2", "ó", "o3", "ǒ", "o4", "ò",
		// e
		"e1", "ē", "e2", "é", "e3", "ě", "e4", "è",
		// i
		"i1", "ī", "i2", "í", "i3", "ǐ", "i4", "ì",
		// u
		"u1", "ū", "u2", "ú", "u3", "ǔ", "u4", "ù",
		// v
		"v1", "ǖ", "v2", "ǘ", "v3", "ǚ", "v4", "ǜ",

		// ai
		"ai1", "āi", "ai2", "ái", "ai3", "ǎi", "ai4", "ài",
		// ei
		"ei1", "ēi", "ei2", "éi", "ei3", "ěi", "ei4", "èi",
		// ui
		"ui1", "uī", "ui2", "uí", "ui3", "uǐ", "ui4", "uì",
		// ao
		"ao1", "āo", "ao2", "áo", "ao3", "ǎo", "ao4", "ào",
		// ou
		"ou1", "ōu", "ou2", "óu", "ou3", "ǒu", "ou4", "òu",
		// iu
		"iu1", "īu", "iu2", "íu", "iu3", "ǐu", "iu4", "ìu",

		// ie
		"ie1", "iē", "ie2", "ié", "ie3", "iě", "ie4", "iè",
		// ve
		"ue1", "üē", "ue2", "üé", "ue3", "üě", "ue4", "üè",
		// er
		"er1", "ēr", "er2", "ér", "er3", "ěr", "er4", "èr",

		// an
		"an1", "ān", "an2", "án", "an3", "ǎn", "an4", "àn",
		// en
		"en1", "ēn", "en2", "én", "en3", "ěn", "en4", "èn",
		// in
		"in1", "īn", "in2", "ín", "in3", "ǐn", "in4", "ìn",
		// un/vn
		"un1", "ūn", "un2", "ún", "un3", "ǔn", "un4", "ùn",

		// ang
		"ang1", "āng", "ang2", "áng", "ang3", "ǎng", "ang4", "àng",
		// eng
		"eng1", "ēng", "eng2", "éng", "eng3", "ěng", "eng4", "èng",
		// ing
		"ing1", "īng", "ing2", "íng", "ing3", "ǐng", "ing4", "ìng",
		// ong
		"ong1", "ōng", "ong2", "óng", "ong3", "ǒng", "ong4", "òng",
	}
)

// -----------------------------------------------------------------------------

// ConvertResult 转换结果
type ConvertResult string

// NewConvertResult 创建转换结果对象
func NewConvertResult(s string) *ConvertResult {
	cr := ConvertResult(s)
	return &cr
}

// ASCII 带数字的声调
// mei3 hao3
func (r *ConvertResult) ASCII() string {
	return string(*r)
}

// Unicode Unicode声调
// měi hǎo
func (r *ConvertResult) Unicode() string {
	s := string(*r)
	for i := len(finals) - 1; i >= 0; i -= 2 {
		s = strings.Replace(s, finals[i-1], finals[i], -1)
	}
	return s
}

// None 不带声调输出
// mei hao
func (r *ConvertResult) None() string {
	s := string(*r)
	re := regexp.MustCompile(`[1-4]{1}`)
	s = re.ReplaceAllString(s, "")
	return s
}

// String 与 ASCII 相同
func (r *ConvertResult) String() string { return r.ASCII() }

// -----------------------------------------------------------------------------

// Dict 拼音词典
type Dict struct{}

// NewDict 新建拼音词典对象
func NewDict() *Dict {
	return new(Dict)
}

// Convert 中文转换为拼音, 不保留标点符号
func (p *Dict) Convert(s string, sep string) (result *ConvertResult) {
	s = p.romanize(s, false)

	split := ToSlice(s)

	result = NewConvertResult(strings.Join(split, sep))
	return
}

// Sentence 中文转换为拼音, 保留标点符号
func (p *Dict) Sentence(s string) (result *ConvertResult) {
	s = p.romanize(s, false)

	r := regexp.QuoteMeta(strings.Join(punctuations, ""))
	r = strings.Replace(r, " ", "", -1)
	re := regexp.MustCompile("[^a-zA-Z0-9" + r + `\s_]+`)
	s = re.ReplaceAllString(s, "")

	for i := 0; i < len(punctuations); i += 2 {
		s = strings.Replace(s, punctuations[i], punctuations[i+1], -1)
	}

	result = NewConvertResult(s)
	return
}

// Name 转换人名
func (p *Dict) Name(s string, sep string) (result *ConvertResult) {
	s = p.romanize(s, true)

	split := ToSlice(s)

	result = NewConvertResult(strings.Join(split, sep))
	return
}

// Abbr 获取拼音的首字符
func (p *Dict) Abbr(s string, sep string) string {
	s = p.romanize(s, false)

	var abbr []string
	for _, item := range ToSlice(s) {
		abbr = append(abbr, item[0:1])
	}

	return strings.Join(abbr, sep)
}

func (p *Dict) prepare(s string) string {
	var re *regexp.Regexp

	re = regexp.MustCompile(`[a-zA-Z0-9_-]+`)
	s = re.ReplaceAllStringFunc(s, func(repl string) string {
		return "\t" + repl
	})

	re = regexp.MustCompile(`[^\p{Han}\p{P}\p{Z}\p{M}\p{N}\p{L}\t]`)
	s = re.ReplaceAllString(s, "")

	return s
}

func (p *Dict) romanize(s string, convertName bool) string {
	s = p.prepare(s)

	if convertName {
		for i := 0; i < len(surnames); i += 2 {
			if strings.Index(s, surnames[i]) == 0 {
				s = strings.Replace(s, surnames[i], surnames[i+1], 1)
			}
		}
	}

	for i := 0; i < len(dict); i += 2 {
		s = strings.Replace(s, dict[i], dict[i+1], -1)
	}

	s = strings.Replace(s, "\t", " ", -1)
	s = strings.Replace(s, "  ", " ", -1)
	s = strings.TrimSpace(s)

	return s
}

// ToSlice 转换为字符串数组
func ToSlice(s string) []string {
	var split []string
	re := regexp.MustCompile(`[^a-zA-Z1-4]+`)
	for _, str := range re.Split(s, -1) {
		if str != "" {
			split = append(split, str)
		}
	}
	return split
}
