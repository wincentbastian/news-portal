package domain

type NewsArticle struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	Source      struct {
		Name string `json:"name"`
	} `json:"source"`
}
