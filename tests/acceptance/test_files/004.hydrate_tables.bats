load "$LIB_BATS_ASSERT/load.bash"
load "$LIB_BATS_SUPPORT/load.bash"

@test "test fatal_error in hydrate call" {
  run steampipe query "select fatal_error from chaos.chaos_hydrate_errors"
  assert_failure
}

# @test "status" {
#   run steampipe service status
#   assert_output --partial 'not running'
# }

@test "test retryable_error in hydrate call" {
  export STEAMPIPE_CACHE=FALSE
  run steampipe query "select retryable_error from chaos.chaos_hydrate_errors"
  assert_failure
}

# @test "status" {
#   run steampipe service status
#   assert_output --partial 'not running'
# }

@test "test ignorable_error in hydrate call" {
  run steampipe query "select ignorable_error from chaos.chaos_hydrate_errors"
  assert_success
}

# @test "status" {
#   run steampipe service status
#   assert_output --partial 'not running'
# }

@test "test delay in hydrate call [DISABLED]" {
  # run steampipe query --output json "select delay from chaos.chaos_hydrate_errors order by id"
  # assert_equal "$output" "$(cat $TEST_DATA_DIR/output_hydrate_delay.json)"
}

# @test "status" {
#   run steampipe service status
#   assert_output --partial 'not running'
# }

@test "test panic in hydrate call" {
  run steampipe query --output json "select panic from chaos.chaos_hydrate_errors"
  assert_failure
  run steampipe service stop --force
}

# @test "status" {
#   run steampipe service status
#   assert_output --partial 'not running'
# }