sampler
=======

A simple package to help make sampling HTTP performance of remote websites a little easier.

## Usage

```golang
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
```
