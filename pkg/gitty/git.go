package gitty

import (
	"context"
	"fmt"
	"time"

	"github.com/curlyotter/schettsFragen-be/pkg/environment"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/google/go-github/github"
	"github.com/rs/zerolog/log"
)

// Init initializes the git and github workflow. It retrieves a github client and a given config to
// clone the questions repo, add the questions yaml, commits and pushes it and creates a PR
func Init(ctx context.Context, ghClient *github.Client, config map[string]string) error {
	now := time.Now()
	timestamp := now.Unix()

	log.Info().Msg("clone questions repository")
	r, err := cloneQuestionsRepo(config)
	if err != nil {
		return err
	}

	commit, err := commitQuestions(r, config, now)
	if err != nil {
		return err
	}
	fmt.Println(commit)

	// TODO Push commit to github

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

	log.Info().Msg("create pull request to remote origin")
	_, _, err = ghClient.PullRequests.Create(
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

func cloneQuestionsRepo(config map[string]string) (*git.Repository, error) {
	r, err := git.Clone(memory.NewStorage(), memfs.New(), &git.CloneOptions{
		URL: config[environment.GithubQuestionsRepoURL],
	})
	if err != nil {
		return nil, err
	}

	return r, nil
}

func commitQuestions(r *git.Repository, c map[string]string, now time.Time) (string, error) {
	w, err := r.Worktree()
	if err != nil {
		return "", err
	}

	log.Info().Msg("add questions yaml to worktree")
	path := c[environment.GithubQuestionsPathToYAML]
	_, err = w.Add(path)
	if err != nil {
		return "", err
	}

	log.Info().Msg("commit changes")
	msg := fmt.Sprintf("add questions to %s", path)
	commit, err := w.Commit(msg, &git.CommitOptions{
		Author: &object.Signature{
			Name:  "curlyBOTter",
			Email: "john@doe.org",
			When:  now,
		},
	})
	if err != nil {
		return "", err
	}

	return commit.String(), nil
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
