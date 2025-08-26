package xledger_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/omniboost/go-xledger"
)

func TestCompanyGetAll(t *testing.T) {

	var result []xledger.Company
	var after *string
	variables := map[string]interface{}{
		"after": after,
	}
	for {
		q := struct {
			Data xledger.QLQuery[xledger.Company] `graphql:"companies(objectStatus: OPEN, first: 5000, after: $after)"`
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

		// there a are to many companies, so we limit to first 5000
		//if variables["after"] == "" {
		break
		//}

	}

	t.Log(len(result))
}

func TestCompanyGet(t *testing.T) {
	variables := map[string]interface{}{
		"taxNumber": "NO997331508",
	}
	q := struct {
		Data xledger.QLQuery[xledger.Company] `graphql:"companies(first: 1, filter: {taxNumber: $taxNumber})"`
	}{}
	err := client.GraphQLClient().Query(context.Background(), &q, variables)
	if err != nil {
		t.Fatal(err)
	}
	if len(q.Data.Edges) == 0 {
		t.Fatal("no company found")
	}
	t.Log(q.Data.Edges[0].Node)

	d, _ := json.Marshal(q.Data.Edges[0].Node)
	t.Log(string(d))
}
