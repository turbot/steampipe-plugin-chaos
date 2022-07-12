# Table: chaos_all_column_types

Chaos table to test all columns of different types. This table has columns for all possible data types used in steampipe.

This is used in `@test "test all columns of different types" in 001.query.bats`

This table currently returns 100 rows.

## Examples

### Basic info

```sql
select
  boolean_column,
  date_time_column,
  ipaddress_column,
  json_column,
  long_string_column,
  epoch_column_seconds,
  string_to_array_column,
  inet_column,
  ltree_column
from
  chaos_all_column_types;
```
