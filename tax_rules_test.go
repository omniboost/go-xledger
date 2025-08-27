package xledger_test

import (
	"context"
	"encoding/json"
	"strconv"
	"testing"

	"github.com/omniboost/go-xledger"
)

func TestTaxRulesGet(t *testing.T) {
	var kinds []xledger.ObjectKind
	variables := map[string]interface{}{
		"after": (*string)(nil),
	}
	for {
		q := struct {
			Data xledger.QLQueryPaginated[xledger.ObjectKind] `graphql:"objectKinds(first: 500, after: $after)"`
		}{}
		err := client.GraphQLClient().Query(context.Background(), &q, variables)
		if err != nil {
			t.Fatal(err)
		}
		variables["after"] = (*string)(nil)
		for _, edge := range q.Data.Edges {
			kinds = append(kinds, edge.Node)
			variables["after"] = edge.Cursor
		}
		if variables["after"] == (*string)(nil) {
			break
		}
	}

	taxRuleId := int64(0)
	for _, k := range kinds {
		if k.Name == "TAX_RULE" {
			taxRuleId, _ = strconv.ParseInt(k.DBId, 10, 64)
		}
	}
	if taxRuleId == 0 {
		t.Fatal("tax rule object kind not found")
	}
	variables["objectKindId"] = taxRuleId
	t.Log("tax rule object kind id:", taxRuleId)
	var result []xledger.ObjectValue
	for {
		q := struct {
			Data xledger.QLQueryPaginated[xledger.ObjectValue] `graphql:"objectValues(filter: {objectKindDbId: $objectKindId}, first: 500, after: $after)"`
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
