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
		`以下の内容で%s小説を執筆してください

%s

・登場人物

%s

%s

・物語の流れ

%s
`
	text := fmt.Sprintf(template,
		t.plot.GetGenre(),
		t.plot.BuildPrefixPrompt(),
		t.targetCharacter.BuildPrompt(),
		t.heroCharacter.BuildPrompt(),
		embedCharactersInPrompt(t.targetCharacter, t.heroCharacter, t.plot),
	)

	return text
}
