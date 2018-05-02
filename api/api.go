package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Article is a struct
type Article struct {
	Name    string `json:"Name"`
	Author  string `json:"Author"`
	Website string `json:"Website"`
}

// GetArticles is a function
func GetArticles() []Article {
	a := []Article{}
	res, err := http.Get("http://localhost:5050/articles")
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	err = json.Unmarshal(body, &a)
	if err != nil {
		fmt.Println("unmarshall error: ", err)
	}
	return a
}
