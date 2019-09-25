package main

import (
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)



func main() {
	request esapi.DeleteByQueryRequest{Index:"openshift*",Query:}
	esapi.DeleteByQuery
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://elasticsearch.bilyoner.com:9200",
		},
	}
	es, _ := elasticsearch.NewClient(cfg)
	log.Println(elasticsearch.Version)
	log.Println(es.Info())
	es.DeleteByQuery.WithQuery()


}
