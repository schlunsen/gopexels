package pexels

import (
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
	results := client.SimpleQuery("horse", 10)
	photos := results.Get("photos")

	for i, photo := range photos.MustArray() {
		ndata, _ := photo.(map[string]interface{})
		n2data, _ := ndata["src"].(map[string]interface{})
		log.Println(i, n2data["original"])

	}

}
