# Table: chaos_plugin_crash

Chaos table to print 50 rows and do an os.Exit(-1) to simulate a plugin crash.
This table is used in some steampipe acceptance tests to simulate a plugin crash.

Querying this table will always result in a `Error: rpc error: code = Unavailable desc = error reading from server: EOF (SQLSTATE HV000)`


## Examples

### Basic info

```sql
select
  id
from
  chaos_plugin_crash;
```