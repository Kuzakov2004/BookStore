package repo

import (
	"BookStore/internal/book"
	"context"
	"database/sql"
	"fmt"
	"log"
)

type bookRepo struct {
	connPool *sql.DB

	queryBooksStmt     *sql.Stmt
	getBooksCntStmt    *sql.Stmt
	queryFindBooksStmt *sql.Stmt
	getBookInfoStmt    *sql.Stmt
}

func NewBookRepo(db *sql.DB) (BookRepo, error) {

	r := &bookRepo{connPool: db}
	var e error

	r.queryBooksStmt, e = db.Prepare(`SELECT b.id, b.isbn, b.title, b.price, p.name publisher, a.first_name || ' ' || a.last_name author, b.publication_year, b.genre 
											FROM store.books b 
											JOIN store.publishers p ON (b.publisher_id = p.id)
											JOIN store.authors a ON (b.author_id = a.id)
											WHERE b.genre := $1 ORDER BY b.title DESC
    										LIMIT $2 OFFSET $3`)
	if e != nil {
		return nil, e
	}

	r.getBooksCntStmt, e = db.Prepare(`SELECT count(*) 
											FROM store.books b 
											JOIN store.publishers p ON (b.publisher_id = p.id)
											JOIN store.authors a ON (b.author_id = a.id)
											WHERE b.genre := $1`)
	if e != nil {
		return nil, e
	}

	r.queryFindBooksStmt, e = db.Prepare(`SELECT b.id, b.isbn, b.title, b.price, p.name publisher, a.first_name || ' ' || a.last_name author, b.publication_year, b.genre 
											FROM store.books b 
											JOIN store.publishers p ON (b.publisher_id = p.id)
											JOIN store.authors a ON (b.author_id = a.id)
											WHERE UPPER(b.genre) like $1 OR UPPER(b.isbn) like $2 OR UPPER(b.title) like $3 OR UPPER(p.name) like $4 OR UPPER(a.first_name || ' ' || a.last_name) like $5 
											ORDER BY b.title DESC
    										LIMIT 100`)
	if e != nil {
		return nil, e
	}

	r.getBookInfoStmt, e = db.Prepare(`SELECT b.id, b.isbn, b.title, b.price, p.name publisher, a.first_name || ' ' || a.last_name author, b.publication_year, b.genre,
       												b.descr, (select count(*) from store.warehouse_books wb where wb.book_id = b.id)
											FROM store.books b 
											JOIN store.publishers p ON (b.publisher_id = p.id)
											JOIN store.authors a ON (b.author_id = a.id)
											WHERE b.id := $1 
											ORDER BY b.title DESC
    										LIMIT 1`)
	if e != nil {
		return nil, e
	}

	return r, nil
}

func (r *bookRepo) GetBooks(ctx context.Context, genre string, page, count int) (lst []*book.Book, e error) {
	lst = make([]*book.Book, 0, count)

	if count < 10 {
		count = 10
	}

	if page < 0 {
		page = 0
	}

	rows, e := r.queryBooksStmt.Query(genre, page, count)
	if e != nil {
		return nil, e
	}

	defer rows.Close()

	for rows.Next() {
		//b.id, b.isbn, b.title, b.price, publisher, author, b.publication_year, b.genre
		b := &book.Book{}
		if e = rows.Scan(&b.ID, &b.ISBN, &b.Title, &b.Price, &b.Publisher, &b.Author, &b.PublicationYear, &b.Genre); e != nil {
			log.Println("GetBooks", "Error get books [", e, "]")
			return nil, fmt.Errorf("error get books %w", e)
		}
		lst = append(lst, b)
	}

	return lst, nil
}

func (r *bookRepo) GetBooksCnt(ctx context.Context, genre string) (total int, e error) {
	if e = r.getBooksCntStmt.QueryRow(genre).Scan(&total); e != nil {
		log.Println("GetBooks", "Error get books total [", e, "]")
		return 0, fmt.Errorf("error get books total %w", e)
	}
	return total, e
}

func (r *bookRepo) Find(ctx context.Context, findStr string) (lst []*book.Book, e error) {
	lst = make([]*book.Book, 0, 100)

	findStr = "%" + findStr + "%"
	rows, e := r.queryFindBooksStmt.Query(findStr, findStr, findStr, findStr, findStr)
	if e != nil {
		return nil, e
	}

	defer rows.Close()

	for rows.Next() {
		//b.id, b.isbn, b.title, b.price, publisher, author, b.publication_year, b.genre
		b := &book.Book{}
		if e = rows.Scan(&b.ID, &b.ISBN, &b.Title, &b.Price, &b.Publisher, &b.Author, &b.PublicationYear, &b.Genre); e != nil {
			log.Println("Find", "Error find books [", e, "]")
			return nil, fmt.Errorf("error find books %w", e)
		}
		lst = append(lst, b)
	}

	return lst, nil
}
func (r *bookRepo) GetBook(ctx context.Context, id int64) (*book.FullInfo, error) {
	var b book.FullInfo
	var descr sql.NullString
	e := r.getBookInfoStmt.QueryRow(id).Scan(&b.ID, &b.ISBN, &b.Title, &b.Price, &b.Publisher, &b.Author, &b.PublicationYear, &b.Genre,
		&descr, &b.Qty)
	if e != nil {
		log.Println("Find", "Error get books [", e, "]")
		return nil, fmt.Errorf("error get books %w", e)
	}
	if descr.Valid {
		b.Descr = descr.String
	}

	return &b, nil
}
