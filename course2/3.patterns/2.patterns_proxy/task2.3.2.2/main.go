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

type GHAdapterWithCache struct {
	ghadapter  GithubAdapter
	cacheGists map[string][]Item
	cacheRepos map[string][]Item
}

func NewGithub(client *github.Client) *GithubAdapter {
	return &GithubAdapter{client: client}
}

func (gha *GHAdapterWithCache) GetGists(ctx context.Context, username string) ([]Item, error) {
	if res, ok := gha.cacheGists[username]; ok {
		fmt.Println("Используется кэш1")
		return res, nil
	}
	gists, _, err := gha.ghadapter.client.Gists.List(ctx, username, nil)
	if err != nil {
		return nil, err
	}
	var res []Item
	for _, gist := range gists {
		r := Item{Title: *gist.ID, Описание: *gist.Description, Link: *gist.HTMLURL}
		res = append(res, r)
	}
	gha.cacheGists[username] = res
	return res, err
}

func (gha *GHAdapterWithCache) GetRepos(ctx context.Context, username string) ([]Item, error) {
	//opt := &github.RepositoryListOptions{Type: "public"}
	if res, ok := gha.cacheRepos[username]; ok {
		fmt.Println("Используется кэш2")
		return res, nil
	}
	repos, _, err := gha.ghadapter.client.Repositories.List(ctx, username, nil)
	if err != nil {
		return nil, err
	}
	var res []Item
	for _, repo := range repos {
		r := Item{Title: *repo.Name, Описание: *repo.FullName, Link: *repo.URL}
		res = append(res, r)
	}
	gha.cacheRepos[username] = res
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
	gha := &GHAdapterWithCache{ghadapter: *g, cacheGists: make(map[string][]Item), cacheRepos: make(map[string][]Item)}

	fmt.Println(gha.GetGists(context.Background(), "ptflp"))
	fmt.Println(gha.GetGists(context.Background(), "ptflp"))
	fmt.Println(gha.GetRepos(context.Background(), "ptflp"))
	fmt.Println(gha.GetRepos(context.Background(), "ptflp"))
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
