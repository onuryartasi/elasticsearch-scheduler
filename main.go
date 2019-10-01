package main

import (
	"elastic/pkg/elasticsearch"
	"elastic/pkg/schduler"
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
)




func main() {
	var messages= make(chan string)
	c := schduler.Cron()
	var rules schduler.Rulesfile
	c.Start()
	applyRule(&rules,c,messages)

	// Listen messages channel
	for data := range messages{
		log.Printf("%s",data)
	}


}


func applyRule(rules *schduler.Rulesfile,c *cron.Cron,messages chan string){

	schduler.GetRule(rules)

	for _, instance := range rules.Rules.DeleteByQuery {
		go func(instance elasticsearch.DeleteByQuery) {
			c.AddFunc(instance.Cron, func() {
				messages <- fmt.Sprintf("%s %s", instance.Name, instance.Run())
			})
		}(instance)
	}
}


func removeJobs(c *cron.Cron){
	c.Stop()
	for _,entry := range c.Entries(){
		c.Remove(entry.ID)
	}
	c.Start()

}


//for _,rule := range rules.Rule {
//go func(jobName string,index []string,cron string,body string,queryType string) {
//c.AddFunc(cron, func() {
//messages <- fmt.Sprintf("%s %s",jobName, elasticsearch.RunQuery(queryType,index,body))
//})
//}(rule.Name,rule.Index,rule.Cron,rule.Body,rule.Type)
//}