load "$LIB_BATS_ASSERT/load.bash"
load "$LIB_BATS_SUPPORT/load.bash"

@test "test the fatal_error in get call" {
  run steampipe query "select fatal_error from chaos.chaos_get_errors where id=0"
  assert_output --partial 'fatalError'
}

@test "test the retryable_error in get call" {
  run steampipe query "select retryable_error from chaos.chaos_get_errors where id=0"
  assert_output --partial 'retriableError'
}

@test "test the ignorable_error in get call" {
  run steampipe query "select ignorable_error from chaos.chaos_get_errors where id=0"
  assert_success
}

@test "test the delay in get call" {
  run steampipe query --output json "select delay from chaos.chaos_get_errors where id=0 order by id"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_get_delay.json)"
}

@test "test the panic in get call" {
  run steampipe query --output json "select panic from chaos.chaos_get_errors where id=0"
   assert_output --partial 'Panic'
}

@test "test the retryable_error_default_config in case of non fatal error in get call" {
  run steampipe query "select retryable_error_default_config from chaos.chaos_get_errors_default_config where id=0"
  assert_success
}

@test "test the ignorable_error_default_config in case of non fatal error in get call" {
  run steampipe query "select ignorable_error_default_config from chaos.chaos_get_errors_default_config where id=0"
  assert_success
}