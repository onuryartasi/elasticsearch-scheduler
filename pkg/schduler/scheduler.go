package schduler

import (
	"github.com/robfig/cron/v3"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)


func Cron() *cron.Cron {
	var c *cron.Cron
	DEBUG := os.Getenv("DEBUG")

	if DEBUG == "true"{
		c = cron.New(
			cron.WithLogger(
				cron.VerbosePrintfLogger(log.New(os.Stdout, "Schedule: ", log.LstdFlags))))
	}else {
		c = cron.New()
	}

	return c
}


func GetRule() Rules {
	var rules Rules
	data, err := ioutil.ReadFile("config/rule.yml")
	if err != nil {
		log.Fatalf("Can't read Rulefile, %s", err)
	}

	err = yaml.Unmarshal([]byte(data), &rules)
	if err != nil {
		log.Fatalf("%s",err)
	}
	return rules
}
