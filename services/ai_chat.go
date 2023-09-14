package services

import (
	"time"

	"github.com/dfaw20/backend-ai-plot/configuration"
)

type ChatGenerator interface {
	Generate(prompt string) (string, error)
}

type ApiAiChatGenerator struct{}

func (g *ApiAiChatGenerator) Generate(prompt string) (string, error) {
	return "Chat AI Impl", nil
}

type LocalStubGenerator struct{}

func (g *LocalStubGenerator) Generate(prompt string) (string, error) {
	time.Sleep(10 * time.Second)
	return `The test of a first-rate intelligence is the ability to hold two opposed ideas in mind at the same time and still retain the ability to function.`, nil
}

func NewChatGenerator(config configuration.Config) ChatGenerator {
	if config.Production {
		return &ApiAiChatGenerator{}
	} else {
		return &LocalStubGenerator{}
	}
}
