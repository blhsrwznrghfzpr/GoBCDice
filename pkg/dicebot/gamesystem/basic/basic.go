/*
ダイスボットの基本機能のパッケージ。
*/
package basic

import (
	"fmt"
	"github.com/raa0121/GoBCDice/pkg/core/command"
	"github.com/raa0121/GoBCDice/pkg/core/evaluator"
	"github.com/raa0121/GoBCDice/pkg/dicebot"
)

const (
	// 基本的なダイスボットのゲーム識別子
	GAME_ID = "DiceBot"
)

// 基本的なダイスボット。
type Basic struct {
}

// New は新しいダイスボットを構築する。
func New() dicebot.DiceBot {
	return &Basic{}
}

// GameID はゲーム識別子を返す。
func (b *Basic) GameID() string {
	return "DiceBot"
}

// GameName はゲームシステム名を返す。
func (b *Basic) GameName() string {
	return "ダイスボット (指定無し)"
}

// Usage はダイスボットの使用法の説明を返す。
func (b *Basic) Usage() string {
	return `【ダイスボット】チャットにダイス用の文字を入力するとダイスロールが可能
入力例）２ｄ６＋１　攻撃！
出力例）2d6+1　攻撃！
　　　　  diceBot: (2d6) → 7
上記のようにダイス文字の後ろに空白を入れて発言する事も可能。
以下、使用例
　3D6+1>=9 ：3d6+1で目標値9以上かの判定
　1D100<=50 ：D100で50％目標の下方ロールの例
　3U6[5] ：3d6のダイス目が5以上の場合に振り足しして合計する(上方無限)
　3B6 ：3d6のダイス目をバラバラのまま出力する（合計しない）
　10B6>=4 ：10d6を振り4以上のダイス目の個数を数える
　(8/2)D(4+6)<=(5*3)：個数・ダイス・達成値には四則演算も使用可能
　C(10-4*3/2+2)：C(計算式）で計算だけの実行も可能
　choice[a,b,c]：列挙した要素から一つを選択表示。ランダム攻撃対象決定などに
　S3d6 ： 各コマンドの先頭に「S」を付けると他人結果の見えないシークレットロール
　3d6/2 ： ダイス出目を割り算（切り捨て）。切り上げは /2U、四捨五入は /2R。
　D66 ： D66ダイス。順序はゲームに依存。D66N：そのまま、D66S：昇順。`
}

// ExecuteCommand は指定されたコマンドを実行する。
//
// 基本のダイスボットには特別なコマンドが存在しないため、必ずエラーを返す。
func (b *Basic) ExecuteCommand(
	_ string,
	_ *evaluator.Evaluator,
) (*command.Result, error) {
	return nil, fmt.Errorf("no game-system-specific command")
}
