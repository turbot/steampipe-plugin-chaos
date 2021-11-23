load "$LIB_BATS_ASSERT/load.bash"
load "$LIB_BATS_SUPPORT/load.bash"

@test "test the from_json_tag transform function" {
  run steampipe query --output json "select from_json_tag from chaos_transforms order by id"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_chaos_transform_from_json_tag.json)"
}

@test "test the transform_errors during executing transform functions" {
  run steampipe query "select error from chaos_transform_errors"
  assert_output --partial 'TRANSFORM ERROR'
}

@test "test the from_qual_column transform function" {
  run steampipe query --output json "select from_qual_column from chaos_transforms where id=2"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_chaos_transform_from_qual.json)"
}

@test "test the from_optional_qual_column transform function" {
  run steampipe query --output json "select from_optional_qual_column from chaos_transforms where optional_key_column='optional-column-1'"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_chaos_transform_from_optional_qual.json)"
}

@test "test the from_field_column_single transform function" {
  run steampipe query --output json "select from_field_column_single from chaos_transforms order by id"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_chaos_transform_from_field_column_single.json)"
}

@test "test the from_field_column_multiple transform function" {
  run steampipe query --output json "select from_field_column_multiple from chaos_transforms where id=2"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_chaos_transform_from_field_column_multiple.json)"
}

@test "test the transform_method_column transform function" {
  run steampipe query --output json "select transform_method_column from chaos_transforms order by id"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_chaos_transform_from_method.json)"
}

@test "test the from_value_column transform function" {
  run steampipe query --output json "select from_value_column from chaos_transforms order by id"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_chaos_transform_from_value.json)"
}

@test "test the from_tag_column transform function" {
  run steampipe query --output json "select from_tag_column from chaos_transforms order by id"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_chaos_transform_from_tag.json)"
}

@test "test the from_constant_column transform function" {
  run steampipe query --output json "select from_constant_column from chaos_transforms order by id"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_chaos_transform_from_constant.json)"
}

@test "test the from_transform_column transform function" {
  run steampipe query --output json "select from_transform_column from chaos_transforms order by id"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_chaos_transform_from.json)"
}

@test "test the from_matrix_item_column transform function" {
  run steampipe query --output json "select from_matrix_item_column from chaos_transforms order by id"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_chaos_transform_from_matrix_item.json)"
}