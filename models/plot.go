package models

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type Support struct {
	gorm.Model
	PlotID      uint
	CharacterID uint
}

// イベント
type Plot struct {
	gorm.Model
	UserID       uint      // 製作者
	Title        string    // イベントのタイトル
	Description  string    // イベントを選ぶときに表示される説明
	Prompt       string    // AIに入力するプロンプト(自由入力)
	Location     string    // イベントの場所
	Season       string    // イベントの季節
	Genre        string    // イベントのジャンル 例:ラブコメ,冒険,ホラーなど
	OutputFormat string    // 出力形式の指定 例:小説,会話など
	ShowWarning  bool      // 警告文の出力のフラグ
	Supports     []Support `gorm:"many2many:supports;" json:"-"` // イベントの登場人物
}

func (plot *Plot) getWarning() string {
	if plot.ShowWarning {
		return "※文頭に警告文を出力"
	}
	return ""
}

func (plot *Plot) getOutputFormat() string {
	if len(strings.TrimSpace(plot.OutputFormat)) == 0 {
		return ""
	}

	return "※出力形式: " + plot.OutputFormat
}

func (plot *Plot) getTitle() string {
	if len(strings.TrimSpace(plot.Title)) == 0 {
		return ""
	}

	return "タイトル: " + plot.Title
}

func (plot *Plot) getGenre() string {
	if len(strings.TrimSpace(plot.Genre)) == 0 {
		return "ジャンル: ラブコメ"
	}

	return "ジャンル: " + plot.Genre
}

func (plot *Plot) getLocation() string {
	if len(strings.TrimSpace(plot.Location)) == 0 {
		return ""
	}

	return "物語の舞台: " + plot.Location
}

func (plot *Plot) getSeason() string {
	if len(strings.TrimSpace(plot.Season)) == 0 {
		return ""
	}

	return "季節: " + plot.Season
}

func (plot *Plot) BuildPrefixPrompt() string {

	template := `
		%s
		%s

		%s
		%s
		%s
		%s
	`
	text := fmt.Sprintf(template,
		plot.getWarning(),
		plot.getOutputFormat(),
		plot.getTitle(),
		plot.getGenre(),
		plot.getLocation(),
		plot.getSeason(),
	)

	return text
}
