load "$LIB_BATS_ASSERT/load.bash"
load "$LIB_BATS_SUPPORT/load.bash"

@test "test fatal_error in hydrate call" {
  run steampipe query "select fatal_error from chaos.chaos_hydrate_errors"
  assert_output --partial 'fatalError'
}

@test "test retryable_error in hydrate call" {
  export STEAMPIPE_CACHE=FALSE
  run steampipe query "select retryable_error from chaos.chaos_hydrate_errors"
  assert_success
}

@test "test ignorable_error in hydrate call" {
  run steampipe query "select ignorable_error from chaos.chaos_hydrate_errors"
  assert_success
}

@test "test delay in hydrate call" {
  steampipe query "select delay from chaos.chaos_hydrate_errors order by id" > output.txt
  assert_equal "$(cat output.txt)" "$(cat $TEST_DATA_DIR/output_hydrate_delay.txt)"
  rm -f output.txt
}

@test "test panic in hydrate call" {
  run steampipe query --output json "select panic from chaos.chaos_hydrate_errors"
  assert_output --partial 'failed with panic'
  run steampipe service stop --force
}
