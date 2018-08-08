package pinyin

import (
	"testing"
)

var (
	filename = "cedict.lofanmi"
)

func getTestDict(t *testing.T) *Dict {
	return NewDict()
}

func getTestConvertResult(s string) *ConvertResult {
	return NewConvertResult(s)
}

func isEqual(cr1 *ConvertResult, cr2 *ConvertResult) bool {
	return cr1.ASCII() == cr2.ASCII() &&
		cr1.Unicode() == cr2.Unicode() &&
		cr1.None() == cr2.None()
}

func TestNewDict(t *testing.T) {
	t.Run("dict", func(t *testing.T) {
		if got := NewDict(); got == nil {
			t.Errorf("NewDict() = %v, want not nil", got)
		}
	})
}

func TestDict_Convert(t *testing.T) {
	dict := getTestDict(t)
	type args struct {
		s   string
		sep string
	}
	tests := []struct {
		name       string
		p          *Dict
		args       args
		wantResult *ConvertResult
	}{
		{"test_1", dict, args{`我，何时能暴富？`, " "}, getTestConvertResult("wo3 he2 shi2 neng2 bao4 fu4")},
		{"test_2", dict, args{`嗯 en 好的`, " "}, getTestConvertResult("en4 en hao3 de")},
		{"test_3", dict, args{`马`, " "}, getTestConvertResult("ma3")},
		{"test_4", dict, args{`馬`, " "}, getTestConvertResult("ma3")},
		// 康熙来了
		{"overtrue.me", dict, args{`康熙来了`, " "}, getTestConvertResult("kang1 xi1 lai2 le")},
		// 带着希望去旅行，比到达终点更美好
		{"overtrue.me", dict, args{`带着希望去旅行，比到达终点更美好`, " "}, getTestConvertResult("dai4 zhe xi1 wang4 qu4 lv3 xing2 bi3 dao4 da2 zhong1 dian3 geng4 mei3 hao3")},
		// 多音字
		// 了
		{"overtrue.me", dict, args{`了然`, " "}, getTestConvertResult("liao3 ran2")},
		{"overtrue.me", dict, args{`来了`, " "}, getTestConvertResult("lai2 le")},
		// 还
		{"overtrue.me", dict, args{`还有`, " "}, getTestConvertResult("hai2 you3")},
		{"overtrue.me", dict, args{`交还`, " "}, getTestConvertResult("jiao1 huan2")},
		// 什
		{"overtrue.me", dict, args{`什么`, " "}, getTestConvertResult("shen2 me")},
		{"overtrue.me", dict, args{`什锦`, " "}, getTestConvertResult("shi2 jin3")},
		// 便
		{"overtrue.me", dict, args{`便当`, " "}, getTestConvertResult("bian4 dang1")},
		{"overtrue.me", dict, args{`便宜`, " "}, getTestConvertResult("pian2 yi2")},
		// 剥
		{"overtrue.me", dict, args{`剥皮`, " "}, getTestConvertResult("bao1 pi2")},
		{"overtrue.me", dict, args{`剥皮器`, " "}, getTestConvertResult("bao1 pi2 qi4")},
		// 不
		{"overtrue.me", dict, args{`赔不是`, " "}, getTestConvertResult("pei2 bu2 shi4")},
		{"overtrue.me", dict, args{`跑了和尚，跑不了庙`, " "}, getTestConvertResult("pao3 le he2 shang4 pao3 bu4 liao3 miao4")},
		// 降
		{"overtrue.me", dict, args{`降温`, " "}, getTestConvertResult("jiang4 wen1")},
		{"overtrue.me", dict, args{`投降`, " "}, getTestConvertResult("tou2 xiang2")},
		// 都
		{"overtrue.me", dict, args{`首都`, " "}, getTestConvertResult("shou3 du1")},
		{"overtrue.me", dict, args{`都什么年代了`, " "}, getTestConvertResult("dou1 shen2 me nian2 dai4 le")},
		// 乐
		{"overtrue.me", dict, args{`快乐`, " "}, getTestConvertResult("kuai4 le4")},
		{"overtrue.me", dict, args{`音乐`, " "}, getTestConvertResult("yin1 yue4")},
		// 长
		{"overtrue.me", dict, args{`成长`, " "}, getTestConvertResult("cheng2 zhang3")},
		{"overtrue.me", dict, args{`长江`, " "}, getTestConvertResult("chang2 jiang1")},
		// 难
		{"overtrue.me", dict, args{`难民`, " "}, getTestConvertResult("nan4 min2")},
		{"overtrue.me", dict, args{`难过`, " "}, getTestConvertResult("nan2 guo4")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := tt.p.Convert(tt.args.s, tt.args.sep); !isEqual(gotResult, tt.wantResult) {
				t.Errorf("Dict.Convert() = %v, want %v", gotResult.ASCII(), tt.wantResult.ASCII())
			}
		})
	}
}

func TestDict_Sentence(t *testing.T) {
	dict := getTestDict(t)
	type args struct {
		s string
	}
	tests := []struct {
		name       string
		p          *Dict
		args       args
		wantResult *ConvertResult
	}{
		{"test", dict, args{`我，何时能暴富？`}, getTestConvertResult("wo3, he2 shi2 neng2 bao4 fu4?")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := tt.p.Sentence(tt.args.s); !isEqual(gotResult, tt.wantResult) {
				t.Errorf("Dict.Sentence() = %v, want %v", gotResult.ASCII(), tt.wantResult.ASCII())
			}
		})
	}
}

func TestDict_Name(t *testing.T) {
	dict := getTestDict(t)
	type args struct {
		s   string
		sep string
	}
	tests := []struct {
		name       string
		p          *Dict
		args       args
		wantResult *ConvertResult
	}{
		{"test", dict, args{`万俟沃喜欢吃酸奶`, " "}, getTestConvertResult("mo4 qi2 wo4 xi3 huan1 chi1 suan1 nai3")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := tt.p.Name(tt.args.s, tt.args.sep); !isEqual(gotResult, tt.wantResult) {
				t.Errorf("Dict.Name() = %v, want %v", gotResult.ASCII(), tt.wantResult.ASCII())
			}
		})
	}
}

func TestDict_Abbr(t *testing.T) {
	dict := getTestDict(t)
	type args struct {
		s   string
		sep string
	}
	tests := []struct {
		name       string
		p          *Dict
		args       args
		wantResult string
	}{
		{"test", dict, args{`万俟沃喜欢吃酸奶`, "-"}, "m-q-w-x-h-c-s-n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := tt.p.Abbr(tt.args.s, tt.args.sep); gotResult != tt.wantResult {
				t.Errorf("Dict.Abbr() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
