package main

import (
	"context"
	"log"

	"github.com/curlyotter/schettsFragen-be/pkg/environment"
	"github.com/curlyotter/schettsFragen-be/pkg/gitty"
	"github.com/curlyotter/schettsFragen-be/pkg/question"
	"github.com/curlyotter/schettsFragen-be/pkg/writer"
	"github.com/google/go-github/github"
)

func main() {
	ctx := context.Background()

	config, err := environment.GetEnvvars()
	if err != nil {
		log.Fatalln(err)
	}

	ghClient := github.NewClient(nil)

	question.Questions = question.Add(
		"Wie viele Einwohner hat Andalusien (in millionen)",
		8,
	)

	if err = writer.QuestionsToYAML(question.Questions); err != nil {
		log.Fatalln(err)
	}

	if err = gitty.Init(ctx, ghClient, config); err != nil {
		log.Fatalln(err)
	}
}
