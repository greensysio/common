package elasticsearch

import (
	cmContext "github.com/greensysio/common/context"
	"github.com/greensysio/common/log"
	"context"
	"os"
	"strconv"
)

func CreateObject(c cmContext.CustomContext, indexName, objectID string, ob interface{}) (ok bool) {
	if elasticMode, err := strconv.ParseBool(os.Getenv("ELASTIC_MODE")); err == nil && elasticMode {
		log.InfofCtx(&c, "ElasticSearch! Create object: %+v", ob)
		_, err := Client().Index().
			Index(indexName).
			Id(objectID).
			BodyJson(ob).
			Do(context.Background())
		if err != nil {
			log.ErrorfCtx(&c, "ElasticSearch! Create object fail. Error : %s", err.Error())
			return
		}
		return true
	}
	return
}

func UpdateObject(c cmContext.CustomContext, indexName, objectID string, ob interface{}) (ok bool) {
	if elasticMode, err := strconv.ParseBool(os.Getenv("ELASTIC_MODE")); err == nil && elasticMode {
		log.InfofCtx(&c, "ElasticSearch! Update object: %+v", ob)
		_, err := Client().Update().
			Index(indexName).
			Id(objectID).
			Doc(ob).
			DetectNoop(true).
			Do(context.Background())
		if err != nil {
			log.ErrorfCtx(&c, "ElasticSearch! Update object fail. Error : %s", err.Error())
			return
		}
		return true
	}
	return
}

func DeleteObject(c cmContext.CustomContext, indexName, objectID string) (ok bool) {
	if elasticMode, err := strconv.ParseBool(os.Getenv("ELASTIC_MODE")); err == nil && elasticMode {
		log.InfofCtx(&c, "ElasticSearch! Delete object ID: %s", objectID)
		_, err := Client().Delete().
			Index(indexName).
			Id(objectID).
			Do(context.Background())
		if err != nil {
			log.ErrorfCtx(&c, "Delete user! Error: %+v", err)
			return
		}
		return true
	}
	return
}
