package pexels

type PexelAPIInterFace interface {
}

type PexelAPI struct {
	ApiKey         string
	PopularService *PopularService
}

type Result struct {
	url string
}

func NewClient(apiKey string) *PexelAPI {
	client := &PexelAPI{ApiKey: apiKey}
	client.PopularService = &PopularService{Client: client}

	return client
}
