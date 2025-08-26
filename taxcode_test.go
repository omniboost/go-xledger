package xledger_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/hasura/go-graphql-client"
	"github.com/omniboost/go-xledger"
	"github.com/pkg/errors"
)

func TestTaxCodeGetAll(t *testing.T) {
	q := struct {
		AddGLImportItems struct {
			Edges []struct {
				Node struct {
					DBId int `graphql:"dbId"`
				} `graphql:"node"`
			} `graphql:"edges"`
		} `graphql:"addGLImportItems(inputs: { node: { taxRule: { showAlternatives: YES } } })"` // Replace with actual query
	}{}

	var taxCodes []xledger.TaxCode

	err := client.GraphQLClient().Mutate(context.Background(), &q, nil)
	if err != nil {
		var gerr graphql.Errors
		if errors.As(err, &gerr) {
			for _, e := range gerr {
				exampleValidInputs, ok := e.Extensions["exampleValidInputs"]
				if !ok {
					continue
				}
				for _, v := range exampleValidInputs.([]any) {
					vd := v.(map[string]any)
					taxCodes = append(taxCodes, xledger.TaxCode{
						Code:      vd["code"].(string),
						DBId:      int(vd["dbId"].(float64)),
						OwnerDbId: int(vd["ownerDbId"].(float64)),
						Text:      vd["text"].(string),
					})
				}
			}
		} else {
			t.Fatal(err)
		}
	}
	d, _ := json.MarshalIndent(taxCodes, "", "  ")
	t.Log(string(d))
}
