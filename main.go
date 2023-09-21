package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/ai/azopenai"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/cli/go-gh/v2/pkg/api"
	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

func getGitDiff() (string, error) {
	cmd := exec.Command("git", "diff")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("error running git diff: %v", err)
	}
	return string(output), nil
}

func getChatCompletionResponse() (string, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Errorf(".env file not found: %v", err)
	}
	keyCredential, err := azopenai.NewKeyCredential(os.Getenv("OPENAI_API_KEY"))
	if err != nil {
		fmt.Errorf("export OPENAI_API_KEY=<api_key> #execute this in your terminal and try again")
		return "", fmt.Errorf("error creating Azure OpenAI client: %v", err)
	}
	url := os.Getenv("OPENAI_URL")
	model := os.Getenv("OPENAI_MODEL")
	var client *azopenai.Client

	if strings.Contains(url, "azure") {
		client, err = azopenai.NewClientWithKeyCredential(url, keyCredential, nil)
		if err != nil {
			return "", fmt.Errorf("error creating Azure OpenAI client: %v", err)
		}
	} else {
		client, err = azopenai.NewClientForOpenAI(url, keyCredential, nil)
		if err != nil {
			return "", fmt.Errorf("error creating Azure OpenAI client: %v", err)
		}

	}
	if model == "" {
		model = openai.GPT4
	}
	diff, err := getGitDiff()
	if err != nil {
		return "", fmt.Errorf("error getting git diff: %v", err)
	}

	messages := []azopenai.ChatMessage{
		{Role: to.Ptr(azopenai.ChatRoleSystem), Content: to.Ptr("You will examine and explain the given code changes and provide a commit message.")},
		{Role: to.Ptr(azopenai.ChatRoleUser), Content: to.Ptr(diff)},
		{Role: to.Ptr(azopenai.ChatRoleSystem), Content: to.Ptr("Enter commit message:")},
	}

	resp, err := client.GetChatCompletions(
		context.Background(),
		azopenai.ChatCompletionsOptions{
			Messages:   messages,
			Deployment: model,
		},
		nil,
	)

	if err != nil {
		return "", fmt.Errorf("Completion error: %v", err)
	}

	//for _, choice := range resp.Choices {
	//	fmt.Fprintf(os.Stderr, "Content[%d]: %s\n", *choice.Index, *choice.Message.Content)
	//}

	return *resp.Choices[0].Message.Content, nil
}

func main() {
	fmt.Println("examining code changes in the commit")
	client, err := api.DefaultRESTClient()
	if err != nil {
		fmt.Println(err)
		return
	}
	response := struct{ Login string }{}
	err = client.Get("user", &response)
	if err != nil {
		fmt.Println(err)
		return
	}
	completionResponse, err := getChatCompletionResponse()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println(completionResponse)
}
