package xledger_test

import (
	"context"
	"testing"

	"github.com/omniboost/go-xledger"
)

func TestAccountStructure(t *testing.T) {
	r := xledger.NewGraphQLRequest[any](client)
	body := r.RequestBody()
	body.Query = `{
  __schema {
    types {
      name
      fields {
        name
        description
		type {
          name
          kind
		  fields {
			name
			description
			type {
			  name
			  kind
			}
		  }
        }
		 
      }
    }
  }
}`
	_, err := r.Do(context.Background())
	if err != nil {
		t.Error(err)
	}
	//t.Log(res)
}

func TestAccountGetAll(t *testing.T) {
	var result []xledger.Account
	var after *string
	variables := map[string]interface{}{
		"after": after,
	}
	for {
		q := struct {
			Data xledger.QLQuery[xledger.Account] `graphql:"accounts(ownerSet: LOWER, first: 500, after: $after)"`
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

	//d, _ := json.Marshal(result)
	//t.Log(string(d))
}
