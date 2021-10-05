package kvhub

import (
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

func (r *Hub) Get(key string) ([]byte, error) {
	return internal.Get(r.repo, r.token, r.scope, key)
}

func (r *Hub) Set(key string, val []byte) error {
	return internal.Set(r.repo, r.token, r.scope, key, val)
}
