package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/kevin8428/hackernews/domain"
)

// // Article is a struct
// type Article struct {
// 	Name     string `json:"Name"`
// 	Author   string `json:"Author"`
// 	Website  string `json:"Website"`
// 	Category string `json:"Category"`
// }

// GetArticles is a function
func GetArticles() []domain.Article {
	a := []domain.Article{}
	res, err := http.Get("http://localhost:5050/articles")
	if err != nil {
		fmt.Println("error1: ", err)
		panic(err.Error())
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(body, &a)
	if err != nil {
		fmt.Println("unmarshall error: ", err)
	}
	return a
}
