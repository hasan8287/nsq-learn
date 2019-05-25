package main

import (
	"log"
	"os"

	"github.com/bitly/go-nsq"
)

func main() {

	// simple
	config := nsq.NewConfig()
	w, _ := nsq.NewProducer(os.Getenv("NSQD_HOST"), config)

	err := w.Publish("write_test", []byte("test"))
	if err != nil {
		log.Panic("Could not connect")
	}

	w.Stop()
}
