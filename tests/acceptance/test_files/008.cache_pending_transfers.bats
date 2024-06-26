load "$LIB_BATS_ASSERT/load.bash"
load "$LIB_BATS_SUPPORT/load.bash"

setup() {
  steampipe service start > /dev/null
}

teardown() {
  rm -f output?.json
  steampipe service stop
}

######## PENDING TRANSFERS TESTS ###########

@test "check cache functionality when querying same columns(first query in background)" {
  skip "TODO - verify behavior and re-enable - https://github.com/turbot/steampipe-plugin-sdk/issues/710"
  steampipe query "select unique_col, a, b, c, delay from chaos_cache_check" --output json &> output1.json &
  sleep 1
  steampipe query "select unique_col, a, b, c, delay from chaos_cache_check" --output json &> output2.json

  # store the unique number from 1st query in `content`
  content=$(cat output1.json | jq '.rows[0].unique_col')
  # store the unique number from 2nd query in `new_content`
  new_content=$(cat output2.json | jq '.rows[0].unique_col')

  echo $content
  echo $new_content

  # verify that `content` and `new_content` are the same
  assert_equal "$new_content" "$content"
}

@test "check cache functionality when the second query's columns is a subset of the first(first query in background)" {
  skip "TODO - verify behavior and re-enable - https://github.com/turbot/steampipe-plugin-sdk/issues/710"
  steampipe query "select unique_col, a, b, c, delay from chaos_cache_check" --output json &> output1.json &
  sleep 1
  steampipe query "select unique_col, a, b, delay from chaos_cache_check" --output json &> output2.json

  # store the unique number from 1st query in `content`
  content=$(cat output1.json | jq '.rows[0].unique_col')
  # store the unique number from 2nd query in `new_content`
  new_content=$(cat output2.json | jq '.rows[0].unique_col')

  echo $content
  echo $new_content

  # verify that `content` and `new_content` are the same
  assert_equal "$new_content" "$content"
}

@test "check cache functionality multiple queries with same columns(first query in background)" {
  skip "TODO - verify behavior and re-enable - https://github.com/turbot/steampipe-plugin-sdk/issues/710"
  steampipe query "select unique_col, a, b, c, delay from chaos_cache_check" --output json &> output1.json &
  sleep 1
  steampipe query "select unique_col, a, b, c, delay from chaos_cache_check" --output json &> output2.json
  steampipe query "select unique_col, a, b, c, delay from chaos_cache_check" --output json &> output3.json
  steampipe query "select unique_col, a, b, c, delay from chaos_cache_check" --output json &> output4.json

  # store the unique number from 1st query in `content`
  content=$(cat output1.json | jq '.rows[0].unique_col')
  # store the unique number from 2nd query in `new_content`
  content2=$(cat output2.json | jq '.rows[0].unique_col')
  # store the unique number from 3rd query in `new_content`
  content3=$(cat output3.json | jq '.rows[0].unique_col')
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

@test "check cache functionality when multiple query's columns are a subset of the first(first query in background)" {
  skip "TODO - verify behavior and re-enable - https://github.com/turbot/steampipe-plugin-sdk/issues/710"
  steampipe query "select unique_col, a, b, c, delay from chaos_cache_check" --output json &> output1.json &
  sleep 1
  steampipe query "select unique_col, a, b from chaos_cache_check" --output json &> output2.json
  steampipe query "select unique_col, a, b from chaos_cache_check" --output json &> output3.json
  steampipe query "select unique_col, a, b from chaos_cache_check" --output json &> output4.json

  # store the unique number from 1st query in `content`
  content=$(cat output1.json | jq '.rows[0].unique_col')
  # store the unique number from 2nd query in `new_content`
  content2=$(cat output2.json | jq '.rows[0].unique_col')
  # store the unique number from 3rd query in `new_content`
  content3=$(cat output3.json | jq '.rows[0].unique_col')
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

@test "check cache functionality when the second query has more columns than the first(first query in background)" {
  skip "TODO - verify behavior and re-enable - https://github.com/turbot/steampipe-plugin-sdk/issues/710"
  steampipe query "select unique_col, a, b, delay from chaos_cache_check" --output json &> output1.json &
  sleep 1
  steampipe query "select unique_col, a, b, delay, c from chaos_cache_check" --output json &> output2.json

  # store the unique number from 1st query in `content`
  content=$(cat output1.json | jq '.rows[0].unique_col')
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

@test "check cache functionality when the both the queries have same limits(first query in background)" {
  skip "TODO - verify behavior and re-enable - https://github.com/turbot/steampipe-plugin-sdk/issues/710"
  steampipe query "select unique_col, a, b, c, delay from chaos_cache_check limit 3" --output json &> output1.json &
  sleep 1
  steampipe query "select unique_col, a, b, c, delay from chaos_cache_check limit 3" --output json &> output2.json

  # store the unique number from 1st query in `content`
  content=$(cat output1.json | jq '.rows[0].unique_col')
  # store the unique number from 2nd query in `new_content`
  new_content=$(cat output2.json | jq '.rows[0].unique_col')

  echo $content
  echo $new_content

  # verify that `content` and `new_content` are the same
  assert_equal "$new_content" "$content"
}

@test "check cache functionality when first query has no limit but second query has a limit(first query in background)" {
  skip "TODO - verify behavior and re-enable - https://github.com/turbot/steampipe-plugin-sdk/issues/710"
  steampipe query "select unique_col, a, b, c, delay from chaos_cache_check" --output json &> output1.json &
  sleep 1
  steampipe query "select unique_col, a, b, c, delay from chaos_cache_check limit 3" --output json &> output2.json

  # store the unique number from 1st query in `content`
  content=$(cat output1.json | jq '.rows[0].unique_col')
  # store the unique number from 2nd query in `new_content`
  new_content=$(cat output2.json | jq '.rows[0].unique_col')

  echo $content
  echo $new_content

  # verify that `content` and `new_content` are the same
  assert_equal "$new_content" "$content"
}

@test "check cache functionality when second query has no limit but first query has a limit(first query in background)" {
  skip "TODO - verify behavior and re-enable - https://github.com/turbot/steampipe-plugin-sdk/issues/710"
  steampipe query "select unique_col, a, b, c, delay from chaos_cache_check limit 3" --output json &> output1.json &
  sleep 1
  steampipe query "select unique_col, a, b, c, delay from chaos_cache_check" --output json &> output2.json

  # store the unique number from 1st query in `content`
  content=$(cat output1.json | jq '.rows[0].unique_col')
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

@test "check cache functionality when second query has lower limit than first(first query in background)" {
  skip "TODO - verify behavior and re-enable - https://github.com/turbot/steampipe-plugin-sdk/issues/710"
  steampipe query "select unique_col, a, b, c, delay from chaos_cache_check limit 4" --output json &> output1.json &
  sleep 1
  steampipe query "select unique_col, a, b, c, delay from chaos_cache_check limit 3" --output json &> output2.json

  # store the unique number from 1st query in `content`
  content=$(cat output1.json | jq '.rows[0].unique_col')
  # store the unique number from 2nd query in `new_content`
  new_content=$(cat output2.json | jq '.rows[0].unique_col')

  echo $content
  echo $new_content

  # verify that `content` and `new_content` are the same
  assert_equal "$new_content" "$content"
}

@test "check cache functionality when second query has higher limit than first(first query in background)" {
  skip "TODO - verify behavior and re-enable - https://github.com/turbot/steampipe-plugin-sdk/issues/710"
  steampipe query "select unique_col, a, b, c, delay from chaos_cache_check limit 3" --output json &> output1.json &
  sleep 1
  steampipe query "select unique_col, a, b, c, delay from chaos_cache_check limit 4" --output json &> output2.json

  # store the unique number from 1st query in `content`
  content=$(cat output1.json | jq '.rows[0].unique_col')
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

######## ERROR AND TIMEOUT TESTS ###########

@test "check cache functionality when first query returns error, other queries should not cache(first query in background)" {
  skip "TODO - verify behavior and re-enable - https://github.com/turbot/steampipe-plugin-sdk/issues/710"
  steampipe query "select unique_col, a, b, c, error_after_delay from chaos_cache_check" --output json &> output1.json &
  sleep 1
  steampipe query "select unique_col, a, b from chaos_cache_check" --output json &> output2.json
  steampipe query "select unique_col, a, b, c from chaos_cache_check" --output json &> output3.json

  # store the unique number from 2nd query in `new_content`
  content2=$(cat output2.json | jq '.rows[0].unique_col')
  # store the unique number from 3rd query in `new_content`
  content3=$(cat output3.json | jq '.rows[0].unique_col')

  echo $content2
  echo $content3

  # verify that `content2` and `content3` are not the same
  if [[ "$content2" == "$content3" ]]; then
    flag=1
  else
    flag=0
  fi
  assert_equal "$flag" "0"
}

@test "check cache functionality when first query times out, other queries should cache(first query in background)" {
  skip

  steampipe query "select unique_col, a, b, c, long_delay from chaos_cache_check" --output json &> output1.json &
  sleep 1
  steampipe query "select unique_col, a, b from chaos_cache_check" --output json &> output2.json
  steampipe query "select unique_col, a, b from chaos_cache_check" --output json &> output3.json

  # store the unique number from 2nd query in `new_content`
  content2=$(cat output2.json | jq '.rows[0].unique_col')
  # store the unique number from 3rd query in `new_content`
  content3=$(cat output3.json | jq '.rows[0].unique_col')

  echo $content2
  echo $content3

  # verify that `content2` and `content3` are the same
  assert_equal "$content2" "$content3"
}

@test "check cache functionality when first query times out, other queries should not cache(first query in background)" {
  skip 

  steampipe query "select unique_col, a, b, c, long_delay from chaos_cache_check" --output json &> output1.json &
  sleep 1
  steampipe query "select unique_col, a, b from chaos_cache_check" --output json &> output2.json
  steampipe query "select unique_col, a, b, c from chaos_cache_check" --output json &> output3.json

  # store the unique number from 2nd query in `new_content`
  content2=$(cat output2.json | jq '.rows[0].unique_col')
  # store the unique number from 3rd query in `new_content`
  content3=$(cat output3.json | jq '.rows[0].unique_col')

  echo $content2
  echo $content3

  # verify that `content2` and `content3` are not the same
  if [[ "$content2" == "$content3" ]]; then
    flag=1
  else
    flag=0
  fi
  assert_equal "$flag" "0"
}

# Test to validate the issue where after waiting for pending cache data, query cache is failing to load the data for the completed pending item
# This is because the cache key of the pending item does not match the cache key used to write data which satisfies the pending request
# https://github.com/turbot/steampipe-plugin-sdk/issues/512
@test "check cache functionality(edge case)" {
  skip "TODO - verify behavior and re-enable - https://github.com/turbot/steampipe-plugin-sdk/issues/710"
  steampipe query "select unique_col, a, b, c, d, delay from chaos_cache_with_delay_quals where delay=10" --output json &> output1.json &
  sleep 1
  steampipe query "select unique_col, a, b from chaos_cache_with_delay_quals where delay=10" --output json &> output2.json

  # store the unique number from 1st query in `content`
  content=$(cat output1.json | jq '.rows[0].unique_col')
  # store the unique number from 2nd query in `new_content`
  content2=$(cat output2.json | jq '.rows[0].unique_col')

  echo $content
  echo $content2

  # verify that `content` and `new_content` are the same
  assert_equal "$content2" "$content"
}

# Test to validate the issue where completed cache requests with quals are incorrectly being identified as satisfying pending cache transfers
# https://github.com/turbot/steampipe-plugin-sdk/issues/517
@test "check cache functionality 2(edge case)" {
  skip "TODO - verify behavior and re-enable - https://github.com/turbot/steampipe-plugin-sdk/issues/710"
  # skip
  steampipe query "select unique_col, a, b, c, delay from chaos_cache_with_delay_quals where delay=10" --output json &> output1.json &
  sleep 1
  steampipe query "select unique_col, a, b from chaos_cache_with_delay_quals where delay=10" --output json &> output2.json
  steampipe query "select unique_col, a, b, c, d from chaos_cache_with_delay_quals where unique_col > 10 and delay=10" --output json &> output3.json

  # store the unique number from 1st query in `content`
  content=$(cat output1.json | jq '.rows[0].unique_col')
  # store the unique number from 2nd query in `new_content`
  content2=$(cat output2.json | jq '.rows[0].unique_col')
  # store the unique number from 3rd query in `new_content`
  content3=$(cat output3.json | jq '.rows[0].unique_col')

  echo $content
  echo $content2
  echo $content3

  # verify that `content` and `content2` are the same
  assert_equal "$content2" "$content"

  # verify that `content` and `content3` are not the same
  if [[ "$content" == "$content3" ]]; then
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
