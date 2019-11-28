package elasticsearch
//
//import (
//	"context"
//	"fmt"
//	"github.com/elastic/go-elasticsearch/v8/esapi"
//	"io/ioutil"
//)
//
//
//
//func  RunSearch()  string {
//	req := esapi.SearchRequest{Index:[]string{"articles"},Query:"Alice",}
//
//	response,err := req.Do(context.Background(),es)
//	if err!=nil{
//		return fmt.Sprintf("Search error, %s",err)
//	}
//	responseBody,err:= ioutil.ReadAll(response.Body)
//	if err != nil {
//		return fmt.Sprintf("Convert Response Body error from DeleteByQuery, %s",err)
//
//	}
//	defer response.Body.Close()
//	return fmt.Sprintf("%s",string(responseBody))
//
//}