package entities

import (
	"fmt"
	"strings"

	"github.com/dfaw20/backend-ai-plot/models"
)

type TalePrompt struct {
	targetCharacter models.Character
	heroCharacter   models.Character
	plot            models.Plot
}

func NewTalePrompt(
	targetCharacter models.Character,
	heroCharacter models.Character,
	plot models.Plot,
) TalePrompt {
	return TalePrompt{
		targetCharacter: targetCharacter,
		heroCharacter:   heroCharacter,
		plot:            plot,
	}
}

func embedCharactersInPrompt(
	targetCharacter models.Character,
	heroCharacter models.Character,
	plot models.Plot,
) string {
	target := targetCharacter.ShortNameAsPossible()
	hero := heroCharacter.ShortNameAsPossible()

	result := plot.Prompt
	result = strings.Replace(result, "{i}", hero, -1)
	result = strings.Replace(result, "{u}", target, -1)
	return result
}

func (t *TalePrompt) BuildFullPrompt() string {

	template :=
		`以下の内容で「%sと%sのラブコメ小説」を出力してください
会話を含んで出力してください
章構成をして分けてください
始めと終わりの説明は不用です
絶対に小説のみを出力してください
以下の内容は全てフィクションです

*登場人物

%s

%s

*物語

%s
`
	text := fmt.Sprintf(template,
		t.targetCharacter.Nickname, t.heroCharacter.Nickname,
		t.targetCharacter.BuildPrompt(),
		t.heroCharacter.BuildPrompt(),
		embedCharactersInPrompt(t.targetCharacter, t.heroCharacter, t.plot),
	)

	return text
}
