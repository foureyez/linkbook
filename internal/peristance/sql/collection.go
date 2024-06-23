package sql

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"

	"github.com/foureyez/linkbook/internal/logger"
	persistence "github.com/foureyez/linkbook/internal/peristance"
)

type collectionStore struct {
	db *sqlx.DB
}

func NewCollectionStore(db *sqlx.DB) persistence.CollectionStore {
	return &collectionStore{
		db: db,
	}
}

func (c *collectionStore) Create(ctx context.Context, name string) (*persistence.Collection, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("INSERT into collection (id, name) VALUES (?,?)", id, name)
	if err != nil {
		return nil, err
	}

	return &persistence.Collection{
		Id:   id,
		Name: name,
	}, nil
}

func (c *collectionStore) GetAll(ctx context.Context) ([]persistence.Collection, error) {
	rows, err := c.db.Queryx("SELECT * from collection")
	if err != nil {
		return nil, err
	}

	var collections []persistence.Collection
	for rows.Next() {
		var c persistence.Collection
		err = rows.StructScan(&c)
		if err != nil {
			return nil, err
		}
		collections = append(collections, c)
	}
	return collections, nil
}

func (c *collectionStore) GetByName(ctx context.Context, name string) (*persistence.Collection, error) {
	row := c.db.QueryRowx("SELECT * from collection WHERE name=?", name)
	if row == nil {
		return nil, persistence.ErrNoEntityFound
	}

	var collection persistence.Collection
	err := row.StructScan(&collection)
	if err != nil {
		return nil, err
	}

	return &collection, nil
}

func (c *collectionStore) SearchByName(ctx context.Context, query string) ([]persistence.Collection, error) {
	rows, err := c.db.Queryx("SELECT * from collection WHERE name like ?", query+"%")
	if err != nil {
		logger.Get().Error(err)
		return nil, err
	}

	var collections []persistence.Collection
	for rows.Next() {
		var c persistence.Collection
		err = rows.StructScan(&c)
		if err != nil {
			return nil, err
		}
		collections = append(collections, c)
	}
	return collections, nil
}
