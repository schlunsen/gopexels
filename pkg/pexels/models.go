package pexels

type Photo struct {
	ID               int64
	Width            int16
	Height           int16
	URL              string
	Photographer     string
	Photographer_URL string
	Photographer_ID  string
	Avg_Color        string
	src              *SrcObj
}

type SrcObj struct {
	Original  string
	Large     string
	Large2x   string
	Medium    string
	Small     string
	Portrait  string
	Landscape string
	Tiny      string
}

type QueryObj struct {
	Query string
}
