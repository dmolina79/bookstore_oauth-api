package db

import (
	"github.com/dmolina79/bookstore_oauth-api/src/domain/access_token"
	"github.com/dmolina79/bookstore_oauth-api/src/utils/errors"
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct {
}

func New() DbRepository {
	return &dbRepository{}
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	return nil, errors.NewInternalServer("database connection not implemented yet")
}
