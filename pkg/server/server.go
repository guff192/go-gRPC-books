package server

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/guff192/go-gRPC-books/pkg/pb"
	"github.com/guff192/go-gRPC-books/pkg/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	*pb.UnimplementedBookServiceServer
	service *service.Service
}

// Gets book title as a request and returning list of its authors
func (s *server) GetBookAuthors(ctx context.Context, baRequest *pb.BookAuthorsRequest) (*pb.BookAuthorsResponse, error) {
	authors, err := s.service.GetAuthors(baRequest.BookTitle)
	if err != nil {
		return nil, err
	}

	return &pb.BookAuthorsResponse{Authors: authors}, nil
}

// Gets author's first and last name as a request and returning list of his books
func (s *server) GetAuthorBooks(ctx context.Context, abRequest *pb.AuthorBooksRequest) (*pb.AuthorBooksRespponse, error) {
	books, err := s.service.GetBooks(abRequest.FirstName, abRequest.LastName)
	if err != nil {
		return nil, err
	}

	return &pb.AuthorBooksRespponse{Books: books}, nil
}

type Config struct {
	Host string
	Port string
}

// Runs server on tcp host:port, given by config
func RunServer(cfg Config, service *service.Service) error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Host, cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb.RegisterBookServiceServer(grpcServer, &server{service: service})

	if err := grpcServer.Serve(lis); err != nil {
		return err
	}

	return nil
}
