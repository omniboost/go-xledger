package xledger_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/omniboost/go-xledger"
)

func TestCurrencyGetAll(t *testing.T) {
	var result []xledger.Currency
	var after *string
	variables := map[string]interface{}{
		"after": after,
	}
	for {
		q := struct {
			Data xledger.QLQuery[xledger.Currency] `graphql:"currencies(first: 500, after: $after)"`
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

	t.Log(len(result))

	d, _ := json.Marshal(result)
	t.Log(string(d))
}
