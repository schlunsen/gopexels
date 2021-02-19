package pexels

import (
	"log"
	"testing"
)

var test_apiKey string

func setup() {
	test_apiKey = ""
}

// func Test_ClientCreation(t *testing.T) {
// 	setup()
// 	client := NewClient("Horse")
// 	log.Println((client))
// }

// func Test_GetPopular(t *testing.T) {
// 	setup()
// 	client := PexelAPI{ApiKey: test_apiKey}
// 	fmt.Println(client.PopularService.GetPopular())
// }

func Test_Query(t *testing.T) {
	setup()
	client := PexelAPI{ApiKey: test_apiKey}
	//fmt.Println(client.Query(&QueryObj{Query: "Horse"}), "Horse")
	client.Query(map[string][]string{
		"Query": {"horse"},
	})
}

func Test_SimpleQuery(t *testing.T) {
	setup()
	client := PexelAPI{ApiKey: test_apiKey}
	//fmt.Println(client.Query(&QueryObj{Query: "Horse"}), "Horse")
	results := client.SimpleQuery("horse")

	photos := results.Get("photos")
	//value := gjson.Get(results, "photos")
	for i, photo := range photos.MustArray() {
		ndata, _ := photo.(map[string]interface{})
		n2data, _ := ndata["src"].(map[string]interface{})
		log.Println(i, n2data["original 	"])

	}

	//log.Println(photos, "JPRs")
}
