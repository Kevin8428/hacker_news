package domain

type User struct {
	ID         int
	FirstName  string
	LastName   string
	Password   string
	IsLoggedIn bool
	Articles   []Article
}
