package elasticsearch

import (
	"context"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"io/ioutil"
	"log"
	"strings"
)




//cfg := elasticsearch.Config{
//	Addresses: []string{
//		"http://elasticsearch.bilyoner.com:9200",
//	},
//}
//es, _ := elasticsearch.NewClient(cfg)

var es,_ = elasticsearch.NewDefaultClient()

func deleteByQuery(index []string,body string) string{
	//todo: deletebyquery
//	body := `{
//	"query": {
//		"match":{"author.first_name":"John"}
//	}
//}`
	deleteReq := esapi.DeleteByQueryRequest{Index:index,Body:strings.NewReader(fmt.Sprintf("%s",body))}
	deleteRes,err := deleteReq.Do(context.Background(),es)

	if err != nil {
		log.Fatalf("%s",err)
	}

	response,err:= ioutil.ReadAll(deleteRes.Body)
	if err != nil {
		log.Fatalf("%s",err)
	}
	defer deleteRes.Body.Close()
	return fmt.Sprintf("%s",string(response))
}


func RunQuery(queryType string,index []string,body string) string{
	switch queryType {
	case "deletebyquery":
		return fmt.Sprintf("%s",deleteByQuery(index,body))
	default:
		return fmt.Sprintf("Type error, type: %s, Please check rulefile.",queryType)
	}
}


