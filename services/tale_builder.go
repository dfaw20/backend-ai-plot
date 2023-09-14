package services

import (
	"fmt"

	"github.com/dfaw20/backend-ai-plot/models"
)

type TaleBuilder struct{}

func getWarning(plot models.Plot) string {
	if (plot.ShowWarning) {
		return "文頭に警告文を出力"
	}
}

func buildPrompt(
	targetCharacter models.Character,
	heroCharacter models.Character,
	plot models.Plot,
) string {
	template := `
		%s
		%s
		%s
	`
	text := fmt.Sprintf(template,
		ge
	)

	return text
}

func (b *TaleBuilder) BuildTale(
	targetCharacter models.Character,
	heroCharacter models.Character,
	plot models.Plot,
) {
}
