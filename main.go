package main

import (
	"context"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
)



func main() {

	//cfg := elasticsearch.Config{
	//	Addresses: []string{
	//		"http://elasticsearch.bilyoner.com:9200",
	//	},
	//}
	//es, _ := elasticsearch.NewClient(cfg)
	es, _ := elasticsearch.NewDefaultClient()


	//todo: index create with spesific shard number

	req := esapi.IndexRequest{Index:"test",Body:strings.NewReader("{\"title\":\"test\"}"),WaitForActiveShards:"1"}
	res, err := req.Do(context.Background(),es)
	if err != nil {
		log.Fatalf("%s",err)
	}
	defer res.Body.Close()
	log.Printf("%s",res)

	////todo: delete index
	//deleteReq := esapi.DeleteRequest{Index:"test"}
	//res, err = deleteReq.Do(context.Background(),es)
	//
	//if err != nil {
	//	log.Fatalf("%s",err)
	//}
	//defer res.Body.Close()
	//fmt.Printf("%s",res)

}
