package xledger_test

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/omniboost/go-xledger"
)

var client *xledger.Client

func TestMain(m *testing.M) {
	var err error

	baseURLString := os.Getenv("BASE_URL")
	token := os.Getenv("TOKEN")
	debug := os.Getenv("DEBUG")
	var baseURL *url.URL

	if baseURLString != "" {
		baseURL, err = url.Parse(baseURLString)
		if err != nil {
			log.Fatal(err)
		}
	}

	client = xledger.NewClient(http.DefaultClient)
	client.SetToken(token)
	if debug != "" {
		client.SetDebug(true)
	}

	if baseURL != nil {
		client.SetBaseURL(*baseURL)
	}

	client.SetDisallowUnknownFields(true)
	m.Run()
}
