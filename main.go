package main

import (
	"fmt"
	"openai-cli-go/client"
	"openai-cli-go/config"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		panic("not found query param!")
	}

	openAiInfo := config.NewOpenAIInfo("application.yml")
	openAIClient := client.NewOpenAIClient(openAiInfo)
	answers, err := openAIClient.Query(args[1])

	if err != nil {
		panic(err)
	}

	if answers.Error != nil {
		panic(answers.Error.Message)
	}

	fmt.Println(answers.Choices[0].ToText())
}
