package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/v56/github"
	"golang.org/x/oauth2"
	"log"
)

type GeneralGithub struct {
	client *github.Client
}

type GithubGist struct {
	client *github.Client
}

func (gg *GithubGist) GetItems(ctx context.Context, username string) ([]Item, error) {
	gists, _, err := gg.client.Gists.List(ctx, username, nil)
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

type GithubRepo struct {
	client *github.Client
}

func (gr *GithubRepo) GetItems(ctx context.Context, username string) ([]Item, error) {
	repos, _, err := gr.client.Repositories.List(ctx, username, nil)
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

func NewGeneralGithub(client *github.Client) *GeneralGithub {
	return &GeneralGithub{client: client}
}

func (g *GeneralGithub) GetItems(ctx context.Context, username string, strategy GithubLister) ([]Item, error) {
	return strategy.GetItems(ctx, username)
}

type GithubLister interface {
	GetItems(ctx context.Context, username string) ([]Item, error)
}

type GeneralGithubLister interface {
	GetItems(ctx context.Context, username string, strategy GithubLister) ([]Item, error)
}

func NewGithubGist(client *github.Client) *GithubGist {
	return &GithubGist{client: client}
}

func NewGithubRepo(client *github.Client) *GithubRepo {
	return &GithubRepo{client: client}
}

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "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"},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	gist := NewGithubGist(client)
	repo := NewGithubRepo(client)

	gg := NewGeneralGithub(client)

	data, err := gg.GetItems(context.Background(), "ptflp", gist)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(data)

	data, err = gg.GetItems(context.Background(), "ptflp", repo)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(data)
}

type Githuber interface {
	GetGists(ctx context.Context, username string) ([]Item, error)
	GetRepos(ctx context.Context, username string) ([]Item, error)
}

type Item struct {
	Title    string
	Описание string
	Link     string
}
