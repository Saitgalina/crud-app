package repository

import (
	"github.com/Saitgalina/crud-app/internal/core/model"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(login, password string) (model.User, error)
}

type Book interface {
	CreateBook(userId int, book model.Book) (int, error)
	GetAllBooks() ([]model.Book, error)
	GetByIdBook(idBook int) (model.Book, error)
	GetByNameBook(nameBook string) ([]model.Book, error)
	GetByYearBook(yearBook string) ([]model.Book, error)
	GetByAuthorBook(authorBook string) ([]model.Book, error)
	GetSortDescBook(valueSort string) ([]model.Book, error)
	GetSortAscBook(valueSort string) ([]model.Book, error)
	AddFavouritesBook(idBook, idUser int) (string, error)
	GetFavouritesBooks(idUser int) ([]model.Book, error)
}

type Repository struct {
	Authorization
	Book
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Book:          NewBookPostgres(db),
	}
}
