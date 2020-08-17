require "redis"
require "json"

class Colour
  include JSON::Serializable

  property red : UInt8
  property green : UInt8
  property blue : UInt8
  property alpha : Float32
end

class BenchData
  include JSON::Serializable

  property user : UInt64
  property name : String
  property money : UInt32
  property xp : UInt32
  property pvpwins : UInt32
  property money_booster : UInt32
  property luck_booster : UInt32
  property time_booster : UInt32
  property marriage : UInt64
  property background : String
  property guild : UInt32
  @[JSON::Field(key: "class")]
  property class_ : Array(String)
  property deaths : UInt32
  property completed : UInt32
  property lovescore : UInt32
  property guildrank : String
  property backgrounds : Array(String)?
  property puzzles : UInt8
  property atkmultiply : String
  property defmultiply : String
  property crates_common : UInt32
  property crates_uncommon : UInt32
  property crates_rare : UInt32
  property crates_magic : UInt32
  property crates_legendary : UInt32
  property luck : String
  property god : String?
  property favor : UInt16
  property race : String
  property cv : UInt8
  property reset_points : UInt8
  property chocolates : UInt16
  property trickortreat : UInt16
  property eastereggs : UInt16
  property colour : Colour
end

module Crystal
  bench_string = "{\"user\":356091260429402122,\"name\":\"Why are you reading\",\"money\":9164,\"xp\":6000000,\"pvpwins\":14,\"money_booster\":0,\"time_booster\":0,\"luck_booster\":0,\"marriage\":463318425901596672,\"background\":\"https://i.imgur.com/LRV2QCK.png\",\"guild\":15306,\"class\":[\"Paragon\",\"White Sorcerer\"],\"deaths\":0,\"completed\":0,\"lovescore\":647,\"guildrank\":\"Leader\",\"backgrounds\":null,\"puzzles\":0,\"atkmultiply\":\"10.0\",\"defmultiply\":\"10.0\",\"crates_common\":30,\"crates_uncommon\":2,\"crates_rare\":1,\"crates_magic\":0,\"crates_legendary\":0,\"luck\":\"1.0\",\"god\":null,\"favor\":0,\"race\":\"Elf\",\"cv\":2,\"reset_points\":2,\"chocolates\":0,\"trickortreat\":0,\"eastereggs\":0,\"colour\":{\"red\":255,\"green\":255,\"blue\":255,\"alpha\":0.8}}"

  redis = Redis.new
  redis.set("bench", bench_string)

  100000.times do
    val = redis.get "bench"
    parsed = BenchData.from_json val.not_nil!
    parsed.crates_common += 1
    parsed.crates_uncommon += 1
    encoded = parsed.to_json
    redis.set("bench", encoded)
  end
end
