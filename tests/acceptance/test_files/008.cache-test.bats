load "$LIB_BATS_ASSERT/load.bash"
load "$LIB_BATS_SUPPORT/load.bash"

######## ERROR AND TIMEOUT TESTS ###########

@test "check cache functionality when first query returns error, other queries should not cache(first 2 queries in background)" {
  run steampipe plugin install chaos
  run steampipe service start

  steampipe query "select unique_col, a, b, c, error_after_delay from chaos_cache_check" --output json &> output1.json &
  sleep .5s
  steampipe query "select unique_col, a, b from chaos_cache_check" --output json &> output2.json &
  sleep .5s
  steampipe query "select unique_col, a, b from chaos_cache_check" --output json &> output3.json &
  sleep .5s
  steampipe query "select unique_col, a, b from chaos_cache_check" --output json &> output4.json &
  sleep .5s
  steampipe query "select unique_col, a, b from chaos_cache_check" --output json &> output5.json &
  sleep .5s
  steampipe query "select unique_col, a, b from chaos_cache_check" --output json &> output6.json &
  sleep .5s
  steampipe query "select unique_col, a, b from chaos_cache_check" --output json &> output7.json &
  sleep .5s
  steampipe query "select unique_col, a, b from chaos_cache_check" --output json &> output8.json &
  sleep .5s
  steampipe query "select unique_col, a, b from chaos_cache_check" --output json &> output9.json &
  sleep .5s
  steampipe query "select unique_col, a, b from chaos_cache_check" --output json &> output10.json

  # store the time from 2nd query in `content2`
  content2=$(cat output2.json | jq '.[0].unique_col')
  # store the time from 3rd query in `content3`
  content3=$(cat output3.json | jq '.[0].unique_col')
  content4=$(cat output4.json | jq '.[0].unique_col')
  content5=$(cat output5.json | jq '.[0].unique_col')
  content6=$(cat output6.json | jq '.[0].unique_col')
  content7=$(cat output7.json | jq '.[0].unique_col')
  content8=$(cat output8.json | jq '.[0].unique_col')
  content9=$(cat output9.json | jq '.[0].unique_col')
  content10=$(cat output10.json | jq '.[0].unique_col')

  echo $content2
  echo $content3
  echo $content4
  echo $content5
  echo $content6
  echo $content7
  echo $content8
  echo $content9
  echo $content10

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
