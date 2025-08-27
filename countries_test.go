package xledger_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/omniboost/go-xledger"
)

func TestCountriesGetAll(t *testing.T) {

	var result []xledger.Country
	var after *string
	variables := map[string]interface{}{
		"after": after,
	}
	for {
		q := struct {
			Data xledger.QLQueryPaginated[xledger.Country] `graphql:"countries(first: 5000, after: $after)"`
		}{}
		err := client.GraphQLClient().Query(context.Background(), &q, variables)
		if err != nil {
			t.Fatal(err)
		}
		variables["after"] = ""
		for _, edge := range q.Data.Edges {
			result = append(result, edge.Node)
			variables["after"] = edge.Cursor
		}

		if variables["after"] == "" {
			break
		}
	}

	d, _ := json.MarshalIndent(result, "", "  ")
	t.Log(string(d))
}
