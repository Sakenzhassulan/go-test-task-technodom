package server

import (
	"github.com/Sakenzhassulan/go-test-task-technodom/db"
	"github.com/Sakenzhassulan/go-test-task-technodom/internal/config"
	"github.com/Sakenzhassulan/go-test-task-technodom/store"
	"github.com/go-playground/validator/v10"
)

type Client struct {
	Validator *validator.Validate
	Config    *config.Config
	Cache     *store.Store
	DB        *db.Instance
}

func NewClient(conf *config.Config, db *db.Instance, cache *store.Store) (*Client, error) {
	return &Client{
		Validator: validator.New(),
		Config:    conf,
		DB:        db,
		Cache:     cache,
	}, nil
}
