package xledger_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/hasura/go-graphql-client"
	"github.com/omniboost/go-xledger"
)

func TestSubLedgerGetAll(t *testing.T) {
	var result []xledger.SubLedger
	var after *string
	variables := map[string]interface{}{
		"after": after,
	}
	for {
		q := struct {
			Data xledger.QLQueryPaginated[xledger.SubLedger] `graphql:"subledgers(first: 500, after: $after)"`
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

func TestSubLedgerGet(t *testing.T) {
	wanted := 24934231
	variables := map[string]interface{}{
		"dbId": graphql.ToID(wanted),
	}
	q := struct {
		Data *xledger.SubLedger `graphql:"subledger(dbId: $dbId)"`
	}{}
	err := client.GraphQLClient().Query(context.Background(), &q, variables)
	if err != nil {
		t.Fatal(err)
	}
	if q.Data.DBId != wanted {
		t.Fatalf("subledger not found, wanted %d, got %d", wanted, q.Data.DBId)
	}
	d, _ := json.Marshal(q.Data)
	t.Log(string(d))
}
