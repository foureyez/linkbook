package persistence

import "context"

type CollectionStore interface {
	Create(ctx context.Context, name string) (*Collection, error)
	GetAll(ctx context.Context) ([]Collection, error)
	GetByName(ctx context.Context, name string) (*Collection, error)
	SearchByName(ctx context.Context, query string) ([]Collection, error)
}
