# Results

## Redis Benchmark

On an Intel Core i7 6700T running Redis 6.0.6 in a container, the results are as follows:

| name           | real   | user   | sys   |
|----------------|--------|--------|-------|
| python-asyncio | 18.52s | 12.22s | 4.14s |
| golang         | 11.85s | 8.27s  | 2.92s |
| rust-async-std | 6.03s  | 2.87s  | 2.26s |
| crystal        | 7.03s  | 3.68s  | 2.40s |

This is with each benchmark running 5 times, taking the best result by "real" value.
