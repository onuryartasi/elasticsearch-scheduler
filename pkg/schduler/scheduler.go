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
	var DefaultLogger cron.Logger = cron.PrintfLogger(log.New(os.Stdout, "cron: ", log.LstdFlags))
	DEBUG := os.Getenv("DEBUG")
	if DEBUG == "true"{
		c = cron.New(
			cron.WithLogger(
				cron.VerbosePrintfLogger(log.New(os.Stdout, "Schedule: ", log.LstdFlags))))
	}else {
		c = cron.New((cron.WithChain(
			cron.SkipIfStillRunning(DefaultLogger),
		)))
	}
	return c
}


func GetRule(rules *Rulesfile) {
	data, err := ioutil.ReadFile("config/rule.yml")
	if err != nil {
		log.Fatalf("Can't read Rulefile, %s", err)
	}
	err = yaml.Unmarshal([]byte(data), rules)
	if err != nil {
		log.Fatalf("%s",err)
	}
}
