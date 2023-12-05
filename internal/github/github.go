package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v32/github"
	"golang.org/x/oauth2"
)

// AddLabelToPR adds a label to a pull request.
func AddLabelToPR(token, owner, repo string, prNumber int, labelToAdd string) error {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	_, _, err := client.Issues.AddLabelsToIssue(ctx, owner, repo, prNumber, []string{labelToAdd})
	if err != nil {
		return fmt.Errorf("error adding label to PR: %v", err)
	}

	return nil
}
