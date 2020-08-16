import asyncio

import aioredis
import orjson

BENCH_VALUE = b'{"user":356091260429402122,"name":"Why are you reading","money":9164,"xp":6000000,"pvpwins":14,"money_booster":0,"time_booster":0,"luck_booster":0,"marriage":463318425901596672,"background":"https://i.imgur.com/LRV2QCK.png","guild":15306,"class":["Paragon","White Sorcerer"],"deaths":0,"completed":0,"lovescore":647,"guildrank":"Leader","backgrounds":null,"puzzles":0,"atkmultiply":"10.0","defmultiply":"10.0","crates_common":30,"crates_uncommon":2,"crates_rare":1,"crates_magic":0,"crates_legendary":0,"luck":"1.0","god":null,"favor":0,"race":"Elf","cv":2,"reset_points":2,"chocolates":0,"trickortreat":0,"eastereggs":0,"colour":{"red":255,"green":255,"blue":255,"alpha":0.8}}'


async def main():
    conn = await aioredis.create_connection("redis://localhost")
    await conn.execute("SET", "bench", BENCH_VALUE)

    for _ in range(100_000):
        val = await conn.execute("GET", "bench")
        decoded = orjson.loads(val)
        decoded["crates_common"] += 1
        decoded["crates_uncommon"] += 1
        encoded = orjson.dumps(decoded)
        await conn.execute("SET", "bench", encoded)

    conn.close()
    await conn.wait_closed()


asyncio.run(main())
