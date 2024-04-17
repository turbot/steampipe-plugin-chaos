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
  steampipe query "select retryable_error from chaos.chaos_list_errors order by id limit 5" > output.txt
  assert_equal "$(cat output.txt)" "$(cat $TEST_DATA_DIR/output_list_retryable_error.txt)"
  rm -f output.txt
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
  steampipe query "select should_ignore_error_after_streaming from chaos.chaos_list_errors order by id" > output.txt
  assert_equal "$(cat output.txt)" "$(cat $TEST_DATA_DIR/output_list_should_ignore_error_after_streaming.txt)"
  rm -f output.txt
}

@test "test the delay in list call" {
  steampipe query "select delay from chaos.chaos_list_errors order by id limit 5" > output.txt
  assert_equal "$(cat output.txt)" "$(cat $TEST_DATA_DIR/output_list_delay.txt)"
  rm -f output.txt
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
  steampipe query "select parent_retryable_error from chaos.chaos_list_parent_child order by id limit 5" > output.txt
  assert_equal "$(cat output.txt)" "$(cat $TEST_DATA_DIR/output_list_parent_retryable_error.txt)"
  rm -f output.txt
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
  steampipe query "select parent_should_ignore_error_after_streaming from chaos.chaos_list_parent_child order by id" > output.txt
  assert_equal "$(cat output.txt)" "$(cat $TEST_DATA_DIR/output_list_parent_should_ignore_error_after_streaming.txt)"
}

@test "test the parent_delay in list call" {
  steampipe query "select parent_delay from chaos.chaos_list_parent_child order by id limit 5" > output.txt
  assert_equal "$(cat output.txt)" "$(cat $TEST_DATA_DIR/output_list_parent_delay.txt)"
  rm -f output.txt
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
  run steampipe query "select child_retryable_error from chaos.chaos_list_parent_child"
  assert_output --partial 'retriableError'
}

@test "test the child_retryable_error (after streaming) in list call" {
  steampipe query "select child_retryable_error_after_streaming from chaos.chaos_list_parent_child order by id" > output.txt
  assert_equal "$(cat output.txt)" "$(cat $TEST_DATA_DIR/output_list_child_retryable_error_after_streaming.txt)"
  rm -f output.txt
}

@test "test the child_should_ignore_error in list call" {
  run steampipe query "select child_should_ignore_error from chaos.chaos_list_parent_child"
  assert_success
}

@test "test the child_should_ignore_error(after streaming) in list call" {
  steampipe query "select child_should_ignore_error_after_streaming from chaos.chaos_list_parent_child order by id" > output.txt
  assert_equal "$(cat output.txt)" "$(cat $TEST_DATA_DIR/output_list_child_should_ignore_error_after_streaming.txt)"
  rm -f output.txt
}

@test "test the child_delay in list call" {
  steampipe query "select child_delay from chaos.chaos_list_parent_child order by id" > output.txt
  assert_equal "$(cat output.txt)" "$(cat $TEST_DATA_DIR/output_list_child_delay.txt)"
  rm -f output.txt
}

@test "test the child_panic in list call" {
  # skip "bats unable to recover from panic"
  run steampipe query "select child_panic from chaos.chaos_list_parent_child"
  assert_output --partial 'Panic'
}

@test "service stop" {
  run steampipe service stop --force
}
