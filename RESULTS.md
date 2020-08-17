# Results

## Redis Benchmark

On an Intel Core i7 6700T running Redis 6.0.6 in a container, the results are as follows:

| name           | real   | user   | sys   |
|----------------|--------|--------|-------|
| python-asyncio | 18.52s | 12.22s | 4.14s |
| golang         | 7.11s  | 3.82s  | 2.40s |
| rust-async-std | 6.03s  | 2.87s  | 2.26s |
| crystal        | 7.03s  | 3.68s  | 2.40s |
| nodejs         | 7.32s  | 3.82s  | 2.14s |

This is with each benchmark running 5 times, taking the best result by "real" value.

## PostgreSQL Benchmark

On an Intel Core i7 6700T running PostgreSQL 13beta3 in a container, the results are as follows:

| name           | real      | user   | sys    |
|----------------|-----------|--------|--------|
| python-asyncio | 2m 43.86s | 23.45s | 10.94s |

This is with each benchmark running 5 times, taking the best result by "real" value.
