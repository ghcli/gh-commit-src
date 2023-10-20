package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Error loading .env file. If environment variables are set, tests will still run.")
	}
	os.Exit(m.Run())
}

func TestStatsFlag(t *testing.T) {
	flagSet := flag.NewFlagSet("TestStatsFlag", flag.ContinueOnError)
	stats := flagSet.Bool("stats", false, "get stats")

	err := flagSet.Parse([]string{"-stats"})
	if err != nil {
		t.Fatal("Error parsing flags:", err)
	}

	if *stats != true {
		t.Error("Expected stats flag to be true, but it was false")
	}
}

func TestAskFlag(t *testing.T) {
	os.Args = []string{"gh-commit", "-ask", "What's the weather like?"}

	// Capture the output
	var buf bytes.Buffer
	run(&buf)

	// Check the output
	got := buf.String()
	if got == "" {
		t.Errorf("No output for ask flag")
	}
}

func run(out io.Writer) {
	stats := flag.Bool("stats", false, "display stats")
	ask := flag.String("ask", "", "ask a question")
	flag.Parse()

	if flag.NFlag() == 0 {
		fmt.Fprintln(out, "No flags were passed. Please provide a flag.")
		return
	}

	if *ask != "" {
		fmt.Fprintf(out, "You've asked: %s\n", *ask)
	}

	if *stats {
		fmt.Fprintln(out, "Usage: gh-commit")
		// rest of your code...
	}

	// rest of your code...
}
