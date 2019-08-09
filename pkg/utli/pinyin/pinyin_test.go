package pinyin

import (
	//"framework/pkg/utli"
	"testing"
)

func TestConvert(t *testing.T) {
	str, err := Char2Pinyin("我是中国人").Split("").Mode(InitialsInCapitals).Convert()
	if err != nil {
		t.Error(err)
	} else {
		t.Log(str)
	}

	str, err = Char2Pinyin("我是中国人").Split(" ").Mode(WithoutTone).Convert()
	if err != nil {
		t.Error(err)
	} else {
		t.Log(str)
	}

	str, err = Char2Pinyin("我是中国人").Split("-").Mode(Tone).Convert()
	if err != nil {
		t.Error(err)
	} else {
		t.Log(str)
	}

	str, err = Char2Pinyin("我是中国人").Convert()
	if err != nil {
		t.Error(err)
	} else {
		t.Log(str)
	}
}
