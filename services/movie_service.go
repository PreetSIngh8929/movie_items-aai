package services

import (
	"github.com/PreetSIngh8929/movie_items-aai/domain/movies"
	"github.com/PreetSIngh8929/movie_items-aai/domain/queries"
	"github.com/PreetSIngh8929/movie_utils-go/rest_errors"
)

var (
	MoviesService moviesServiceInterface = &moviesService{}
)

type moviesServiceInterface interface {
	Create(movies.Movie) (*movies.Movie, rest_errors.RestErr)
	Get(string) (*movies.Movie, rest_errors.RestErr)
	Search(queries.EsQuery) ([]movies.Movie, rest_errors.RestErr)
}

type moviesService struct {
}

func (s *moviesService) Create(item movies.Movie) (*movies.Movie, rest_errors.RestErr) {
	if err := item.Save(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *moviesService) Get(id string) (*movies.Movie, rest_errors.RestErr) {
	item := movies.Movie{Id: id}

	if err := item.Get(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *moviesService) Search(query queries.EsQuery) ([]movies.Movie, rest_errors.RestErr) {
	dao := movies.Movie{}
	return dao.Search(query)
}
