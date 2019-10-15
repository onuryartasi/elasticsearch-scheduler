package elasticsearch

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"io/ioutil"
	"strings"
)


type SearchCount struct {
	Shards *Shards `json:"_shards"`
	Count int `json:"count"`


}
type Shards struct {
	Total int `json:"total"`
	Successful int `json:"successful"`
	Skipped int `json:"skipped"`
	Failed int `json:"failed"`
}



func  RunCount()  string {
	query := `{"query": {
    "bool": {
      "must": [
        {
          "range": {
            "published": {
              "format": "strict_date_optional_time",
              "gte": "now-%s"
            }
          }
        }
      ],
      "filter": [
        {
          "multi_match": {
            "type": "best_fields",
            "query": "%s",
            "lenient": true
          }
        }
      ]
    }
  }
}
`

	count := &SearchCount{Shards:&Shards{},}
	req := esapi.CountRequest{Index:[]string{"articles"},Body:strings.NewReader(fmt.Sprintf(query,"15m","Alice"))}


	response,err := req.Do(context.Background(),es)
	if err!=nil{
		return fmt.Sprintf("Search error, %s",err)
	}
	responseBody,err:= ioutil.ReadAll(response.Body)
	if err != nil {
		return fmt.Sprintf("Convert Response Body error from DeleteByQuery, %s",err)

	}
	err = json.Unmarshal(responseBody,count)
	if err !=nil {
		fmt.Println(err)
	}
	fmt.Println(string(responseBody))
	defer response.Body.Close()
	fmt.Println(count.Count)
	return fmt.Sprintf("%d",count)
}