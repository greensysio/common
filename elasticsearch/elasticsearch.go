package elasticsearch

import (
	"bitbucket.org/greensys-tech/common/log"
	"context"
	"errors"
	"github.com/olivere/elastic/v7"
	"os"
	"strconv"
)

var elasticsearchClient *elastic.Client

func Client() *elastic.Client {
	return elasticsearchClient
}

func InitElasticSearch(indices []string) {
	if elasticMode, err := strconv.ParseBool(os.Getenv("ELASTIC_MODE")); err == nil && elasticMode {
		client, err := NewElasticClient(context.Background())
		if err != nil {
			log.Error("Init Elasticsearch failed!")
			os.Exit(1)
		}
		elasticsearchClient = client
		log.Info("Init Elasticsearch successfully!")

		// Create indices
		for _, indexName := range indices {
			if err := CreateIndexIfDoesNotExist(context.Background(), elasticsearchClient, indexName); err != nil {
				log.Errorf("Create Index %s if does not exist fail! Err: %+v", indexName, err)
				os.Exit(1)
			}
		}
		log.Info("Create indices if does not exist successfully!")
	}
}

// NewElasticClient ...
func NewElasticClient(ctx context.Context) (*elastic.Client, error) {
	url := os.Getenv("ELASTIC_SEARCH_URL")
	client, err := elastic.NewClient(elastic.SetURL(url), elastic.SetSniff(false))
	if err != nil {
		log.Errorf("New ElasticSearch Client fail! Err: %+v", err)
		return nil, err
	}

	err = ping(ctx, client, url)
	if err != nil {
		log.Errorf("Ping ElasticSearch Server fail! Err: %+v", err)
		return nil, err
	}
	return client, nil
}

// CreateIndexIfDoesNotExist ...
func CreateIndexIfDoesNotExist(ctx context.Context, client *elastic.Client, indexName string) error {
	exists, err := client.IndexExists(indexName).Do(ctx)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}

	res, err := client.CreateIndex(indexName).Do(ctx)
	if err != nil {
		return err
	}

	if !res.Acknowledged {
		return errors.New("CreateIndex was not acknowledged. Check that timeout value is correct.")
	}
	log.Infof("Created index: %s", indexName)
	return nil
}

// Ping method
func ping(ctx context.Context, client *elastic.Client, url string) error {
	// Ping the Elasticsearch server to get HttpStatus, version number
	if client != nil {
		info, _, err := client.Ping(url).Do(ctx)
		if err != nil {
			return err
		}
		log.Infof("Elasticsearch version is %s", info.Version.Number)
		return nil
	}
	return errors.New("elasticsearch client is nil")
}
