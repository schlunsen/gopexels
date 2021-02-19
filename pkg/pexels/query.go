package pexels

import (
	"fmt"
	"net/http"

	"github.com/bitly/go-simplejson"
	"github.com/mozillazg/request"
)

func (s *PexelAPI) Query(queryObj map[string][]string) *simplejson.Json {

	c := new(http.Client)
	req := request.NewRequest(c)
	req.Headers = map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer 563492ad6f917000010000018cb81cc7612a426c905f8bd9d1eac770",
	}
	url := fmt.Sprintf("https://api.pexels.com/v1/search?query=%s&per_page=%s", queryObj["Query"][0], queryObj["PerPage"][0])
	fmt.Println(url)
	resp, _ := req.Get(url)
	j, _ := resp.Json()
	defer resp.Body.Close()
	return j
}

func (s *PexelAPI) SimpleQuery(query string, perPage int) *simplejson.Json {
	queryMap := map[string][]string{
		"Query":   {query},
		"PerPage": {fmt.Sprintf("%d", perPage)},
	}
	return s.Query(queryMap)
}
