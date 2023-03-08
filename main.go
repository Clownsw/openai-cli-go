package main

import (
	"openai-cli-go/client"
	"openai-cli-go/config"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		panic("not found query param!")
	}

	var query strings.Builder

	for i := 2; i < len(args); i++ {
		query.WriteString(args[i])
	}

	openAiInfo := config.NewOpenAIInfo("application.yml")
	c := client.NewClient(args[1], openAiInfo)
	result, err := c.Query(query.String())
	if err != nil {
		panic(err)
	}
	println(strings.TrimSpace(result))
}
