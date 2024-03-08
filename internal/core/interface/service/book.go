package service

import (
	"github.com/Saitgalina/crud-app/internal/core/interface/repository"
	"github.com/Saitgalina/crud-app/internal/core/model"
)

type BookService struct {
	repo repository.Book
}

func NewBookService(repo repository.Book) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) CreateBook(userId int, book model.Book) (int, error) {
	return s.repo.CreateBook(userId, book)
}

func (s *BookService) GetAllBooks() ([]model.Book, error) {
	return s.repo.GetAllBooks()
}
func (s *BookService) GetByIdBook(idBook int) (model.Book, error) {
	return s.repo.GetByIdBook(idBook)
}

func (s *BookService) GetByNameBook(nameBook string) ([]model.Book, error) {
	return s.repo.GetByNameBook(nameBook)
}

func (s *BookService) GetByYearBook(yearBook string) ([]model.Book, error) {
	return s.repo.GetByYearBook(yearBook)
}

func (s *BookService) GetByAuthorBook(authorBook string) ([]model.Book, error) {
	return s.repo.GetByAuthorBook(authorBook)
}
func (s *BookService) GetSortDescBook(valueSort string) ([]model.Book, error) {
	return s.repo.GetSortDescBook(valueSort)
}
func (s *BookService) GetSortAscBook(valueSort string) ([]model.Book, error) {
	return s.repo.GetSortAscBook(valueSort)
}

func (s *BookService) AddFavouritesBook(idBook, idUser int) (string, error) {
	return s.repo.AddFavouritesBook(idBook, idUser)
}

func (s *BookService) GetFavouritesBooks(idUser int) ([]model.Book, error) {
	return s.repo.GetFavouritesBooks(idUser)
}
