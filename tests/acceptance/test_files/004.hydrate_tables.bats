load "$LIB_BATS_ASSERT/load.bash"
load "$LIB_BATS_SUPPORT/load.bash"

@test "test ignorable_error in hydrate call" {
  run steampipe query "select ignorable_error from chaos.chaos_hydrate_errors"
  assert_success
}

@test "test fatal_error in hydrate call" {
  run steampipe query "select fatal_error from chaos.chaos_hydrate_errors"
  assert_failure
}

@test "test retryable_error in hydrate call [DISABLED]" {
  export STEAMPIPE_CACHE=FALSE
  run steampipe query "select retryable_error from chaos.chaos_hydrate_errors"
  assert_failure
}

@test "status" {
  steampipe service status
  ps -ef | grep steampipe
  assert_failure
}

@test "test delay in hydrate call [DISABLED]" {
  skip
  run steampipe query --output json "select delay from chaos.chaos_hydrate_errors order by id"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_hydrate_delay.json)"
}

@test "test panic in hydrate call [DISABLED]" {
  skip
  run steampipe query --output json "select panic from chaos.chaos_hydrate_errors"
  assert_failure
  run steampipe service stop --force
}
