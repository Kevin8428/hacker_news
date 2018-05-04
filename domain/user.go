package domain

type User struct {
	id        int
	FirstName string
	LastName  string
	Password  string
	Articles  []Article
}
