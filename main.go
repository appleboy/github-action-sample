package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/google/go-github/v50/github"
	"github.com/gregjones/httpcache"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tc := &http.Client{
		Transport: &oauth2.Transport{
			Base:   httpcache.NewMemoryCacheTransport(),
			Source: ts,
		},
	}
	client := github.NewClient(tc)

	repos, _, err := client.Repositories.List(ctx, "", &github.RepositoryListOptions{
		Sort: "updated",
		Type: "owner",
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, repo := range repos {
		log.Println(repo.GetName())
	}
}
