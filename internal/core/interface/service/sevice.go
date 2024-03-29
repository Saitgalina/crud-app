package service

import (
	"github.com/Saitgalina/crud-app/internal/core/interface/repository"
	"github.com/Saitgalina/crud-app/internal/core/model"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(login, password string) (string, error)
	ParseToken(token string) (int, error)
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

type Sevice struct {
	Authorization
	Book
}

func NewService(repos *repository.Repository) *Sevice {
	return &Sevice{
		Authorization: NewAuthService(repos.Authorization),
		Book:          NewBookService(repos.Book),
	}
}
