load "$LIB_BATS_ASSERT/load.bash"
load "$LIB_BATS_SUPPORT/load.bash"

@test "check limit passed in query" {
  run steampipe query "select limit_value from chaos_limit_verify_rows_remaining where total_row_count=1 limit 10" --output=json
  limit=$(echo $output | jq .[].limit_value)
  assert_equal "$limit" "10"
}

@test "check limit returns null when distinct passed in query" {
  run steampipe query "select distinct limit_value from chaos_limit_verify_rows_remaining where total_row_count=1 limit 10" --output=json
  limit=$(echo $output | jq .[].limit_value)
  assert_equal "$limit" "null"
}

@test "check limit returns null when order by passed in query" {
  run steampipe query "select limit_value from chaos_limit_verify_rows_remaining order by total_row_count limit 10" --output=json
  limit=$(echo $output | jq .[0].limit_value)
  assert_equal "$limit" "null"
}

@test "check limit returns null when there is a non-key column passed in query" {
  run steampipe query "select limit_value from chaos.chaos_limit_verify_rows_remaining where rows_streamed=6 limit 10" --output=json

  # limit is returned as null since rows_streamed is a not a key column 
  limit=$(echo $output | jq .[0].limit_value)
  assert_equal "$limit" "null"
}

@test "check limit returns null when there is a key column passed in query but with wrong operator" {
  run steampipe query "select limit_value from chaos.chaos_limit_verify_rows_remaining where total_row_count > 4 limit 10" --output=json

  # limit is returned as null since c2 does not support = operator 
  limit=$(echo $output | jq .[0].limit_value)
  assert_equal "$limit" "null"
}

@test "check limit when a key column passed in query" {
  run steampipe query "select limit_value from chaos.chaos_limit_verify_rows_remaining where total_row_count=4 limit 10" --output=json

  # limit is not returned as null since total_row_count is a key column 
  limit=$(echo $output | jq .[0].limit_value)
  assert_equal "$limit" "10"
}
