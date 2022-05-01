package service

import (
	"github.com/guff192/go-gRPC-books/pkg/pb"
	"github.com/guff192/go-gRPC-books/pkg/repository"
)

type AuthorService struct {
	repo repository.Author
}

func NewAuthorService(repo repository.Author) *AuthorService {
	return &AuthorService{repo: repo}
}

// Service method for getting books by author
func (s *AuthorService) GetBooks(first_name, last_name string) ([]*pb.Book, error) {
	return s.repo.GetBooks(first_name, last_name)
}
