package app

import "context"

// ApplicationInterface main app interface
type Application interface {
	Run(ctx context.Context)
}
