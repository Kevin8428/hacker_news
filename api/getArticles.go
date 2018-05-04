package api

import (
	"encoding/json"
	"net/http"

	"github.com/kevin8428/hackernews/domain"
)

// type ShowAllArticles handler function
func ShowAllArticles() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		a := []domain.Article{
			{
				Title:    "article 1 some long title name here",
				Author:   "kevin",
				Website:  "cnn",
				Category: "politics",
				URL:      "https://www.google.com",
			},
			{
				Title:    "article 2 some name",
				Author:   "matt",
				Website:  "new york times",
				Category: "health",
				URL:      "https://www.google.com",
			},
			{
				Title:    "article 3 some long title name here",
				Author:   "dave",
				Website:  "forbes",
				Category: "finance",
				URL:      "https://www.google.com",
			},
			{
				Title:    "article 4 something",
				Author:   "ben",
				Website:  "the wall street journal",
				Category: "finance",
				URL:      "https://www.google.com",
			},
			{
				Title:    "article 5 some long title name here",
				Author:   "kevin",
				Website:  "cnn",
				Category: "entertainment",
				URL:      "https://www.google.com",
			},
			{
				Title:    "article 6 some name",
				Author:   "matt",
				Website:  "new york times",
				Category: "current affairs",
				URL:      "https://www.google.com",
			},
			{
				Title:    "article 7 some long title name here",
				Author:   "dave",
				Website:  "forbes",
				Category: "finance",
				URL:      "https://www.google.com",
			},
			{
				Title:    "article 8 some long title name here",
				Author:   "ben",
				Website:  "the wall street journal",
				Category: "finance",
				URL:      "https://www.google.com",
			},
		}
		json.NewEncoder(w).Encode(a)
	})
}
