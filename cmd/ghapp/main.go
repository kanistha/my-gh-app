package main

import (
	"log"
	"os"

	"github.com/kanistha/my-gh-app/internal/github"
)

func main() {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		log.Fatal("GitHub token not found. Set GITHUB_TOKEN environment variable.")
	}

	owner := "username"
	repo := "repository"
	prNumber := 1
	labelToAdd := "label"

	err := github.AddLabelToPR(token, owner, repo, prNumber, labelToAdd)
	if err != nil {
		log.Fatalf("Error adding label to PR: %v", err)
	}

	log.Printf("Label '%s' added to PR #%d in %s/%s.", labelToAdd, prNumber, owner, repo)
}
