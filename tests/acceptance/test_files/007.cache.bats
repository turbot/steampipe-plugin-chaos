load "$LIB_BATS_ASSERT/load.bash"
load "$LIB_BATS_SUPPORT/load.bash"

@test "check cache functionality when querying same columns" {
  run steampipe service stop --force
  run steampipe plugin install chaos
  run steampipe service start

  steampipe query "select unique_col, a, b, c from chaos_cache_check" --output json &> output1.json
  # store the time from 1st query in `content`
  content=$(cat output1.json | jq '.[0].unique_col')

  steampipe query "select unique_col, a, b, c from chaos_cache_check" --output json &> output2.json
  # store the time from 2nd query in `new_content`
  new_content=$(cat output2.json | jq '.[0].unique_col')

  echo $content
  echo $new_content

  # verify that `content` and `new_content` are the same
  assert_equal "$new_content" "$content"

  rm -f output?.json
  run steampipe service stop
}

@test "check cache functionality when the second query's columns is a subset of the first" {
  run steampipe plugin install chaos
  run steampipe service start

  steampipe query "select unique_col, a, b, c from chaos_cache_check" --output json &> output1.json
  # store the time from 1st query in `content`
  content=$(cat output1.json | jq '.[0].unique_col')

  steampipe query "select unique_col, a, b from chaos_cache_check" --output json &> output2.json
  # store the time from 2nd query in `new_content`
  new_content=$(cat output2.json | jq '.[0].unique_col')

  echo $content
  echo $new_content

  # verify that `content` and `new_content` are the same
  assert_equal "$new_content" "$content"

  rm -f output?.json
  run steampipe service stop
}

@test "check cache functionality multiple queries with same columns" {
  run steampipe plugin install chaos
  run steampipe service start

  run steampipe query "select unique_col, a, b, c from chaos_cache_check" --output json &> output1.json
  # store the time from 1st query in `content`
  content=$(cat output1.json | jq '.[0].unique_col')

  run steampipe query "select unique_col, a, b, c from chaos_cache_check" --output json &> output2.json
  # store the time from 2nd query in `new_content`
  content2=$(cat output2.json | jq '.[0].unique_col')

  run steampipe query "select unique_col, a, b, c from chaos_cache_check" --output json &> output3.json
  # store the time from 3rd query in `new_content`
  content3=$(cat output3.json | jq '.[0].unique_col')

  run steampipe query "select unique_col, a, b, c from chaos_cache_check" --output json &> output4.json
  # store the time from 4th query in `new_content`
  content4=$(cat output4.json | jq '.[0].unique_col')

  echo $content
  echo $content2
  echo $content3
  echo $content4

  # verify that `content` from 1st query is same as the next queries
  assert_equal "$content2" "$content"
  assert_equal "$content3" "$content"
  assert_equal "$content4" "$content"

  rm -f output?.json
  run steampipe service stop
}

@test "check cache functionality when multiple query's columns are a subset of the first" {
  run steampipe plugin install chaos
  run steampipe service start

  run steampipe query "select unique_col, a, b, c from chaos_cache_check" --output json &> output1.json
  # store the time from 1st query in `content`
  content=$(cat output1.json | jq '.[0].unique_col')

  run steampipe query "select unique_col, a, b from chaos_cache_check" --output json &> output2.json
  # store the time from 2nd query in `new_content`
  content2=$(cat output2.json | jq '.[0].unique_col')

  run steampipe query "select unique_col, a, b from chaos_cache_check" --output json &> output3.json
  # store the time from 3rd query in `new_content`
  content3=$(cat output3.json | jq '.[0].unique_col')

  run steampipe query "select unique_col, a, b from chaos_cache_check" --output json &> output4.json
  # store the time from 4th query in `new_content`
  content4=$(cat output4.json | jq '.[0].unique_col')

  echo $content
  echo $content2
  echo $content3
  echo $content4

  # verify that `content` from 1st query is same as the next queries
  assert_equal "$content2" "$content"
  assert_equal "$content3" "$content"
  assert_equal "$content4" "$content"

  rm -f output?.json
  run steampipe service stop
}

@test "check cache functionality when the second query has more columns than the first" {
  run steampipe plugin install chaos
  run steampipe service start

  steampipe query "select unique_col, a, b from chaos_cache_check" --output json &> output1.json
  # store the time from 1st query in `content`
  content=$(cat output1.json | jq '.[0].unique_col')

  steampipe query "select unique_col, a, b, c from chaos_cache_check" --output json &> output2.json
  # store the time from 2nd query in `new_content`
  new_content=$(cat output2.json | jq '.[0].unique_col')

  echo $content
  echo $new_content

  # verify that `content` and `new_content` are not the same
  if [[ "$content" == "$new_content" ]]; then
    flag=1
  else
    flag=0
  fi
  assert_equal "$flag" "0"

  rm -f output?.json
  run steampipe service stop
}

@test "check cache functionality when the both the queries have same limits" {
  run steampipe plugin install chaos
  run steampipe service start

  steampipe query "select unique_col, a, b, c from chaos_cache_check limit 3" --output json &> output1.json
  # store the time from 1st query in `content`
  content=$(cat output1.json | jq '.[0].unique_col')

  steampipe query "select unique_col, a, b, c from chaos_cache_check limit 3" --output json &> output2.json
  # store the time from 2nd query in `new_content`
  new_content=$(cat output2.json | jq '.[0].unique_col')

  echo $content
  echo $new_content

  # verify that `content` and `new_content` are the same
  assert_equal "$new_content" "$content"

  rm -f output?.json
  run steampipe service stop
}

@test "check cache functionality when first query has no limit but second query has a limit" {
  run steampipe plugin install chaos
  run steampipe service start

  steampipe query "select unique_col, a, b, c from chaos_cache_check" --output json &> output1.json
  # store the time from 1st query in `content`
  content=$(cat output1.json | jq '.[0].unique_col')

  steampipe query "select unique_col, a, b, c from chaos_cache_check limit 3" --output json &> output2.json
  # store the time from 2nd query in `new_content`
  new_content=$(cat output2.json | jq '.[0].unique_col')

  echo $content
  echo $new_content

  # verify that `content` and `new_content` are the same
  assert_equal "$new_content" "$content"

  rm -f output?.json
  run steampipe service stop
}

@test "check cache functionality when second query has lower limit than first" {
  run steampipe plugin install chaos
  run steampipe service start

  steampipe query "select unique_col, a, b, c from chaos_cache_check limit 4" --output json &> output1.json
  # store the time from 1st query in `content`
  content=$(cat output1.json | jq '.[0].unique_col')

  steampipe query "select unique_col, a, b, c from chaos_cache_check limit 3" --output json &> output2.json
  # store the time from 2nd query in `new_content`
  new_content=$(cat output2.json | jq '.[0].unique_col')

  echo $content
  echo $new_content

  # verify that `content` and `new_content` are the same
  assert_equal "$new_content" "$content"

  rm -f output?.json
  run steampipe service stop
}

@test "check cache functionality when second query has higher limit than first" {
  run steampipe plugin install chaos
  run steampipe service start

  steampipe query "select unique_col, a, b, c from chaos_cache_check limit 3" --output json &> output1.json
  # store the time from 1st query in `content`
  content=$(cat output1.json | jq '.[0].unique_col')

  steampipe query "select unique_col, a, b, c from chaos_cache_check limit 4" --output json &> output2.json
  # store the time from 2nd query in `new_content`
  new_content=$(cat output2.json | jq '.[0].unique_col')

  echo $content
  echo $new_content

  # verify that `content` and `new_content` are not the same
  if [[ "$content" == "$new_content" ]]; then
    flag=1
  else
    flag=0
  fi
  assert_equal "$flag" "0"

  rm -f output?.json
  run steampipe service stop
}

######## BACKGROUND TESTS ###########

@test "check cache functionality when querying same columns(first query in background)" {
  run steampipe plugin install chaos
  run steampipe service start

  steampipe query "select unique_col, a, b, c, delay from chaos_cache_check" --output json &> output1.json &
  sleep .5s
  steampipe query "select unique_col, a, b, c, delay from chaos_cache_check" --output json &> output2.json

  # store the time from 1st query in `content`
  content=$(cat output1.json | jq '.[0].unique_col')
  # store the time from 2nd query in `new_content`
  new_content=$(cat output2.json | jq '.[0].unique_col')

  echo $content
  echo $new_content

  # verify that `content` and `new_content` are the same
  assert_equal "$new_content" "$content"

  rm -f output?.json
  run steampipe service stop
}

@test "check cache functionality when the second query's columns is a subset of the first(first query in background)" {
  run steampipe plugin install chaos
  run steampipe service start

  steampipe query "select unique_col, a, b, c, delay from chaos_cache_check" --output json &> output1.json &
  sleep .5s
  steampipe query "select unique_col, a, b, delay from chaos_cache_check" --output json &> output2.json

  # store the time from 1st query in `content`
  content=$(cat output1.json | jq '.[0].unique_col')
  # store the time from 2nd query in `new_content`
  new_content=$(cat output2.json | jq '.[0].unique_col')

  echo $content
  echo $new_content

  # verify that `content` and `new_content` are the same
  assert_equal "$new_content" "$content"

  rm -f output?.json
  run steampipe service stop
}

@test "check cache functionality multiple queries with same columns(first query in background)" {
  run steampipe plugin install chaos
  run steampipe service start

  steampipe query "select unique_col, a, b, c, delay from chaos_cache_check" --output json &> output1.json &
  sleep .5s
  steampipe query "select unique_col, a, b, c, delay from chaos_cache_check" --output json &> output2.json
  steampipe query "select unique_col, a, b, c, delay from chaos_cache_check" --output json &> output3.json
  steampipe query "select unique_col, a, b, c, delay from chaos_cache_check" --output json &> output4.json

  # store the time from 1st query in `content`
  content=$(cat output1.json | jq '.[0].unique_col')
  # store the time from 2nd query in `new_content`
  content2=$(cat output2.json | jq '.[0].unique_col')
  # store the time from 3rd query in `new_content`
  content3=$(cat output3.json | jq '.[0].unique_col')
  # store the time from 4th query in `new_content`
  content4=$(cat output4.json | jq '.[0].unique_col')

  echo $content
  echo $content2
  echo $content3
  echo $content4

  # verify that `content` from 1st query is same as the next queries
  assert_equal "$content2" "$content"
  assert_equal "$content3" "$content"
  assert_equal "$content4" "$content"

  rm -f output?.json
  run steampipe service stop
}

@test "check cache functionality when multiple query's columns are a subset of the first(first query in background)" {
  run steampipe plugin install chaos
  run steampipe service start

  steampipe query "select unique_col, a, b, c, delay from chaos_cache_check" --output json &> output1.json &
  sleep .5s
  steampipe query "select unique_col, a, b from chaos_cache_check" --output json &> output2.json
  steampipe query "select unique_col, a, b from chaos_cache_check" --output json &> output3.json
  steampipe query "select unique_col, a, b from chaos_cache_check" --output json &> output4.json

  # store the time from 1st query in `content`
  content=$(cat output1.json | jq '.[0].unique_col')
  # store the time from 2nd query in `new_content`
  content2=$(cat output2.json | jq '.[0].unique_col')
  # store the time from 3rd query in `new_content`
  content3=$(cat output3.json | jq '.[0].unique_col')
  # store the time from 4th query in `new_content`
  content4=$(cat output4.json | jq '.[0].unique_col')

  echo $content
  echo $content2
  echo $content3
  echo $content4

  # verify that `content` from 1st query is same as the next queries
  assert_equal "$content2" "$content"
  assert_equal "$content3" "$content"
  assert_equal "$content4" "$content"

  rm -f output?.json
  run steampipe service stop
}

@test "check cache functionality when the second query has more columns than the first(first query in background)" {
  run steampipe plugin install chaos
  run steampipe service start

  steampipe query "select unique_col, a, b, delay from chaos_cache_check" --output json &> output1.json &
  sleep .5s
  steampipe query "select unique_col, a, b, delay, c from chaos_cache_check" --output json &> output2.json

  # store the time from 1st query in `content`
  content=$(cat output1.json | jq '.[0].unique_col')
  # store the time from 2nd query in `new_content`
  new_content=$(cat output2.json | jq '.[0].unique_col')

  echo $content
  echo $new_content

  # verify that `content` and `new_content` are not the same
  if [[ "$content" == "$new_content" ]]; then
    flag=1
  else
    flag=0
  fi
  assert_equal "$flag" "0"

  rm -f output?.json
  run steampipe service stop
}

@test "check cache functionality when the both the queries have same limits(first query in background)" {
  run steampipe plugin install chaos
  run steampipe service start

  steampipe query "select unique_col, a, b, c, delay from chaos_cache_check limit 3" --output json &> output1.json &
  sleep .5s
  steampipe query "select unique_col, a, b, c, delay from chaos_cache_check limit 3" --output json &> output2.json

  # store the time from 1st query in `content`
  content=$(cat output1.json | jq '.[0].unique_col')
  # store the time from 2nd query in `new_content`
  new_content=$(cat output2.json | jq '.[0].unique_col')

  echo $content
  echo $new_content

  # verify that `content` and `new_content` are the same
  assert_equal "$new_content" "$content"

  rm -f output?.json
  run steampipe service stop
}

@test "check cache functionality when first query has no limit but second query has a limit(first query in background)" {
  run steampipe plugin install chaos
  run steampipe service start

  steampipe query "select unique_col, a, b, c, delay from chaos_cache_check" --output json &> output1.json &
  sleep .5s
  steampipe query "select unique_col, a, b, c, delay from chaos_cache_check limit 3" --output json &> output2.json

  # store the time from 1st query in `content`
  content=$(cat output1.json | jq '.[0].unique_col')
  # store the time from 2nd query in `new_content`
  new_content=$(cat output2.json | jq '.[0].unique_col')

  echo $content
  echo $new_content

  # verify that `content` and `new_content` are the same
  assert_equal "$new_content" "$content"

  rm -f output?.json
  run steampipe service stop
}

@test "check cache functionality when second query has no limit but first query has a limit(first query in background)" {
  run steampipe plugin install chaos
  run steampipe service start

  steampipe query "select unique_col, a, b, c, delay from chaos_cache_check limit 3" --output json &> output1.json &
  sleep .5s
  steampipe query "select unique_col, a, b, c, delay from chaos_cache_check" --output json &> output2.json

  # store the time from 1st query in `content`
  content=$(cat output1.json | jq '.[0].unique_col')
  # store the time from 2nd query in `new_content`
  new_content=$(cat output2.json | jq '.[0].unique_col')

  echo $content
  echo $new_content

  # verify that `content` and `new_content` are not the same
  if [[ "$content" == "$new_content" ]]; then
    flag=1
  else
    flag=0
  fi
  assert_equal "$flag" "0"

  rm -f output?.json
  run steampipe service stop
}

@test "check cache functionality when second query has lower limit than first(first query in background)" {
  run steampipe plugin install chaos
  run steampipe service start

  steampipe query "select unique_col, a, b, c, delay from chaos_cache_check limit 4" --output json &> output1.json &
  sleep .5s
  steampipe query "select unique_col, a, b, c, delay from chaos_cache_check limit 3" --output json &> output2.json

  # store the time from 1st query in `content`
  content=$(cat output1.json | jq '.[0].unique_col')
  # store the time from 2nd query in `new_content`
  new_content=$(cat output2.json | jq '.[0].unique_col')

  echo $content
  echo $new_content

  # verify that `content` and `new_content` are the same
  assert_equal "$new_content" "$content"

  rm -f output?.json
  run steampipe service stop
}

@test "check cache functionality when second query has higher limit than first(first query in background)" {
  run steampipe plugin install chaos
  run steampipe service start

  steampipe query "select unique_col, a, b, c, delay from chaos_cache_check limit 3" --output json &> output1.json &
  sleep .5s
  steampipe query "select unique_col, a, b, c, delay from chaos_cache_check limit 4" --output json &> output2.json

  # store the time from 1st query in `content`
  content=$(cat output1.json | jq '.[0].unique_col')
  # store the time from 2nd query in `new_content`
  new_content=$(cat output2.json | jq '.[0].unique_col')

  echo $content
  echo $new_content

  # verify that `content` and `new_content` are not the same
  if [[ "$content" == "$new_content" ]]; then
    flag=1
  else
    flag=0
  fi
  assert_equal "$flag" "0"

  rm -f output?.json
  run steampipe service stop
}

######## ERROR AND TIMEOUT TESTS ###########

@test "check cache functionality when first query returns error, other queries should not cache(first query in background)" {
  run steampipe plugin install chaos
  run steampipe service start

  steampipe query "select unique_col, a, b, c, error_after_delay from chaos_cache_check" --output json &> output1.json &
  sleep .5s
  steampipe query "select unique_col, a, b from chaos_cache_check" --output json &> output2.json
  steampipe query "select unique_col, a, b, c from chaos_cache_check" --output json &> output3.json

  # store the time from 2nd query in `new_content`
  content2=$(cat output2.json | jq '.[0].unique_col')
  # store the time from 3rd query in `new_content`
  content3=$(cat output3.json | jq '.[0].unique_col')

  echo $content2
  echo $content3

  # verify that `content2` and `content3` are not the same
  if [[ "$content2" == "$content3" ]]; then
    flag=1
  else
    flag=0
  fi
  assert_equal "$flag" "0"

  rm -f output?.json
  run steampipe service stop --force
}

@test "check cache functionality when first query times out, other queries should cache(first query in background)" {
  run steampipe plugin install chaos
  run steampipe service start
  export STEAMPIPE_CACHE_PENDING_QUERY_TIMEOUT=10

  steampipe query "select unique_col, a, b, c, long_delay from chaos_cache_check" --output json &> output1.json &
  sleep .5s
  steampipe query "select unique_col, a, b from chaos_cache_check" --output json &> output2.json
  steampipe query "select unique_col, a, b from chaos_cache_check" --output json &> output3.json

  # store the time from 2nd query in `new_content`
  content2=$(cat output2.json | jq '.[0].unique_col')
  # store the time from 3rd query in `new_content`
  content3=$(cat output3.json | jq '.[0].unique_col')

  echo $content2
  echo $content3

  # verify that `content2` and `content3` are the same
  assert_equal "$content2" "$content3"

  rm -f output?.json
  run steampipe service stop --force
}

@test "check cache functionality when first query times out, other queries should not cache(first query in background)" {
  run steampipe plugin install chaos
  run steampipe service start
  export STEAMPIPE_CACHE_PENDING_QUERY_TIMEOUT=10

  steampipe query "select unique_col, a, b, c, long_delay from chaos_cache_check" --output json &> output1.json &
  sleep .5s
  steampipe query "select unique_col, a, b from chaos_cache_check" --output json &> output2.json
  steampipe query "select unique_col, a, b, c from chaos_cache_check" --output json &> output3.json

  # store the time from 2nd query in `new_content`
  content2=$(cat output2.json | jq '.[0].unique_col')
  # store the time from 3rd query in `new_content`
  content3=$(cat output3.json | jq '.[0].unique_col')

  echo $content2
  echo $content3

  # verify that `content2` and `content3` are not the same
  if [[ "$content2" == "$content3" ]]; then
    flag=1
  else
    flag=0
  fi
  assert_equal "$flag" "0"

  rm -f output?.json
  run steampipe service stop --force
}

@test "check cache functionality when second query has higher limit than first(first query in background)" {
  run steampipe plugin install chaos
  run steampipe service start

  steampipe query "select unique_col, a, b, c, delay from chaos_cache_check limit 3" --output json &> output1.json &
  sleep .5s
  steampipe query "select unique_col, a, b, c, delay from chaos_cache_check limit 4" --output json &> output2.json

  # store the time from 1st query in `content`
  content=$(cat output1.json | jq '.[0].unique_col')
  # store the time from 2nd query in `new_content`
  new_content=$(cat output2.json | jq '.[0].unique_col')

  echo $content
  echo $new_content

  # verify that `content` and `new_content` are not the same
  if [[ "$content" == "$new_content" ]]; then
    flag=1
  else
    flag=0
  fi
  assert_equal "$flag" "0"

  rm -f output?.json
  run steampipe service stop
}
