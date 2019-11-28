package elasticsearch

import "fmt"


type Query interface {
	GetQuery()
}
func (instance Alert)GetQuery() string{
	query := `{"query": {
    "bool": {
      "must": [
        {
          "range": {
            "%s": {
              "format": "strict_date_optional_time",
              "gte": "now-%s"
            }
          }
        }
      ]
		%s
    }
  }
}
`

	filter := `,
      "filter": [
        {
          "query_string": {
            "type": "phrase",
            "query": "/.*%s.*/"
          }
        }
      ]`

	if len(instance.TimeField)==0{
		instance.TimeField = "@timestamp"
	}
	if len(instance.Since)==0{
		instance.Since = "15m"
	}
	if len(instance.Query)>0{
		filter =fmt.Sprintf(filter,instance.Query)
	}else {
		filter=""
	}
	return fmt.Sprintf(query,instance.TimeField,instance.Since,filter)
}
func (instance DeleteByQuery)GetQuery() string{
	query := `{"query": {
    "bool": {
      "must": [
        {
          "range": {
            "%s": {
              "format": "strict_date_optional_time",
              "gte": "now-%s"
            }
          }
        }
      ]
		%s
    }
  }
}
`

	filter := `,
      "filter": [
        {
          "query_string": {
            "type": "phrase",
            "query": "/.*%s.*/"
          }
        }
      ]`

	if len(instance.TimeField)==0{
		instance.TimeField = "@timestamp"
	}
	if len(instance.Since)==0{
		instance.Since = "15m"
	}
	if len(instance.Query)>0{
		filter =fmt.Sprintf(filter,instance.Query)
	}else {
		filter=""
	}
	return fmt.Sprintf(query,instance.TimeField,instance.Since,filter)
}