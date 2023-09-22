package main

import (
	"flag"
	"fmt"
)

func main() {
	stats := flag.Bool("stats", false, "display stats")
	ask := flag.String("ask", "", "ask a question")
	flag.Parse()

	if *ask != "" {
        fmt.Printf("You've asked: %s\n", *ask)
        // rest of your code...
    }

	if *stats {
		fmt.Println("Usage: gh-commit")
		numCommits, wordCount, err := getCommitStats()
		hoursSaved := calculateTimeSaved(numCommits, wordCount)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		} else {
			emoji := "🤖"
			message := fmt.Sprintf(
				"Git commit stats:\n"+
					"Number of commits: %d\n"+
					"Number of words in the commit message: %d\n"+
					"Format the git commit stats data and figure out profound insights on how writing these commit messages by AI is saving human hours? Use relevant emojis, use real-world stats in calculations and explain.\n"+
					"If all commit messages were written by %s, you would have saved %.1f hours! %s",
				numCommits, wordCount, "AI", hoursSaved, emoji)

			prompt := getPrompt(message)
			completionResponse, err := getChatCompletionResponse(prompt)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				return
			}
			fmt.Println(completionResponse)
		}
		return
	}

	diff, err := getGitDiff()
	prompt := getDiffPrompt(diff)
	completionResponse, err := getChatCompletionResponse(prompt)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println(completionResponse)
}


