package pinyin

import (
	"bufio"
	"framework/pkg/utli"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

var (
	tones = [][]rune{
		{'ā', 'ē', 'ī', 'ō', 'ū', 'ǖ', 'Ā', 'Ē', 'Ī', 'Ō', 'Ū', 'Ǖ'},
		{'á', 'é', 'í', 'ó', 'ú', 'ǘ', 'Á', 'É', 'Í', 'Ó', 'Ú', 'Ǘ'},
		{'ǎ', 'ě', 'ǐ', 'ǒ', 'ǔ', 'ǚ', 'Ǎ', 'Ě', 'Ǐ', 'Ǒ', 'Ǔ', 'Ǚ'},
		{'à', 'è', 'ì', 'ò', 'ù', 'ǜ', 'À', 'È', 'Ì', 'Ò', 'Ù', 'Ǜ'},
	}
	neutrals = []rune{'a', 'e', 'i', 'o', 'u', 'v', 'A', 'E', 'I', 'O', 'U', 'V'}
)

var (
	// 从带声调的声母到对应的英文字符的映射
	tonesMap map[rune]rune

	// 从汉字到声调的映射
	numericTonesMap map[rune]int

	// 从汉字到拼音的映射（带声调）
	pinyinMap map[rune]string

	initialized bool
)

type Mode int

const (
	WithoutTone        Mode = iota + 1 // 默认模式，例如：guo
	Tone                               // 带声调的拼音 例如：guó
	InitialsInCapitals                 // 首字母大写不带声调，例如：Guo
)

type pinyin struct {
	origin string
	split  string
	mode   Mode
}

func init() {
	tonesMap = make(map[rune]rune)
	numericTonesMap = make(map[rune]int)
	pinyinMap = make(map[rune]string)
	for i, runes := range tones {
		for j, tone := range runes {
			tonesMap[tone] = neutrals[j]
			numericTonesMap[tone] = i + 1
		}
	}

	f, err := getFileContent()
	defer f.Close()
	if err != nil {
		initialized = false
		return
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		strs := strings.Split(scanner.Text(), "=>")
		if len(strs) < 2 {
			continue
		}
		i, err := strconv.ParseInt(strs[0], 16, 32)
		if err != nil {
			continue
		}
		pinyinMap[rune(i)] = strs[1]
	}
	initialized = true
}

func getFileContent() (io.ReadCloser, error) {
	var path = ""
	var ok bool
	if path, ok = utli.CurrentFilePath(); ok {
		path = utli.Dirname(path)
	}

	//文件完整路径
	if utli.IsExist(path + "/pinyin.txt") {
		inputFile, inputError := os.Open(path + "/pinyin.txt")
		return inputFile, inputError
	} else {
		resp, err := http.Get("https://raw.githubusercontent.com/chain-zhang/pinyin/master/pinyin.txt")
		f, _ := os.OpenFile(path+"/pinyin.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
		}
		defer f.Close()
		io.Copy(f, resp.Body)
		return resp.Body, err
	}
}

func Char2Pinyin(origin string) *pinyin {
	return &pinyin{
		origin: origin,
		split:  " ",
		mode:   WithoutTone,
	}
}

func (py *pinyin) Split(split string) *pinyin {
	py.split = split
	return py
}

func (py *pinyin) Mode(mode Mode) *pinyin {
	py.mode = mode
	return py
}

var isWordReg = regexp.MustCompile(`([a-zA-Z0-9]+)`)

func (py *pinyin) Convert() (string, error) {
	if !initialized {
		return "", ErrInitialize
	}

	//py.origin = strings.ReplaceAll(py.origin, " ", py.split)
	sr := []rune(py.origin)
	words := make([]string, 0)
	for _, s := range sr {
		word, err := getPinyin(s, py.mode)
		if err != nil {
			return "", err
		}

		if len(word) > 0 {
			words = append(words, word)
		}
	}

	matches := isWordReg.FindAllString(py.origin, -1)
	words = append(words, matches...)
	return strings.Join(words, py.split), nil
}

func getPinyin(hanzi rune, mode Mode) (string, error) {
	if !initialized {
		return "", ErrInitialize
	}
	switch mode {
	case Tone:
		return getTone(hanzi), nil
	case InitialsInCapitals:
		return getInitialsInCapitals(hanzi), nil
	default:
		return getDefault(hanzi), nil
	}
}

func getTone(hanzi rune) string {
	return pinyinMap[hanzi]
}

func getDefault(hanzi rune) string {
	tone := getTone(hanzi)

	if tone == "" {
		return tone
	}

	output := make([]rune, utf8.RuneCountInString(tone))

	count := 0
	for _, t := range tone {
		neutral, found := tonesMap[t]
		if found {
			output[count] = neutral
		} else {
			output[count] = t
		}
		count++
	}
	return string(output)
}

func getInitialsInCapitals(hanzi rune) string {
	def := getDefault(hanzi)
	if def == "" {
		return def
	}
	sr := []rune(def)
	if sr[0] > 32 {
		sr[0] = sr[0] - 32
	}
	return string(sr)
}
