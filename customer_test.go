package xledger_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/omniboost/go-xledger"
)

func TestCustomerGetAll(t *testing.T) {

	var result []xledger.Customer
	var after *string
	variables := map[string]interface{}{
		"after": after,
	}
	for {
		q := struct {
			Data xledger.QLQuery[xledger.Customer] `graphql:"customers(ownerSet: CURRENT, objectStatus: OPEN, first: 500, after: $after)"`
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
}

func TestCustomerGet(t *testing.T) {
	variables := map[string]interface{}{
		"code": "10000133",
	}
	q := struct {
		Data xledger.QLQuery[xledger.Customer] `graphql:"customers(first: 1, filter: {code: $code})"`
	}{}
	err := client.GraphQLClient().Query(context.Background(), &q, variables)
	if err != nil {
		t.Fatal(err)
	}
	if len(q.Data.Edges) == 0 {
		t.Fatal("no customer found")
	}
	t.Log(q.Data.Edges[0].Node)

	d, _ := json.Marshal(q.Data.Edges[0].Node)
	t.Log(string(d))
}
