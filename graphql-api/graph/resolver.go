package graph

import (
	"sync"

	"graphql-api/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	profiles  map[int32]*model.Profile
	companies map[int32]*model.Company
	mu        sync.RWMutex
	nextID    struct {
		profile int32
		company int32
	}
}

func NewResolver() *Resolver {
	return &Resolver{
		profiles:  make(map[int32]*model.Profile),
		companies: make(map[int32]*model.Company),
	}
}
