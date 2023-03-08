package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"openai-cli-go/config"
	"openai-cli-go/domain/openai"
)

type OpenAIClient struct {
	Config *ClientConfig
}

func (openAIClient *OpenAIClient) Query(query string) (string, error) {
	question := openai.NewOpenAIWithDefault("text-davinci-003", query, "bigduu")
	endpoint := fmt.Sprintf("%s/v1/completions", openAIClient.Config.ApiDomain)

	response, err := openAIClient.Config.ReqClient.R().SetBodyJsonMarshal(question).Post(endpoint)
	if err != nil {
		return config.EmptyString, err
	}

	answers := openai.Answers{}
	err = json.Unmarshal([]byte(response.String()), &answers)
	if err != nil {
		return config.EmptyString, err
	}

	if answers.Error != nil {
		return config.EmptyString, errors.New(answers.Error.Message)
	}

	return answers.Choices[0].ToText(), nil
}

func (openAIClient *OpenAIClient) SetConfig(config *ClientConfig) {
	openAIClient.Config = config
}
