package pexels

import (
	"log"
	"net/http"

	"github.com/mozillazg/request"
)

var ()

type PopularService struct {
	Client *PexelAPI
}

func (s *PopularService) GetPopular() []Result {

	c := new(http.Client)
	req := request.NewRequest(c)
	req.Headers = map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer 563492ad6f917000010000018cb81cc7612a426c905f8bd9d1eac770",
	}
	resp, err := req.Get("https://api.pexels.com/v1/search?query=people")
	j, err := resp.Json()
	defer resp.Body.Close()

	log.Println(resp, err, j)

	return make([]Result, 10)
}
