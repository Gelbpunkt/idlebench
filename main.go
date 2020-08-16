package main

import (
	"flag"
	"log"

	"github.com/domterion/idlebench/redis"
)

func main() {
	testFlag := flag.String("test", "", "Set which test to run")

	flag.Parse()

	switch test := *testFlag; test {
	case "redis":
		redis.DoTest()
	case "psql":
		log.Println("test")
	default:
		log.Fatal("Please add a `--test=[psql|redis]` flag.")
	}
}
