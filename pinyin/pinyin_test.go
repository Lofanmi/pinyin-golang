package pinyin

import (
	"testing"
)

var (
	filename = "cedict.lofanmi"
)

func getTestDict(t *testing.T) *Dict {
	dict, err := NewDict(filename)
	if err != nil {
		t.Errorf("getTestDict() error = %v", err)
	}
	return dict
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
	type args struct {
		dict string
	}
	tests := []struct {
		name    string
		args    args
		wantD   *Dict
		wantErr bool
	}{
		{"invalid_file_name", args{""}, nil, true},
		{"dict", args{filename}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotD, err := NewDict(tt.args.dict)
			if tt.name == "invalid_file_name" && err != nil && !tt.wantErr {
				t.Errorf("NewDict() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.name == "dict" && gotD == nil {
				t.Errorf("NewDict() = %v, want not nil. %v", gotD, err)
			}
		})
	}
}

func TestDict_Convert(t *testing.T) {
	dict := getTestDict(t)
	type args struct {
		s      string
		sep    string
		option Language
	}
	tests := []struct {
		name       string
		p          *Dict
		args       args
		wantResult *ConvertResult
	}{
		{"test_1", dict, args{`我，何时能暴富？`, " ", All}, getTestConvertResult("wo3 he2 shi2 neng2 bao4 fu4")},
		{"test_2", dict, args{`嗯 en 好的`, " ", All}, getTestConvertResult("en4 en hao3 de")},
		{"test_3", dict, args{`马`, " ", All}, getTestConvertResult("ma3")},
		{"test_4", dict, args{`馬`, " ", All}, getTestConvertResult("ma3")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := tt.p.Convert(tt.args.s, tt.args.sep, tt.args.option); !isEqual(gotResult, tt.wantResult) {
				t.Errorf("Dict.Convert() = %v, want %v", gotResult.ASCII(), tt.wantResult.ASCII())
			}
		})
	}
}

func TestDict_Sentence(t *testing.T) {
	dict := getTestDict(t)
	type args struct {
		s      string
		option Language
	}
	tests := []struct {
		name       string
		p          *Dict
		args       args
		wantResult *ConvertResult
	}{
		{"test", dict, args{`我，何时能暴富？`, All}, getTestConvertResult("wo3, he2 shi2 neng2 bao4 fu4?")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := tt.p.Sentence(tt.args.s, tt.args.option); !isEqual(gotResult, tt.wantResult) {
				t.Errorf("Dict.Sentence() = %v, want %v", gotResult.ASCII(), tt.wantResult.ASCII())
			}
		})
	}
}

func TestDict_Name(t *testing.T) {
	dict := getTestDict(t)
	type args struct {
		s      string
		sep    string
		option Language
	}
	tests := []struct {
		name       string
		p          *Dict
		args       args
		wantResult *ConvertResult
	}{
		{"test", dict, args{`万俟沃喜欢吃酸奶`, " ", All}, getTestConvertResult("mo4 qi2 wo4 xi3 huan1 chi1 suan1 nai3")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := tt.p.Name(tt.args.s, tt.args.sep, tt.args.option); !isEqual(gotResult, tt.wantResult) {
				t.Errorf("Dict.Name() = %v, want %v", gotResult.ASCII(), tt.wantResult.ASCII())
			}
		})
	}
}

func TestDict_Abbr(t *testing.T) {
	dict := getTestDict(t)
	type args struct {
		s      string
		sep    string
		option Language
	}
	tests := []struct {
		name       string
		p          *Dict
		args       args
		wantResult string
	}{
		{"test", dict, args{`万俟沃喜欢吃酸奶`, "-", All}, "m-q-w-x-h-c-s-n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := tt.p.Abbr(tt.args.s, tt.args.sep, tt.args.option); gotResult != tt.wantResult {
				t.Errorf("Dict.Abbr() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
