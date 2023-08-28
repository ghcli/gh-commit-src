package main

import (
	"context"
	"fmt"
	"github.com/cli/go-gh/v2/pkg/api"
	openai "github.com/sashabaranov/go-openai"
	"os"
	"os/exec"
)

func getGitDiff() (string, error) {
	cmd := exec.Command("git", "diff")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("error running git diff: %v", err)
	}
	return string(output), nil
}

func getCommitMessagePrompt() (string, error) {
	diff, err := getGitDiff()
	if err != nil {
		return "", fmt.Errorf("error getting git diff: %v", err)
	}

	prompt := fmt.Sprintf("Please review the following changes:\n\n%s\n\nEnter commit message:", diff)
	return prompt, nil
}

func getChatCompletionResponse(prompt string) (string, error) {
	token := os.Getenv("OPENAI_API_KEY")
	if token == "" {
		return "", fmt.Errorf("OPENAI_API_KEY environment variable not set")
	}
	client := openai.NewClient(token)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)

	if err != nil {
		return "", fmt.Errorf("Completion error: %v", err)
	}

	return resp.Choices[0].Message.Content, nil
}

func main() {
	fmt.Println("hi world, this is the gh-commit extension!")
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
	fmt.Printf("running as %s\n", response.Login)
	prompt, err := getCommitMessagePrompt()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	completionResponse, err := getChatCompletionResponse(prompt)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println(completionResponse)
}
