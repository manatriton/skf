package api

import (
	"context"

	"github.com/shurcooL/graphql"
)

type Variable struct {
	ID          graphql.String
	WorkspaceID graphql.String
	Key         graphql.String
	Value       graphql.String
	Sensitive   graphql.Boolean
}

type Variables struct {
	client *graphql.Client
}

type CreateWorkspaceVariableInput struct {
	WorkspaceID graphql.String
	Key         graphql.String
	Value       graphql.String
	Sensitive   graphql.Boolean
}

type UpdateVariableInput struct {
	ID    string
	Key   string
	Value string
}

func (vs *Variables) GetVariable(id string) (*Variable, error) {
	var q struct {
		Node struct {
			ID                graphql.String
			WorkspaceVariable Variable `graphql:"... on WorkspaceVariable"`
		} `graphql:"node(id: $id)"`
	}

	variables := map[string]interface{}{
		"id": id,
	}

	if err := vs.client.Query(context.Background(), &q, variables); err != nil {
		return nil, err
	}

	v := q.Node.WorkspaceVariable
	v.ID = q.Node.ID
	return &v, nil
}

func (vs *Variables) Create(workspaceId, key, value string, sensitive bool) (*Variable, error) {
	variables := map[string]interface{}{
		"input": CreateWorkspaceVariableInput{
			WorkspaceID: graphql.String(workspaceId),
			Key:         graphql.String(key),
			Value:       graphql.String(value),
			Sensitive:   graphql.Boolean(sensitive),
		},
	}

	var m struct {
		CreateWorkspaceVariable struct {
			WorkspaceVariable *Variable
		} `graphql:"createWorkspace(input: $input)"`
	}

	if err := vs.client.Mutate(context.Background(), &m, variables); err != nil {
		return nil, err
	}

	return m.CreateWorkspaceVariable.WorkspaceVariable, nil
}
