package gpt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"cvpchatalert/pkg/jsonlog"

	"github.com/joho/godotenv"

	cvslack "cvpchatalert/pkg/slack"
)

type ResponseJson struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Text         string      `json:"text"`
		Index        int         `json:"index"`
		Logprobs     interface{} `json:"logprobs"`
		FinishReason string      `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

type RequestJson struct {
	Model       string  `json:"model"`
	Prompt      string  `json:"prompt"`
	Maxtokens   int     `json:"max_tokens"`
	Temperature float32 `json:"temperature"`
}

func AskGpt(question string) {
	godotenv.Load()

	apiKey := os.Getenv("CHATGPTTOK")
	if apiKey == "" {
		log.Fatalln("Missing API KEY")
	}
	httpposturl := "https://api.openai.com/v1/completions"

	JsonRequest := RequestJson{
		Model:       "text-davinci-003",
		Prompt:      "How do I fix this on my Arista switch " + question,
		Maxtokens:   400,
		Temperature: 0.5,
	}

	j, err := json.Marshal(JsonRequest)
	if err != nil {
		fmt.Println(err)
	}

	request, error := http.NewRequest("POST", httpposturl, bytes.NewBuffer(j))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	models := ResponseJson{}
	_ = json.Unmarshal(body, &models)

	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	var gptresponse string

	for _, i := range models.Choices {
		logger.PrintInfo(i.Text+"\n", nil)
		gptresponse = i.Text
	}
	cvslack.SendtoSlack(question, gptresponse)
}
