load "$LIB_BATS_ASSERT/load.bash"
load "$LIB_BATS_SUPPORT/load.bash"

@test "test the delay in get call" {
  run steampipe query --output json "select delay from chaos.chaos_get_errors where id=0 order by id"
  assert_equal "$output" "$(cat $TEST_DATA_DIR/output_get_delay.json)"
}

@test "test the panic in get call" {
  run steampipe query --output json "select panic from chaos.chaos_get_errors where id=0"
   assert_output --partial 'Panic'
}
