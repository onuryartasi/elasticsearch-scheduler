package elasticsearch

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"os"
)




//cfg := elasticsearch.Config{
//	Addresses: []string{
//		"http://elasticsearch.bilyoner.com:9200",
//	},
//}
//es, _ := elasticsearch.NewClient(cfg)

var es,_ = elasticsearch.NewDefaultClient()
func EsClient(){
	os.Setenv("ES_HOST","http://es02.bilyoner.com:9200")
	if len(os.Getenv("ES_HOST")) > 0 {
		fmt.Println("Ok")
		var  cfg = elasticsearch.Config{Addresses:[]string{fmt.Sprintf(os.Getenv("ES_HOST"))}}
		es,_ = elasticsearch.NewClient(cfg)
	}

}

//func deleteByQuery(index []string,body string) string{
//	//todo: deletebyquery
////	body := `{
////	"query": {
////		"match":{"author.first_name":"John"}
////	}
////}`
//
//
//	deleteReq := esapi.DeleteByQueryRequest{
//		Index:index,
//		Body:strings.NewReader(fmt.Sprintf("%s",body)),
//
//
//	}
//	deleteRes,err := deleteReq.Do(context.Background(),es)
//
//	if err != nil {
//		log.Fatalf("%s",err)
//	}
//
//	response,err:= ioutil.ReadAll(deleteRes.Body)
//	if err != nil {
//		log.Fatalf("%s",err)
//	}
//	defer deleteRes.Body.Close()
//	return fmt.Sprintf("%s",string(response))
//}


//func RunQuery(queryType string,index []string,body string) string{
//	switch queryType {
//	case "deletebyquery":
//		return fmt.Sprintf("%s",deleteByQuery(index,body))
//	default:
//		return fmt.Sprintf("Type error, type: %s, Please check rulefile.",queryType)
//	}
//}
//

