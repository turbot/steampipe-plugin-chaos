load "$LIB_BATS_ASSERT/load.bash"
load "$LIB_BATS_SUPPORT/load.bash"

@test "steampipe query list 1" {
  run steampipe query "select fatal_error from chaos.chaos_list"
  assert_output --partial 'fatalError'
}

@test "steampipe query list 2" {
  run steampipe query "select fatal_error_after_streaming from chaos.chaos_list"
  assert_output --partial 'fatalError'
}

@test "steampipe query list 3" {
  run steampipe query --output json "select retryable_error from chaos.chaos_list order by id"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_retryable_error.json)"
}

@test "steampipe query list 4" {
  run steampipe query "select retryable_error_after_streaming from chaos.chaos_list"
  assert_output --partial 'retriableError'
}

@test "steampipe query list 5" {
  run steampipe query "select should_ignore_error from chaos.chaos_list"
  assert_success
}

@test "steampipe query list 6" {
  run steampipe query --output json "select should_ignore_error_after_streaming from chaos.chaos_list order by id"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_should_ignore_error_after_streaming.json)"
}

@test "steampipe query list 7" {
  run steampipe query --output json "select delay from chaos.chaos_list order by id"
   assert_equal "$output" "$(cat $TEST_DATA_DIR/output_delay.json)"
}

@test "steampipe query list 8" {
  run steampipe query "select panic from chaos.chaos_list"
  assert_output --partial 'Panic'
}

@test "steampipe query list 9" {
  run steampipe query "select panic_after_streaming from chaos.chaos_list"
  assert_output --partial 'Panic'
}

@test "steampipe query list 10" {
  run steampipe query "select parent_fatal_error from chaos.chaos_list_parent_child"
  assert_output --partial 'fatalError'
}

@test "steampipe query list 11" {
  run steampipe query "select parent_fatal_error_after_streaming from chaos.chaos_list_parent_child"
  assert_output --partial 'fatalError'
}

@test "steampipe query list 12" {
  run steampipe query --output json "select parent_retryable_error from chaos.chaos_list_parent_child order by id"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_parent_retryable_error.json)"
}

@test "steampipe query list 13" {
  run steampipe query "select parent_retryable_error_after_streaming from chaos.chaos_list_parent_child"
  assert_output --partial 'retriableError'
}

@test "steampipe query list 14" {
  run steampipe query "select parent_should_ignore_error from chaos.chaos_list_parent_child"
  assert_success
}

@test "steampipe query list 15" {
  run steampipe query --output json "select parent_should_ignore_error_after_streaming from chaos.chaos_list_parent_child order by id"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_parent_should_ignore_error_after_streaming.json)"
}

@test "steampipe query list 16" {
  run steampipe query --output json "select parent_delay from chaos.chaos_list_parent_child order by id"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_parent_delay.json)"
}

@test "steampipe query list 17" {
  run steampipe query "select parent_panic from chaos.chaos_list_parent_child"
  assert_output --partial 'Panic'
}

@test "steampipe query list 18" {
  run steampipe query "select child_fatal_error from chaos.chaos_list_parent_child"
  assert_output --partial 'fatalError'
}

@test "steampipe query list 19" {
  run steampipe query "select child_fatal_error_after_streaming from chaos.chaos_list_parent_child"
  assert_output --partial 'fatalError'
}

@test "steampipe query list 20" {
  run steampipe query "select child_retryable_error from chaos.chaos_list_parent_child"
  assert_output --partial 'retriableError'
}

@test "steampipe query list 21" {
  run steampipe query --output json "select child_retryable_error_after_streaming from chaos.chaos_list_parent_child order by id"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_child_retryable_error_after_streaming.json)"
}

@test "steampipe query list 22" {
  run steampipe query "select child_should_ignore_error from chaos.chaos_list_parent_child"
  assert_success
}

@test "steampipe query list 23" {
  run steampipe query --output json "select child_should_ignore_error_after_streaming from chaos.chaos_list_parent_child order by id"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_child_should_ignore_error_after_streaming.json)"
}

@test "steampipe query list 24" {
  run steampipe query --output json "select child_delay from chaos.chaos_list_parent_child order by id"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_child_delay.json)"
}

@test "steampipe query list 25" {
  run steampipe query "select child_panic from chaos.chaos_list_parent_child"
  assert_output --partial 'Panic'
}