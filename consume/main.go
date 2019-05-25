package main

import (
	"log"
	"os"
	"sync"

	"github.com/bitly/go-nsq"
)

func main() {

	wg := &sync.WaitGroup{}
	wg.Add(1)

	config := nsq.NewConfig()
	q, _ := nsq.NewConsumer("write_test", "ch", config)
	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Printf("Got a message: %v", message)
		wg.Done()
		return nil
	}))
	err := q.ConnectToNSQD(os.Getenv("NSQD_HOST"))
	if err != nil {
		log.Panic("Could not connect")
	}
	wg.Wait()

}
