package repo

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

type adminRepo struct {
	connPool *sql.DB

	loginStmt *sql.Stmt
}

func NewAdminRepo(db *sql.DB) (AdminRepo, error) {

	r := &adminRepo{connPool: db}
	var e error

	r.loginStmt, e = db.Prepare(`select t.ID from store.admins t where encode(sha256(($1 || t.salt || $2)::bytea), 'hex') = t.pass`)
	if e != nil {
		return nil, e
	}

	return r, nil
}

func (r *adminRepo) Login(ctx context.Context, user, pass string) (int64, error) {

	var id int64
	if e := r.loginStmt.QueryRow(user, pass).Scan(&id); e != nil {
		log.Println("Login", "Error login [", e, "]")
		return 0, fmt.Errorf("error login %w", e)
	}
	return id, nil
}

func (r *adminRepo) Logout(ctx context.Context) error {
	return nil
}
