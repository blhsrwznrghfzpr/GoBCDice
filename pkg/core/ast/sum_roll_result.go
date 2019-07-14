package ast

import (
	"github.com/raa0121/GoBCDice/pkg/core/die"
	"strings"
)

// 加算ロールの結果を表すノード。
// 一次式。
type SumRollResult struct {
	NodeImpl

	// 振られたダイスの配列
	Dice []die.Die
}

// SumRollResult がNodeを実装していることの確認。
var _ Node = (*SumRollResult)(nil)

// Type はノードの種類を返す。
func (n *SumRollResult) Type() NodeType {
	return SUM_ROLL_RESULT_NODE
}

// Value は出目の合計を返す。
func (n *SumRollResult) Value() int {
	sum := 0

	for _, d := range n.Dice {
		sum += d.Value
	}

	return sum
}

// SExp はノードのS式を返す。
func (n *SumRollResult) SExp() string {
	diceStrs := []string{}

	for _, d := range n.Dice {
		diceStrs = append(diceStrs, d.SExp())
	}

	return "(SumRollResult " + strings.Join(diceStrs, " ") + ")"
}

// IsPrimaryExpression は一次式かどうかを返す。
// SumRollResultではtrueを返す。
func (n *SumRollResult) IsPrimaryExpression() bool {
	return true
}

// IsVariable は可変ノードかどうかを返す。
// SumRollResultではfalseを返す。
func (n *SumRollResult) IsVariable() bool {
	return false
}

// NewSumRollResult は新しい加算ロール結果のノードを返す。
//
// dice: 振られたダイスのスライス。
func NewSumRollResult(dice []die.Die) *SumRollResult {
	r := &SumRollResult{
		Dice: make([]die.Die, len(dice)),
	}

	copy(r.Dice, dice)

	return r
}
