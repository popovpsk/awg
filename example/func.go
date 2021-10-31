package example

//go:generate awg

import (
	"context"
	"encoding/json"
	err "errors"
	"math/rand"
)

type Service struct {
}

type Object struct {
}

type Closer interface {
	close()
}

// GetById ...
// awg: future
func (s *Service) GetById(ctx context.Context, number []*json.Number, str []string, n *int, q map[string]*rand.Rand, cl Closer, i interface{}) ([]*Object, error) {
	json.Marshal(&Object{})
	return []*Object{}, err.New("error")
}
