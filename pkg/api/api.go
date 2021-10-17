package api

import (
	"errors"

	"github.com/shurcooL/graphql"
)

var (
	ErrNotExist = errors.New("resource does not exist")
)

type API struct {
	Workspaces *Workspaces
	Variables  *Variables
	Runs       *Runs
}

func NewAPI(url string) *API {
	client := graphql.NewClient(url, nil)
	return &API{
		Workspaces: &Workspaces{client},
		Variables:  &Variables{client},
		Runs:       &Runs{client},
	}
}
