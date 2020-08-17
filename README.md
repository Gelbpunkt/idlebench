# idlebench
IdleRPG programming language database benchmark suite.

## Benchmark suite

The benchmark consists of 2 benchmarks, one for Redis clients and one for PostgreSQL clients.

Each benchmark should be put in the folder for the database server and a folder for the programming language.

For example:

* postgresql/go/ for a PostgreSQL Go implementation
* redis/python/ for a Redis Python implementation
* redis/python-asyncio/ for a Redis Python implementation based on AsyncIO

Benchmarks should ship with a Dockerfile that will be used to benchmark the implementation with the Unix "time" command-line tool. As this benchmark suite is aimed at container applications, alpine-based Dockerfiles would be preferred.

Benchmarks should be single-threaded, async programming and microthreads would be favorable.

## The Redis Benchmark

The program should:
* Connect to the Redis server at the default port on localhost
* Set the key "bench" to the value `'{"user":356091260429402122,"name":"Why are you reading","money":9164,"xp":6000000,"pvpwins":14,"money_booster":0,"time_booster":0,"luck_booster":0,"marriage":463318425901596672,"background":"https://i.imgur.com/LRV2QCK.png","guild":15306,"class":["Paragon","White Sorcerer"],"deaths":0,"completed":0,"lovescore":647,"guildrank":"Leader","backgrounds":null,"puzzles":0,"atkmultiply":"10.0","defmultiply":"10.0","crates_common":30,"crates_uncommon":2,"crates_rare":1,"crates_magic":0,"crates_legendary":0,"luck":"1.0","god":null,"favor":0,"race":"Elf","cv":2,"reset_points":2,"chocolates":0,"trickortreat":0,"eastereggs":0,"colour":{"red":255,"green":255,"blue":255,"alpha":0.8}}'`
* Then, loop 100,000 times:
    * Get the key "bench"
    * Load the JSON
    * Increment the values for the keys "crates_common" and "crates_uncommon" by 1
    * Encode the JSON again
    * Set the key "bench" to the encoded JSON
* Finally, close the connection gracefully

## The PostgreSQL Benchmark

The example row data used for the benchmark is identical to the Redis JSON as a SQL row:

|               user | name                |   money |      xp |   pvpwins |   money_booster |   time_booster |   luck_booster |           marriage | background                      |   guild | class                         |   deaths |   completed |   lovescore | guildrank   | backgrounds   |   puzzles |   atkmultiply |   defmultiply |   crates_common |   crates_uncommon |   crates_rare |   crates_magic |   crates_legendary |   luck | god   |   favor | race   |   cv |   reset_points |   chocolates |   trickortreat |   eastereggs | colour                                                              |
|--------------------|---------------------|---------|---------|-----------|-----------------|----------------|----------------|--------------------|---------------------------------|---------|-------------------------------|----------|-------------|-------------|-------------|---------------|-----------|---------------|---------------|-----------------|-------------------|---------------|----------------|--------------------|--------|-------|---------|--------|------|----------------|--------------|----------------|--------------|---------------------------------------------------------------------|
| 356091260429402122 | Why are you reading |    9164 | 6000000 |        14 |               0 |              0 |              0 | 463318425901596672 | https://i.imgur.com/LRV2QCK.png |   15306 | ['Paragon', 'White Sorcerer'] |        0 |           0 |         647 | Leader      |               |         0 |            10 |            10 |              30 |                 2 |             1 |              0 |                  0 |      1 |       |       0 | Elf    |    2 |              2 |            0 |              0 |            0 | {'red': 255, 'green': 255, 'blue': 255, 'alpha': 0.800000011920929} |

color is a custom type `public.rgba` with keys and values like the JSON.

The full schema can be seen in the schema.sql file.

The program should:
* Connect to the PostgreSQL server at the default port with the user name "postgres" and password "postgres" on database "postgres"
* Then, loop 1,000 times:
    * Insert the example row data into the table public.profile 100 times
    * Select one row where "user" equals 356091260429402122
    * Update all rows where "user" equals 356091260429402122 by increasing "crates_common" and "crates_uncommon" by 1
    * Delete all rows where "user" equals 356091260429402122
* Finally, close the connection gracefully
