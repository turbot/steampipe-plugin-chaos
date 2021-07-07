load "$LIB_BATS_ASSERT/load.bash"
load "$LIB_BATS_SUPPORT/load.bash"

@test "steampipe query key columns 1" {
  run steampipe query --output json "select * from chaos_list_single_key_columns where id=2"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_list_single_key_columns.json)"
}

@test "steampipe query key columns 2" {
  run steampipe query --output json "select * from chaos_list_single_key_columns"
  assert_failure
}

@test "steampipe query key columns 3" {
  run steampipe query --output json "select * from chaos_list_all_key_columns where id=2 and column_a='column-1'"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_list_all_key_columns.json)"
}

@test "steampipe query key columns 4" {
  run steampipe query --output json "select * from chaos_list_all_key_columns where id=2"
  assert_failure
}

@test "steampipe query key columns 5" {
  run steampipe query --output json "select * from chaos_list_any_key_columns where id=2 and column_a='column-1'"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_list_any_key_columns.json)"
}

@test "steampipe query key columns 6" {
  run steampipe query --output json "select * from chaos_get_single_key_columns where id=2"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_get_single_key_columns.json)"
}

@test "steampipe query key columns 7" {
  run steampipe query --output json "select * from chaos_get_single_key_columns"
  assert_failure
}

@test "steampipe query key columns 8" {
  run steampipe query --output json "select * from chaos_get_all_key_columns where id=2 and column_a='column-1'"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_get_all_key_columns.json)"
}

@test "steampipe query key columns 9" {
  run steampipe query --output json "select * from chaos_get_all_key_columns"
  assert_failure
}

@test "steampipe query key columns 10" {
  run steampipe query --output json "select * from chaos_get_any_key_columns where id=2 and column_a='column-1'"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_get_any_key_columns.json)"
}

@test "steampipe query key columns 11" {
  run steampipe query --output json "select * from chaos_list_key_column_single_equal where id=2"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_list_key_column_single_equal.json)"
}

@test "steampipe query key columns 12" {
  run steampipe query --output json "select * from chaos_list_key_columns_any_multiple_operator where id<5 and col_1>92 order by id"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_list_key_columns_any_multiple_operator.json)"
}

@test "steampipe query key columns 13" {
  run steampipe query --output json "select * from chaos_list_key_columns_all_multiple_operator where id<5 and col_1>92 order by id"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_list_key_columns_all_multiple_operator.json)"
}

@test "steampipe query key columns 14" {
  run steampipe query "select * from chaos_list_key_columns_all_multiple_operator where id=5 and col_1=92 order by id"
  assert_failure
}

@test "steampipe query key columns 15" {
  run steampipe query --output json "select * from chaos_get_key_column_single_equal where id=2"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_get_key_column_single_equal.json)"
}

@test "steampipe query key columns 16" {
  run steampipe query --output json "select * from chaos_get_key_columns_any_multiple_operator where id=2"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_get_key_columns_any_multiple_operator.json)"
}

@test "steampipe query key columns 17" {
  run steampipe query --output json "select * from chaos_get_key_columns_any_multiple_operator where id>2"
  assert_failure
}

@test "steampipe query key columns 18" {
  run steampipe query --output json "select * from chaos_get_key_columns_all_multiple_operator where id=2 and col_1=92"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_get_key_column_all_multiple_operator.json)"
}

@test "steampipe query key columns 19" {
  run steampipe query --output json "select * from chaos_get_key_columns_all_multiple_operator where id<2 and col_1>92"
  assert_failure
}

@test "steampipe query key columns 20" {
  run steampipe query --output json "select * from chaos_list_single_key_columns where id in (1,2,3) order by id"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_list_single_key_columns_in.json)"
}

@test "steampipe query key columns 21" {
  run steampipe query --output json "select * from chaos_get_single_key_columns where id in (1,2,3) order by id"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_get_single_key_columns_in.json)"
}

@test "steampipe query key columns 22" {
  run steampipe query --output json "select * from chaos_get_single_key_columns where id in (1,2,3) order by id"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_get_single_key_columns_in.json)"
}