package main

import (
	"flag"
	"log"

	psql "github.com/domterion/idlebench/postgres/golang"
	redis "github.com/domterion/idlebench/redis/golang"
)

func main() {
	testFlag := flag.String("test", "", "Set which test to run")

	flag.Parse()

	switch test := *testFlag; test {
	case "redis":
		redis.DoTest()
	case "psql":
		psql.DoTest()
	default:
		log.Fatal("Please add a `--test=[psql|redis]` flag.")
	}
}
