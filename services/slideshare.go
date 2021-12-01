package services

type BinResp struct {
	Slideshow Slideshow
}

// auto case-conversion
// override as needed
type Slideshow struct {
	Presenter string `json:"author"`
	Date      string
	Slides    []Slide
}

type Slide struct {
	Title string
	Type  string
}
