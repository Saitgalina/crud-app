package repository

import (
	"errors"
	"fmt"
	"github.com/Saitgalina/crud-app/internal/core/model"
	"github.com/jmoiron/sqlx"
)

type BookPostgres struct {
	db *sqlx.DB
}

func NewBookPostgres(db *sqlx.DB) *BookPostgres {
	return &BookPostgres{db: db}
}

func (r *BookPostgres) CreateBook(userId int, book model.Book) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var isAdm bool
	//скорее всего проверка на админа должна быть на слое сервис
	checkAdminQuery := fmt.Sprintf("SELECT isAdmin FROM %s WHERE id=$1", usersTable)
	rowIsAdm := tx.QueryRow(checkAdminQuery, userId)
	if err := rowIsAdm.Scan(&isAdm); err != nil {
		tx.Rollback()
		return 0, err
	}
	if !isAdm {
		fmt.Println("USER IS NOT ADMIN")
		return 0, errors.New("USER IS NOT ADMIN")
	}
	var idBook int
	createBookQuery := fmt.Sprintf("INSERT INTO %s (name, author, year, description, source) VALUES ($1, $2, $3,$4, $5) RETURNING id", booksTable)
	rowCreateBook := tx.QueryRow(createBookQuery, book.Name, book.Author, book.Year, book.Description, book.Source)
	if err := rowCreateBook.Scan(&idBook); err != nil {
		tx.Rollback()
		return 0, err
	}
	return idBook, tx.Commit()
}

func (r *BookPostgres) GetAllBooks() ([]model.Book, error) {
	var books []model.Book
	query := fmt.Sprintf("SELECT id, name, author,year, description, source FROM %s", booksTable)
	err := r.db.Select(&books, query)
	return books, err
}

func (r *BookPostgres) GetByIdBook(idBook int) (model.Book, error) {
	var book model.Book
	query := fmt.Sprintf("SELECT id, name, author,year, description, source FROM %s WHERE id=$1", booksTable)
	err := r.db.Get(&book, query, idBook)
	return book, err
}
