package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/v56/github"
	"golang.org/x/oauth2"
)

type Github struct {
	item   Item
	client *github.Client
}
type GithubAdapter struct {
	client *github.Client
}

func NewGithub(client *github.Client) *GithubAdapter {
	return &GithubAdapter{client: client}
}

func (ga *GithubAdapter) GetGists(ctx context.Context, username string) ([]Item, error) {
	gists, _, err := ga.client.Gists.List(ctx, username, nil)
	if err != nil {
		return nil, err
	}
	var res []Item
	for _, gist := range gists {
		r := Item{Title: *gist.ID, Описание: *gist.Description, Link: *gist.HTMLURL}
		res = append(res, r)
	}
	return res, err
}

func (ga *GithubAdapter) GetRepos(ctx context.Context, username string) ([]Item, error) {
	//opt := &github.RepositoryListOptions{Type: "public"}
	repos, _, err := ga.client.Repositories.List(ctx, username, nil)
	if err != nil {
		return nil, err
	}
	var res []Item
	for _, repo := range repos {
		r := Item{Title: *repo.Name, Описание: *repo.FullName, Link: *repo.URL}
		res = append(res, r)
	}
	return res, err
}

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "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"},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	g := NewGithub(client)

	fmt.Println(g.GetGists(context.Background(), "ptflp"))
	fmt.Println(g.GetRepos(context.Background(), "ptflp"))
}

//go:generate mockery --name Githuber --inpackage
type Githuber interface {
	GetGists(ctx context.Context, username string) ([]Item, error)
	GetRepos(ctx context.Context, username string) ([]Item, error)
}

type Item struct {
	Title    string
	Описание string
	Link     string
}
