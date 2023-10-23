package main

import (
	"reflect"
	"testing"
	"github.com/Azure/azure-sdk-for-go/sdk/ai/azopenai"
)

func Test_getGitDiff(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getGitDiff()
			if (err != nil) != tt.wantErr {
				t.Errorf("getGitDiff() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getGitDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateTimeSaved(t *testing.T) {
	type args struct {
		numCommits int
		wordCount  int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateTimeSaved(tt.args.numCommits, tt.args.wordCount); got != tt.want {
				t.Errorf("calculateTimeSaved() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getCommitStats(t *testing.T) {
	tests := []struct {
		name    string
		want    int
		want1   int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := getCommitStats()
			if (err != nil) != tt.wantErr {
				t.Errorf("getCommitStats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getCommitStats() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getCommitStats() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_getDiffPrompt(t *testing.T) {
	type args struct {
		diff string
	}
	tests := []struct {
		name string
		args args
		want []azopenai.ChatMessage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDiffPrompt(tt.args.diff); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getDiffPrompt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getPrompt(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
		want []azopenai.ChatMessage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPrompt(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getPrompt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getChatCompletionResponse(t *testing.T) {
	type args struct {
		messages []azopenai.ChatMessage
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getChatCompletionResponse(tt.args.messages)
			if (err != nil) != tt.wantErr {
				t.Errorf("getChatCompletionResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getChatCompletionResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getUserName(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Test_getUserName",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getUserName()
		})
	}
}

func Test_formatResponse(t *testing.T) {
	type args struct {
		response string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test with pattern1",
			args: args{
				response: "```Hello World```",
			},
			want: "Hello World",
		},
		{
			name: "Test with pattern2",
			args: args{
				response: "```bashHello World```",
			},
			want: "Hello World",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatResponse(tt.args.response); got != tt.want {
				t.Errorf("formatResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
