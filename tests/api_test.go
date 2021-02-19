package pexels

import (
	"fmt"
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
	results := client.SimpleQuery("horse", 10)
	fmt.Println(results)

}
