package models

import (
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
