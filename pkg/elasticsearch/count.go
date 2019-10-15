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
type Alert struct {
	Name  string   `yaml:"name,omitempty"`
	Cron  string   `yaml:"cron"`
	Index []string `yaml:"index,omitempty"`
	Since string `yaml:"since"`
	Query string `yaml:"query"`
}

func (instance Alert) Run()  string {
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
	req := esapi.CountRequest{Index:instance.Index,Body:strings.NewReader(fmt.Sprintf(query,instance.Since,instance.Query))}


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
	defer response.Body.Close()
	return fmt.Sprintf("%d",count.Count)
}