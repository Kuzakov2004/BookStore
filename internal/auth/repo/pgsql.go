package repo

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

type authRepo struct {
	connPool *sql.DB

	loginStmt       *sql.Stmt
	clientLoginStmt *sql.Stmt
}

func NewAuthRepo(db *sql.DB) (AuthRepo, error) {

	r := &authRepo{connPool: db}
	var e error

	r.loginStmt, e = db.Prepare(`select t.ID from store.admins t where encode(sha256(($1 || t.salt || $2)::bytea), 'hex') = t.pass`)
	if e != nil {
		return nil, e
	}

	r.clientLoginStmt, e = db.Prepare(`select t.ID from store.clients t where encode(sha256(($1 || t.salt || $2)::bytea), 'hex') = t.password`)
	if e != nil {
		return nil, e
	}

	return r, nil
}

func (r *authRepo) Login(ctx context.Context, user, pass string) (int64, error) {

	var id int64
	if e := r.loginStmt.QueryRow(user, pass).Scan(&id); e != nil {
		log.Println("Login", "Error login [", e, "]")
		return 0, fmt.Errorf("error login %w", e)
	}
	return id, nil
}

func (r *authRepo) ClientLogin(ctx context.Context, user, pass string) (int64, error) {

	var id int64
	if e := r.clientLoginStmt.QueryRow(user, pass).Scan(&id); e != nil {
		log.Println("Login", "Error client login [", e, "]")
		return 0, fmt.Errorf("error client login %w", e)
	}
	return id, nil
}

func (r *authRepo) Logout(ctx context.Context) error {
	return nil
}

func (r *authRepo) ClientLogout(ctx context.Context) error {
	return nil
}
