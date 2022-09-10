package movies

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/PreetSIngh8929/movie_items-aai/clients/elasticsearch"
	"github.com/PreetSIngh8929/movie_items-aai/domain/queries"
	"github.com/PreetSIngh8929/movie_utils-go/rest_errors"
)

const (
	indexItems = "moviesapi"
	typeItem   = "_doc"
)

func (i *Movie) Save() rest_errors.RestErr {
	result, err := elasticsearch.Client.Index(indexItems, typeItem, i)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to save item", errors.New("database_error"))
	}
	i.Id = result.Id
	return nil
}

func (i *Movie) Get() rest_errors.RestErr {
	itemId := i.Id
	result, err := elasticsearch.Client.Get(indexItems, typeItem, i.Id)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			return rest_errors.NewNotFoundError(fmt.Sprintf("no item found with if %s", i.Id))
		}
		return rest_errors.NewInternalServerError(fmt.Sprintf("error when trying to get id %s", i.Id), errors.New("database_error"))
	}

	bytes, err := result.Source.MarshalJSON()
	if err != nil {
		return rest_errors.NewInternalServerError("error ehrn trying to parse database response", errors.New("database_error"))
	}
	if err := json.Unmarshal(bytes, &i); err != nil {
		return rest_errors.NewInternalServerError("error ehrn trying to parse database response", errors.New("database_error"))
	}
	i.Id = itemId
	return nil
}

func (i *Movie) Search(query queries.EsQuery) ([]Movie, rest_errors.RestErr) {
	result, err := elasticsearch.Client.Search(indexItems, query.Build())
	if err != nil {
		return nil, rest_errors.NewInternalServerError("error when trying to search documents", errors.New("database error"))
	}

	items := make([]Movie, result.TotalHits())
	for index, hit := range result.Hits.Hits {
		bytes, _ := hit.Source.MarshalJSON()
		var item Movie
		if err := json.Unmarshal(bytes, &item); err != nil {
			return nil, rest_errors.NewInternalServerError("error when trying to parse response", errors.New("database error"))
		}
		item.Id = hit.Id
		items[index] = item
	}

	if len(items) == 0 {
		return nil, rest_errors.NewNotFoundError("no items found matching given criteria")
	}
	return items, nil
}
