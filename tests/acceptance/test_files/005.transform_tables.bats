load "$LIB_BATS_ASSERT/load.bash"
load "$LIB_BATS_SUPPORT/load.bash"

@test "steampipe query transform 1" {
  run steampipe query "select error from chaos_transform_errors"
  assert_output --partial 'TRANSFORM ERROR'
}

@test "steampipe query transform 2" {
  run steampipe query "select panic from chaos_transform_errors"
  assert_output --partial 'TRANSFORM PANIC'
}

@test "steampipe query transform 3" {
  run steampipe query "select delay from chaos_transform_errors"
  assert_success
}

@test "steampipe query transform 4" {
  run steampipe query --output json "select from_json_tag from chaos_transforms order by id"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_chaos_transform_from_json_tag.json)"
}

@test "steampipe query transform 5" {
  run steampipe query --output json "select from_qual_column from chaos_transforms where id=2"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_chaos_transform_from_qual.json)"
}

@test "steampipe query transform 6" {
  run steampipe query --output json "select from_optional_qual_column from chaos_transforms where optional_key_column='optional-column-1'"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_chaos_transform_from_optional_qual.json)"
}

@test "steampipe query transform 7" {
  run steampipe query --output json "select from_field_column_single from chaos_transforms order by id"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_chaos_transform_from_field_column_single.json)"
}

@test "steampipe query transform 8" {
  run steampipe query --output json "select from_field_column_multiple from chaos_transforms where id=2"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_chaos_transform_from_field_column_multiple.json)"
}

@test "steampipe query transform 9" {
  run steampipe query --output json "select transform_method_column from chaos_transforms order by id"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_chaos_transform_from_method.json)"
}

@test "steampipe query transform 10" {
  run steampipe query --output json "select from_value_column from chaos_transforms order by id"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_chaos_transform_from_value.json)"
}

@test "steampipe query transform 11" {
  run steampipe query --output json "select from_tag_column from chaos_transforms order by id"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_chaos_transform_from_tag.json)"
}

@test "steampipe query transform 12" {
  run steampipe query --output json "select from_constant_column from chaos_transforms order by id"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_chaos_transform_from_constant.json)"
}

@test "steampipe query transform 13" {
  run steampipe query --output json "select from_transform_column from chaos_transforms order by id"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_chaos_transform_from.json)"
}

@test "steampipe query transform 14" {
  run steampipe query --output json "select from_matrix_item_column from chaos_transforms order by id"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_chaos_transform_from_matrix_item.json)"
}