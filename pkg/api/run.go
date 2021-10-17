package api

import (
	"context"

	"github.com/shurcooL/graphql"
)

type Run struct {
}

type Runs struct {
	client *graphql.Client
}

func (rs *Runs) GetPlanOutput(id string) (string, error) {
	var q struct {
		Node *struct {
			ID  graphql.String
			Run struct {
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

	if q.Node == nil {
		return "", ErrNotExist
	}

	return q.Node.Run.PlanOutput, nil
}
