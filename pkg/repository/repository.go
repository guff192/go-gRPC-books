package repository

import (
	"github.com/guff192/go-gRPC-books/pkg/pb"
	"github.com/jmoiron/sqlx"
)

type Book interface {
	GetAuthors(bookId int64) ([]*pb.Author, error)
}

type Author interface {
	GetBooks(authorId int64) ([]*pb.Book, error)
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
