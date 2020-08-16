package golang

import (
	"context"
	"encoding/json"
	"log"

	"github.com/go-redis/redis/v8"
)

var (
	ctx = context.Background()
)

const (
	bench = `{"user":356091260429402122,"name":"Why are you reading","money":9164,"xp":6000000,"pvpwins":14,"money_booster":0,"time_booster":0,"luck_booster":0,"marriage":463318425901596672,"background":"https://i.imgur.com/LRV2QCK.png","guild":15306,"class":["Paragon","White Sorcerer"],"deaths":0,"completed":0,"lovescore":647,"guildrank":"Leader","backgrounds":null,"puzzles":0,"atkmultiply":"10.0","defmultiply":"10.0","crates_common":30,"crates_uncommon":2,"crates_rare":1,"crates_magic":0,"crates_legendary":0,"luck":"1.0","god":null,"favor":0,"race":"Elf","cv":2,"reset_points":2,"chocolates":0,"trickortreat":0,"eastereggs":0,"colour":{"red":255,"green":255,"blue":255,"alpha":0.8}}`
)

func DoTest() {
	redis := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	defer redis.Close()

	err := redis.Set(ctx, "bench", bench, 0).Err()
	if err != nil {
		log.Fatal("Error when running Redis bench, ", err)
	}

	for i := 0; i < 100000; i++ {
		val, err := redis.Get(ctx, "bench").Result()

		if err != nil {
			log.Printf("Error on iter %d, %s", i, err)
			break
		}

		var result map[string]interface{}

		// Unmarshal or Decode the JSON to the interface.
		json.Unmarshal([]byte(val), &result)

		result["crates_common"] = result["crates_common"].(float64) + 1
		result["crates_uncommon"] = result["crates_uncommon"].(float64) + 1

		e, _ := json.Marshal(result)

		err = redis.Set(ctx, "bench", e, 0).Err()
		if err != nil {
			log.Fatal("Error when running Redis bench, ", err)
		}
	}

	log.Println("Done")
}
