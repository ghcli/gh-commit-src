package main

import (
	"context"
	"fmt"
	"math"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"github.com/Azure/azure-sdk-for-go/sdk/ai/azopenai"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/cli/go-gh/v2/pkg/api"
	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

const MaxDiffLength = 30000 // set to 30k since large model has maximum context length is 32768 tokens.

func getGitDiff() (string, error) {
	cmd := exec.Command("git", "diff")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("error running git diff: %v", err)
	}
	diff := string(output)
    runes := []rune(diff)
	size := len(runes)
    if size > MaxDiffLength {
        runes = runes[:MaxDiffLength]
        return string(runes), fmt.Errorf("the total length was %d and only first 30k were used", size)
    }
	return string(output), nil
}

func calculateTimeSaved(numCommits int, wordCount int) float64 {

	// Assuming an average typing speed of 40 words per minute
	wordsPerMinute := 40.0
	hoursSaved := float64(wordCount) / wordsPerMinute / 60
	return math.Round(hoursSaved*10) / 10
}

func getCommitStats() (int, int, error) {
	cmd := exec.Command("git", "log", "--oneline")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return 0, 0, err
	}
	if err := cmd.Start(); err != nil {
		return 0, 0, err
	}
	defer cmd.Wait()
	cmd = exec.Command("wc", "-lw")
	cmd.Stdin = stdout
	output, err := cmd.Output()
	if err != nil {
		return 0, 0, err
	}
	fields := strings.Fields(string(output))
	numLines, err := strconv.Atoi(fields[0])
	if err != nil {
		return 0, 0, err
	}
	numWords, err := strconv.Atoi(fields[1])
	if err != nil {
		return numLines, 0, err
	}
	return numLines, numWords, nil
}

func getDiffPrompt(diff string) []azopenai.ChatMessage {

	prompt := os.Getenv("PROMPT_OVERRIDE")
	if prompt == "" {
		prompt = `You will examine and explain the given code changes and write a commit message in Conventional Commits format. 
		The first line of the commit message should be a 20 word Title summary include a type, optional scope, subject in text, seperated by a newline and the following body. 
		The types should be one of:
			- fix: for a bug fix
			- feat: for a new feature 
			- perf: for a performance improvement
			- revert: to revert a previous commit
		The body will explain the code change. Body will be formatted in well structured beautifully rendered and use relevant emojis
		if no code changes are detected, you will reply with no code change detected message.`
	}
	messages := []azopenai.ChatMessage{
		{Role: to.Ptr(azopenai.ChatRoleSystem), Content: to.Ptr(prompt)},
		{Role: to.Ptr(azopenai.ChatRoleUser), Content: to.Ptr(diff)},
		{Role: to.Ptr(azopenai.ChatRoleSystem), Content: to.Ptr("Commit message as follows:")},
	}
	return messages
}

func getPrompt(message string) []azopenai.ChatMessage {
	messages := []azopenai.ChatMessage{
		{Role: to.Ptr(azopenai.ChatRoleSystem), Content: to.Ptr(message)},
	}
	return messages
}

func getChatCompletionResponse(messages []azopenai.ChatMessage) (string, error) {
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

func getUserName() {
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
}
