package main

import (
	"io"
	"bytes"
	"log"
	"os"
	"testing"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    os.Exit(m.Run())
}

func TestStatsFlag(t *testing.T) {
    os.Args = []string{"cmd", "-stats"}

    // Capture the output
    var buf bytes.Buffer
    run(&buf)

    // Check the output
    got := buf.String()
    if got == "" {
        t.Errorf("No output for stats flag")
    }
}

func TestAskFlag(t *testing.T) {
    os.Args = []string{"cmd", "-ask", "What's the weather like?"}

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