package service

import (
	"github.com/guff192/go-gRPC-books/pkg/pb"
	"github.com/guff192/go-gRPC-books/pkg/repository"
)

type Book interface {
	GetAuthors(book *pb.Book) ([]*pb.Author, error)
}

type Author interface {
	GetBooks(author *pb.Author) ([]*pb.Book, error)
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
