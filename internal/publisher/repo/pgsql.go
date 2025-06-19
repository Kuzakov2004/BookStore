package repo

import (
	"BookStore/internal/publisher"
	"context"
	"database/sql"
	"fmt"
	"log"
)

type publisherRepo struct {
	connPool *sql.DB

	queryPublishersStmt  *sql.Stmt
	getPublishersCntStmt *sql.Stmt
	getPublisherStmt     *sql.Stmt
}

func NewPublisherRepo(db *sql.DB) (PublisherRepo, error) {

	r := &publisherRepo{connPool: db}
	var e error

	r.queryPublishersStmt, e = db.Prepare(`SELECT p.id, p.name, p.country, p.phone 
											FROM store.publishers p
    										LIMIT $1 OFFSET $2`)
	if e != nil {
		return nil, e
	}

	r.getPublishersCntStmt, e = db.Prepare(`SELECT count(*) FROM store.publishers`)
	if e != nil {
		return nil, e
	}

	r.getPublisherStmt, e = db.Prepare(`SELECT p.id, p.name, p.country, p.phone
											FROM store.publishers p 
											WHERE p.id = $1 
    										LIMIT 1`)
	if e != nil {
		return nil, e
	}

	return r, nil
}

func (r *publisherRepo) GetPublishers(ctx context.Context, page, count int) (lst []*publisher.Publisher, e error) {

	if count < 10 {
		count = 10
	}

	lst = make([]*publisher.Publisher, 0, count)

	if page < 0 {
		page = 0
	}

	offset := page * count

	rows, e := r.queryPublishersStmt.Query(count, offset)
	if e != nil {
		return nil, e
	}

	defer rows.Close()

	for rows.Next() {
		//p.id, p.name, p.country, p.phone
		p := &publisher.Publisher{}
		if e = rows.Scan(&p.ID, &p.Name, &p.Country, &p.Phone); e != nil {
			log.Println("GetPublishers", "Error get publishers [", e, "]")
			return nil, fmt.Errorf("error get publishers %w", e)
		}
		lst = append(lst, p)
	}

	return lst, nil
}

func (r *publisherRepo) GetPublishersCnt(ctx context.Context) (total int, e error) {
	if e = r.getPublishersCntStmt.QueryRow().Scan(&total); e != nil {
		log.Println("GetPublishersCnt", "Error get publishers total [", e, "]")
		return 0, fmt.Errorf("error get publishers total %w", e)
	}
	return total, e
}

func (r *publisherRepo) GetPublisher(ctx context.Context, id int64) (*publisher.Publisher, error) {
	var p publisher.Publisher
	e := r.getPublisherStmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Country, &p.Phone)
	if e != nil {
		log.Println("GetPublisher", "Error get publisher [", e, "]")
		return nil, fmt.Errorf("error get publisher %w", e)
	}

	return &p, nil
}
