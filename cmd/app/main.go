package main

import (
	"context"

	"github.com/curlyotter/schettsFragen-be/pkg/environment"
	"github.com/curlyotter/schettsFragen-be/pkg/gitty"
	"github.com/curlyotter/schettsFragen-be/pkg/question"
	"github.com/curlyotter/schettsFragen-be/pkg/writer"
	"github.com/google/go-github/github"
	"github.com/rs/zerolog/log"
)

func main() {
	ctx := context.Background()

	log.Info().Msg("read env vars from config")
	config, err := environment.GetEnvvars()
	if err != nil {
		log.Error().Msg(err.Error())
	}

	ghClient := github.NewClient(nil)

	question.Questions = question.Add(
		"Wie viele Einwohner hat Andalusien (in millionen)",
		8,
	)

	log.Info().Msg("add questions to yaml file")
	if err = writer.QuestionsToYAML(question.Questions); err != nil {
		log.Error().Msg(err.Error())
	}

	log.Info().Msg("initialize git flow")
	if err = gitty.Init(ctx, ghClient, config); err != nil {
		log.Error().Msg(err.Error())
	}
}
