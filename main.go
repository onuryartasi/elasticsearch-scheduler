package main

import (
	"elastic/pkg/elasticsearch"
	"elastic/pkg/schduler"
	"fmt"
)

func main() {

	var messages = make(chan string)
	c := schduler.Cron()
	c.AddFunc("@every 10s", func(){
		 messages <- elasticsearch.RunDeleteByQuery()
	})
	c.Start()


	for data := range messages{
		fmt.Println(data)
	}

}
