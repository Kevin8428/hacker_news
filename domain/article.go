package domain

// Article is a struct
type Article struct {
	UserID   int    `json:"UserID"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Website  string `json:"Website"`
	Category string `json:"Category"`
	URL      string `json:"url"`
}

type ParentArticle struct {
	Status   string `json:"status"`
	Articles []struct {
		Author string `json:"author"`
		Title  string `json:"title"`
		URL    string `json:url`
	} `json:"articles"`
}
