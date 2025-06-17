package repo

import (
	"BookStore/internal/book"
	"context"
	"database/sql"
	"fmt"
	"log"
)

type adminRepo struct {
	connPool *sql.DB

	createStmt *sql.Stmt
	updateStmt *sql.Stmt
}

func NewAdminRepo(db *sql.DB) (AdminRepo, error) {

	r := &adminRepo{connPool: db}
	var e error

	r.createStmt, e = db.Prepare(`
		insert into store.books (isbn, title, descr, price, publisher_id, author_id, publication_year, genre)
		values ($1, $2, $3, $4, $5, $6, $7, $8)
		returning id`)
	if e != nil {
		return nil, e
	}

	r.updateStmt, e = db.Prepare(`
		update store.books set isbn=$1, title=$2, descr=$3, price=$4, publisher_id=$5, author_id=$6, publication_year=$7, genre=$8
		where id=$9`)
	if e != nil {
		return nil, e
	}
	return r, nil
}

func (r *adminRepo) CreateBook(ctx context.Context, info *book.FullInfo) (int64, error) {

	var id int64
	//isbn, title, descr, price, publisher_id, author_id, publication_year, genre
	if e := r.createStmt.QueryRow(info.ISBN, info.Title, info.Descr, info.Price, 1, 1, info.PublicationYear, info.Genre).Scan(&id); e != nil {
		log.Println("Login", "Error update book [", e, "]")
		return 0, fmt.Errorf("error update book %w", e)
	}
	return id, nil
}

func (r *adminRepo) UpdateBook(ctx context.Context, info *book.FullInfo) error {

	//isbn, title, descr, price, publisher_id, author_id, publication_year, genre
	if _, e := r.updateStmt.Exec(info.ISBN, info.Title, info.Descr, info.Price, 1, 1, info.PublicationYear, info.Genre, info.ID); e != nil {
		log.Println("Login", "Error update book [", e, "]")
		return fmt.Errorf("error update book %w", e)
	}
	return nil
}
