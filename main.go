package main

import (
	"crypto/sha256"
	"elastic/pkg/elasticsearch"
	"elastic/pkg/schduler"
	"flag"
	"fmt"
	"github.com/robfig/cron/v3"
	"io"
	"log"
	"os"
	"time"
)


var configPath string
var esHost string
var port string
func main() {
	flag.StringVar(&configPath,"config","config/rule.yml","--config config/path ")
	flag.StringVar(&esHost,"host","http://localhost:9200","--config example.com:9200")
	flag.Parse()
	elasticsearch.EsClient(esHost)
	var messages= make(chan string)
	var sig = make(chan bool)
	c := schduler.Cron()
	var rules schduler.Rulesfile
	c.Start()
	//applyRule(&rules,c,messages)
	go fileChecker(sig)
	go listener(messages)
	for{
		if <-sig == true{
			fmt.Println("Rulefile changed!")
			removeJobs(c)
			applyRule(&rules,c,messages)
		}
	}

}


func applyRule(rules *schduler.Rulesfile,c *cron.Cron,messages chan string){

	schduler.GetRule(rules)

	for _, instance := range rules.Rules.DeleteByQuery {
		go func(instance elasticsearch.DeleteByQuery) {
			log.Printf("%s, Thread Starting...!",instance.Name)
			c.AddFunc(instance.Cron, func() {
				messages <- fmt.Sprintf("%s %s", instance.Name, instance.Run())
			})
		}(instance)
	}

	for _, instance := range rules.Rules.Alert {
		go func(instance elasticsearch.Alert) {
			c.AddFunc(instance.Cron, func() {
				messages <- fmt.Sprintf("%s %s", instance.Name, instance.Run())
				//slack.SendMessage("test")
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

func fileChecker(sig chan<- bool){
	var hashes string
	for {
		f, err := os.Open("config/rule.yml")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		h := sha256.New()
		if _, err := io.Copy(h, f); err != nil {
			log.Fatal(err)
		}
		newhash := string(h.Sum(nil))
		if hashes != newhash{
			hashes = newhash
			sig<- true
		}else{
			sig<- false
		}
		time.Sleep(5*time.Second)
	}
}


func listener(messages chan string){
	for data := range messages{
		log.Printf("%s",data)
		//slack.SendMessage(data)
	}
}
