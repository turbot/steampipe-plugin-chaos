load "$LIB_BATS_ASSERT/load.bash"
load "$LIB_BATS_SUPPORT/load.bash"

setup() {
  steampipe service start > /dev/null
}

teardown() {
  rm -f output?.json
  steampipe service stop
}

@test "check cache functionality when querying same columns" {
  skip "TODO - verify behavior and re-enable - https://github.com/turbot/steampipe-plugin-sdk/issues/710"
  steampipe query "select unique_col, a, b, c from chaos_cache_check" --output json &> output1.json
  # store the unique number from 1st query in `content`
  content=$(cat output1.json | jq '.rows[0].unique_col')

  steampipe query "select unique_col, a, b, c from chaos_cache_check" --output json &> output2.json
  # store the unique number from 2nd query in `new_content`
  new_content=$(cat output2.json | jq '.rows[0].unique_col')

  echo $content
  echo $new_content

  # verify that `content` and `new_content` are the same
  assert_equal "$new_content" "$content"
}

@test "check cache functionality when the second query's columns is a subset of the first" {
  skip "TODO - verify behavior and re-enable - https://github.com/turbot/steampipe-plugin-sdk/issues/710"
  steampipe query "select unique_col, a, b, c from chaos_cache_check" --output json &> output1.json
  # store the unique number from 1st query in `content`
  content=$(cat output1.json | jq '.rows[0].unique_col')

  steampipe query "select unique_col, a, b from chaos_cache_check" --output json &> output2.json
  # store the unique number from 2nd query in `new_content`
  new_content=$(cat output2.json | jq '.rows[0].unique_col')

  echo $content
  echo $new_content

  # verify that `content` and `new_content` are the same
  assert_equal "$new_content" "$content"
}

@test "check cache functionality multiple queries with same columns" {
  skip "TODO - verify behavior and re-enable - https://github.com/turbot/steampipe-plugin-sdk/issues/710"
  run steampipe query "select unique_col, a, b, c from chaos_cache_check" --output json &> output1.json
  # store the unique number from 1st query in `content`
  content=$(cat output1.json | jq '.rows[0].unique_col')

  run steampipe query "select unique_col, a, b, c from chaos_cache_check" --output json &> output2.json
  # store the unique number from 2nd query in `new_content`
  content2=$(cat output2.json | jq '.rows[0].unique_col')

  run steampipe query "select unique_col, a, b, c from chaos_cache_check" --output json &> output3.json
  # store the unique number from 3rd query in `new_content`
  content3=$(cat output3.json | jq '.rows[0].unique_col')

  run steampipe query "select unique_col, a, b, c from chaos_cache_check" --output json &> output4.json
  # store the unique number from 4th query in `new_content`
  content4=$(cat output4.json | jq '.rows[0].unique_col')

  echo $content
  echo $content2
  echo $content3
  echo $content4

  # verify that `content` from 1st query is same as the next queries
  assert_equal "$content2" "$content"
  assert_equal "$content3" "$content"
  assert_equal "$content4" "$content"
}

@test "check cache functionality when multiple query's columns are a subset of the first" {
  skip "TODO - verify behavior and re-enable - https://github.com/turbot/steampipe-plugin-sdk/issues/710"
  run steampipe query "select unique_col, a, b, c from chaos_cache_check" --output json &> output1.json
  # store the unique number from 1st query in `content`
  content=$(cat output1.json | jq '.rows[0].unique_col')

  run steampipe query "select unique_col, a, b from chaos_cache_check" --output json &> output2.json
  # store the unique number from 2nd query in `new_content`
  content2=$(cat output2.json | jq '.rows[0].unique_col')

  run steampipe query "select unique_col, a, b from chaos_cache_check" --output json &> output3.json
  # store the unique number from 3rd query in `new_content`
  content3=$(cat output3.json | jq '.rows[0].unique_col')

  run steampipe query "select unique_col, a, b from chaos_cache_check" --output json &> output4.json
  # store the unique number from 4th query in `new_content`
  content4=$(cat output4.json | jq '.rows[0].unique_col')

  echo $content
  echo $content2
  echo $content3
  echo $content4

  # verify that `content` from 1st query is same as the next queries
  assert_equal "$content2" "$content"
  assert_equal "$content3" "$content"
  assert_equal "$content4" "$content"
}

@test "check cache functionality when the second query has more columns than the first" {
  skip "TODO - verify behavior and re-enable - https://github.com/turbot/steampipe-plugin-sdk/issues/710"
  steampipe query "select unique_col, a, b from chaos_cache_check" --output json &> output1.json
  # store the unique number from 1st query in `content`
  content=$(cat output1.json | jq '.rows[0].unique_col')

  steampipe query "select unique_col, a, b, c from chaos_cache_check" --output json &> output2.json
  # store the unique number from 2nd query in `new_content`
  new_content=$(cat output2.json | jq '.rows[0].unique_col')

  echo $content
  echo $new_content

  # verify that `content` and `new_content` are not the same
  if [[ "$content" == "$new_content" ]]; then
    flag=1
  else
    flag=0
  fi
  assert_equal "$flag" "0"
}

@test "check cache functionality when the both the queries have same limits" {
  skip "TODO - verify behavior and re-enable - https://github.com/turbot/steampipe-plugin-sdk/issues/710"
  steampipe query "select unique_col, a, b, c from chaos_cache_check limit 3" --output json &> output1.json
  # store the unique number from 1st query in `content`
  content=$(cat output1.json | jq '.rows[0].unique_col')

  steampipe query "select unique_col, a, b, c from chaos_cache_check limit 3" --output json &> output2.json
  # store the unique number from 2nd query in `new_content`
  new_content=$(cat output2.json | jq '.rows[0].unique_col')

  echo $content
  echo $new_content

  # verify that `content` and `new_content` are the same
  assert_equal "$new_content" "$content"
}

@test "check cache functionality when first query has no limit but second query has a limit" {
  skip "TODO - verify behavior and re-enable - https://github.com/turbot/steampipe-plugin-sdk/issues/710"
  steampipe query "select unique_col, a, b, c from chaos_cache_check" --output json &> output1.json
  # store the unique number from 1st query in `content`
  content=$(cat output1.json | jq '.rows[0].unique_col')

  steampipe query "select unique_col, a, b, c from chaos_cache_check limit 3" --output json &> output2.json
  # store the unique number from 2nd query in `new_content`
  new_content=$(cat output2.json | jq '.rows[0].unique_col')

  echo $content
  echo $new_content

  # verify that `content` and `new_content` are the same
  assert_equal "$new_content" "$content"
}

@test "check cache functionality when second query has lower limit than first" {
  skip "TODO - verify behavior and re-enable - https://github.com/turbot/steampipe-plugin-sdk/issues/710"
  steampipe query "select unique_col, a, b, c from chaos_cache_check limit 4" --output json &> output1.json
  # store the unique number from 1st query in `content`
  content=$(cat output1.json | jq '.rows[0].unique_col')

  steampipe query "select unique_col, a, b, c from chaos_cache_check limit 3" --output json &> output2.json
  # store the unique number from 2nd query in `new_content`
  new_content=$(cat output2.json | jq '.rows[0].unique_col')

  echo $content
  echo $new_content

  # verify that `content` and `new_content` are the same
  assert_equal "$new_content" "$content"
}

@test "check cache functionality when second query has higher limit than first" {
  skip "TODO - verify behavior and re-enable - https://github.com/turbot/steampipe-plugin-sdk/issues/710"
  steampipe query "select unique_col, a, b, c from chaos_cache_check limit 3" --output json &> output1.json
  # store the unique number from 1st query in `content`
  content=$(cat output1.json | jq '.rows[0].unique_col')

  steampipe query "select unique_col, a, b, c from chaos_cache_check limit 4" --output json &> output2.json
  # store the unique number from 2nd query in `new_content`
  new_content=$(cat output2.json | jq '.rows[0].unique_col')

  echo $content
  echo $new_content

  # verify that `content` and `new_content` are not the same
  if [[ "$content" == "$new_content" ]]; then
    flag=1
  else
    flag=0
  fi
  assert_equal "$flag" "0"
}

@test "stop service" {
  skip "TODO - verify behavior and re-enable - https://github.com/turbot/steampipe-plugin-sdk/issues/710"
  steampipe service stop --force
}
