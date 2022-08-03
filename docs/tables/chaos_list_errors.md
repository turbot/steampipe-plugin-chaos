# Table: chaos_list_errors

Chaos table to test the List calls with all the possible scenarios like errors, panics and delays.
This table contains columns that test specific scenarios.

| Column  | Purpose  |
|---------|----------|
| fatal_error  | to test the table with fatal error  |
| fatal_error_after_streaming  | to test the table with fatal error after streaming some rows  |
| retryable_error  | to test the List function with retry config in case of non fatal error  |
| retryable_error_after_streaming  | to test the List function with retry config in case of non fatal errors occured after streaming a few rows  |
| should_ignore_error  | to test the List function with Ignorable errors  |
| should_ignore_error_after_streaming  | to test the List function with Ignorable errors occuring after already streaming some rows  |
| delay  | to test delay in List function  |
| panic  | to test panicking List function  |

## Examples

### Basic info

```sql
select
  id,
  fatal_error_after_streaming
from
  chaos_list_errors;
```