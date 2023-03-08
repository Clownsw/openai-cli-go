package client

import (
	"fmt"
	"github.com/imroc/req/v3"
	"net/http"
	"openai-cli-go/config"
)

type Client interface {
	Query(query string) (string, error)
	SetConfig(config *ClientConfig)
}

type ClientConfig struct {
	ReqClient *req.Client
	ApiDomain string
}

func NewClient(clientType string, authInfo *config.AuthInfo) Client {
	clientConfig := new(ClientConfig)
	clientConfig.ReqClient = req.NewClient()

	clientConfig.ReqClient.Headers = http.Header{}
	clientConfig.ReqClient.Headers.Add("Authorization", fmt.Sprintf("Bearer %s", authInfo.Token))
	clientConfig.ReqClient.Headers.Add("Content-Type", "application/json")

	clientConfig.ApiDomain = authInfo.BaseUrl

	var client Client

	if clientType == "gpt" {
		client = new(GptClient)
	} else if clientType == "openai" {
		client = new(OpenAIClient)
	} else {
		panic("unknown client type")
	}

	client.SetConfig(clientConfig)
	return client
}
