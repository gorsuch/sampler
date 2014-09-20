package main

import (
	"log"
	"time"

	"github.com/gorsuch/sampler"
)

func main() {
	s := sampler.New(10 * time.Second)
	sample, err := s.Sample("http://www.canary.io")
	if err != nil {
		log.Fatal(err)
	}

	log.Print(sample)
}
