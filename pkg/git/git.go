package git

import (
	"context"
	"fmt"
	"time"

	"github.com/curlyotter/schettsFragen-be/pkg/environment"
	"github.com/google/go-github/github"
)

// CreatePullRequest retrieves a github client and a given config to
// create a pull request which adds questions to the configured questions repo
func CreatePullRequest(ctx context.Context, ghClient *github.Client, config map[string]string) error {
	now := time.Now()
	timestamp := now.Unix()

	pullTitle := createPullTitle(timestamp)
	pullBranch := createPullBranch(timestamp)
	pullBase := config[environment.GithubQuestionsRepoBase]
	pullBody := createBullBody()
	modifiable := true

	pull := github.NewPullRequest{
		Title:               &pullTitle,
		Head:                &pullBranch,
		Base:                &pullBase,
		Body:                &pullBody,
		MaintainerCanModify: &modifiable,
	}

	_, _, err := ghClient.PullRequests.Create(
		ctx,
		config[environment.GithubQuestionsRepoOwner],
		config[environment.GithubQuestionsRepoURL],
		&pull,
	)
	if err != nil {
		return err
	}

	return nil
}

// createPullTitle creates a title for the questions pull request
// based on a unix timestamp as id
func createPullTitle(t int64) string {
	return fmt.Sprintf("%d: Questions Contribution", t)
}

// createPullHead creates a branch name based on a unix timestamp
// eg.: 1666454491/new-questions
func createPullBranch(t int64) string {
	return fmt.Sprintf("%d/new-questions", t)
}

// createBullBody creates a body for the pr
func createBullBody() string {
	// TODO add variables like the username who requested the pull to the body
	return fmt.Sprintf("pls review my awesome questions and pull them in! they are good!")
}
