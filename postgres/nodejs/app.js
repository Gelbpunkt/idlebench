const { Client } = require("pg");

const TIMES = 1000;
const TIMES_INSERT = 100;

const VALUES = [
  356091260429402122,
  "Why are you reading",
  9164,
  6000000,
  14,
  0,
  0,
  0,
  463318425901596672,
  "https://i.imgur.com/LRV2QCK.png",
  15306,
  ["Paragon", "White Sorcerer"],
  0,
  0,
  647,
  "Leader",
  null,
  0,
  "10.0",
  "10.0",
  30,
  2,
  1,
  0,
  0,
  "1.0",
  null,
  0,
  "Elf",
  2,
  2,
  0,
  0,
  0,
  "'{ red: 255, green: 255, blue: 255, alpha: 0.8 }'",
];

const STATIC_INSERT_QUERY =
  'INSERT INTO public.profile ("user", "name", "money", "xp", "pvpwins","money_booster", "time_booster", "luck_booster", "marriage", "background", "guild", "class", "deaths", "completed", "lovescore", "guildrank", "backgrounds", "puzzles", "atkmultiply", "defmultiply", "crates_common", "crates_uncommon", "crates_rare", "crates_magic", "crates_legendary", "luck", "god", "favor", "race", "cv", "reset_points", "chocolates", "trickortreat", "eastereggs", "colour") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35::public.rgba);';
const STATIC_UPDATE_QUERY =
  'UPDATE public.profile SET "crates_common"="crates_common"+1, "crates_uncommon"="crates_uncommon"+1 WHERE "user"=$1;';

const client = new Client({
  user: "postgres",
  password: "postgres",
  database: "postgres",
  host: "127.0.0.1",
  port: 5432,
});

async function main() {
  await client.connect();
  for (let i = 0; i < TIMES; i++) {
    for (let x = 0; x < TIMES_INSERT; x++) {
      await client.query(STATIC_INSERT_QUERY, VALUES);
    }

    await client.query(
      'SELECT * FROM public.profile WHERE "user"=356091260429402122;'
    );
    await client.query(STATIC_UPDATE_QUERY, [356091260429402122]);

    await client.query(
      'DELETE FROM public.profile WHERE "user"=356091260429402122;'
    );
  }

  await client.end();
}

main();
