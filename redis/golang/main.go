package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

type colour struct {
	Red   uint    `json:"red"`
	Blue  uint    `json:"blue"`
	Green uint    `json:"green"`
	Alpha float32 `json:"alpha"`
}

type Profile struct {
	User            uint64   `json:"user"`
	Name            string   `json:"name"`
	Money           uint32   `json:"money"`
	Xp              uint32   `json:"xp"`
	Pvpwins         uint32   `json:"pvpwins"`
	Moneybooster    uint32   `json:"money_booster"`
	Timebooster     uint32   `json:"time_booster"`
	Luckbooster     uint32   `json:"luck_booster"`
	Marriage        uint64   `json:"marriage"`
	Background      string   `json:"background"`
	Guild           uint32   `json:"guild"`
	Class           []string `json:"class"`
	Deaths          uint32   `json:"deaths"`
	Completed       uint32   `json:"completed"`
	Lovescore       uint32   `json:"lovescore"`
	Guildrank       string   `json:"guildrank"`
	Backgrounds     []string `json:"backgrounds"`
	Puzzles         uint     `json:"puzzles"`
	Atkmultiply     string   `json:"atkmultiply"`
	Defmultiply     string   `json:"defmultiply"`
	Cratescommon    uint32   `json:"crates_common"`
	Cratesuncommon  uint32   `json:"crates_uncommon"`
	Cratesrare      uint32   `json:"crates_rare"`
	Cratesmagic     uint32   `json:"crates_magic"`
	Crateslegendary uint32   `json:"crates_legendary"`
	Luck            string   `json:"string"`
	God             string   `json:"god"`
	Favor           uint16   `json:"favor"`
	Race            string   `json:"race"`
	Cv              uint     `json:"cv"`
	Resetpoints     uint     `json:"reset_points"`
	Chocolates      uint16   `json:"chocolates"`
	Trickortreat    uint16   `json:"trickortreat"`
	Eastereggs      uint16   `json:"eastereggs"`
	Colour          colour   `json:"colour"`
}

var (
	ctx = context.Background()
)

const (
	bench = `{"user":356091260429402122,"name":"Why are you reading","money":9164,"xp":6000000,"pvpwins":14,"money_booster":0,"time_booster":0,"luck_booster":0,"marriage":463318425901596672,"background":"https://i.imgur.com/LRV2QCK.png","guild":15306,"class":["Paragon","White Sorcerer"],"deaths":0,"completed":0,"lovescore":647,"guildrank":"Leader","backgrounds":null,"puzzles":0,"atkmultiply":"10.0","defmultiply":"10.0","crates_common":30,"crates_uncommon":2,"crates_rare":1,"crates_magic":0,"crates_legendary":0,"luck":"1.0","god":null,"favor":0,"race":"Elf","cv":2,"reset_points":2,"chocolates":0,"trickortreat":0,"eastereggs":0,"colour":{"red":255,"green":255,"blue":255,"alpha":0.8}}`
)

func main() {
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

	start := time.Now()
	for i := 0; i < 100000; i++ {
		val, err := redis.Get(ctx, "bench").Result()

		if err != nil {
			log.Printf("Error on iter %d, %s", i, err)
			break
		}

		var profile Profile

		json.Unmarshal([]byte(val), &profile)

		profile.Cratescommon++
		profile.Cratesuncommon++

		e, _ := json.Marshal(profile)

		err = redis.Set(ctx, "bench", e, 0).Err()
		if err != nil {
			log.Fatal("Error when running Redis bench, ", err)
		}
	}
	elapsed := time.Since(start)

	log.Println("Done, ", elapsed)
}
