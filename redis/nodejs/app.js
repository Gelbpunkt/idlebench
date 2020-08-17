const redis = require("promise-redis");

const client = redis().createClient();
const BENCH_VALUE =
  '{"user":356091260429402122,"name":"Why are you reading","money":9164,"xp":6000000,"pvpwins":14,"money_booster":0,"time_booster":0,"luck_booster":0,"marriage":463318425901596672,"background":"https://i.imgur.com/LRV2QCK.png","guild":15306,"class":["Paragon","White Sorcerer"],"deaths":0,"completed":0,"lovescore":647,"guildrank":"Leader","backgrounds":null,"puzzles":0,"atkmultiply":"10.0","defmultiply":"10.0","crates_common":30,"crates_uncommon":2,"crates_rare":1,"crates_magic":0,"crates_legendary":0,"luck":"1.0","god":null,"favor":0,"race":"Elf","cv":2,"reset_points":2,"chocolates":0,"trickortreat":0,"eastereggs":0,"colour":{"red":255,"green":255,"blue":255,"alpha":0.8}}';
const TIMES = 100000;

async function main() {
  await client.set("bench", BENCH_VALUE);
  for (let i = 0; i < TIMES; i++) {
    const response = await client.get("bench");

    let json = JSON.parse(response);

    json.crates_common += 1;
    json.crates_uncommon += 1;

    await client.set("bench", JSON.stringify(json));
  }
  await client.quit();
}

main();
