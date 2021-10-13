package api

import (
	"context"
	"time"

	"github.com/shurcooL/graphql"
)

type Workspace struct {
	ID        string
	Name      string
	CreatedAt time.Time
}

type Workspaces struct {
	client *graphql.Client
}

type CreateWorkspaceInput struct {
	Name string `json:"name"`
}

func (ws *Workspaces) WorkspaceByName(name string) (*Workspace, error) {
	if name == "" {
		return nil, nil
	}

	variables := map[string]interface{}{
		"name": graphql.String(name),
	}

	var q struct {
		WorkspaceByName *Workspace `graphql:"workspaceByName(name: $name)"`
	}

	if err := ws.client.Query(context.Background(), &q, variables); err != nil {
		return nil, err
	}

	return q.WorkspaceByName, nil
}

func (ws *Workspaces) CreateWorkspace(name string) (*Workspace, error) {
	variables := map[string]interface{}{
		"input": CreateWorkspaceInput{
			Name: name,
		},
	}

	var m struct {
		CreateWorkspace struct {
			Workspace *Workspace
		} `graphql:"createWorkspace(input: $input)"`
	}

	if err := ws.client.Mutate(context.Background(), &m, variables); err != nil {
		return nil, err
	}

	return m.CreateWorkspace.Workspace, nil
}
