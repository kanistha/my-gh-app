package github

import (
	"testing"

	"github.com/kanistha/my-gh-app/internal/github"
)

func TestAddLabelToPR(t *testing.T) {
	token := "token"
	owner := "test-owner"
	repo := "test-repo"
	prNumber := 42
	labelToAdd := "test-label"

	err := github.AddLabelToPR(token, owner, repo, prNumber, labelToAdd)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}
