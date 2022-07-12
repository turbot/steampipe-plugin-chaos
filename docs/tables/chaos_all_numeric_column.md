# Table: chaos_all_numeric_column

Chaos table to test all flavours of integer and float data types. This table has numeric columns for all possible numeric data types used in steampipe. 

This is used in the acceptance test `test all flavours of integer and float data types(001.query.bats)`

This table currently returns 10 rows.

## Examples

### Basic info

```sql
select
  int_data,
  int8_data,
  int16_data,
  int32_data,
  uint_data,
  uint16_data,
  uint32_data,
  float64_data,
  float32_data
from
  chaos_all_numeric_column;
```
