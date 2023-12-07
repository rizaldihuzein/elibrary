package types

type (
	BookAPIResponse struct {
		Key   string    `json:"key"`
		Name  string    `json:"name"`
		Books []RawBook `json:"works"`
	}

	RawBook struct {
		Key     string      `json:"key"`
		Title   string      `json:"title"`
		Edition string      `json:"cover_edition_key"`
		Authors []RawAuthor `json:"authors"`
	}

	RawAuthor struct {
		Key  string `json:"key"`
		Name string `json:"name"`
	}
)
