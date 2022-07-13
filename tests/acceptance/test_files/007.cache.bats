load "$LIB_BATS_ASSERT/load.bash"
load "$LIB_BATS_SUPPORT/load.bash"

setup() {
  steampipe service start > /dev/null
}

teardown() {
  rm -f output?.json
  steampipe service stop
}

# Test the cache functionality. This and the subsequent tests use the unique_col in `chaos_cache_check` table, which 
# gives a random value for every hydrate call.
# Querying the same columns should result in a CACHE HIT.
@test "check cache functionality when querying same columns" {
  steampipe query "select unique_col, a, b, c from chaos_cache_check" --output json &> output1.json
  # store the unique number from 1st query in `content`
  content=$(cat output1.json | jq '.[0].unique_col')

  steampipe query "select unique_col, a, b, c from chaos_cache_check" --output json &> output2.json
  # store the unique number from 2nd query in `new_content`
  new_content=$(cat output2.json | jq '.[0].unique_col')

  echo $content
  echo $new_content

  # verify that `content` and `new_content` are the same
  assert_equal "$new_content" "$content"
}

# Test the cache functionality, when the second query's columns is a subset of the first. When the
# second query's columns are a subset of the columns queried in the first query, it should be a
# CACHE HIT.
@test "check cache functionality when the second query's columns is a subset of the first" {
  steampipe query "select unique_col, a, b, c from chaos_cache_check" --output json &> output1.json
  # store the unique number from 1st query in `content`
  content=$(cat output1.json | jq '.[0].unique_col')

  steampipe query "select unique_col, a, b from chaos_cache_check" --output json &> output2.json
  # store the unique number from 2nd query in `new_content`
  new_content=$(cat output2.json | jq '.[0].unique_col')

  echo $content
  echo $new_content

  # verify that `content` and `new_content` are the same
  assert_equal "$new_content" "$content"
}

# Test the cache functionality, when querying the same columns multiple times. This should result in a
# CACHE HIT.
@test "check cache functionality multiple queries with same columns" {
  run steampipe query "select unique_col, a, b, c from chaos_cache_check" --output json &> output1.json
  # store the unique number from 1st query in `content`
  content=$(cat output1.json | jq '.[0].unique_col')

  run steampipe query "select unique_col, a, b, c from chaos_cache_check" --output json &> output2.json
  # store the unique number from 2nd query in `new_content`
  content2=$(cat output2.json | jq '.[0].unique_col')

  run steampipe query "select unique_col, a, b, c from chaos_cache_check" --output json &> output3.json
  # store the unique number from 3rd query in `new_content`
  content3=$(cat output3.json | jq '.[0].unique_col')

  run steampipe query "select unique_col, a, b, c from chaos_cache_check" --output json &> output4.json
  # store the unique number from 4th query in `new_content`
  content4=$(cat output4.json | jq '.[0].unique_col')

  echo $content
  echo $content2
  echo $content3
  echo $content4

  # verify that `content` from 1st query is same as the next queries
  assert_equal "$content2" "$content"
  assert_equal "$content3" "$content"
  assert_equal "$content4" "$content"
}

# Test the cache functionality, when the subsequent query's columns are a subset of the first. When the
# subsequent query's columns are a subset of the columns queried in the first query, it should be a
# CACHE HIT.
@test "check cache functionality when multiple query's columns are a subset of the first" {
  run steampipe query "select unique_col, a, b, c from chaos_cache_check" --output json &> output1.json
  # store the unique number from 1st query in `content`
  content=$(cat output1.json | jq '.[0].unique_col')

  run steampipe query "select unique_col, a, b from chaos_cache_check" --output json &> output2.json
  # store the unique number from 2nd query in `new_content`
  content2=$(cat output2.json | jq '.[0].unique_col')

  run steampipe query "select unique_col, a, b from chaos_cache_check" --output json &> output3.json
  # store the unique number from 3rd query in `new_content`
  content3=$(cat output3.json | jq '.[0].unique_col')

  run steampipe query "select unique_col, a, b from chaos_cache_check" --output json &> output4.json
  # store the unique number from 4th query in `new_content`
  content4=$(cat output4.json | jq '.[0].unique_col')

  echo $content
  echo $content2
  echo $content3
  echo $content4

  # verify that `content` from 1st query is same as the next queries
  assert_equal "$content2" "$content"
  assert_equal "$content3" "$content"
  assert_equal "$content4" "$content"
}

# Test the cache functionality, when the second query's columns is NOT a subset of the first. When the
# second query is querying more columns i.e. not a subset of the columns queried in the first query, it 
# should be a CACHE MISS.
@test "check cache functionality when the second query has more columns than the first" {
  steampipe query "select unique_col, a, b from chaos_cache_check" --output json &> output1.json
  # store the unique number from 1st query in `content`
  content=$(cat output1.json | jq '.[0].unique_col')

  steampipe query "select unique_col, a, b, c from chaos_cache_check" --output json &> output2.json
  # store the unique number from 2nd query in `new_content`
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
}

# Test the cache functionality, when the second query queries the exact same columns and has the same 
# limit as the first query, it should be a CACHE HIT.
@test "check cache functionality when the both the queries have same limits" {
  steampipe query "select unique_col, a, b, c from chaos_cache_check limit 3" --output json &> output1.json
  # store the unique number from 1st query in `content`
  content=$(cat output1.json | jq '.[0].unique_col')

  steampipe query "select unique_col, a, b, c from chaos_cache_check limit 3" --output json &> output2.json
  # store the unique number from 2nd query in `new_content`
  new_content=$(cat output2.json | jq '.[0].unique_col')

  echo $content
  echo $new_content

  # verify that `content` and `new_content` are the same
  assert_equal "$new_content" "$content"
}

# Test the cache functionality, when the second query queries the exact same columns and has a limit,
# but the first query has no limit, it should be a CACHE HIT.
@test "check cache functionality when first query has no limit but second query has a limit" {
  steampipe query "select unique_col, a, b, c from chaos_cache_check" --output json &> output1.json
  # store the unique number from 1st query in `content`
  content=$(cat output1.json | jq '.[0].unique_col')

  steampipe query "select unique_col, a, b, c from chaos_cache_check limit 3" --output json &> output2.json
  # store the unique number from 2nd query in `new_content`
  new_content=$(cat output2.json | jq '.[0].unique_col')

  echo $content
  echo $new_content

  # verify that `content` and `new_content` are the same
  assert_equal "$new_content" "$content"
}

# Test the cache functionality, when the second query queries the exact same columns and has a limit less
# than the limit in but the first query, it should be a CACHE HIT.
@test "check cache functionality when second query has lower limit than first" {
  steampipe query "select unique_col, a, b, c from chaos_cache_check limit 4" --output json &> output1.json
  # store the unique number from 1st query in `content`
  content=$(cat output1.json | jq '.[0].unique_col')

  steampipe query "select unique_col, a, b, c from chaos_cache_check limit 3" --output json &> output2.json
  # store the unique number from 2nd query in `new_content`
  new_content=$(cat output2.json | jq '.[0].unique_col')

  echo $content
  echo $new_content

  # verify that `content` and `new_content` are the same
  assert_equal "$new_content" "$content"
}

# Test the cache functionality, when the second query queries the exact same columns and has a limit more
# than the limit in but the first query, it should be a CACHE MISS.
@test "check cache functionality when second query has higher limit than first" {
  steampipe query "select unique_col, a, b, c from chaos_cache_check limit 3" --output json &> output1.json
  # store the unique number from 1st query in `content`
  content=$(cat output1.json | jq '.[0].unique_col')

  steampipe query "select unique_col, a, b, c from chaos_cache_check limit 4" --output json &> output2.json
  # store the unique number from 2nd query in `new_content`
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
}

######## BACKGROUND TESTS ###########

# Test the cache functionality when the first query is still running in the background. This and the subsequent
# tests use the unique_col in `chaos_cache_check` table, which gives a random value for every hydrate call, and 
# also queries the delay column from the table which causes a 10s delay during the hydrate call, so that the 
# first query still runs in the background when the second query is fired.
# Querying the same columns should result in a CACHE HIT.
@test "check cache functionality when querying same columns(first query in background)" {
  steampipe query "select unique_col, a, b, c, delay from chaos_cache_check" --output json &> output1.json &
  sleep 1
  steampipe query "select unique_col, a, b, c, delay from chaos_cache_check" --output json &> output2.json

  # store the unique number from 1st query in `content`
  content=$(cat output1.json | jq '.[0].unique_col')
  # store the unique number from 2nd query in `new_content`
  new_content=$(cat output2.json | jq '.[0].unique_col')

  echo $content
  echo $new_content

  # verify that `content` and `new_content` are the same
  assert_equal "$new_content" "$content"
}

# Test the cache functionality, when the second query's columns is a subset of the first query(running in the 
# background). When the second query's columns are a subset of the columns queried in the first query, it should be a
# CACHE HIT.
@test "check cache functionality when the second query's columns is a subset of the first(first query in background)" {
  steampipe query "select unique_col, a, b, c, delay from chaos_cache_check" --output json &> output1.json &
  sleep 1
  steampipe query "select unique_col, a, b, delay from chaos_cache_check" --output json &> output2.json

  # store the unique number from 1st query in `content`
  content=$(cat output1.json | jq '.[0].unique_col')
  # store the unique number from 2nd query in `new_content`
  new_content=$(cat output2.json | jq '.[0].unique_col')

  echo $content
  echo $new_content

  # verify that `content` and `new_content` are the same
  assert_equal "$new_content" "$content"
}

# Test the cache functionality, when querying the same columns multiple times while the first query is running in
# the background. This should result in a CACHE HIT.
@test "check cache functionality multiple queries with same columns(first query in background)" {
  steampipe query "select unique_col, a, b, c, delay from chaos_cache_check" --output json &> output1.json &
  sleep 1
  steampipe query "select unique_col, a, b, c, delay from chaos_cache_check" --output json &> output2.json
  steampipe query "select unique_col, a, b, c, delay from chaos_cache_check" --output json &> output3.json
  steampipe query "select unique_col, a, b, c, delay from chaos_cache_check" --output json &> output4.json

  # store the unique number from 1st query in `content`
  content=$(cat output1.json | jq '.[0].unique_col')
  # store the unique number from 2nd query in `new_content`
  content2=$(cat output2.json | jq '.[0].unique_col')
  # store the unique number from 3rd query in `new_content`
  content3=$(cat output3.json | jq '.[0].unique_col')
  # store the unique number from 4th query in `new_content`
  content4=$(cat output4.json | jq '.[0].unique_col')

  echo $content
  echo $content2
  echo $content3
  echo $content4

  # verify that `content` from 1st query is same as the next queries
  assert_equal "$content2" "$content"
  assert_equal "$content3" "$content"
  assert_equal "$content4" "$content"
}

# Test the cache functionality, when the subsequent query's columns are a subset of the first query(running in the 
# background). When the subsequent query's columns are a subset of the columns queried in the first query, it should be a
# CACHE HIT.
@test "check cache functionality when multiple query's columns are a subset of the first(first query in background)" {
  steampipe query "select unique_col, a, b, c, delay from chaos_cache_check" --output json &> output1.json &
  sleep 1
  steampipe query "select unique_col, a, b from chaos_cache_check" --output json &> output2.json
  steampipe query "select unique_col, a, b from chaos_cache_check" --output json &> output3.json
  steampipe query "select unique_col, a, b from chaos_cache_check" --output json &> output4.json

  # store the unique number from 1st query in `content`
  content=$(cat output1.json | jq '.[0].unique_col')
  # store the unique number from 2nd query in `new_content`
  content2=$(cat output2.json | jq '.[0].unique_col')
  # store the unique number from 3rd query in `new_content`
  content3=$(cat output3.json | jq '.[0].unique_col')
  # store the unique number from 4th query in `new_content`
  content4=$(cat output4.json | jq '.[0].unique_col')

  echo $content
  echo $content2
  echo $content3
  echo $content4

  # verify that `content` from 1st query is same as the next queries
  assert_equal "$content2" "$content"
  assert_equal "$content3" "$content"
  assert_equal "$content4" "$content"
}

# Test the cache functionality, when the second query's columns is NOT a subset of the first query(running in the 
# background). When the second query is querying more columns i.e. not a subset of the columns queried in the first query, 
# it should be a CACHE MISS.
@test "check cache functionality when the second query has more columns than the first(first query in background)" {
  steampipe query "select unique_col, a, b, delay from chaos_cache_check" --output json &> output1.json &
  sleep 1
  steampipe query "select unique_col, a, b, delay, c from chaos_cache_check" --output json &> output2.json

  # store the unique number from 1st query in `content`
  content=$(cat output1.json | jq '.[0].unique_col')
  # store the unique number from 2nd query in `new_content`
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
}

# Test the cache functionality, when the second query queries the exact same columns and has the same 
# limit as the first query(running in the background), it should be a CACHE HIT.
@test "check cache functionality when the both the queries have same limits(first query in background)" {
  steampipe query "select unique_col, a, b, c, delay from chaos_cache_check limit 3" --output json &> output1.json &
  sleep 1
  steampipe query "select unique_col, a, b, c, delay from chaos_cache_check limit 3" --output json &> output2.json

  # store the unique number from 1st query in `content`
  content=$(cat output1.json | jq '.[0].unique_col')
  # store the unique number from 2nd query in `new_content`
  new_content=$(cat output2.json | jq '.[0].unique_col')

  echo $content
  echo $new_content

  # verify that `content` and `new_content` are the same
  assert_equal "$new_content" "$content"
}

# Test the cache functionality, when the second query queries the exact same columns and has a limit,
# but the first query(running in the background) has no limit, it should be a CACHE HIT.
@test "check cache functionality when first query has no limit but second query has a limit(first query in background)" {
  steampipe query "select unique_col, a, b, c, delay from chaos_cache_check" --output json &> output1.json &
  sleep 1
  steampipe query "select unique_col, a, b, c, delay from chaos_cache_check limit 3" --output json &> output2.json

  # store the unique number from 1st query in `content`
  content=$(cat output1.json | jq '.[0].unique_col')
  # store the unique number from 2nd query in `new_content`
  new_content=$(cat output2.json | jq '.[0].unique_col')

  echo $content
  echo $new_content

  # verify that `content` and `new_content` are the same
  assert_equal "$new_content" "$content"
}

# Test the cache functionality, when the second query queries the exact same columns but has no limit, but the first
# query(running in the background) has a limit, it should be a CACHE MISS.
@test "check cache functionality when second query has no limit but first query has a limit(first query in background)" {
  steampipe query "select unique_col, a, b, c, delay from chaos_cache_check limit 3" --output json &> output1.json &
  sleep 1
  steampipe query "select unique_col, a, b, c, delay from chaos_cache_check" --output json &> output2.json

  # store the unique number from 1st query in `content`
  content=$(cat output1.json | jq '.[0].unique_col')
  # store the unique number from 2nd query in `new_content`
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
}

# Test the cache functionality, when the second query queries the exact same columns and has a limit less
# than the limit in but the first query(running in the background), it should be a CACHE HIT.
@test "check cache functionality when second query has lower limit than first(first query in background)" {
  steampipe query "select unique_col, a, b, c, delay from chaos_cache_check limit 4" --output json &> output1.json &
  sleep 1
  steampipe query "select unique_col, a, b, c, delay from chaos_cache_check limit 3" --output json &> output2.json

  # store the unique number from 1st query in `content`
  content=$(cat output1.json | jq '.[0].unique_col')
  # store the unique number from 2nd query in `new_content`
  new_content=$(cat output2.json | jq '.[0].unique_col')

  echo $content
  echo $new_content

  # verify that `content` and `new_content` are the same
  assert_equal "$new_content" "$content"
}

# Test the cache functionality, when the second query queries the exact same columns and has a limit more
# than the limit in but the first query(running in the background), it should be a CACHE MISS.
@test "check cache functionality when second query has higher limit than first(first query in background)" {
  steampipe query "select unique_col, a, b, c, delay from chaos_cache_check limit 3" --output json &> output1.json &
  sleep 1
  steampipe query "select unique_col, a, b, c, delay from chaos_cache_check limit 4" --output json &> output2.json

  # store the unique number from 1st query in `content`
  content=$(cat output1.json | jq '.[0].unique_col')
  # store the unique number from 2nd query in `new_content`
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
}

######## ERROR AND TIMEOUT TESTS ###########

# Test the cache functionality when the first query runs in the background, but returns an ERROR. This test uses
# the unique_col in `chaos_cache_check` table, which gives a random value for every hydrate call, and also 
# queries the error_after_delay column from the table which causes a 10s delay during  the hydrate call and 
# returns an ERROR.
# If the first query returns error, it should be a CACHE MISS.
@test "check cache functionality when first query returns error, other queries should not cache(first query in background)" {
  steampipe query "select unique_col, a, b, c, error_after_delay from chaos_cache_check" --output json &> output1.json &
  sleep 1
  steampipe query "select unique_col, a, b from chaos_cache_check" --output json &> output2.json
  steampipe query "select unique_col, a, b, c from chaos_cache_check" --output json &> output3.json

  # store the unique number from 2nd query in `new_content`
  content2=$(cat output2.json | jq '.[0].unique_col')
  # store the unique number from 3rd query in `new_content`
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
}

# Test the cache functionality when the first query runs in the background, but TIMES OUT. This test uses
# the unique_col in `chaos_cache_check` table, which gives a random value for every hydrate call, and also 
# queries the long_delay column from the table which causes a 10hrs delay during the hydrate call and the
# query times out due to the env variable `STEAMPIPE_CACHE_PENDING_QUERY_TIMEOUT` set to 10.
# If the first query running in the background TIMES OUT, but the other queries query the exact same columns
# it should be a CACHE HIT.
@test "check cache functionality when first query times out, other queries should cache(first query in background)" {
  export STEAMPIPE_CACHE_PENDING_QUERY_TIMEOUT=10

  steampipe query "select unique_col, a, b, c, long_delay from chaos_cache_check" --output json &> output1.json &
  sleep 1
  steampipe query "select unique_col, a, b from chaos_cache_check" --output json &> output2.json
  steampipe query "select unique_col, a, b from chaos_cache_check" --output json &> output3.json

  # store the unique number from 2nd query in `new_content`
  content2=$(cat output2.json | jq '.[0].unique_col')
  # store the unique number from 3rd query in `new_content`
  content3=$(cat output3.json | jq '.[0].unique_col')

  echo $content2
  echo $content3

  # verify that `content2` and `content3` are the same
  assert_equal "$content2" "$content3"
}

# If the first query running in the background TIMES OUT, and the other queries DO NOT query the exact same 
# columns it should be a CACHE MISS.
@test "check cache functionality when first query times out, other queries should not cache(first query in background)" {
  export STEAMPIPE_CACHE_PENDING_QUERY_TIMEOUT=10 

  steampipe query "select unique_col, a, b, c, long_delay from chaos_cache_check" --output json &> output1.json &
  sleep 1
  steampipe query "select unique_col, a, b from chaos_cache_check" --output json &> output2.json
  steampipe query "select unique_col, a, b, c from chaos_cache_check" --output json &> output3.json

  # store the unique number from 2nd query in `new_content`
  content2=$(cat output2.json | jq '.[0].unique_col')
  # store the unique number from 3rd query in `new_content`
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
}

@test "stop service" {
  steampipe service stop --force
}
