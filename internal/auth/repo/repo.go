package repo

import "context"

type AdminInfo struct {
	ID int64
}
type AuthRepo interface {
	Login(ctx context.Context, user, pass string) (int64, error)
	Logout(ctx context.Context) error

	ClientLogin(ctx context.Context, user, pass string) (int64, error)
	ClientLogout(ctx context.Context) error
}
