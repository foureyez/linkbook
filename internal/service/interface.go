package service

import "context"

type CollectionService interface {
	GetAll(ctx context.Context) ([]Collection, error)
	GetByName(ctx context.Context, name string) (*Collection, error)
	SearchByName(ctx context.Context, query string) ([]Collection, error)
	Create(ctx context.Context, name string) (*Collection, error)
}
