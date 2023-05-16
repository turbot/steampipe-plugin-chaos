load "$LIB_BATS_ASSERT/load.bash"
load "$LIB_BATS_SUPPORT/load.bash"

@test "test the fatal_error in list call" {
  run steampipe query "select fatal_error from chaos.chaos_list_errors"
  assert_output --partial 'fatalError'
}

@test "test the fatal_error(after streaming) in list call" {
  run steampipe query "select fatal_error_after_streaming from chaos.chaos_list_errors"
  assert_output --partial 'fatalError'
}

@test "test the retryable_error in list call" {
  skip
  run steampipe query --output json "select retryable_error from chaos.chaos_list_errors order by id limit 5"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_list_retryable_error.json)"
}

@test "test the retryable_error(after streaming) in list call" {
  run steampipe query "select retryable_error_after_streaming from chaos.chaos_list_errors"
  assert_output --partial 'retriableError'
}

@test "test the should_ignore_error in list call" {
  run steampipe query "select should_ignore_error from chaos.chaos_list_errors"
  assert_success
}

@test "test the should_ignore_error(after streaming) in list call" {
  run steampipe query --output json "select should_ignore_error_after_streaming from chaos.chaos_list_errors order by id"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_list_should_ignore_error_after_streaming.json)"
}

@test "test the delay in list call" {
  run steampipe query --output json "select delay from chaos.chaos_list_errors order by id limit 5"
   assert_equal "$output" "$(cat $TEST_DATA_DIR/output_list_delay.json)"
}

@test "test the panic in list call" {
  run steampipe query "select panic from chaos.chaos_list_errors"
  assert_output --partial 'Panic'
}

@test "test the panic(after streaming) in list call" {
  run steampipe query "select panic_after_streaming from chaos.chaos_list_errors"
  assert_output --partial 'Panic'
}

@test "test the parent_fatal_error in list call" {
  run steampipe query "select parent_fatal_error from chaos.chaos_list_parent_child"
  assert_output --partial 'fatalError'
}

@test "test the parent_fatal_error(after streaming) in list call" {
  run steampipe query "select parent_fatal_error_after_streaming from chaos.chaos_list_parent_child"
  assert_output --partial 'fatalError'
}

@test "test the parent_retryable_error in list call" {
  skip
  run steampipe query --output json "select parent_retryable_error from chaos.chaos_list_parent_child order by id limit 5"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_list_parent_retryable_error.json)"
}

@test "test the parent_retryable_error(after streaming) in list call" {
  run steampipe query "select parent_retryable_error_after_streaming from chaos.chaos_list_parent_child"
  assert_output --partial 'retriableError'
}

@test "test the parent_should_ignore_error in list call" {
  run steampipe query "select parent_should_ignore_error from chaos.chaos_list_parent_child"
  assert_success
}

@test "test the parent_should_ignore_error(after streaming) in list call" {
  run steampipe query --output json "select parent_should_ignore_error_after_streaming from chaos.chaos_list_parent_child order by id"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_list_parent_should_ignore_error_after_streaming.json)"
}

@test "test the parent_delay in list call" {
  run steampipe query --output json "select parent_delay from chaos.chaos_list_parent_child order by id limit 5"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_list_parent_delay.json)"
}

@test "test the parent_panic in list call" {
  run steampipe query "select parent_panic from chaos.chaos_list_parent_child"
  assert_output --partial 'Panic'
}

@test "test the child_fatal_error in list call" {
  run steampipe query "select child_fatal_error from chaos.chaos_list_parent_child"
  assert_output --partial 'fatalError'
}

@test "test the child_fatal_error(after streaming) in list call" {
  run steampipe query "select child_fatal_error_after_streaming from chaos.chaos_list_parent_child"
  assert_output --partial 'fatalError'
}

@test "test the child_retryable_error in list call" {
  skip
  run steampipe query "select child_retryable_error from chaos.chaos_list_parent_child"
  assert_output --partial 'retriableError'
}

@test "test the child_retryable_error (after streaming) in list call" {
  skip
  run steampipe query --output json "select child_retryable_error_after_streaming from chaos.chaos_list_parent_child order by id"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_list_child_retryable_error_after_streaming.json)"
}

@test "test the child_should_ignore_error in list call" {
  run steampipe query "select child_should_ignore_error from chaos.chaos_list_parent_child"
  assert_success
}

@test "test the child_should_ignore_error(after streaming) in list call" {
  run steampipe query --output json "select child_should_ignore_error_after_streaming from chaos.chaos_list_parent_child order by id"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_list_child_should_ignore_error_after_streaming.json)"
}

@test "test the child_delay in list call" {
  run steampipe query --output json "select child_delay from chaos.chaos_list_parent_child order by id"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_list_child_delay.json)"
}

@test "test the child_panic in list call" {
  # skip "bats unable to recover from panic"
  run steampipe query "select child_panic from chaos.chaos_list_parent_child"
  assert_output --partial 'Panic'
}

@test "service stop" {
  run steampipe service stop --force
}
