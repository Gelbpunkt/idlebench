# idlebench
IdleRPG programming language database benchmark suite.

## Benchmark suite

The benchmark consists of 2 benchmarks, one for Redis clients and one for PostgreSQL clients.

Each benchmark should be put in the folder for the database server and a folder for the programming language.

For example:

* postgresql/go/ for a PostgreSQL Go implementation
* redis/python/ for a Redis Python implementation
* redis/python-asyncio/ for a Redis Python implementation based on AsyncIO

Benchmarks should ship with a Dockerfile that will be used to benchmark the implementation with the Unix "time" command-line tool.

## The Redis Benchmark

The program should:
- Connect to the Redis server at the default port on localhost
- Set the key "bench" to the value `'{"user":356091260429402122,"name":"Why are you reading","money":9164,"xp":6000000,"pvpwins":14,"money_booster":0,"time_booster":0,"luck_booster":0,"marriage":463318425901596672,"background":"https://i.imgur.com/LRV2QCK.png","guild":15306,"class":["Paragon","White Sorcerer"],"deaths":0,"completed":0,"lovescore":647,"guildrank":"Leader","backgrounds":null,"puzzles":0,"atkmultiply":"10.0","defmultiply":"10.0","crates_common":30,"crates_uncommon":2,"crates_rare":1,"crates_magic":0,"crates_legendary":0,"luck":"1.0","god":null,"favor":0,"race":"Elf","cv":2,"reset_points":2,"chocolates":0,"trickortreat":0,"eastereggs":0,"colour":{"red":255,"green":255,"blue":255,"alpha":0.8}}'`
- Then, loop 100,000 times:
    - Get the key "bench"
    - Load the JSON
    - Increment the values for the keys "crates_common" and "crates_uncommon" by 1
    - Encode the JSON again
    - Set the key "bench" to the encoded JSON
- Finally, close the connection gracefully
