package test

import (
	"github.com/bytedance/sonic"
	"openai-cli-go/util"
	"testing"
)

func TestJsonPath(t *testing.T) {
	str := `{"a":[{"name":"123123"}]}`
	r1, err := sonic.Get(util.StringToByteSlice(str), "a")
	if err != nil {
		t.Fatal(err)
	}

	r2, err := r1.ArrayUseNode()
	if err != nil {
		t.Fatal(err)
	}
	println(len(r2))

	r3, err := r2[0].Get("name").String()
	if err != nil {
		t.Fatal(err)
	}
	println(r3)
}
