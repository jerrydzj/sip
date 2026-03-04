package main

import (
	"fmt"
	"os"

	"github.com/jerrydzj/sip/internal/poem"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: sip <mood>")
		os.Exit(1)
	}

	mood := os.Args[1]
	p, ok := poem.RandomByMood(mood)
	if !ok {
		fmt.Fprintf(os.Stderr, "No poem found for mood: %s\n", mood)
		os.Exit(1)
	}

	fmt.Printf("%s\n%s\n\n%s\n", p.Title, p.Author, p.Text)
}
