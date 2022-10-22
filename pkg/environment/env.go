package environment

import (
	"flag"

	"github.com/Clarilab/envi/v2"
	"github.com/pkg/errors"
)

const (
	GithubQuestionsRepoURL   = "GITHUB_QUESTIONS_REPO_URL"
	GithubQuestionsRepoOwner = "GITHUB_QUESTIONS_REPO_OWNER"
	GithubQuestionsRepoBase  = "GITHUB_QUESTIONS_REPO_BASE"
)

func GetEnvvars() (map[string]string, error) {
	const errMsg = "failed to read env vars"

	configPath := flag.String("config", "./config/config.yaml", "path/to/config.yaml")

	enviLoader := envi.NewEnvi()

	err := enviLoader.LoadYAMLFiles(*configPath)
	if err != nil {
		return nil, errors.Wrap(err, errMsg)
	}

	err = enviLoader.EnsureVars(
		GithubQuestionsRepoURL,
		GithubQuestionsRepoOwner,
		GithubQuestionsRepoBase,
	)

	if err != nil {
		return nil, errors.Wrap(err, errMsg)
	}

	return enviLoader.ToMap(), nil
}
