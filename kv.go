package kvhub

import (
	"github.com/chyroc/kvhub/internal"
)

type Hub struct{}

func New() *Hub {
	return &Hub{}
}

func (r *Hub) Get(key string) (string, error) {
	return internal.Get(key)
}

func (r *Hub) Set(key, val string) error {
	return internal.Set(key, val)
}
