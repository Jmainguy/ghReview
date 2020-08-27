package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/v32/github"
	"github.com/slack-go/slack"
	"golang.org/x/oauth2"
	"os"
	"strings"
	"time"
)

func main() {

	slackToken := os.Getenv("slackToken")
	slackUserID := os.Getenv("slackUserID")
	githubName := os.Getenv("githubName")
	subscribedRepos := strings.Split(os.Getenv("subscribedRepos"), " ")

	api := slack.New(slackToken)
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("ghReviewToken")},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	alreadyPinged := make(map[string]bool)

	// Never ending loop, unless app crashes, restart after 5 minute sleep
	for {
		// List of repos to watch
		for _, ownerRepo := range subscribedRepos {
			owner := strings.Split(ownerRepo, "/")[0]
			repo := strings.Split(ownerRepo, "/")[1]
			// List of pull requests for specific repo
			pullRequests, _, err := client.PullRequests.List(ctx, owner, repo, nil)
			if err != nil {
				panic(err)
			}
			for _, pr := range pullRequests {
				// We only care about open PR's
				if *pr.State == "open" {
					// If it has a requestedReviewer, that has not reviewed, print the PR and the reviewer.
					for _, reviewer := range pr.RequestedReviewers {
						if *reviewer.Login == githubName {
							if !alreadyPinged[*pr.HTMLURL] {
								// The PR has a requested Reviewer, it is open, it is assigned to our user, and we have not sent an message about it yet.
								// Slack
								msg := fmt.Sprintf("You have been assigned to review %s\n", *pr.HTMLURL)
								_, _, err := api.PostMessage(slackUserID, slack.MsgOptionText(msg, false))
								if err != nil {
									panic(err)
								}
								// Mark this as sent, so it is not sent again. (unless app crashes)
								alreadyPinged[*pr.HTMLURL] = true
							}
						}
					}
				}
			}
		}
		time.Sleep(5 * time.Minute)
	}

}
