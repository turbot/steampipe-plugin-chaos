load "$LIB_BATS_ASSERT/load.bash"
load "$LIB_BATS_SUPPORT/load.bash"

@test "steampipe query get 1" {
  run steampipe query "select fatal_error from chaos.chaos_get_errors where id=0"
  assert_output --partial 'fatalError'
}

@test "steampipe query get 2" {
  run steampipe query "select retryable_error from chaos.chaos_get_errors where id=0"
  assert_output --partial 'retriableError'
}

@test "steampipe query get 3" {
  run steampipe query "select ignorable_error from chaos.chaos_get_errors where id=0"
  assert_success
}

@test "steampipe query get 4" {
  run steampipe query --output json "select delay from chaos.chaos_get_errors where id=0 order by id"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_get_delay.json)"
}

@test "steampipe query get 5" {
  run steampipe query --output json "select panic from chaos.chaos_get_errors where id=0"
   assert_output --partial 'Panic'
}

@test "steampipe query get 6" {
  run steampipe query "select retryable_error from chaos.chaos_get_errors_default_config where id=0"
  assert_success
}

@test "steampipe query get 7" {
  run steampipe query "select ignorable_error from chaos.chaos_get_errors_default_config where id=0"
  assert_success
}