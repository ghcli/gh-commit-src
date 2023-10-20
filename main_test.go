package main

import (
	"flag"
	"testing"
)

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

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
