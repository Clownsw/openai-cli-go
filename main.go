package main

import (
	"fmt"
	"openai-cli-go/client"
	"openai-cli-go/config"
)

func main() {
	openAiInfo := config.NewOpenAIInfo("C:\\Users\\Administrator\\Downloads\\windows-latest\\application.yml")
	openAIClient := client.NewOpenAIClient(openAiInfo)
	answers, err := openAIClient.Query("傻蛋, 你在吗")
	if err != nil {
		panic(err)
	}

	fmt.Println(answers.Choices[0].ToText())
}
