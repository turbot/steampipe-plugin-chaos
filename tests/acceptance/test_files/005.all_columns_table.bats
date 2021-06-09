load "$LIB_BATS_ASSERT/load.bash"
load "$LIB_BATS_SUPPORT/load.bash"

@test "steampipe query 24" {
  run steampipe query --output json  "select id, string_column, date_time_column from chaos.chaos_all_column_types where id='0'"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output012.json)"
}

@test "steampipe query 25" {
  run steampipe query --output json  "select id, string_column, date_time_column from chaos.chaos_all_column_types where string_column='stringValuesomething-0'"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output013.json)"
}

@test "steampipe query 26" {
  run steampipe query --output json  "select id, string_column, date_time_column from chaos.chaos_all_column_types where boolean_column=true order by id"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output014.json)"
}

@test "steampipe query 27" {
  run steampipe query --output json  "select id, string_column, date_time_column from chaos.chaos_all_column_types where date_time_column='2001-02-20 01:28:00' order by id"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output015.json)"
}

@test "steampipe query 28" {
  run steampipe query --output json  "select id, string_column, date_time_column from chaos.chaos_all_column_types where double_column='0' order by id"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output016.json)"
}

@test "steampipe query 29" {
  run steampipe query --output json  "select id, string_column, date_time_column from chaos.chaos_all_column_types where ipaddress_column='10.0.1.4' order by id"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output017.json)"
}

@test "steampipe query 30" {
  run steampipe query --output json  "select id, string_column, date_time_column from chaos.chaos_all_column_types where cidr_column='10.0.0.0/24' order by id"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output018.json)"
}

@test "steampipe query 31" {
  run steampipe query --output json  "select id, string_column, date_time_column from chaos.chaos_all_column_types where array_element='{\"Key\":\"stringValuesomething-0\",\"Value\":\"value\"}' order by id"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output019.json)"
}

@test "steampipe query 32" {
  run steampipe query --output json  "select id, string_column, date_time_column from chaos.chaos_all_column_types where epoch_column_seconds='2021-01-19 13:39:39' order by id"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output020.json)"
}

@test "steampipe query 33" {
  run steampipe query --output json  "select id, string_column, date_time_column from chaos.chaos_all_column_types where epoch_column_milliseconds='2021-01-19 11:53:18' order by id"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output021.json)"
}

@test "steampipe query 34" {
  run steampipe query --output json  "select id, string_column, date_time_column from chaos.chaos_all_column_types where json_column='{\"Id\":0,\"Name\":\"stringValuesomething-0\",\"Statement\":{\"Action\":\"iam:GetContextKeysForCustomPolicy\",\"Effect\":\"Allow\"}}' order by id"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output022.json)"
}

