package main

import (
	"fmt"
	"openai-cli-go/client"
	"openai-cli-go/config"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		panic("not found query param!")
	}
	
	var query strings.Builder

	for i := 1; i < len(args); i++ {
		query.WriteString(args[i])
	}

	openAiInfo := config.NewOpenAIInfo("application.yml")
	openAIClient := client.NewOpenAIClient(openAiInfo)
	answers, err := openAIClient.Query(query.String())

	if err != nil {
		panic(err)
	}

	if answers.Error != nil {
		panic(answers.Error.Message)
	}

	fmt.Println(answers.Choices[0].ToText())
}
