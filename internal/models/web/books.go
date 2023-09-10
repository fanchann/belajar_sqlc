package web

type (
	BookCreateForm struct {
		Title  string `json:"title" validate:"required"`
		Author string `json:"author" validate:"required"`
	}

	BookUpdateForm struct {
		Id     int
		Title  string `json:"title" validate:"required"`
		Author string `json:"author" validate:"required"`
	}
)
