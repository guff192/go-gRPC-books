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

func (s *server) GetBookAuthors(ctx context.Context, bookRequest *pb.BookAuthorsRequest) (*pb.BookAuthorsResponse, error) {
	authors, err := s.service.GetAuthors(bookRequest.Book)
	if err != nil {
		return nil, err
	}

	return &pb.BookAuthorsResponse{Authors: authors}, nil
}

func (s *server) GetAuthorBooks(ctx context.Context, authorRequest *pb.AuthorBooksRequest) (*pb.AuthorBooksRespponse, error) {
	books, err := s.service.GetBooks(authorRequest.Author)
	if err != nil {
		return nil, err
	}

	return &pb.AuthorBooksRespponse{Books: books}, nil
}

type Config struct {
	Host string
	Port string
}

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
