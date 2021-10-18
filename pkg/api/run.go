package api

import (
	"context"

	"github.com/shurcooL/graphql"
)

type Run struct {
	ID        graphql.String
	Status    graphql.String
	CreatedAt graphql.String
}

type Runs struct {
	client *graphql.Client
}

func (rs *Runs) GetById(id string) (*Run, error) {
	variables := map[string]interface{}{
		"id": id,
	}

	var q struct {
		Node *struct {
			ID  graphql.String
			Run Run `graphql:"... on Run"`
		} `graphql:"node(id: $id)"`
	}

	if err := rs.client.Query(context.Background(), &q, variables); err != nil {
		return nil, err
	}

	if q.Node == nil || q.Node.Run.Status == "" {
		return nil, nil
	}

	return &q.Node.Run, nil
}

func (rs *Runs) GetPlanOutputById(id string) (string, error) {
	var q struct {
		Node *struct {
			Run struct {
				Status     string
				PlanOutput string
			} `graphql:"... on Run"`
		} `graphql:"node(id: $id)"`
	}

	variables := map[string]interface{}{
		"id": id,
	}

	if err := rs.client.Query(context.Background(), &q, variables); err != nil {
		return "", err
	}

	if q.Node == nil || q.Node.Run.Status == "" {
		return "", ErrNotExist
	}

	return q.Node.Run.PlanOutput, nil
}
