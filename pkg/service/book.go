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

// Service method for getting authors by book
func (s *BookService) GetAuthors(book_title string) ([]*pb.Author, error) {
	return s.repo.GetAuthors(book_title)
}
