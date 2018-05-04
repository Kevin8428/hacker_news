package domain

// Article is a struct
type Article struct {
	UserID   int    `json:"UserID"`
	Title    string `json:"Title"`
	Author   string `json:"Author"`
	Website  string `json:"Website"`
	Category string `json:"Category"`
	URL      string `json:"url"`
}
