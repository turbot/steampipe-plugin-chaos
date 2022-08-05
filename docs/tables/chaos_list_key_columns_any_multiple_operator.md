# Table: chaos_list_key_columns_any_multiple_operator

Chaos table to test a list function supporting optional key columns(which supports only multiple operators).

NOTE: The columns defined as optional key columns must be used as quals in the query.

## Examples

### Basic info

```sql
select 
  id, 
  col_1 
from 
  chaos_list_key_columns_any_multiple_operator 
where 
  id > 5;
```