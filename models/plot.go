package models

import (
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
	UserID       uint   // 製作者
	Title        string // イベントのタイトル
	Description  string // イベントを選ぶときに表示される説明
	Prompt       string // AIに入力するプロンプト(自由入力)
	Location     string // イベントの場所
	Season       string // イベントの季節
	Genre        string // イベントのジャンル 例:ラブコメ,冒険,ホラーなど
	OutputFormat string // 出力形式の指定 例:小説,会話など
	ShowWarning  bool   // 警告文の出力のフラグ
	Sensitive    bool
	Supports     []Support `gorm:"many2many:supports;" json:"-"` // イベントの登場人物
}

func (plot *Plot) getWarning() string {
	if plot.ShowWarning {
		return "※文頭に警告文を出力"
	}
	return ""
}

func (plot *Plot) GetOutputFormat() string {
	if len(strings.TrimSpace(plot.OutputFormat)) == 0 {
		return "小説"
	}

	return strings.TrimSpace(plot.OutputFormat)
}

func (plot *Plot) getTitle() string {
	if len(strings.TrimSpace(plot.Title)) == 0 {
		return ""
	}

	return "タイトル: " + strings.TrimSpace(plot.Title)
}

func (plot *Plot) GetGenre() string {
	if len(strings.TrimSpace(plot.Genre)) == 0 {
		return ""
	}

	return strings.TrimSpace(plot.Genre)
}

func (plot *Plot) getLocation() string {
	if len(strings.TrimSpace(plot.Location)) == 0 {
		return ""
	}

	return "物語の舞台: " + strings.TrimSpace(plot.Location)
}

func (plot *Plot) getSeason() string {
	if len(strings.TrimSpace(plot.Season)) == 0 {
		return ""
	}

	return "季節: " + strings.TrimSpace(plot.Season)
}

func (plot *Plot) BuildPrefixPrompt() string {

	var lines []string

	warning := plot.getWarning()
	if len(warning) > 0 {
		lines = append(lines, warning)
	}

	lines = append(lines, "")

	location := plot.getLocation()
	if len(location) > 0 {
		lines = append(lines, location)
	}

	season := plot.getSeason()
	if len(season) > 0 {
		lines = append(lines, season)
	}

	return strings.Join(lines, "\n")
}
