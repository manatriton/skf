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
	WorkspaceID graphql.String  `json:"workspaceId"`
	Key         graphql.String  `json:"key"`
	Value       graphql.String  `json:"value"`
	Sensitive   graphql.Boolean `json:"sensitive"`
}

type DeleteWorkspaceVariableInput struct {
	ID graphql.ID `json:"id"`
}

type UpdateWorkspaceVariableInput struct {
	ID    *graphql.String `json:"id,omitempty"`
	Key   *graphql.String `json:"key,omitempty"`
	Value *graphql.String `json:"value,omitempty"`
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
		} `graphql:"createWorkspaceVariable(input: $input)"`
	}

	if err := vs.client.Mutate(context.Background(), &m, variables); err != nil {
		return nil, err
	}

	return m.CreateWorkspaceVariable.WorkspaceVariable, nil
}

func (vs *Variables) DeleteById(id string) (bool, error) {
	variables := map[string]interface{}{
		"input": DeleteWorkspaceVariableInput{
			ID: id,
		},
	}

	var m struct {
		DeleteWorkspaceVariable struct {
			DeletedWorkspaceVariableId graphql.ID
		} `graphql:"deleteWorkspaceVariable(input: $input)"`
	}

	if err := vs.client.Mutate(context.Background(), &m, variables); err != nil {
		return false, err
	}
	if m.DeleteWorkspaceVariable.DeletedWorkspaceVariableId == "" {
		return false, nil
	}

	return true, nil
}

func (vs *Variables) UpdateById(id, key, value string, sensitive bool) (*Variable, error) {
	var (
		idPtr    *graphql.String
		keyPtr   *graphql.String
		valuePtr *graphql.String
	)
	if id != "" {
		idPtr = (*graphql.String)(&id)
	}

	if key != "" {
		keyPtr = (*graphql.String)(&key)
	}

	if value != "" {
		valuePtr = (*graphql.String)(&value)
	}

	variables := map[string]interface{}{
		"input": UpdateWorkspaceVariableInput{
			ID:    idPtr,
			Key:   keyPtr,
			Value: valuePtr,
		},
	}

	var m struct {
		UpdateWorkspaceVariable struct {
			WorkspaceVariable *Variable
		} `graphql:"updateWorkspaceVariable(input: $input)"`
	}

	if err := vs.client.Mutate(context.Background(), &m, variables); err != nil {
		return nil, err
	}

	return m.UpdateWorkspaceVariable.WorkspaceVariable, nil
}
