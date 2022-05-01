package service

import (
	"github.com/guff192/go-gRPC-books/pkg/pb"
	"github.com/guff192/go-gRPC-books/pkg/repository"
)

// Service for working with books
type Book interface {
	GetAuthors(book_title string) ([]*pb.Author, error)
}

// Service for working with authors
type Author interface {
	GetBooks(first_name, last_name string) ([]*pb.Book, error)
}

type Service struct {
	Book
	Author
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Book:   NewBookService(repos.Book),
		Author: NewAuthorService(repos.Author),
	}
}
