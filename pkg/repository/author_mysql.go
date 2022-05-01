package repository

import (
	"fmt"

	"github.com/guff192/go-gRPC-books/pkg/pb"
	"github.com/jmoiron/sqlx"
)

type AuthorMysql struct {
	db *sqlx.DB
}

func NewAuthorMysql(db *sqlx.DB) *AuthorMysql {
	return &AuthorMysql{db: db}
}

// Repository method for querying books by author's first and last name
func (r *AuthorMysql) GetBooks(first_name, last_name string) ([]*pb.Book, error) {
	var books []*pb.Book

	query := fmt.Sprintf(`SELECT b.id, b.title, b.summary FROM %s a 
	INNER JOIN %s ba ON a.id=ba.author_id
		INNER JOIN %s b ON ba.book_id=b.id 
		WHERE a.first_name = ? AND a.last_name = ?;`,
		authorsTable, booksAuthorsTable, booksTable)
	err := r.db.Select(&books, query, first_name, last_name)

	return books, err
}
