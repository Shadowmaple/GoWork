package main

import (
	"flag"
	"fmt"
	"time"
)

var period = flag.Duration("period", 1*time.Second, "sleep perid")

func main() {
	flag.Parse()
	fmt.Printf("Sleeping for %v\n", *period)
	time.Sleep(*period)
}
