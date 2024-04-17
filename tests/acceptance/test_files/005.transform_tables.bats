load "$LIB_BATS_ASSERT/load.bash"
load "$LIB_BATS_SUPPORT/load.bash"

@test "test the from_json_tag transform function" {
  steampipe query "select from_json_tag from chaos_transforms order by id" > output.txt
  assert_equal "$(cat output.txt)" "$(cat $TEST_DATA_DIR/output_chaos_transform_from_json_tag.txt)"
  rm -f output.txt
}

@test "test the transform_errors during executing transform functions" {
  skip "TODO - verify behavior and re-enable - https://github.com/turbot/steampipe-plugin-sdk/issues/710"
  run steampipe query "select error from chaos_transform_errors"
  assert_output --partial 'TRANSFORM ERROR'
}

@test "test the from_qual_column transform function" {
  steampipe query "select from_qual_column from chaos_transforms where id=2" > output.txt
  assert_equal "$(cat output.txt)" "$(cat $TEST_DATA_DIR/output_chaos_transform_from_qual.txt)"
  rm -f output.txt
}

@test "test the from_optional_qual_column transform function" {
  steampipe query "select from_optional_qual_column from chaos_transforms where optional_key_column='optional-column-1'" > output.txt
  assert_equal "$(cat output.txt)" "$(cat $TEST_DATA_DIR/output_chaos_transform_from_optional_qual.txt)"
  rm -f output.txt
}

@test "test the from_field_column_single transform function" {
  steampipe query "select from_field_column_single from chaos_transforms order by id" > output.txt
  assert_equal "$(cat output.txt)" "$(cat $TEST_DATA_DIR/output_chaos_transform_from_field_column_single.txt)"
  rm -f output.txt
}

@test "test the from_field_column_multiple transform function" {
  steampipe query "select from_field_column_multiple from chaos_transforms where id=2" > output.txt
  assert_equal "$(cat output.txt)" "$(cat $TEST_DATA_DIR/output_chaos_transform_from_field_column_multiple.txt)"
  rm -f output.txt
}

@test "test the transform_method_column transform function" {
  steampipe query "select transform_method_column from chaos_transforms order by id" > output.txt
  assert_equal "$(cat output.txt)" "$(cat $TEST_DATA_DIR/output_chaos_transform_from_method.txt)"
  rm -f output.txt
}

@test "test the from_value_column transform function" {
  steampipe query "select from_value_column from chaos_transforms order by id" > output.txt
  assert_equal "$(cat output.txt)" "$(cat $TEST_DATA_DIR/output_chaos_transform_from_value.txt)"
  rm -f output.txt
}

@test "test the from_tag_column transform function" {
  steampipe query "select from_tag_column from chaos_transforms order by id" > output.txt
  assert_equal "$(cat output.txt)" "$(cat $TEST_DATA_DIR/output_chaos_transform_from_tag.txt)"
  rm -f output.txt
}

@test "test the from_constant_column transform function" {
  steampipe query "select from_constant_column from chaos_transforms order by id" > output.txt
  assert_equal "$(cat output.txt)" "$(cat $TEST_DATA_DIR/output_chaos_transform_from_constant.txt)"
  rm -f output.txt
}

@test "test the from_transform_column transform function" {
  steampipe query "select from_transform_column from chaos_transforms order by id" > output.txt
  assert_equal "$(cat output.txt)" "$(cat $TEST_DATA_DIR/output_chaos_transform_from.txt)"
  rm -f output.txt
}

@test "test the from_matrix_item_column transform function" {
  steampipe query "select from_matrix_item_column from chaos_transforms order by id" > output.txt
  assert_equal "$(cat output.txt)" "$(cat $TEST_DATA_DIR/output_chaos_transform_from_matrix_item.txt)"
  rm -f output.txt
}