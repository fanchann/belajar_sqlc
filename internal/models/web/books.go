package web

type (
	BookCreateForm struct {
		Title  string `json:"title"`
		Author string `json:"author"`
	}

	BookUpdateForm struct {
		Id     int
		Title  string `json:"title"`
		Author string `json:"author"`
	}
)
