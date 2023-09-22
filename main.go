package main

import (
	"context"
	"flag"
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

func getGitDiff() (string, error) {
	cmd := exec.Command("git", "diff")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("error running git diff: %v", err)
	}
	return string(output), nil
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

func getPrompt(diff string) []azopenai.ChatMessage {
	messages := []azopenai.ChatMessage{
		{Role: to.Ptr(azopenai.ChatRoleSystem), Content: to.Ptr("You will examine and explain the given code changes and provide a commit message. The first line of the response will be a 20 word Title summary ending with a newline in plain text. The subsequent lines will have a detailed commit message. You will write the commit message in well structured beautiful markdown and use relevant emojis")},
		{Role: to.Ptr(azopenai.ChatRoleUser), Content: to.Ptr(diff)},
		{Role: to.Ptr(azopenai.ChatRoleSystem), Content: to.Ptr("Enter commit message:")},
	}
	return messages
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
	fmt.Sprintf(" %s", output)
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

func calculateTimeSaved(numCommits int, wordCount int) float64 {

	// Assuming an average typing speed of 40 words per minute
	wordsPerMinute := 40.0
	hoursSaved := float64(wordCount) / wordsPerMinute / 60
	return math.Round(hoursSaved*10) / 10
}

func main() {
	help := flag.Bool("help", false, "display help message")
	flag.Parse()

	if *help {
		fmt.Println("Usage: gh-commit")
		numCommits, wordCount, err := getCommitStats()
		hoursSaved := calculateTimeSaved(numCommits, wordCount)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		} else {
			fmt.Printf("Number of commits: %d\n", numCommits)
			fmt.Printf("Number of words in the commit message: %d\n", wordCount)
			message := fmt.Sprintf("Could you please format the git commit stats data (number of commits: %d, word count: %d) and figure out profound insights on how writing these commit messages by AI is saving human hours? use relevant emojis, use real world stats in calculations and explain", numCommits, wordCount)

			emoji := "🤖"
			hourssavedmsg := fmt.Sprintf("If all commit messages were written by %s, you would saved %.1f hours! %s", "AI", hoursSaved, emoji)

			messages := []azopenai.ChatMessage{
				{Role: to.Ptr(azopenai.ChatRoleSystem), Content: to.Ptr(message)},
				{Role: to.Ptr(azopenai.ChatRoleSystem), Content: to.Ptr(hourssavedmsg)},
			}
			completionResponse, err := getChatCompletionResponse(messages)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				return
			}
			fmt.Println(completionResponse)
		}
		return
	}

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

	diff, err := getGitDiff()

	messages := getPrompt(diff)

	completionResponse, err := getChatCompletionResponse(messages)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println(completionResponse)
}
