package entities

import (
	"fmt"

	"github.com/dfaw20/backend-ai-plot/models"
)

type TalePrompt struct {
	targetCharacter models.Character
	heroCharacter   models.Character
	plot            models.Plot
}

func buildFullPrompt(
	targetCharacter models.Character,
	heroCharacter models.Character,
	plot models.Plot,
) string {

	template := `
		以下の内容で小説を執筆してください

		%s
	`
	text := fmt.Sprintf(template,
		plot.BuildPrefixPrompt(),
	)

	return text
}

func (t *TalePrompt) BuildTale() {
}
