package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"os"
)


func main() {
	var c *cron.Cron
	message := make(chan string)

	DEBUG := os.Getenv("DEBUG")

	if DEBUG == "true"{
		c = cron.New(
			cron.WithLogger(
				cron.VerbosePrintfLogger(log.New(os.Stdout, "cron: ", log.LstdFlags))))
	}else {
		c = cron.New()
	}
	c.AddFunc("@every 1s", func(){
		text := task()
		message<- text
	})
	c.Start()
	for data := range message{
		fmt.Println(data)
	}

}

func task() string{
	return "dangalak"
}

func taskWithParams(a int, b string) {
	fmt.Println(a, b)
}