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