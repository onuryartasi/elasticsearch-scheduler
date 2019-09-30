package main

import (
	"elastic/pkg/elasticsearch"
	"elastic/pkg/schduler"
	"fmt"
	"log"
)




func main() {
	var messages = make(chan string)
	c := schduler.Cron()
	rules := schduler.GetRule()
	c.Start()
	for _,rule := range rules.Rule {
		go func(jobName string,index []string,cron string,body string,queryType string) {
			c.AddFunc(cron, func() {
				messages <- fmt.Sprintf("%s %s",jobName, elasticsearch.RunQuery(queryType,index,body))
			})
		}(rule.Name,rule.Index,rule.Cron,rule.Body,rule.Type)
	}

	// Listen messages channel
	for data := range messages{
		log.Printf("%s",data)
	}

}
