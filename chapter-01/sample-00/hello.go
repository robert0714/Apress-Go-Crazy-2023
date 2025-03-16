package main

import (
	"context"
	"fmt"
	"log"

	"github.com/PullRequestInc/go-gpt3"
)

func main() {
	apiKey := "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"
	ctx := context.Background()
	client := gpt3.NewClient(apiKey)

	request := gpt3.CompletionRequest{
		Prompt: []string{"How many cups of coffee should I drink per day?"},
	}
	resp, err := client.Completion(ctx, request)
	if err != nil {
		log.Fatalf("Failed to get completion: %v", err)
	}

	fmt.Print(resp.Choices[0].Text)
}
