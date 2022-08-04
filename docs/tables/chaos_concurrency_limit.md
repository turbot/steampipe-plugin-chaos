# Table: chaos_concurrency_limit

Chaos table to test the concurrency limit of hydrate calls.

Hydrate function `hydrateCallColumn1` has a `MaxConcurrency` of 20.
Hydrate function `hydrateCallColumn2` has a `MaxConcurrency` of 10.
Hydrate function `totalHydrateCallsColumn` has a `MaxConcurrency` of 5.

All these functions use the `doHydrateCall` function which increments the hydrate count for the given name, do some work(sleep), decrement hydrate count and returns the number of instances of this hydrate function running, and total number of hydrate calls running.
 
## Examples

### Basic info

```sql
select 
  id, 
  hydrate_call_1, 
  hydrate_call_2, 
  total_calls 
from
  chaos_concurrency_limit  limit 5;
```