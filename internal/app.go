package internal

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"

	"github.com/foureyez/linkbook/config"
	"github.com/foureyez/linkbook/internal/api"
	"github.com/foureyez/linkbook/internal/http"
	"github.com/foureyez/linkbook/internal/peristance/sql"
	"github.com/foureyez/linkbook/internal/service"
)

type App struct{}

func NewApp() *App {
	return &App{}
}

func (a *App) Run(ctx context.Context) error {
	cfg, err := a.initConfig()
	if err != nil {
		return err
	}

	db, err := sqlx.Open("sqlite3", "data/db.sqlite")
	if err != nil {
		return err
	}

	handlers, err := a.getApiHandlers(db)
	if err != nil {
		return err
	}

	if err := http.StartServer(ctx, &cfg.Server, handlers); err != nil {
		return err
	}
	return nil
}

func (a *App) getApiHandlers(db *sqlx.DB) ([]api.Handler, error) {
	handlers := make([]api.Handler, 0)
	collectionStore := sql.NewCollectionStore(db)
	collectionService := service.NewCollectionService(collectionStore)

	handlers = append(handlers, api.NewCollectionHandler(collectionService))
	return handlers, nil
}

func (a *App) initConfig() (*config.Config, error) {
	v := viper.New()
	v.AddConfigPath("./config")
	v.AddConfigPath(".")
	v.SetConfigName("linkbook")
	v.SetConfigType("yaml")
	v.AutomaticEnv()
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var cfg *config.Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
