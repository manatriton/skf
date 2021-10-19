package api

import (
	"context"
	"errors"
	"net/http"

	"github.com/shurcooL/graphql"
	"golang.org/x/oauth2"
)

var (
	ErrNotExist = errors.New("resource does not exist")
)

type API struct {
	Workspaces *Workspaces
	Variables  *Variables
	Runs       *Runs
}

func NewAPI(url, token string) *API {
	var httpClient *http.Client
	if token != "" {
		src := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		httpClient = oauth2.NewClient(context.Background(), src)
	}

	client := graphql.NewClient(url, httpClient)
	return &API{
		Workspaces: &Workspaces{client},
		Variables:  &Variables{client},
		Runs:       &Runs{client},
	}
}
