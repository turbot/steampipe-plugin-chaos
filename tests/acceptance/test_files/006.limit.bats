load "$LIB_BATS_ASSERT/load.bash"
load "$LIB_BATS_SUPPORT/load.bash"

@test "check limit passed in query" {
  run steampipe query "select limit_value from chaos_limit where c1=2 limit 10" --output=json
  limit=$(echo $output | jq .[].limit_value)
  assert_equal "$limit" "10"
}

@test "check limit returns null when distinct passed in query" {
  run steampipe query "select distinct limit_value from chaos_limit where c1=2 limit 10" --output=json
  limit=$(echo $output | jq .[].limit_value)
  assert_equal "$limit" "null"
}

@test "check limit returns null when order by passed in query" {
  run steampipe query "select limit_value from chaos_limit order by c1 limit 10" --output=json
  limit=$(echo $output | jq .[0].limit_value)
  assert_equal "$limit" "null"
}