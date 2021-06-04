load "$LIB_BATS_ASSERT/load.bash"
load "$LIB_BATS_SUPPORT/load.bash"

@test "steampipe query hydrate 1" {
  run steampipe query "select fatal_error from chaos.chaos_hydrate_errors"
  assert_output --partial 'fatalError'
}

@test "steampipe query hydrate 2" {
  run steampipe query "select retryable_error from chaos.chaos_hydrate_errors"
  assert_output --partial 'retriableError'
}

@test "steampipe query hydrate 3" {
  run steampipe query "select ignorable_error from chaos.chaos_hydrate_errors"
  assert_success
}

@test "steampipe query hydrate 4" {
  run steampipe query --output json "select delay from chaos.chaos_hydrate_errors order by id"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_hydrate_delay.json)"
}

@test "steampipe query hydrate 5" {
  run steampipe query --output json "select panic from chaos.chaos_hydrate_errors"
   assert_output --partial 'Panic'
}