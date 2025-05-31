package graph

import (
	"context"
	"graphql-api/graph/model"
)

type ProfileResolver interface {
	Company(ctx context.Context, obj *model.Profile) (*model.Company, error)
}

type CompanyResolver interface {
	Employees(ctx context.Context, obj *model.Company) ([]*model.Profile, error)
}
