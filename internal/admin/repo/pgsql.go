package repo

import (
	"BookStore/internal/book"
	"BookStore/internal/publisher"
	"context"
	"database/sql"
	"fmt"
	"log"
)

type adminRepo struct {
	connPool *sql.DB

	createStmt *sql.Stmt
	updateStmt *sql.Stmt
	deleteStmt *sql.Stmt

	createPublisherStmt *sql.Stmt
	updatePublisherStmt *sql.Stmt
	deletePublisherStmt *sql.Stmt
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

	r.deleteStmt, e = db.Prepare(`
		delete from  store.books where id=$1`)
	if e != nil {
		return nil, e
	}

	r.createPublisherStmt, e = db.Prepare(`
		insert into store.publishers (name, country, phone)
		values ($1, $2, $3)
		returning id`)
	if e != nil {
		return nil, e
	}

	r.updatePublisherStmt, e = db.Prepare(`
		update store.publishers set name=$1, country=$2, phone=$3
		where id=$4`)
	if e != nil {
		return nil, e
	}

	r.deletePublisherStmt, e = db.Prepare(`
		delete from  store.publishers where id=$1`)
	if e != nil {
		return nil, e
	}

	return r, nil
}

func (r *adminRepo) CreateBook(ctx context.Context, info *book.FullInfo) (int64, error) {

	var id int64
	//isbn, title, descr, price, publisher_id, author_id, publication_year, genre
	if e := r.createStmt.QueryRow(info.ISBN, info.Title, info.Descr, info.Price, info.PublisherID, info.AuthorID, info.PublicationYear, info.Genre).Scan(&id); e != nil {
		log.Println("Error update book [", e, "]")
		return 0, fmt.Errorf("error update book %w", e)
	}
	return id, nil
}

func (r *adminRepo) UpdateBook(ctx context.Context, info *book.FullInfo) error {

	//isbn, title, descr, price, publisher_id, author_id, publication_year, genre
	if _, e := r.updateStmt.Exec(info.ISBN, info.Title, info.Descr, info.Price, info.PublisherID, info.AuthorID, info.PublicationYear, info.Genre, info.ID); e != nil {
		log.Println("Error update book [", e, "]")
		return fmt.Errorf("error update book %w", e)
	}
	return nil
}

func (r *adminRepo) DeleteBook(ctx context.Context, id int64) error {

	//id
	if _, e := r.deleteStmt.Exec(id); e != nil {
		log.Println("Error delete book [", e, "]")
		return fmt.Errorf("error delete book %w", e)
	}
	return nil
}

func (r *adminRepo) UpdatePublisher(ctx context.Context, p *publisher.Publisher) error {
	if _, e := r.updatePublisherStmt.Exec(p.Name, p.Country, p.Phone, p.ID); e != nil {
		log.Println("Error update publisher [", e, "]")
		return fmt.Errorf("error update publisher %w", e)
	}
	return nil
}
func (r *adminRepo) CreatePublisher(ctx context.Context, p *publisher.Publisher) (int64, error) {
	var id int64

	if e := r.createPublisherStmt.QueryRow(p.Name, p.Country, p.Phone).Scan(&id); e != nil {
		log.Println("Error update publisher [", e, "]")
		return 0, fmt.Errorf("error update publisher %w", e)
	}
	return id, nil
}
func (r *adminRepo) DeletePublisher(ctx context.Context, id int64) error {
	if _, e := r.deletePublisherStmt.Exec(id); e != nil {
		log.Println("Error delete publisher [", e, "]")
		return fmt.Errorf("error delete publisher %w", e)
	}
	return nil
}
