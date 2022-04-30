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

func (r *AuthorMysql) GetBooks(authorId int64) ([]*pb.Book, error) {
	var books []*pb.Book

	query := fmt.Sprintf(`SELECT b.id, b.title, b.summary FROM %s a 
	INNER JOIN %s ba ON a.id=ba.author_id
		INNER JOIN %s b ON ba.book_id=b.id WHERE a.id = ?;`,
		authorsTable, booksAuthorsTable, booksTable)
	err := r.db.Select(&books, query, authorId)

	return books, err
}
