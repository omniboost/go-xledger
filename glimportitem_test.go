package xledger_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/omniboost/go-xledger"
)

func TestGLImportItemsGet(t *testing.T) {
	q := struct {
		Data xledger.QLQueryPaginated[xledger.GLImportItem] `graphql:"glImportItems(first: 1)"`
	}{}
	err := client.GraphQLClient().Query(context.Background(), &q, nil)
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

func TestGLImportItemsAdd(t *testing.T) {
	q := struct {
		Data xledger.QLQuery[xledger.GLImportItem] `graphql:"addGLImportItems(inputs: {node :$node})"`
	}{}

	variables := map[string]interface{}{
		"node": xledger.GLImportItemInput{
			Text: "Omniboost TEST",
		},
	}
	err := client.GraphQLClient().Mutate(context.Background(), &q, variables)
	if err != nil {
		t.Fatal(err)
	}

	d, _ := json.Marshal(q.Data.Edges[0].Node)
	t.Log(string(d))
}
