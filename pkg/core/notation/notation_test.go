package notation

import (
	"fmt"
	"testing"

	"github.com/raa0121/GoBCDice/pkg/core/ast"
	"github.com/raa0121/GoBCDice/pkg/core/parser"
)

// 加算ロール式の中置表記の例。
func ExampleInfixNotation_sumRoll() {
	// 構文解析する
	r, parseErr := parser.Parse("ExampleInfixNotation_sumRoll", []byte("(2*3-4)d6-1d4+1"))
	if parseErr != nil {
		return
	}

	node := r.(ast.Node)

	// 中置表記を生成する
	infixNotation, notationErr := InfixNotation(node, true)
	if notationErr != nil {
		return
	}

	// 中置表記を出力する
	fmt.Println(infixNotation)
	// Output: (2*3-4)D6-1D4+1
}

func ExampleParenthesize() {
	parenthesized := Parenthesize("1+2")
	fmt.Println(parenthesized)
	// Output: (1+2)
}

// 演算子の優先順位を考慮した中置表記の生成例。
// "C((1-(2*3))/4)" を入力した場合：
// 演算子の優先順位が "-" < "*" であるため、"2*3" には括弧は不要。
// 演算子の優先順位が "-" < "/" であるため、"/" の左側には括弧が必要。
func ExampleInfixNotation_operatorPrecedence() {
	// 構文解析する
	r, parseErr := parser.Parse("ExampleInfixNotation_operatorPrecedence", []byte("C((1-(2*3))/4)"))
	if parseErr != nil {
		return
	}

	node := r.(ast.Node)

	// 中置表記を生成する
	infixNotation, notationErr := InfixNotation(node, true)
	if notationErr != nil {
		return
	}

	// 中置表記を出力する
	fmt.Println(infixNotation)
	// Output: C((1-2*3)/4)
}

// 演算子の結合性を考慮した中置表記の生成例。
// "C(1+(2-(3-4)))" を入力した場合：
// "+" は右結合性であるため、"(2-(3-4))" の部分全体には括弧は不要。
// "-" は右結合性でないため、"-(3-4)" では括弧が必要。
func ExampleInfixNotation_associativity() {
	// 構文解析する
	r, parseErr := parser.Parse("ExampleInfixNotation_associativity", []byte("C(1+(2-(3-4)))"))
	if parseErr != nil {
		return
	}

	node := r.(ast.Node)

	// 中置表記を生成する
	infixNotation, notationErr := InfixNotation(node, true)
	if notationErr != nil {
		return
	}

	// 中置表記を出力する
	fmt.Println(infixNotation)
	// Output: C(1+2-(3-4))
}

// TestInfixNotationはノードの中置表記をテストする
func TestInfixNotation(t *testing.T) {
	testcase := []struct {
		input    string
		expected string
	}{
		// 計算コマンド
		{"C(1)", "C(1)"},
		{"C(-1)", "C(-1)"},
		{"C(1-(-1))", "C(1-(-1))"},
		{"C(-1+2)", "C(-1+2)"},
		{"C((-1)*2)", "C(-1*2)"},
		{"C((-1)*(-2))", "C(-1*(-2))"},
		{"C(((1+2)+3)+4)", "C(1+2+3+4)"},
		{"C(((1+2)+3)-4)", "C(1+2+3-4)"},
		{"C(((1+2)+3)*4)", "C((1+2+3)*4)"},
		{"C(((1+2)+3)/4)", "C((1+2+3)/4)"},
		{"C(((1+2)-3)+4)", "C(1+2-3+4)"},
		{"C(((1+2)-3)-4)", "C(1+2-3-4)"},
		{"C(((1+2)-3)*4)", "C((1+2-3)*4)"},
		{"C(((1+2)-3)/4)", "C((1+2-3)/4)"},
		{"C(((1+2)*3)+4)", "C((1+2)*3+4)"},
		{"C(((1+2)*3)-4)", "C((1+2)*3-4)"},
		{"C(((1+2)*3)*4)", "C((1+2)*3*4)"},
		{"C(((1+2)*3)/4)", "C((1+2)*3/4)"},
		{"C(((1+2)/3)+4)", "C((1+2)/3+4)"},
		{"C(((1+2)/3)-4)", "C((1+2)/3-4)"},
		{"C(((1+2)/3)*4)", "C((1+2)/3*4)"},
		{"C(((1+2)/3)/4)", "C((1+2)/3/4)"},
		{"C(((1-2)+3)+4)", "C(1-2+3+4)"},
		{"C(((1-2)+3)-4)", "C(1-2+3-4)"},
		{"C(((1-2)+3)*4)", "C((1-2+3)*4)"},
		{"C(((1-2)+3)/4)", "C((1-2+3)/4)"},
		{"C(((1-2)-3)+4)", "C(1-2-3+4)"},
		{"C(((1-2)-3)-4)", "C(1-2-3-4)"},
		{"C(((1-2)-3)*4)", "C((1-2-3)*4)"},
		{"C(((1-2)-3)/4)", "C((1-2-3)/4)"},
		{"C(((1-2)*3)+4)", "C((1-2)*3+4)"},
		{"C(((1-2)*3)-4)", "C((1-2)*3-4)"},
		{"C(((1-2)*3)*4)", "C((1-2)*3*4)"},
		{"C(((1-2)*3)/4)", "C((1-2)*3/4)"},
		{"C(((1-2)/3)+4)", "C((1-2)/3+4)"},
		{"C(((1-2)/3)-4)", "C((1-2)/3-4)"},
		{"C(((1-2)/3)*4)", "C((1-2)/3*4)"},
		{"C(((1-2)/3)/4)", "C((1-2)/3/4)"},
		{"C(((1*2)+3)+4)", "C(1*2+3+4)"},
		{"C(((1*2)+3)-4)", "C(1*2+3-4)"},
		{"C(((1*2)+3)*4)", "C((1*2+3)*4)"},
		{"C(((1*2)+3)/4)", "C((1*2+3)/4)"},
		{"C(((1*2)-3)+4)", "C(1*2-3+4)"},
		{"C(((1*2)-3)-4)", "C(1*2-3-4)"},
		{"C(((1*2)-3)*4)", "C((1*2-3)*4)"},
		{"C(((1*2)-3)/4)", "C((1*2-3)/4)"},
		{"C(((1*2)*3)+4)", "C(1*2*3+4)"},
		{"C(((1*2)*3)-4)", "C(1*2*3-4)"},
		{"C(((1*2)*3)*4)", "C(1*2*3*4)"},
		{"C(((1*2)*3)/4)", "C(1*2*3/4)"},
		{"C(((1*2)/3)+4)", "C(1*2/3+4)"},
		{"C(((1*2)/3)-4)", "C(1*2/3-4)"},
		{"C(((1*2)/3)*4)", "C(1*2/3*4)"},
		{"C(((1*2)/3)/4)", "C(1*2/3/4)"},
		{"C(((1/2)+3)+4)", "C(1/2+3+4)"},
		{"C(((1/2)+3)-4)", "C(1/2+3-4)"},
		{"C(((1/2)+3)*4)", "C((1/2+3)*4)"},
		{"C(((1/2)+3)/4)", "C((1/2+3)/4)"},
		{"C(((1/2)-3)+4)", "C(1/2-3+4)"},
		{"C(((1/2)-3)-4)", "C(1/2-3-4)"},
		{"C(((1/2)-3)*4)", "C((1/2-3)*4)"},
		{"C(((1/2)-3)/4)", "C((1/2-3)/4)"},
		{"C(((1/2)*3)+4)", "C(1/2*3+4)"},
		{"C(((1/2)*3)-4)", "C(1/2*3-4)"},
		{"C(((1/2)*3)*4)", "C(1/2*3*4)"},
		{"C(((1/2)*3)/4)", "C(1/2*3/4)"},
		{"C(((1/2)/3)+4)", "C(1/2/3+4)"},
		{"C(((1/2)/3)-4)", "C(1/2/3-4)"},
		{"C(((1/2)/3)*4)", "C(1/2/3*4)"},
		{"C(((1/2)/3)/4)", "C(1/2/3/4)"},
		{"C((1+(2+3))+4)", "C(1+2+3+4)"},
		{"C((1+(2+3))-4)", "C(1+2+3-4)"},
		{"C((1+(2+3))*4)", "C((1+2+3)*4)"},
		{"C((1+(2+3))/4)", "C((1+2+3)/4)"},
		{"C((1+(2-3))+4)", "C(1+2-3+4)"},
		{"C((1+(2-3))-4)", "C(1+2-3-4)"},
		{"C((1+(2-3))*4)", "C((1+2-3)*4)"},
		{"C((1+(2-3))/4)", "C((1+2-3)/4)"},
		{"C((1+(2*3))+4)", "C(1+2*3+4)"},
		{"C((1+(2*3))-4)", "C(1+2*3-4)"},
		{"C((1+(2*3))*4)", "C((1+2*3)*4)"},
		{"C((1+(2*3))/4)", "C((1+2*3)/4)"},
		{"C((1+(2/3))+4)", "C(1+2/3+4)"},
		{"C((1+(2/3))-4)", "C(1+2/3-4)"},
		{"C((1+(2/3))*4)", "C((1+2/3)*4)"},
		{"C((1+(2/3))/4)", "C((1+2/3)/4)"},
		{"C((1-(2+3))+4)", "C(1-(2+3)+4)"},
		{"C((1-(2+3))-4)", "C(1-(2+3)-4)"},
		{"C((1-(2+3))*4)", "C((1-(2+3))*4)"},
		{"C((1-(2+3))/4)", "C((1-(2+3))/4)"},
		{"C((1-(2-3))+4)", "C(1-(2-3)+4)"},
		{"C((1-(2-3))-4)", "C(1-(2-3)-4)"},
		{"C((1-(2-3))*4)", "C((1-(2-3))*4)"},
		{"C((1-(2-3))/4)", "C((1-(2-3))/4)"},
		{"C((1-(2*3))+4)", "C(1-2*3+4)"},
		{"C((1-(2*3))-4)", "C(1-2*3-4)"},
		{"C((1-(2*3))*4)", "C((1-2*3)*4)"},
		{"C((1-(2*3))/4)", "C((1-2*3)/4)"},
		{"C((1-(2/3))+4)", "C(1-2/3+4)"},
		{"C((1-(2/3))-4)", "C(1-2/3-4)"},
		{"C((1-(2/3))*4)", "C((1-2/3)*4)"},
		{"C((1-(2/3))/4)", "C((1-2/3)/4)"},
		{"C((1*(2+3))+4)", "C(1*(2+3)+4)"},
		{"C((1*(2+3))-4)", "C(1*(2+3)-4)"},
		{"C((1*(2+3))*4)", "C(1*(2+3)*4)"},
		{"C((1*(2+3))/4)", "C(1*(2+3)/4)"},
		{"C((1*(2-3))+4)", "C(1*(2-3)+4)"},
		{"C((1*(2-3))-4)", "C(1*(2-3)-4)"},
		{"C((1*(2-3))*4)", "C(1*(2-3)*4)"},
		{"C((1*(2-3))/4)", "C(1*(2-3)/4)"},
		{"C((1*(2*3))+4)", "C(1*2*3+4)"},
		{"C((1*(2*3))-4)", "C(1*2*3-4)"},
		{"C((1*(2*3))*4)", "C(1*2*3*4)"},
		{"C((1*(2*3))/4)", "C(1*2*3/4)"},
		{"C((1*(2/3))+4)", "C(1*2/3+4)"},
		{"C((1*(2/3))-4)", "C(1*2/3-4)"},
		{"C((1*(2/3))*4)", "C(1*2/3*4)"},
		{"C((1*(2/3))/4)", "C(1*2/3/4)"},
		{"C((1/(2+3))+4)", "C(1/(2+3)+4)"},
		{"C((1/(2+3))-4)", "C(1/(2+3)-4)"},
		{"C((1/(2+3))*4)", "C(1/(2+3)*4)"},
		{"C((1/(2+3))/4)", "C(1/(2+3)/4)"},
		{"C((1/(2-3))+4)", "C(1/(2-3)+4)"},
		{"C((1/(2-3))-4)", "C(1/(2-3)-4)"},
		{"C((1/(2-3))*4)", "C(1/(2-3)*4)"},
		{"C((1/(2-3))/4)", "C(1/(2-3)/4)"},
		{"C((1/(2*3))+4)", "C(1/(2*3)+4)"},
		{"C((1/(2*3))-4)", "C(1/(2*3)-4)"},
		{"C((1/(2*3))*4)", "C(1/(2*3)*4)"},
		{"C((1/(2*3))/4)", "C(1/(2*3)/4)"},
		{"C((1/(2/3))+4)", "C(1/(2/3)+4)"},
		{"C((1/(2/3))-4)", "C(1/(2/3)-4)"},
		{"C((1/(2/3))*4)", "C(1/(2/3)*4)"},
		{"C((1/(2/3))/4)", "C(1/(2/3)/4)"},
		{"C((1+2)+(3+4))", "C(1+2+3+4)"},
		{"C((1+2)+(3-4))", "C(1+2+3-4)"},
		{"C((1+2)+(3*4))", "C(1+2+3*4)"},
		{"C((1+2)+(3/4))", "C(1+2+3/4)"},
		{"C((1+2)-(3+4))", "C(1+2-(3+4))"},
		{"C((1+2)-(3-4))", "C(1+2-(3-4))"},
		{"C((1+2)-(3*4))", "C(1+2-3*4)"},
		{"C((1+2)-(3/4))", "C(1+2-3/4)"},
		{"C((1+2)*(3+4))", "C((1+2)*(3+4))"},
		{"C((1+2)*(3-4))", "C((1+2)*(3-4))"},
		{"C((1+2)*(3*4))", "C((1+2)*3*4)"},
		{"C((1+2)*(3/4))", "C((1+2)*3/4)"},
		{"C((1+2)/(3+4))", "C((1+2)/(3+4))"},
		{"C((1+2)/(3-4))", "C((1+2)/(3-4))"},
		{"C((1+2)/(3*4))", "C((1+2)/(3*4))"},
		{"C((1+2)/(3/4))", "C((1+2)/(3/4))"},
		{"C((1-2)+(3+4))", "C(1-2+3+4)"},
		{"C((1-2)+(3-4))", "C(1-2+3-4)"},
		{"C((1-2)+(3*4))", "C(1-2+3*4)"},
		{"C((1-2)+(3/4))", "C(1-2+3/4)"},
		{"C((1-2)-(3+4))", "C(1-2-(3+4))"},
		{"C((1-2)-(3-4))", "C(1-2-(3-4))"},
		{"C((1-2)-(3*4))", "C(1-2-3*4)"},
		{"C((1-2)-(3/4))", "C(1-2-3/4)"},
		{"C((1-2)*(3+4))", "C((1-2)*(3+4))"},
		{"C((1-2)*(3-4))", "C((1-2)*(3-4))"},
		{"C((1-2)*(3*4))", "C((1-2)*3*4)"},
		{"C((1-2)*(3/4))", "C((1-2)*3/4)"},
		{"C((1-2)/(3+4))", "C((1-2)/(3+4))"},
		{"C((1-2)/(3-4))", "C((1-2)/(3-4))"},
		{"C((1-2)/(3*4))", "C((1-2)/(3*4))"},
		{"C((1-2)/(3/4))", "C((1-2)/(3/4))"},
		{"C((1*2)+(3+4))", "C(1*2+3+4)"},
		{"C((1*2)+(3-4))", "C(1*2+3-4)"},
		{"C((1*2)+(3*4))", "C(1*2+3*4)"},
		{"C((1*2)+(3/4))", "C(1*2+3/4)"},
		{"C((1*2)-(3+4))", "C(1*2-(3+4))"},
		{"C((1*2)-(3-4))", "C(1*2-(3-4))"},
		{"C((1*2)-(3*4))", "C(1*2-3*4)"},
		{"C((1*2)-(3/4))", "C(1*2-3/4)"},
		{"C((1*2)*(3+4))", "C(1*2*(3+4))"},
		{"C((1*2)*(3-4))", "C(1*2*(3-4))"},
		{"C((1*2)*(3*4))", "C(1*2*3*4)"},
		{"C((1*2)*(3/4))", "C(1*2*3/4)"},
		{"C((1*2)/(3+4))", "C(1*2/(3+4))"},
		{"C((1*2)/(3-4))", "C(1*2/(3-4))"},
		{"C((1*2)/(3*4))", "C(1*2/(3*4))"},
		{"C((1*2)/(3/4))", "C(1*2/(3/4))"},
		{"C((1/2)+(3+4))", "C(1/2+3+4)"},
		{"C((1/2)+(3-4))", "C(1/2+3-4)"},
		{"C((1/2)+(3*4))", "C(1/2+3*4)"},
		{"C((1/2)+(3/4))", "C(1/2+3/4)"},
		{"C((1/2)-(3+4))", "C(1/2-(3+4))"},
		{"C((1/2)-(3-4))", "C(1/2-(3-4))"},
		{"C((1/2)-(3*4))", "C(1/2-3*4)"},
		{"C((1/2)-(3/4))", "C(1/2-3/4)"},
		{"C((1/2)*(3+4))", "C(1/2*(3+4))"},
		{"C((1/2)*(3-4))", "C(1/2*(3-4))"},
		{"C((1/2)*(3*4))", "C(1/2*3*4)"},
		{"C((1/2)*(3/4))", "C(1/2*3/4)"},
		{"C((1/2)/(3+4))", "C(1/2/(3+4))"},
		{"C((1/2)/(3-4))", "C(1/2/(3-4))"},
		{"C((1/2)/(3*4))", "C(1/2/(3*4))"},
		{"C((1/2)/(3/4))", "C(1/2/(3/4))"},
		{"C(1+((2+3)+4))", "C(1+2+3+4)"},
		{"C(1+((2+3)-4))", "C(1+2+3-4)"},
		{"C(1+((2+3)*4))", "C(1+(2+3)*4)"},
		{"C(1+((2+3)/4))", "C(1+(2+3)/4)"},
		{"C(1+((2-3)+4))", "C(1+2-3+4)"},
		{"C(1+((2-3)-4))", "C(1+2-3-4)"},
		{"C(1+((2-3)*4))", "C(1+(2-3)*4)"},
		{"C(1+((2-3)/4))", "C(1+(2-3)/4)"},
		{"C(1+((2*3)+4))", "C(1+2*3+4)"},
		{"C(1+((2*3)-4))", "C(1+2*3-4)"},
		{"C(1+((2*3)*4))", "C(1+2*3*4)"},
		{"C(1+((2*3)/4))", "C(1+2*3/4)"},
		{"C(1+((2/3)+4))", "C(1+2/3+4)"},
		{"C(1+((2/3)-4))", "C(1+2/3-4)"},
		{"C(1+((2/3)*4))", "C(1+2/3*4)"},
		{"C(1+((2/3)/4))", "C(1+2/3/4)"},
		{"C(1-((2+3)+4))", "C(1-(2+3+4))"},
		{"C(1-((2+3)-4))", "C(1-(2+3-4))"},
		{"C(1-((2+3)*4))", "C(1-(2+3)*4)"},
		{"C(1-((2+3)/4))", "C(1-(2+3)/4)"},
		{"C(1-((2-3)+4))", "C(1-(2-3+4))"},
		{"C(1-((2-3)-4))", "C(1-(2-3-4))"},
		{"C(1-((2-3)*4))", "C(1-(2-3)*4)"},
		{"C(1-((2-3)/4))", "C(1-(2-3)/4)"},
		{"C(1-((2*3)+4))", "C(1-(2*3+4))"},
		{"C(1-((2*3)-4))", "C(1-(2*3-4))"},
		{"C(1-((2*3)*4))", "C(1-2*3*4)"},
		{"C(1-((2*3)/4))", "C(1-2*3/4)"},
		{"C(1-((2/3)+4))", "C(1-(2/3+4))"},
		{"C(1-((2/3)-4))", "C(1-(2/3-4))"},
		{"C(1-((2/3)*4))", "C(1-2/3*4)"},
		{"C(1-((2/3)/4))", "C(1-2/3/4)"},
		{"C(1*((2+3)+4))", "C(1*(2+3+4))"},
		{"C(1*((2+3)-4))", "C(1*(2+3-4))"},
		{"C(1*((2+3)*4))", "C(1*(2+3)*4)"},
		{"C(1*((2+3)/4))", "C(1*(2+3)/4)"},
		{"C(1*((2-3)+4))", "C(1*(2-3+4))"},
		{"C(1*((2-3)-4))", "C(1*(2-3-4))"},
		{"C(1*((2-3)*4))", "C(1*(2-3)*4)"},
		{"C(1*((2-3)/4))", "C(1*(2-3)/4)"},
		{"C(1*((2*3)+4))", "C(1*(2*3+4))"},
		{"C(1*((2*3)-4))", "C(1*(2*3-4))"},
		{"C(1*((2*3)*4))", "C(1*2*3*4)"},
		{"C(1*((2*3)/4))", "C(1*2*3/4)"},
		{"C(1*((2/3)+4))", "C(1*(2/3+4))"},
		{"C(1*((2/3)-4))", "C(1*(2/3-4))"},
		{"C(1*((2/3)*4))", "C(1*2/3*4)"},
		{"C(1*((2/3)/4))", "C(1*2/3/4)"},
		{"C(1/((2+3)+4))", "C(1/(2+3+4))"},
		{"C(1/((2+3)-4))", "C(1/(2+3-4))"},
		{"C(1/((2+3)*4))", "C(1/((2+3)*4))"},
		{"C(1/((2+3)/4))", "C(1/((2+3)/4))"},
		{"C(1/((2-3)+4))", "C(1/(2-3+4))"},
		{"C(1/((2-3)-4))", "C(1/(2-3-4))"},
		{"C(1/((2-3)*4))", "C(1/((2-3)*4))"},
		{"C(1/((2-3)/4))", "C(1/((2-3)/4))"},
		{"C(1/((2*3)+4))", "C(1/(2*3+4))"},
		{"C(1/((2*3)-4))", "C(1/(2*3-4))"},
		{"C(1/((2*3)*4))", "C(1/(2*3*4))"},
		{"C(1/((2*3)/4))", "C(1/(2*3/4))"},
		{"C(1/((2/3)+4))", "C(1/(2/3+4))"},
		{"C(1/((2/3)-4))", "C(1/(2/3-4))"},
		{"C(1/((2/3)*4))", "C(1/(2/3*4))"},
		{"C(1/((2/3)/4))", "C(1/(2/3/4))"},
		{"C(1+(2+(3+4)))", "C(1+2+3+4)"},
		{"C(1+(2+(3-4)))", "C(1+2+3-4)"},
		{"C(1+(2+(3*4)))", "C(1+2+3*4)"},
		{"C(1+(2+(3/4)))", "C(1+2+3/4)"},
		{"C(1+(2-(3+4)))", "C(1+2-(3+4))"},
		{"C(1+(2-(3-4)))", "C(1+2-(3-4))"},
		{"C(1+(2-(3*4)))", "C(1+2-3*4)"},
		{"C(1+(2-(3/4)))", "C(1+2-3/4)"},
		{"C(1+(2*(3+4)))", "C(1+2*(3+4))"},
		{"C(1+(2*(3-4)))", "C(1+2*(3-4))"},
		{"C(1+(2*(3*4)))", "C(1+2*3*4)"},
		{"C(1+(2*(3/4)))", "C(1+2*3/4)"},
		{"C(1+(2/(3+4)))", "C(1+2/(3+4))"},
		{"C(1+(2/(3-4)))", "C(1+2/(3-4))"},
		{"C(1+(2/(3*4)))", "C(1+2/(3*4))"},
		{"C(1+(2/(3/4)))", "C(1+2/(3/4))"},
		{"C(1-(2+(3+4)))", "C(1-(2+3+4))"},
		{"C(1-(2+(3-4)))", "C(1-(2+3-4))"},
		{"C(1-(2+(3*4)))", "C(1-(2+3*4))"},
		{"C(1-(2+(3/4)))", "C(1-(2+3/4))"},
		{"C(1-(2-(3+4)))", "C(1-(2-(3+4)))"},
		{"C(1-(2-(3-4)))", "C(1-(2-(3-4)))"},
		{"C(1-(2-(3*4)))", "C(1-(2-3*4))"},
		{"C(1-(2-(3/4)))", "C(1-(2-3/4))"},
		{"C(1-(2*(3+4)))", "C(1-2*(3+4))"},
		{"C(1-(2*(3-4)))", "C(1-2*(3-4))"},
		{"C(1-(2*(3*4)))", "C(1-2*3*4)"},
		{"C(1-(2*(3/4)))", "C(1-2*3/4)"},
		{"C(1-(2/(3+4)))", "C(1-2/(3+4))"},
		{"C(1-(2/(3-4)))", "C(1-2/(3-4))"},
		{"C(1-(2/(3*4)))", "C(1-2/(3*4))"},
		{"C(1-(2/(3/4)))", "C(1-2/(3/4))"},
		{"C(1*(2+(3+4)))", "C(1*(2+3+4))"},
		{"C(1*(2+(3-4)))", "C(1*(2+3-4))"},
		{"C(1*(2+(3*4)))", "C(1*(2+3*4))"},
		{"C(1*(2+(3/4)))", "C(1*(2+3/4))"},
		{"C(1*(2-(3+4)))", "C(1*(2-(3+4)))"},
		{"C(1*(2-(3-4)))", "C(1*(2-(3-4)))"},
		{"C(1*(2-(3*4)))", "C(1*(2-3*4))"},
		{"C(1*(2-(3/4)))", "C(1*(2-3/4))"},
		{"C(1*(2*(3+4)))", "C(1*2*(3+4))"},
		{"C(1*(2*(3-4)))", "C(1*2*(3-4))"},
		{"C(1*(2*(3*4)))", "C(1*2*3*4)"},
		{"C(1*(2*(3/4)))", "C(1*2*3/4)"},
		{"C(1*(2/(3+4)))", "C(1*2/(3+4))"},
		{"C(1*(2/(3-4)))", "C(1*2/(3-4))"},
		{"C(1*(2/(3*4)))", "C(1*2/(3*4))"},
		{"C(1*(2/(3/4)))", "C(1*2/(3/4))"},
		{"C(1/(2+(3+4)))", "C(1/(2+3+4))"},
		{"C(1/(2+(3-4)))", "C(1/(2+3-4))"},
		{"C(1/(2+(3*4)))", "C(1/(2+3*4))"},
		{"C(1/(2+(3/4)))", "C(1/(2+3/4))"},
		{"C(1/(2-(3+4)))", "C(1/(2-(3+4)))"},
		{"C(1/(2-(3-4)))", "C(1/(2-(3-4)))"},
		{"C(1/(2-(3*4)))", "C(1/(2-3*4))"},
		{"C(1/(2-(3/4)))", "C(1/(2-3/4))"},
		{"C(1/(2*(3+4)))", "C(1/(2*(3+4)))"},
		{"C(1/(2*(3-4)))", "C(1/(2*(3-4)))"},
		{"C(1/(2*(3*4)))", "C(1/(2*3*4))"},
		{"C(1/(2*(3/4)))", "C(1/(2*3/4))"},
		{"C(1/(2/(3+4)))", "C(1/(2/(3+4)))"},
		{"C(1/(2/(3-4)))", "C(1/(2/(3-4)))"},
		{"C(1/(2/(3*4)))", "C(1/(2/(3*4)))"},
		{"C(1/(2/(3/4)))", "C(1/(2/(3/4)))"},
		{"C(1/2u)", "C(1/2U)"},
		{"C(1/2r)", "C(1/2R)"},
		{"C(100/(1+2)u)", "C(100/(1+2)U)"},
		{"C(100/(1+2)r)", "C(100/(1+2)R)"},
		{"C(-1+-2*-3--4)", "C(-1+(-2)*(-3)-(-4))"},
		{"C(-1+(-2-3)*-4)", "C(-1+(-2-3)*(-4))"},
		{"C((-1+-2)-3*-4)", "C(-1+(-2)-3*(-4))"},
		{"C(-1+-1*2-32/8)", "C(-1+(-1)*2-32/8)"},
		{"C(-1+(-(-1-3))*2-32/8)", "C(-1+(-(-1-3))*2-32/8)"},
		{"C(-1+(-(-1-(1+2)))*2-32/8)", "C(-1+(-(-1-(1+2)))*2-32/8)"},

		// 加算ロール式
		{"2D6", "2D6"},
		{"12D60", "12D60"},
		{"-2D6", "-2D6"},
		{"+2D6", "2D6"},
		{"2D6+1", "2D6+1"},
		{"1+2D6", "1+2D6"},
		{"-2D6+1", "-2D6+1"},
		{"+2D6+1", "2D6+1"},
		{"2d6+1-1-2-3-4", "2D6+1-1-2-3-4"},
		{"2D6+4D10", "2D6+4D10"},
		{"(2D6)", "2D6"},
		{"-(2D6)", "-2D6"},
		{"+(2D6)", "2D6"},
		{"2d6*3", "2D6*3"},
		{"2d6/2", "2D6/2"},
		{"2d6/2u", "2D6/2U"},
		{"2d6/2r", "2D6/2R"},
		{"100/2d6+1", "100/2D6+1"},
		{"100/2d6u+1", "100/2D6U+1"},
		{"100/2d6r+1", "100/2D6R+1"},
		{"100/(2d6+1)+4*5", "100/(2D6+1)+4*5"},
		{"100/(2d6+1)u+4*5", "100/(2D6+1)U+4*5"},
		{"100/(2d6+1)r+4*5", "100/(2D6+1)R+4*5"},
		{"4d10/2d6+1", "4D10/2D6+1"},
		{"4d10/2d6u+1", "4D10/2D6U+1"},
		{"4d10/2d6r+1", "4D10/2D6R+1"},
		{"2d10+3-4", "2D10+3-4"},
		{"2d10+3*4", "2D10+3*4"},
		{"2d10/3+4*5-6", "2D10/3+4*5-6"},
		{"2d10/3u+4*5-6", "2D10/3U+4*5-6"},
		{"2d10/3r+4*5-6", "2D10/3R+4*5-6"},
		{"2d6*3-1d6+1", "2D6*3-1D6+1"},
		{"(2+3)d6-1+3d6+2", "(2+3)D6-1+3D6+2"},
		{"(2*3-4)d6-1d4+1", "(2*3-4)D6-1D4+1"},
		{"((2+3)*4/3)d6*2+5", "((2+3)*4/3)D6*2+5"},
		{"2d(1+5)", "2D(1+5)"},
		{"(8/2)D(4+6)", "(8/2)D(4+6)"},
		{"(2-1)d(8/2)*(1+1)d(3*4/2)+2*3", "(2-1)D(8/2)*(1+1)D(3*4/2)+2*3"},

		// ランダム数値取り出しを含む加算ロール
		{"[1...5]D6", "[1...5]D6"},
		{"([2...4]+2)D10", "([2...4]+2)D10"},
		{"[(2+3)...8]D6", "[(2+3)...8]D6"},
		{"[5...(7+1)]D6", "[5...(7+1)]D6"},
		{"2d[1...5]", "2D[1...5]"},
		{"2d([2...4]+2)", "2D([2...4]+2)"},
		{"2d[(2+3)...8]", "2D[(2+3)...8]"},
		{"2d[5...(7+1)]", "2D[5...(7+1)]"},
		{"[1...5]d(2*3)", "[1...5]D(2*3)"},
		{"(1+1)d[1...5]", "(1+1)D[1...5]"},
		{"([1...4]+1)d([2...4]+2)-1", "([1...4]+1)D([2...4]+2)-1"},

		// 加算ロール式の成功判定
		{"2D6=7", "2D6=7"},
		{"2D6+1=7", "2D6+1=7"},
		{"1+2D6=7", "1+2D6=7"},
		{"2D6<>7", "2D6<>7"},
		{"2D6<7", "2D6<7"},
		{"2D6>7", "2D6>7"},
		{"2D6<=7", "2D6<=7"},
		{"2D6>=7", "2D6>=7"},
		{"-2D6<-7", "-2D6<-7"},
		{"-2D6>-7", "-2D6>-7"},
		{"-2D6<=-7", "-2D6<=-7"},
		{"-2D6>=-7", "-2D6>=-7"},

		// バラバラロール
		{"2b6", "2B6"},
		{"[1...3]b6", "[1...3]B6"},
		{"2b[4...6]", "2B[4...6]"},
		{"[1...3]b[4...6]", "[1...3]B[4...6]"},
		{"(1*2)b6", "(1*2)B6"},
		{"([1...3]+1)b6", "([1...3]+1)B6"},
		{"2b(2+4)", "2B(2+4)"},
		{"2b([3...5]+1)", "2B([3...5]+1)"},
		{"[1...5]b(2*3)", "[1...5]B(2*3)"},
		{"(1+1)b[1...5]", "(1+1)B[1...5]"},
		{"(1*2)b(2+4)", "(1*2)B(2+4)"},
		{"2b6+4b10", "2B6+4B10"},
		{"2b6+3b8+5b12", "2B6+3B8+5B12"},

		// バラバラロールの成功数カウント
		{"2b6=3", "2B6=3"},
		{"2b6<>3", "2B6<>3"},
		{"2b6>3", "2B6>3"},
		{"2b6<3", "2B6<3"},
		{"2b6>=3", "2B6>=3"},
		{"2b6<=3", "2B6<=3"},
		{"2b6>4-1", "2B6>4-1"},
		{"2b6+4b10>4", "2B6+4B10>4"},
		{"2b6>-(-1*3)", "2B6>-(-1*3)"},

		// 個数振り足しロール
		{"3r6>=4", "3R6>=4"},
		{"3r6+2r6<=2", "3R6+2R6<=2"},
		{"(3+2)r6>=5", "(3+2)R6>=5"},
		{"1r(2*3)>=4", "1R(2*3)>=4"},
		{"3r6>1*4", "3R6>1*4"},
		{"2r6", "2R6"},
		{"2r6[5]", "2R6[5]"},
		{"3r6+2r6[2]", "3R6+2R6[2]"},
		{"6R6[6]>=5", "6R6[6]>=5"},
		{"6R6[2*3]>=5", "6R6[2*3]>=5"},

		// 上方無限ロール
		{"3u6", "3U6"},
		{"(1*3)u6", "(1*3)U6"},
		{"3u(5+1)", "3U(5+1)"},
		{"3u6[6]", "3U6[6]"},
		{"3u6[2+4]", "3U6[2+4]"},
		{"3u6+5u6[6]", "3U6+5U6[6]"},
		{"3u6[6]+1", "3U6[6]+1"},
		{"3u6[6]-1", "3U6[6]-1"},
		{"3u6[6]+1*2", "3U6[6]+1*2"},
		{"3u6[6]+5/2", "3U6[6]+5/2"},
		{"1U100[96]+3", "1U100[96]+3"},
		{"3u6[6]=10", "3U6[6]=10"},
		{"3u6[6]<>10", "3U6[6]<>10"},
		{"3u6[6]>10", "3U6[6]>10"},
		{"3u6[6]<10", "3U6[6]<10"},
		{"3u6[6]>=10", "3U6[6]>=10"},
		{"3u6[6]<=10", "3U6[6]<=10"},
		{"3u6[6]>=2+8", "3U6[6]>=2+8"},
		{"3u6[6]+1>=10", "3U6[6]+1>=10"},
		{"3u6+5u6[6]>=7", "3U6+5U6[6]>=7"},
		{"(5+6)u10[10]+5>=8", "(5+6)U10[10]+5>=8"},

		// ランダム選択
		{"choice[A,B,C]どれにしよう", "CHOICE[A,B,C]"},
		{"choice[A,B, ]", "CHOICE[A,B]"},
		{"Choice[ A, B,   C     ,D ]", "CHOICE[A,B,C,D]"},
		{
			input:    "CHOICE[Call of Cthulhu, Sword World, Double Cross]",
			expected: "CHOICE[Call of Cthulhu,Sword World,Double Cross]",
		},
		{
			input:    "CHOICE[日本語, でも,　だいじょうぶ]",
			expected: "CHOICE[日本語,でも,だいじょうぶ]",
		},
		{"choice[1+2, (3*4), 5d6]", "CHOICE[1+2,(3*4),5d6]"},
	}

	for _, test := range testcase {
		t.Run(fmt.Sprintf("%q", test.input), func(t *testing.T) {
			r, parseErr := parser.Parse("test", []byte(test.input))
			if parseErr != nil {
				t.Fatalf("構文エラー: %s", parseErr)
				return
			}

			node := r.(ast.Node)

			actual, notationErr := InfixNotation(node, true)
			if notationErr != nil {
				t.Fatalf("中置表記生成エラー: %s", notationErr)
				return
			}

			if actual != test.expected {
				t.Fatalf("got %q, want %q", actual, test.expected)
			}
		})
	}
}
