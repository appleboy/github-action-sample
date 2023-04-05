package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/google/go-github/v50/github"
	"github.com/gookit/goutil/dump"
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

	repo := lastString(strings.Split(os.Getenv("GITHUB_REPOSITORY"), "/"))
	commit, resp, err := client.Repositories.GetCommit(
		ctx,
		os.Getenv("GITHUB_REPOSITORY_OWNER"),
		repo,
		os.Getenv("GITHUB_SHA"),
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	dump.P(commit)
	dump.P(resp)
}

func lastString(ss []string) string {
	return ss[len(ss)-1]
}
