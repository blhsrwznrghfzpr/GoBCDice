package testcase

import (
	"fmt"
	"github.com/raa0121/GoBCDice/pkg/core/dice"
	"reflect"
	"testing"
)

var parseTestCases = []struct {
	// テストケースのソース
	source string
	// ゲーム識別子
	gameId string
	// テストケース番号
	index int
	// 期待する値
	expected DiceBotTestCase
	// エラーを期待するか
	err bool
}{
	{
		source: "",
		err:    true,
	},
	{
		source: "input:",
		err:    true,
	},
	{
		source: `input:
2d6+1-1-2-3-4
output:
DiceBot : (2D6+1-1-2-3-4) ＞ 5[4,1]+1-1-2-3-4 ＞ -4
rand:4/6,1/6`,
		gameId: "DiceBot",
		index:  1,
		expected: DiceBotTestCase{
			GameId: "DiceBot",
			Index:  1,
			Input:  []string{"2d6+1-1-2-3-4"},
			Output: "DiceBot : (2D6+1-1-2-3-4) ＞ 5[4,1]+1-1-2-3-4 ＞ -4",
			Dice:   []dice.Die{{4, 6}, {1, 6}},
		},
		err: false,
	},
	{
		source: `input:
S2d6
output:
DiceBot : (2D6) ＞ 5[4,1] ＞ 5###secret dice###
rand:4/6,1/6`,
		gameId: "DiceBot",
		index:  2,
		expected: DiceBotTestCase{
			GameId: "DiceBot",
			Index:  2,
			Input:  []string{"S2d6"},
			Output: "DiceBot : (2D6) ＞ 5[4,1] ＞ 5###secret dice###",
			Dice:   []dice.Die{{4, 6}, {1, 6}},
		},
		err: false,
	},
	{
		source: `input:
GETSST
output:
Satasupe : サタスペ作成：ベース部品：「大型の金属製の筒」  アクセサリ部品：「ガスボンベや殺虫剤」
部品効果：「命中：8、ダメージ：5、耐久度3、両手」「爆発3」
完成品：サタスペ  （ダメージ＋5・命中8・射撃、「両手」「爆発3」「サタスペ1」「耐久度3」）
rand:6/6,6/6,6/6`,
		gameId: "Satasupe",
		index:  1,
		expected: DiceBotTestCase{
			GameId: "Satasupe",
			Index:  1,
			Input:  []string{"GETSST"},
			Output: `Satasupe : サタスペ作成：ベース部品：「大型の金属製の筒」  アクセサリ部品：「ガスボンベや殺虫剤」
部品効果：「命中：8、ダメージ：5、耐久度3、両手」「爆発3」
完成品：サタスペ  （ダメージ＋5・命中8・射撃、「両手」「爆発3」「サタスペ1」「耐久度3」）`,
			Dice: []dice.Die{{6, 6}, {6, 6}, {6, 6}},
		},
		err: false,
	},
	{
		source: `input:
CCT
output:
GranCrest : 国特徴・文化表(13) ＞ 禁欲的
あなたの国民は、道徳を重んじ、常に自分の欲望を制限することが理想的だと考えている。
食料＋４、資金－１
rand:1/6,3/6`,
		gameId: "GranCrest",
		index:  1,
		expected: DiceBotTestCase{
			GameId: "GranCrest",
			Index:  1,
			Input:  []string{"CCT"},
			Output: `GranCrest : 国特徴・文化表(13) ＞ 禁欲的
あなたの国民は、道徳を重んじ、常に自分の欲望を制限することが理想的だと考えている。
食料＋４、資金－１`,
			Dice: []dice.Die{{1, 6}, {3, 6}},
		},
		err: false,
	},
}

func TestParse(t *testing.T) {
	for _, test := range parseTestCases {
		t.Run(fmt.Sprintf("%s-%d", test.gameId, test.index), func(t *testing.T) {
			actual, err := Parse(test.source, test.gameId, test.index)
			if err != nil {
				if !test.err {
					t.Fatalf("got err: %v", err)
				}

				return
			}

			if test.err {
				t.Fatal("should err")
				return
			}

			if !reflect.DeepEqual(*actual, test.expected) {
				t.Errorf("got: %+v, want: %+v", *actual, test.expected)
			}
		})
	}
}

var parseDiceTestCases = []struct {
	source   string
	expected []dice.Die
	err      bool
}{
	{
		source: "1",
		err:    true,
	},
	{
		source: "1/6,1",
		err:    true,
	},
	{
		source: "a1/6",
		err:    true,
	},
	{
		source:   "",
		expected: []dice.Die{},
		err:      false,
	},
	{
		source:   "1/6",
		expected: []dice.Die{{1, 6}},
		err:      false,
	},
	{
		source:   "1/6,2/6,3/6",
		expected: []dice.Die{{1, 6}, {2, 6}, {3, 6}},
		err:      false,
	},
	{
		source:   "1/6, 2/6, 3/6",
		expected: []dice.Die{{1, 6}, {2, 6}, {3, 6}},
		err:      false,
	},
}

func TestParseDice(t *testing.T) {
	for _, test := range parseDiceTestCases {
		t.Run(fmt.Sprintf("%q", test.source), func(t *testing.T) {
			actual, err := ParseDice(test.source)

			if err != nil {
				if !test.err {
					t.Fatalf("got err: %v", err)
				}
				return
			}

			if test.err {
				t.Fatal("should err")
				return
			}

			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("got: %+v, want: %+v", actual, test.expected)
			}
		})
	}
}

var parseFileTestCases = []struct {
	filename string
	expected []DiceBotTestCase
}{
	{
		filename: "testdata/DiceBot.txt",
		expected: []DiceBotTestCase{
			{
				GameId: "DiceBot",
				Index:  1,
				Input:  []string{"2d6+1-1-2-3-4"},
				Output: "DiceBot : (2D6+1-1-2-3-4) ＞ 5[4,1]+1-1-2-3-4 ＞ -4",
				Dice:   []dice.Die{{4, 6}, {1, 6}},
			},
			{
				GameId: "DiceBot",
				Index:  2,
				Input:  []string{"S2d6"},
				Output: "DiceBot : (2D6) ＞ 5[4,1] ＞ 5###secret dice###",
				Dice:   []dice.Die{{4, 6}, {1, 6}},
			},
			{
				GameId: "DiceBot",
				Index:  3,
				Input:  []string{"4d10"},
				Output: "4d10 : (4D10) ＞ 18[3,2,5,8] ＞ 18",
				Dice:   []dice.Die{{3, 10}, {2, 10}, {5, 10}, {8, 10}},
			},
			{
				GameId: "DiceBot",
				Index:  4,
				Input:  []string{"2R6"},
				Output: "DiceBot : 2R6 ＞ 条件が間違っています。2R6>=5 あるいは 2R6[5] のように振り足し目標値を指定してください。",
				Dice:   []dice.Die{},
			},
		},
	},
	{
		filename: "testdata/multiline.txt",
		expected: []DiceBotTestCase{
			{
				GameId: "multiline",
				Index:  1,
				Input:  []string{"GETSST"},
				Output: `Satasupe : サタスペ作成：ベース部品：「大型の金属製の筒」  アクセサリ部品：「ガスボンベや殺虫剤」
部品効果：「命中：8、ダメージ：5、耐久度3、両手」「爆発3」
完成品：サタスペ  （ダメージ＋5・命中8・射撃、「両手」「爆発3」「サタスペ1」「耐久度3」）`,
				Dice: []dice.Die{{6, 6}, {6, 6}, {6, 6}},
			},
			{
				GameId: "multiline",
				Index:  2,
				Input:  []string{"CCT"},
				Output: `GranCrest : 国特徴・文化表(13) ＞ 禁欲的
あなたの国民は、道徳を重んじ、常に自分の欲望を制限することが理想的だと考えている。
食料＋４、資金－１`,
				Dice: []dice.Die{{1, 6}, {3, 6}},
			},
		},
	},
}

func TestParseFile(t *testing.T) {
	for _, test := range parseFileTestCases {
		t.Run(fmt.Sprintf("%q", test.filename), func(t *testing.T) {
			loadedTestCases, err := ParseFile(test.filename)

			if err != nil {
				t.Fatalf("got err: %v", err)
				return
			}

			for j, expected := range test.expected {
				t.Run(fmt.Sprintf("%s-%d", expected.GameId, expected.Index), func(t *testing.T) {
					if len(loadedTestCases) <= j {
						t.Fatal("読み込まれたテストケースが不足しています")
						return
					}

					actual := *loadedTestCases[j]

					if !reflect.DeepEqual(actual, expected) {
						t.Errorf("got: %+v, want: %+v", actual, expected)
					}
				})
			}
		})
	}
}
