package web

type (
	BookResponse struct {
		Id     int    `json:"id"`
		Title  string `json:"title"`
		Author string `json:"author"`
	}
)
