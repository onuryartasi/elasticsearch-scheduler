package elasticsearch

import (
	"context"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"io/ioutil"
	"strings"
)

type DeleteByQuery struct {
	Name string `yaml:"name,omitempty"`
	Cron string `yaml:"cron"`

	Index []string `yaml:"index,omitempty"`

	Body string  `yaml:"body,omitempty"`

//	AllowNoIndices      bool `yaml:"allowNoIndices"`
//	Analyzer            string
//	AnalyzeWildcard     *bool
	Conflicts           string  `yaml:"conflicts,omitempty"`
	DefaultOperator     string  `yaml:"defaultOperator,omitempty"`
//	Df                  string  `yaml:"df"`
//	ExpandWildcards     string
	From                int    `yaml:"from,omitempty"`
	IgnoreUnavailable   bool   `yaml:"ignoreUnavailable,omitempty"`
	Lenient             bool   `yaml:"lenient,omitempty"`
	MaxDocs             int    `yaml:"maxDocs,omitempty"`
//	Preference          string  `yaml:"preference,omitempty"`
//	Query               string  `yaml:"query,omitempty"`
//	Refresh             bool   `yaml:"refresh,omitempty"`
//	RequestCache        bool   `yaml:"requestCache,omitempty"`
//	RequestsPerSecond   int    `yaml:"requestsPerSecond,omitempty"`
//	Routing             []string `yaml:"routing,omitempty"`
//	Scroll              string `yaml:"scroll,omitempty"`
//	ScrollSize          int        `yaml:"scrollSize,omitempty"`
//	SearchTimeout       time.Duration `yaml:"searchTimeout,omitempty"`
//	SearchType          string     `yaml:"searchType,omitempty"`
//	Slices              int        `yaml:"slices,omitempty"`
//	Sort                []string    `yaml:"sort,omitempty"`
//	Source              []string    `yaml:"source,omitempty"`
//	SourceExcludes      []string     `yaml:"sourceExcludes,omitempty"`
//	SourceIncludes      []string     `yaml:"sourceIncludes,omitempty"`
//	Stats               []string     `yaml:"stats,omitempty"`
//	TerminateAfter      int         `yaml:"terminateAfter,omitempty"`
//	Timeout             string `yaml:"timeout,omitempty"`
////	Version             *bool
//	WaitForActiveShards string        `yaml:"waitForActiveShards,omitempty"`
//	WaitForCompletion   *bool         `yaml:"waitForCompletion,omitempty"`
//
//
//	Pretty     bool                    `yaml:"pretty,omitempty"`
//	Human      bool                     `yaml:"human,omitempty"`
//	ErrorTrace bool						`yaml:"errorTrace,omitempty"`
//	FilterPath []string					`yaml:"filterPath,omitempty"`

//	Header http.Header

//	ctx context.Context
}

func (instance DeleteByQuery) Run()  string {
	if len(instance.Index) == 0 {
		return fmt.Sprintf("You must define least one index!")
	}
	if len(instance.Body) == 0   {
		return fmt.Sprintf("You must define body!")
	}
	if instance.Conflicts != "proceed"{
		instance.Conflicts = "abort"
	}
	req := esapi.DeleteByQueryRequest{
		Index:instance.Index,
		Body:strings.NewReader(instance.Body),
		Conflicts:instance.Conflicts,
		DefaultOperator:instance.DefaultOperator,
		}

	if instance.From > 0 {
		req.From = &instance.From
	}
	if instance.MaxDocs > 0 {
		req.MaxDocs = &instance.MaxDocs
	}

	response,err := req.Do(context.Background(),es)
	if err != nil {
		return fmt.Sprintf("Request DeleteByQuery error, %s",err)
	}
	responseBody,err:= ioutil.ReadAll(response.Body)
	if err != nil {
		return fmt.Sprintf("Convert Response Body error from DeleteByQuery, %s",err)

	}
	defer response.Body.Close()
	return fmt.Sprintf("%s",string(responseBody))

}