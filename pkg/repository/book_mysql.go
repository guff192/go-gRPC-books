package repository

import (
	"fmt"

	"github.com/guff192/go-gRPC-books/pkg/pb"
	"github.com/jmoiron/sqlx"
)

type BookMysql struct {
	db *sqlx.DB
}

func NewBookMysql(db *sqlx.DB) *BookMysql {
	return &BookMysql{db: db}
}

func (r *BookMysql) GetAuthors(bookId int64) ([]*pb.Author, error) {
	var authors []*pb.Author

	query := fmt.Sprintf(`SELECT a.id, a.first_name, a.last_name FROM %s b 
	INNER JOIN %s ba ON b.id=ba.book_id
		INNER JOIN %s a ON ba.author_id=a.id WHERE b.id = ?;`,
		booksTable, booksAuthorsTable, authorsTable)
	err := r.db.Select(&authors, query, bookId)

	return authors, err
}
