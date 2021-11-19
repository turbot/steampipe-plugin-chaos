load "$LIB_BATS_ASSERT/load.bash"
load "$LIB_BATS_SUPPORT/load.bash"

@test "check cache functionality when the second query fetches columns that are a subset of the first" {
  run steampipe plugin install chaos
  run steampipe service start

  steampipe query "select int_col, a, b, unique_col from chaos_cache_check where int_col < 3 order by id" --output json &> output1.json
  # store the time from 1st query in `content`
  content=$(cat output1.json | jq '.[0].time_now')

  steampipe query "select int_col, a, b, unique_col from chaos_cache_check where int_col < 4 order by id" --output json &> output2.json
  # store the time from 2nd query in `new_content`
  new_content=$(cat output2.json | jq '.[0].time_now')

  echo $content
  echo $new_content

  # verify that `content` and `new_content` are the same
  assert_equal "$new_content" "$content"

  rm -f output?.json
  run steampipe service stop
}

@test "check cache functionality when the second query fetches columns that are a subset of the first" {
  run steampipe plugin install chaos
  run steampipe service start

  steampipe query "select int_col, a, b, unique_col from chaos_cache_check where int_col < 3 order by id" --output json &> output1.json
  # store the time from 1st query in `content`
  content=$(cat output1.json | jq '.[0].time_now')

  steampipe query "select int_col, a, b, unique_col from chaos_cache_check where int_col < 2 order by id" --output json &> output2.json
  # store the time from 2nd query in `new_content`
  new_content=$(cat output2.json | jq '.[0].time_now')

  echo $content
  echo $new_content

  # verify that `content` and `new_content` are the same
  assert_equal "$new_content" "$content"

  rm -f output?.json
  run steampipe service stop
}