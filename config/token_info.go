package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

type OpenAIInfo struct {
	Token   string `yaml:"token"`
	BaseUrl string `yaml:"base_url"`
}

func NewOpenAIInfo(filePath string) *OpenAIInfo {
	openAiInfo := OpenAIInfo{}

	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(fileContent, &openAiInfo)
	if err != nil {
		panic(err)
	}

	return &openAiInfo
}
