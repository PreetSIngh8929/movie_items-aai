package elasticsearch

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/PreetSIngh8929/movie_utils-go/logger"
	"github.com/olivere/elastic"
)

const (
	es_username = "mysql_users_username"
	es_password = "mysql_users_password"
)

var (
	username = os.Getenv(es_username)
	password = os.Getenv(es_password)
)
var (
	Client esClientInterface = &esClient{}
)

type esClient struct {
	client *elastic.Client
}
type esClientInterface interface {
	setClient(*elastic.Client)
	Index(string, string, interface{}) (*elastic.IndexResponse, error)
	Get(string, string, string) (*elastic.GetResult, error)
	Search(string, elastic.Query) (*elastic.SearchResult, error)
	Update(string, string, string) (*elastic.UpdateResponse, error)
}

func Init() {
	log := logger.GetLogger()

	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetBasicAuth(es_username, es_password),
		elastic.SetSniff(false),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetErrorLog(log),
		elastic.SetInfoLog(log),
	)
	if err != nil {

		panic(err)
	}
	Client.setClient(client)
}
func (c *esClient) setClient(client *elastic.Client) {
	c.client = client
}
func (c *esClient) Index(index string, docType string, doc interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	result, err := c.client.Index().
		Index(index).
		Type(docType).
		BodyJson(doc).
		Do(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to index document in index %s", index), err)
		return nil, err
	}
	return result, nil
}

func (c *esClient) Get(index string, docType string, id string) (*elastic.GetResult, error) {
	ctx := context.Background()
	result, err := c.client.Get().
		Index(index).
		Type(docType).
		Id(id).
		Do(ctx)

	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to get id %s", id), err)
		return nil, err
	}
	return result, nil
}

func (c *esClient) Search(index string, query elastic.Query) (*elastic.SearchResult, error) {
	ctx := context.Background()
	result, err := c.client.Search(index).
		Query(query).
		RestTotalHitsAsInt(true).
		Do(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to search documents in index %s", index), err)
		return nil, err
	}
	return result, nil
}
func (c *esClient) Update(index string, docType string, id string) (*elastic.UpdateResponse, error) {
	ctx := context.Background()
	result, err := c.client.Update().
		Index(index).
		Type(docType).
		Id(id).
		Script(elastic.NewScript("ctx._source.available_quantity += params.num").Param("num", 1)).
		Upsert(map[string]interface{}{"available_quantity": 0}).
		Do(ctx)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(result.Id, result.Version)
	return result, err
}
