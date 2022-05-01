package repository

import (
	"github.com/guff192/go-gRPC-books/pkg/pb"
	"github.com/jmoiron/sqlx"
)

// Repository for querying books
type Book interface {
	GetAuthors(book_title string) ([]*pb.Author, error)
}

// Repository for querying authors
type Author interface {
	GetBooks(first_name, last_name string) ([]*pb.Book, error)
}

type Repository struct {
	Book
	Author
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Book:   NewBookMysql(db),
		Author: NewAuthorMysql(db),
	}
}
