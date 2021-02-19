package pexels

import (
	"fmt"
	"log"
	"testing"

	"github.com/schlunsen/gopexels/pkg/pexels"
	"github.com/spf13/viper"
)

var test_apiKey string

func setup() {
	test_apiKey = viper.GetString("APIKEY")
}

func Test_Query(t *testing.T) {
	setup()
	client := pexels.NewClient(test_apiKey)

	client.Query(map[string][]string{
		"Query":   {"horse"},
		"PerPage": {"3"},
	})
}

func Test_SimpleQuery(t *testing.T) {
	setup()
	client := pexels.NewClient(test_apiKey)
	numberOfResults := 10
	results := client.SimpleQuery("horse", numberOfResults)

	// Assert results equals requested
	if len(results) != numberOfResults {
		log.Fatalln("Error with results. Got ", len(results), "Expected ", numberOfResults)
	}
}

func Test_GetPopular(t *testing.T) {
	setup()
	client := pexels.NewClient(test_apiKey)
	results := client.PopularService.GetPopular()
	fmt.Println(results)
}
