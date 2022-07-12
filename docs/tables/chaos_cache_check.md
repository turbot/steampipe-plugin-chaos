# Table: chaos_cache_check

Chaos table to test the cache functionality in steampipe. 

Columns `time_col`, `int_col`, `float_col` are the Optional Key Columns which support operators. These columns are used for cache testing for scenarios which involves quals(`008.cache_quals.bats`)

Column `unique_col` generates a random integer value. This column is specifically used to test caching scenarios. It returns the same random value when caching is turned ON, and changes when turned OFF. This column is used in acceptance tests in(`007.cache.bats` and `008.cache_quals.bats`)

Columns `delay` and `long_delay` are used to add delays in the query. Column `delay` adds a 10sec delay and `long_delay` adds 10hrs delay. 
The `delay` column is used to test the cache functionality when the first query runs in the background in `@test "check cache functionality when querying same columns(first query in background)" in 007.cache.bats`. 
The `long_delay` column is used to simulate a timeout in `@test "check cache functionality when first query times out, other queries should cache(first query in background)" in 007.cache.bats`.

`DO NOT` query `select * from chaos_cache_check` which would result in a timeout and no results.
 
## Examples

### Basic info

```sql
select
  unique_col,
  time_col,
  int_col,
  float_col
from
  chaos_cache_check;
```
