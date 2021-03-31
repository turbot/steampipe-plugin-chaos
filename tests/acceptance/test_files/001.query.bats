load "$LIB_BATS_ASSERT/load.bash"
load "$LIB_BATS_SUPPORT/load.bash"

@test "steampipe query 1" {
  run steampipe query --output json  "select * from chaos.chaos_high_row_count order by column_0 limit 5"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output001.json)"
}

@test "steampipe query 2" {
  run steampipe query --output json  "select * from chaos.chaos_all_column_types order by id limit 5"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output002.json)"
}

@test "steampipe query 3" {
  run steampipe query --output json  "select * from chaos.chaos_high_column_count order by column_0 limit 5"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output003.json)"
}

@test "steampipe query 4" {
  run steampipe query --output json "select * from chaos.chaos_parent_child_dependency order by column_1 limit 5"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output004.json)"
}

@test "steampipe query 5" {
  run steampipe query --output json "select * from chaos.chaos_hydrate_columns_dependency order by id limit 5"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output005.json)"
}

@test "steampipe query 6" {
  run steampipe query "select * from chaos.chaos_list_error"
  assert_output --partial 'LIST ERROR'
}

@test "steampipe query 7" {
  run steampipe query "select * from chaos.chaos_list_panic"
  assert_output --partial 'LIST PANIC'
}

@test "steampipe query 8" {
  run steampipe query "select * from chaos.chaos_get_error where column_0='column_0-3'"
  assert_output --partial 'GET ERROR'
}

@test "steampipe query 9" {
  run steampipe query "select * from chaos.chaos_get_panic where column_0='column_0-3'"
  assert_output --partial 'GET PANIC'
}

@test "steampipe query 10" {
  run steampipe query "select * from chaos.chaos_transform_error"
  assert_output --partial 'TRANSFORM ERROR'
}

@test "steampipe query 11" {
  run steampipe query "select * from chaos.chaos_transform_panic" 
  assert_output --partial 'TRANSFORM PANIC'
}


#@test "steampipe query 12" {
#  run steampipe query "select * from chaos.chaos_hydrate_error"
#  assert_output --partial 'HYDRATE ERROR'
#}

#@test "steampipe query 13" {
#  run steampipe query "select * from chaos.chaos_hydrate_panic" 
#  assert_output --partial 'HYDRATE PANIC'
#}

@test "steampipe query 14" {
  run steampipe query "select * from chaos.chaos_list_delay" 
  assert_success
}

@test "steampipe query 15" {
  run steampipe query "select * from chaos.chaos_get_delay where column_0='column_0-3'" 
  assert_success
}

@test "steampipe query 16" {
  run steampipe query "select * from chaos.chaos_hydrate_delay" 
  assert_success
}

@test "steampipe query 17" {
  run steampipe query "select * from chaos.chaos_transform_delay" 
  assert_success
}

@test "steampipe query 18" {
  run steampipe query --output json "select * from chaos.chaos_parallel_hydrate_columns order by id limit 5"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output007.json)"
}

@test "steampipe query 19" {
  run steampipe query --output json "select * from chaos.chaos_all_numeric_column order by id limit 5"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output008.json)"
}

@test "steampipe query 20" {
  run steampipe query --output json "select * from chaos.chaos_get_test order by id limit 5"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output009.json)"
}

@test "steampipe query 21" {
  run steampipe query --output json "select * from chaos.chaos_default_transform order by id limit 5"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output010.json)"
}

@test "steampipe query 22" {
  run steampipe query "create table all_columns (booleancolumn boolean, textcolumn1 CHAR(20), textcolumn2 VARCHAR(20),  textcolumn3 text, integercolumn1 smallint, integercolumn2 int, integercolumn3 SERIAL, integercolumn4 bigint,  integercolumn5 bigserial, numericColumn numeric(6,4), realColumn real, floatcolumn float,  date1 DATE,  time1 TIME,  timestamp1 TIMESTAMP, interval1 TIMESTAMPTZ, timestamp2 INTERVAL, array1 text[], jsondata jsonb, jsondata2 json, uuidcolumn UUID, ipAddress inet, macAddress macaddr, cidrRange cidr, xmlData xml, currency money)"
  run steampipe query "INSERT INTO all_columns (booleancolumn, textcolumn1, textcolumn2, textcolumn3, integercolumn1, integercolumn2, integercolumn3, integercolumn4, integercolumn5, numericColumn, realColumn, floatcolumn, date1, time1, timestamp1, interval1, timestamp2, array1, jsondata, jsondata2, uuidcolumn, ipAddress, macAddress, cidrRange, xmlData, currency) VALUES (TRUE, 'Yes', 'test for varchar', 'This is a very long text for the PostgreSQL text column', 3278, 21445454, 2147483645, 92233720368547758, 922337203685477580, 23.5141543, 4660.33777, 4.6816421254887534, '1978-02-05', '08:00:00', '2016-06-22 19:10:25-07', '2016-06-22 19:10:25-07', '1 year 2 months 3 days', '{\"(408)-589-5841\"}','{ \"customer\": \"John Doe\", \"items\": {\"product\": \"Beer\",\"qty\": 6}}', '{ \"customer\": \"John Doe\", \"items\": {\"product\": \"Beer\",\"qty\": 6}}', '6948DF80-14BD-4E04-8842-7668D9C001F5', '192.168.0.0', '08:00:2b:01:02:03', '10.1.2.3/32', '<?xml version=\"1.0\"?><book><title>Manual</title><chapter>...</chapter></book>', 922337203685477.57)"
  run steampipe query --output json "select * from all_columns"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output006.json)"
  run steampipe query "drop table all_columns"
}

@test "steampipe query 23" {
  run steampipe query --output json "select * from chaos.chaos_transform_method_test order by id limit 5"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output011.json)"
}
