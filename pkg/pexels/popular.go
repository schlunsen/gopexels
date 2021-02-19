package pexels

import (
	"fmt"
	"net/http"

	"github.com/mozillazg/request"
)

var ()

type PopularService struct {
	Client *PexelAPI
}

func (s *PopularService) GetPopular() []string {

	c := new(http.Client)
	req := request.NewRequest(c)
	req.Headers = map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer 563492ad6f917000010000018cb81cc7612a426c905f8bd9d1eac770",
	}
	resp, _ := req.Get("https://api.pexels.com/v1/curated?per_page=20")
	j, _ := resp.Json()
	defer resp.Body.Close()
	var results []string
	for _, photo := range j.Get("photos").MustArray() {
		ndata, _ := photo.(map[string]interface{})
		n2data, _ := ndata["src"].(map[string]interface{})
		results = append(results, fmt.Sprintf("%s", n2data["large2x"]))
	}

	return results

}
