package kvhub

import (
	"github.com/chyroc/kvhub/internal"
)

type Hub struct {
	repo  string
	token string
}

func New(repo, token string) *Hub {
	return &Hub{
		repo:  repo,
		token: token,
	}
}

func (r *Hub) Get(key string) ([]byte, error) {
	return internal.Get(r.repo, r.token, key)
}

func (r *Hub) Set(key string, val []byte) error {
	return internal.Set(r.repo, r.token, key, val)
}
