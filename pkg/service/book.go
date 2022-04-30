package service

import (
	"github.com/guff192/go-gRPC-books/pkg/pb"
	"github.com/guff192/go-gRPC-books/pkg/repository"
)

type BookService struct {
	repo repository.Book
}

func NewBookService(repo repository.Book) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) GetAuthors(book *pb.Book) ([]*pb.Author, error) {
	return s.repo.GetAuthors(book.Id)
}
