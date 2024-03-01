package main

import (
	"context"
	"github.com/google/go-github/v56/github"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestGithubAdapter_GetGists(t *testing.T) {

	type args struct {
		ctx      context.Context
		username string
	}
	tests := []struct {
		name    string
		args    args
		want    []Item
		wantErr bool
	}{
		{
			name: "Case1",
			args: args{
				ctx:      context.Background(),
				username: "qwe",
			},
			want:    make([]Item, 0),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewMockGithuber(t)
			a.On("GetGists", tt.args.ctx, tt.args.username).Return(tt.want, nil)
			got, err := a.GetGists(tt.args.ctx, tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetGists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGithubAdapter_GetRepos(t *testing.T) {

	type args struct {
		ctx      context.Context
		username string
	}
	tests := []struct {
		name    string
		args    args
		want    []Item
		wantErr bool
	}{
		{
			name: "Case1",
			args: args{
				ctx:      context.Background(),
				username: "asd",
			},
			want:    make([]Item, 0),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewMockGithuber(t)
			g.On("GetRepos", tt.args.ctx, tt.args.username).Return(tt.want, nil)
			got, err := g.GetRepos(tt.args.ctx, tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRepos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestNewGithub(t *testing.T) {
	type args struct {
		cl *github.Client
	}

	tests := []struct {
		name string
		args args
		want *GithubAdapter
	}{
		{
			name: "Case1",
			args: args{cl: &github.Client{}},
			want: &GithubAdapter{
				client: &github.Client{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGithub(tt.args.cl); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGithub() = %v, want %v", got, tt.want)
			}
		})
	}
}
