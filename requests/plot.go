package requests

import (
	"errors"
	"strings"
)

type PlotInput struct {
	Title     string
	Prompt    string
	Sensitive bool
}

func (i *PlotInput) TrimTitle() string {
	return strings.TrimSpace(i.Title)
}

func (i *PlotInput) TrimPrompt() string {
	return strings.TrimSpace(i.Prompt)
}

func (i *PlotInput) Validate() error {
	if len(i.TrimTitle()) == 0 {
		return errors.New("タイトルが未入力です")
	}
	if len(i.TrimTitle()) > 20 {
		return errors.New("タイトルの上限は20文字です")
	}
	if len(i.TrimPrompt()) == 0 {
		return errors.New("プロンプトが未入力です")
	}
	if len(i.Prompt) < 30 {
		return errors.New("プロンプトは30文字以上の入力が必要です")
	}
	if len(i.Prompt) > 200 {
		return errors.New("プロンプトの上限は20文字です")
	}

	return nil
}
