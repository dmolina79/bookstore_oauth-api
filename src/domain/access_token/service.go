package access_token

import (
	"github.com/dmolina79/bookstore_oauth-api/src/repository/rest"
	"github.com/dmolina79/bookstore_oauth-api/src/utils/errors"
	"strings"
)

type Repository interface {
	GetById(string) (*AccessToken, *errors.RestErr)
	Create(AccessToken) *errors.RestErr
	UpdateExpirationTime(AccessToken) *errors.RestErr
}

type Service interface {
	GetById(string) (*AccessToken, *errors.RestErr)
	Create(AccessTokenRequest) (*AccessToken, *errors.RestErr)
	UpdateExpirationTime(AccessToken) *errors.RestErr
}

type service struct {
	dbRepo       Repository
	restUserRepo rest.UsersRepository
}

func NewService(repo Repository, userRepo rest.UsersRepository) Service {
	return &service{
		dbRepo:       repo,
		restUserRepo: userRepo,
	}
}

func (s *service) GetById(accessTokenId string) (*AccessToken, *errors.RestErr) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, errors.NewBadRequest("invalid access token id")
	}
	accessToken, err := s.dbRepo.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}

	return accessToken, nil
}

func (s *service) Create(req AccessTokenRequest) (*AccessToken, *errors.RestErr) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	user, err := s.restUserRepo.LoginUser(req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	// generate new access token
	at := GetNewAccessToken(string(user.Id))
	// save access token to DB cassandra

	if err := s.dbRepo.Create(at); err != nil {
		return nil, err
	}
	return &at, nil
}

func (s *service) UpdateExpirationTime(at AccessToken) *errors.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}

	return s.dbRepo.UpdateExpirationTime(at)
}
