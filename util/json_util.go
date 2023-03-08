package util

import (
	"github.com/bytedance/sonic"
	"openai-cli-go/config"
)

func JsonGetString(json string, path ...interface{}) (string, error) {
	result, err := sonic.Get(StringToByteSlice(json), path...)
	if err != nil {
		return config.EmptyString, err
	}
	return result.String()
}
