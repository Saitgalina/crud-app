package model

type Book struct {
	Id          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Author      string `json:"author" db:"author"`
	Year        string `json:"year" db:"year"`
	Description string `json:"description" db:"description"`
	Source      string `json:"source" db:"source"`
}

type Favourites struct {
	Id     int
	UserId int
	BookId int
}
