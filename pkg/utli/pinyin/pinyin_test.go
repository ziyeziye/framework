package pinyin

import (
	"framework/pkg/utli"
	"testing"
)

func TestConvert(t *testing.T) {
	str, err := utli.Char2Pinyin("我是中国人").Split("").Mode(utli.InitialsInCapitals).Convert()
	if err != nil {
		t.Error(err)
	} else {
		t.Log(str)
	}

	str, err = utli.Char2Pinyin("我是中国人").Split(" ").Mode(utli.WithoutTone).Convert()
	if err != nil {
		t.Error(err)
	} else {
		t.Log(str)
	}

	str, err = utli.Char2Pinyin("我是中国人").Split("-").Mode(utli.Tone).Convert()
	if err != nil {
		t.Error(err)
	} else {
		t.Log(str)
	}

	str, err = utli.Char2Pinyin("我是中国人").Convert()
	if err != nil {
		t.Error(err)
	} else {
		t.Log(str)
	}
}
