package kvhub

import (
	"errors"

	"github.com/chyroc/kvhub/internal"
)

type Hub struct {
	repo  string
	token string
	scope string
}

func New(repo, token, scope string) *Hub {
	return &Hub{
		repo:  repo,
		token: token,
		scope: scope,
	}
}

var ErrNotFound = errors.New("not found")

func (r *Hub) Get(key string) ([]byte, error) {
	res, err := internal.Get(r.repo, r.token, r.scope, key)
	if err != nil {
		if errors.Is(err, internal.ErrNotFound) {
			return nil, ErrNotFound
		}
	}
	return res, err
}

func (r *Hub) Set(key string, val []byte) error {
	err := internal.Set(r.repo, r.token, r.scope, key, val)
	if errors.Is(err, internal.ErrNotFound) {
		return ErrNotFound
	}
	return err
}
