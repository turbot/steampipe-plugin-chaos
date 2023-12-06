load "$LIB_BATS_ASSERT/load.bash"
load "$LIB_BATS_SUPPORT/load.bash"

@test "test all columns of different types" {
  run steampipe query --output json  "select id,string_column,boolean_column,date_time_column,double_column,ipaddress_column,json_column,cidr_column,long_string_column, array_element,epoch_column_seconds,epoch_column_milliseconds,string_to_array_column,array_to_maps_column,empty_hydrate from chaos.chaos_all_column_types order by id limit 5"
  # the expected output has been made considering Daylight Saving Time
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_all_column_types.json)"
}

@test "test dependencies between hydrate functions" {
  run steampipe query --output json "select * from chaos.chaos_hydrate_columns_dependency order by id"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_hydrate_columns_dependency.json)"
}

@test "test the execution of multiple hydrate functions and transform functions asynchronously" {
  run steampipe query --output json "select * from chaos.chaos_parallel_hydrate_columns order by id"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_parallel_hydrate_columns.json)"
}

@test "test all flavours of integer and float data types" {
  run steampipe query --output json "select * from chaos.chaos_all_numeric_column order by id"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_all_numeric_column.json)"
}

@test "test basic sql functionalities" {
  run steampipe query "create table all_columns (booleancolumn boolean, textcolumn1 CHAR(20), textcolumn2 VARCHAR(20),  textcolumn3 text, integercolumn1 smallint, integercolumn2 int, integercolumn3 SERIAL, integercolumn4 bigint,  integercolumn5 bigserial, numericColumn numeric(6,4), realColumn real, floatcolumn float,  date1 DATE,  time1 TIME,  timestamp1 TIMESTAMP, interval1 TIMESTAMPTZ, timestamp2 INTERVAL, array1 text[], jsondata jsonb, jsondata2 json, uuidcolumn UUID, ipAddress inet, macAddress macaddr, cidrRange cidr, xmlData xml, currency money)"
  run steampipe query "INSERT INTO all_columns (booleancolumn, textcolumn1, textcolumn2, textcolumn3, integercolumn1, integercolumn2, integercolumn3, integercolumn4, integercolumn5, numericColumn, realColumn, floatcolumn, date1, time1, timestamp1, interval1, timestamp2, array1, jsondata, jsondata2, uuidcolumn, ipAddress, macAddress, cidrRange, xmlData, currency) VALUES (TRUE, 'Yes', 'test for varchar', 'This is a very long text for the PostgreSQL text column', 3278, 21445454, 2147483645, 92233720368547758, 922337203685477580, 23.5141543, 4660.33777, 4.6816421254887534, '1978-02-05', '08:00:00', '2016-06-22 19:10:25-07', '2016-06-22 19:10:25-07', '1 year 2 months 3 days', '{\"(408)-589-5841\"}','{ \"customer\": \"John Doe\", \"items\": {\"product\": \"Beer\",\"qty\": 6}}', '{ \"customer\": \"John Doe\", \"items\": {\"product\": \"Beer\",\"qty\": 6}}', '6948DF80-14BD-4E04-8842-7668D9C001F5', '192.168.0.0', '08:00:2b:01:02:03', '10.1.2.3/32', '<?xml version=\"1.0\"?><book><title>Manual</title><chapter>...</chapter></book>', 922337203685477.57)"
  run steampipe query --output json "select * from all_columns"
  # the expected output has been made considering Daylight Saving Time
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_all_columns.json)"
  run steampipe query "drop table all_columns"
}
