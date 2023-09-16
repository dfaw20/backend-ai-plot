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
	UserID    uint   // 製作者
	Title     string // イベントのタイトル
	Prompt    string // AIに入力するプロンプト(自由入力)
	Sensitive bool
	Supports  []Support // イベントの登場人物
}

func (plot *Plot) getTitle() string {
	if len(strings.TrimSpace(plot.Title)) == 0 {
		return ""
	}

	return "タイトル: " + strings.TrimSpace(plot.Title)
}
