use redis::{aio::connect_async_std, AsyncCommands, IntoConnectionInfo};
use serde::{Deserialize, Serialize};

#[derive(Deserialize, Serialize)]
struct Colour {
    red: usize,
    green: usize,
    blue: usize,
    alpha: f32,
}

#[derive(Deserialize, Serialize)]
struct BenchData {
    user: u64,
    name: String,
    money: u32,
    xp: u32,
    pvpwins: u32,
    money_booster: u32,
    time_booster: u32,
    luck_booster: u32,
    marriage: u64,
    background: String,
    guild: u32,
    class: Vec<String>,
    deaths: u32,
    completed: u32,
    lovescore: u32,
    guildrank: String,
    backgrounds: Option<Vec<String>>,
    puzzles: usize,
    atkmultiply: String,
    defmultiply: String,
    crates_common: u32,
    crates_uncommon: u32,
    crates_rare: u32,
    crates_magic: u32,
    crates_legendary: u32,
    luck: String,
    god: Option<String>,
    favor: u16,
    race: String,
    cv: usize,
    reset_points: usize,
    chocolates: u16,
    trickortreat: u16,
    eastereggs: u16,
    colour: Colour,
}

const BENCH_VALUE: &[u8] = b"{\"user\":356091260429402122,\"name\":\"Why are you reading\",\"money\":9164,\"xp\":6000000,\"pvpwins\":14,\"money_booster\":0,\"time_booster\":0,\"luck_booster\":0,\"marriage\":463318425901596672,\"background\":\"https://i.imgur.com/LRV2QCK.png\",\"guild\":15306,\"class\":[\"Paragon\",\"White Sorcerer\"],\"deaths\":0,\"completed\":0,\"lovescore\":647,\"guildrank\":\"Leader\",\"backgrounds\":null,\"puzzles\":0,\"atkmultiply\":\"10.0\",\"defmultiply\":\"10.0\",\"crates_common\":30,\"crates_uncommon\":2,\"crates_rare\":1,\"crates_magic\":0,\"crates_legendary\":0,\"luck\":\"1.0\",\"god\":null,\"favor\":0,\"race\":\"Elf\",\"cv\":2,\"reset_points\":2,\"chocolates\":0,\"trickortreat\":0,\"eastereggs\":0,\"colour\":{\"red\":255,\"green\":255,\"blue\":255,\"alpha\":0.8}}";

#[async_std::main]
async fn main() -> Result<(), redis::RedisError> {
    let mut con = connect_async_std(&"redis://127.0.0.1".into_connection_info().unwrap()).await?;

    con.set("bench", BENCH_VALUE).await?;

    for _ in 0i32..99999 {
        let val: Vec<u8> = con.get("bench").await?;
        let mut decoded: BenchData = serde_json::from_slice(&val).unwrap();
        decoded.crates_common += 1;
        decoded.crates_uncommon += 1;
        let encoded = serde_json::to_vec(&decoded).unwrap();
        con.set("bench", encoded).await?;
    }

    Ok(())
}
