load "$LIB_BATS_ASSERT/load.bash"
load "$LIB_BATS_SUPPORT/load.bash"

######## ERROR AND TIMEOUT TESTS ###########

@test "check cache functionality when first query returns error, other queries should not cache(first 2 queries in background)" {
  export STEAMPIPE_LOG=trace
  run steampipe plugin install chaos
  run steampipe service start

  steampipe query "select unique_col, a, b, c, error_after_delay from chaos_cache_check" --output json &> output1.json &
  sleep .5s
  steampipe query "select unique_col, a, b from chaos_cache_check" --output json &> output2.json &
  sleep .5s
  steampipe query "select unique_col, a, b from chaos_cache_check" --output json &> output3.json

  # store the time from 2nd query in `content2`
  content2=$(cat output2.json | jq '.[0].unique_col')
  # store the time from 3rd query in `content3`
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
