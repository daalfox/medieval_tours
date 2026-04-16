package tour

import ()

type Tour struct {
	Title string `json:"title"`
	Desc  string `json:"description"`
}

type TourWithId struct {
	Id   int64
	Tour Tour
}
