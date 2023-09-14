package services

type ChatGenerator interface {
	Generate(prompt string) (string, error)
}

type ApiAiChatGenerator struct{}

func (g *ApiAiChatGenerator) Generate(prompt string) (string, error) {
	return "", nil
}

type LocalSampleGenerator struct{}

func (g *LocalSampleGenerator) Generate(prompt string) (string, error) {
	return "", nil
}
