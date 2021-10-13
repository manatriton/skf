package api

import (
	"github.com/shurcooL/graphql"
)

type API struct {
	Workspaces *Workspaces
	Variables  *Variables
}

func NewAPI(url string) *API {
	client := graphql.NewClient(url, nil)
	return &API{
		Workspaces: &Workspaces{client},
		Variables:  &Variables{client},
	}
}
