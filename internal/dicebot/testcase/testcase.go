package testcase

import (
	"fmt"
	"github.com/raa0121/GoBCDice/internal/die"
	"io/ioutil"
	"path"
	"regexp"
	"strconv"
	"strings"
)

// ダイスボットのテストケースを表す構造体
type DiceBotTestCase struct {
	// ゲーム識別子
	gameId string
	// テストケース番号
	index int
	// 入力文字列
	input []string
	// 出力文字列
	output string
	// 入力するダイス列
	dice []die.Die
}

var (
	// テストケースのソースコードを表す正規表現
	sourceRe = regexp.MustCompile("(?s)\\Ainput:\n(.+)\noutput:(.*)\nrand:(.*)")
	// テストケースのソースコード内のダイス表記を表す正規表現
	diceRe = regexp.MustCompile(`\A\s*(\d+)/(\d+)\s*\z`)
)

// Parseは、テストケースのソースコードを構文解析してDiceBotTestCaseを作り、
// それを指すポインタを返す。失敗するとnilを返す。
//
// gameIdにはゲームの識別子を指定する。
// indexにはテストケース番号を指定する。
func Parse(source string, gameId string, index int) (*DiceBotTestCase, error) {
	matches := sourceRe.FindStringSubmatch(source)
	if matches == nil {
		return nil, fmt.Errorf("Parse: %s#%d: テストケース構文エラー", gameId, index)
	}

	dice, err := ParseDice(matches[3])
	if err != nil {
		return nil, err
	}

	input := strings.Split(matches[1], "\n")
	output := strings.TrimLeft(matches[2], "\n")

	return &DiceBotTestCase{
		gameId: gameId,
		index:  index,
		input:  input,
		output: output,
		dice:   dice,
	}, nil
}

// ParseDiceはテストケースのソースコード内のダイス表記を解析する
// 振られたサイコロの情報の配列を返す
func ParseDice(source string) ([]die.Die, error) {
	dice := []die.Die{}

	if source == "" {
		return dice, nil
	}

	diceStrs := strings.Split(source, ",")
	for i, diceStr := range diceStrs {
		matches := diceRe.FindStringSubmatch(diceStr)
		if matches == nil {
			return nil, fmt.Errorf("ParseDice: #%d: %q: ダイス構文エラー", i+1, diceStr)
		}

		Value, _ := strconv.Atoi(matches[1])
		Sides, _ := strconv.Atoi(matches[2])
		dice = append(dice, die.Die{Value, Sides})
	}

	return dice, nil
}

// ParseFileはテストケースのソースコードファイルを解析し、
// DiceBotTestCaseのポインタの配列を返す。
//
// filenameには、ソースコードファイルのパスを指定する。
func ParseFile(filename string) ([]*DiceBotTestCase, error) {
	contentBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	basename := path.Base(filename)
	gameId := strings.TrimSuffix(basename, path.Ext(basename))

	content := strings.TrimRight(string(contentBytes), "\n")
	testCaseSources := strings.Split(content, "\n============================\n")
	testCases := []*DiceBotTestCase{}

	for i, source := range testCaseSources {
		index := i + 1

		testCase, err := Parse(source, gameId, index)
		if err != nil {
			return nil, err
		}

		testCases = append(testCases, testCase)
	}

	return testCases, nil
}
