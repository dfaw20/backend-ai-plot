package services

type ChatGenerator interface {
	Generate(prompt string) (string, error)
}

type ApiAiChatGenerator struct{}

func (g *ApiAiChatGenerator) Generate(prompt string) (string, error) {
	return "", nil
}

type LocalStubGenerator struct{}

func (g *LocalStubGenerator) Generate(prompt string) (string, error) {
	return `The test of a first-rate intelligence is the ability to hold two opposed ideas in mind at the same time and still retain the ability to function.`, nil
}
