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
	//lenRules := len(rules.Rule)

	for _,rule := range rules.Rule {
		go func(jobName string,cron string,body string) {
			c.AddFunc(cron, func() {
				messages <- fmt.Sprintf("%s %s",jobName, elasticsearch.RunDeleteByQuery(body))
			})
		}(rule.Name,rule.Cron,rule.Body)
	}



	// Listen messages channel
	for data := range messages{
		log.Printf("%s",data)
	}

}
