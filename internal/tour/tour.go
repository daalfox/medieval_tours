package tour

type Tour struct {
	Title string `json:"title"`
	Desc  string `json:"description"`
}

type TourWithId struct {
	Id   int64 `json:"id"`
	Tour Tour  `json:"tour"`
}
