# Table: chaos_limit

Chaos table to check the limit functionality.

Columns `c1`, `c2` and `c3` are Key Quals which accept `=`, `<=` and `>=` operators respectively. Whereas columns `c4` and `c5` are timestamp and string.

Column `limit_value` gives the limit mentioned in the query. It is used in the limit tests in `006.limit.bats`. The tests check how the steampipe queries behave when a limit is passed with a quals.


## Examples

### Basic info

```sql
select
  c1,
  c4,
  c5,
  limit_value
from
  chaos_limit
where
  c1=2 limit 19;
```