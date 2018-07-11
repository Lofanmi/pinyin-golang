package pinyin

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

// -----------------------------------------------------------------------------

// Record 词典记录
type Record struct {
	Traditional string
	Simplified  string
	Pinyin      string
}

// Records 记录列表
type Records []*Record

// -----------------------------------------------------------------------------

// Language 简体中文/繁体中文
type Language int

const (
	// Traditional 繁体中文
	Traditional Language = iota
	// Simplified  简体中文
	Simplified
	// All         转换繁体中文和简体中文
	All
)

// -----------------------------------------------------------------------------

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
		"‘", "'", "’", "'",
		// 左/右双引号
		"“", `"`, "”", `"`,
		// 左/右直角引号
		"「", "[", "」", "]",
		"『", "[", "』", "]",
		// 左/右括号
		"（", "(", "）", ")",
		"〔", "[", "〕", "]",
		"【", "[", "】", "]",
		"{", "}", "}", "}",
		// 省略号
		"……", "...",
		// 破折号
		"——", "-",
		// 连接号
		"—", "-",
		// 左/右斜杆
		"/", "/", "\\", "\\",
		// 波浪线
		"～", "~",
		// 书名号
		"《", "<", "》", ">",
		"〈", "<", "〉", ">",
		// 间隔号
		"·", "·",
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
	// surnames 姓氏
	// https://github.com/overtrue/pinyin/blob/master/data/surnames
	surnames = []string{
		"万俟", "	mo4	qi2",
		"尉迟", "	yu4	chi2",
		"单于", "	chan2	yu2",
		"不", "	fou3",
		"沈", "	shen3",
		"称", "	cheng1",
		"车", "	che1",
		"万", "	wan4",
		"汤", "	tang1",
		"阿", "	a1",
		"丁", "	ding1",
		"强", "	qiang2",
		"仇", "	qiu2",
		"叶", "	ye4",
		"阚", "	kan4",
		"乐", "	yue4",
		"乜", "	nie4",
		"陆", "	lu4",
		"殷", "	yin1",
		"牟", "	mou2",
		"区", "	ou1",
		"宿", "	su4",
		"俞", "	yu2",
		"余", "	yu2",
		"齐", "	qi2",
		"许", "	xu3",
		"信", "	xin4",
		"无", "	wu2",
		"浣", "	wan3",
		"艾", "	ai4",
		"浅", "	qian3",
		"烟", "	yan1",
		"蓝", "	lan2",
		"於", "	yu2",
		"寻", "	xun2",
		"殳", "	shu1",
		"思", "	si1",
		"鸟", "	niao3",
		"卜", "	bu3",
		"单", "	shan4",
		"南", "	nan2",
		"柏", "	bai3",
		"朴", "	piao2",
		"繁", "	po2",
		"曾", "	zeng1",
		"瞿", "	qu2",
		"缪", "	miao4",
		"石", "	shi2",
		"冯", "	feng2",
		"覃", "	qin2",
		"幺", "	yao1",
		"种", "	chong2",
		"折", "	she4",
		"燕", "	yan1",
		"纪", "	ji3",
		"过", "	guo1",
		"华", "	hua4",
		"冼", "	xian3",
		"秘", "	bi4",
		"重", "	chong2",
		"解", "	xie4",
		"那", "	na1",
		"和", "	he2",
		"贾", "	jia3",
		"塔", "	ta3",
		"盛", "	sheng4",
		"查", "	zha1",
		"盖", "	ge3",
		"居", "	ju1",
		"哈", "	ha3",
		"的", "	de1",
		"薄", "	bo2",
		"佴", "	nai4",
		"六", "	lu4",
		"都", "	du1",
		"翟", "	zhai2",
		"扎", "	za1",
		"藏", "	zang4",
		"粘", "	nian4",
		"难", "	nan4",
		"若", "	ruo4",
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

// -----------------------------------------------------------------------------

// Dict 拼音词典
type Dict struct {
	data  Records
	tcmap map[string]*Record
	scmap map[string]*Record
}

// NewDict 新建拼音词典对象
func NewDict(dict string) (d *Dict, err error) {
	f, err := os.OpenFile(dict, os.O_RDONLY, 0755)
	if err != nil {
		return
	}

	d = &Dict{
		data:  make(Records, 0),
		tcmap: make(map[string]*Record),
		scmap: make(map[string]*Record),
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()
		if s[0] == '#' {
			continue
		}
		parsed := strings.Split(s, "`")
		r := &Record{
			Traditional: parsed[0],
			Simplified:  parsed[1],
			Pinyin:      parsed[2],
		}
		d.data = append(d.data, r)
		if _, ok := d.tcmap[r.Traditional]; !ok {
			d.tcmap[r.Traditional] = r
		}
		if _, ok := d.scmap[r.Simplified]; !ok {
			d.scmap[r.Simplified] = r
		}
	}

	return
}

// Convert 中文转换为拼音, 不保留标点符号
func (p *Dict) Convert(s string, sep string, option Language) (result *ConvertResult) {
	s = p.romanize(s, option, false)

	split := ToSlice(s)

	result = NewConvertResult(strings.Join(split, sep))
	return
}

// Sentence 中文转换为拼音, 保留标点符号
func (p *Dict) Sentence(s string, option Language) (result *ConvertResult) {
	s = p.romanize(s, option, false)

	re := regexp.MustCompile("[^a-z0-9" + regexp.QuoteMeta(strings.Join(punctuations, "")) + `\s_]+`)
	s = re.ReplaceAllString(s, "")

	for i := 0; i < len(punctuations); i += 2 {
		s = strings.Replace(s, punctuations[i], punctuations[i+1], -1)
	}

	result = NewConvertResult(s)
	return
}

// Name 转换人名
func (p *Dict) Name(s string, sep string, option Language) (result *ConvertResult) {
	s = p.romanize(s, option, true)

	split := ToSlice(s)

	result = NewConvertResult(strings.Join(split, sep))
	return
}

// Abbr 获取拼音的首字符
func (p *Dict) Abbr(s string, sep string, option Language) string {
	s = p.romanize(s, option, false)

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

func (p *Dict) romanize(s string, option Language, convertName bool) string {
	s = p.prepare(s)

	if convertName {
		for i := 0; i < len(surnames); i += 2 {
			if strings.Index(s, surnames[i]) == 0 {
				s = strings.Replace(s, surnames[i], surnames[i+1], 1)
			}
		}
	}

	if option == All || option == Traditional {
		if r, ok := p.tcmap[s]; ok {
			s = r.Pinyin
		} else {
			for _, record := range p.data {
				s = strings.Replace(s, record.Traditional, record.Pinyin, -1)
			}
		}
	}
	if option == All || option == Simplified {
		if r, ok := p.scmap[s]; ok {
			s = r.Pinyin
		} else {
			for _, record := range p.data {
				s = strings.Replace(s, record.Simplified, record.Pinyin, -1)
			}
		}
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
