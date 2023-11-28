load "$LIB_BATS_ASSERT/load.bash"
load "$LIB_BATS_SUPPORT/load.bash"

setup() {
  steampipe service start > /dev/null
}

teardown() {
  rm -f output?.json
  steampipe service stop
}

##### INT #####

@test "check cache functionality when the second query quals is a subset of the first(operator1: '<'; operator2: '<'; cache hit)" {
  skip "TODO - verify behavior and re-enable - https://github.com/turbot/steampipe-plugin-sdk/issues/710"
  steampipe query "select int_col, a, b, unique_col from chaos_cache_check where int_col < 3 order by id" --output json &> output1.json
  # store the time from 1st query in `content`
  content=$(cat output1.json | jq '.[0].unique_col')

  steampipe query "select int_col, a, b, unique_col from chaos_cache_check where int_col < 2 order by id" --output json &> output2.json
  # store the time from 2nd query in `new_content`
  new_content=$(cat output2.json | jq '.[0].unique_col')

  echo $content
  echo $new_content

  # verify that `content` and `new_content` are the same
  assert_equal "$new_content" "$content"
}

@test "check cache functionality when the second query quals is not a subset of the first(operator1: '<'; operator2: '<'; cache miss)" {
  skip "TODO - verify behavior and re-enable - https://github.com/turbot/steampipe-plugin-sdk/issues/710"
  steampipe query "select int_col, a, b, unique_col from chaos_cache_check where int_col < 3 order by id" --output json &> output1.json
  # store the time from 1st query in `content`
  content=$(cat output1.json | jq '.[0].unique_col')

  steampipe query "select int_col, a, b, unique_col from chaos_cache_check where int_col < 4 order by id" --output json &> output2.json
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
}

@test "check cache functionality when the second query quals is a subset of the first(operator1: '<'; operator2: '<='; cache hit)" {
  skip "TODO - verify behavior and re-enable - https://github.com/turbot/steampipe-plugin-sdk/issues/710"
  steampipe query "select int_col, a, b, unique_col from chaos_cache_check where int_col < 4 order by id" --output json &> output1.json
  # store the time from 1st query in `content`
  content=$(cat output1.json | jq '.[0].unique_col')

  steampipe query "select int_col, a, b, unique_col from chaos_cache_check where int_col <= 3 order by id" --output json &> output2.json
  # store the time from 2nd query in `new_content`
  new_content=$(cat output2.json | jq '.[0].unique_col')

  echo $content
  echo $new_content

  # verify that `content` and `new_content` are the same
  assert_equal "$new_content" "$content"
}

@test "check cache functionality when the second query quals is not a subset of the first(operator1: '<'; operator2: '<='; cache miss)" {
  skip "TODO - verify behavior and re-enable - https://github.com/turbot/steampipe-plugin-sdk/issues/710"
  steampipe query "select int_col, a, b, unique_col from chaos_cache_check where int_col < 5 order by id" --output json &> output1.json
  # store the time from 1st query in `content`
  content=$(cat output1.json | jq '.[0].unique_col')

  steampipe query "select int_col, a, b, unique_col from chaos_cache_check where int_col <= 7 order by id" --output json &> output2.json
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
}

@test "check cache functionality when the second query quals is a subset of the first(operator1: '>'; operator2: '>'; cache hit)" {
  skip "TODO - verify behavior and re-enable - https://github.com/turbot/steampipe-plugin-sdk/issues/710"
  steampipe query "select int_col, a, b, unique_col from chaos_cache_check where int_col > 6 order by id" --output json &> output1.json
  # store the time from 1st query in `content`
  content=$(cat output1.json | jq '.[0].unique_col')

  steampipe query "select int_col, a, b, unique_col from chaos_cache_check where int_col > 6 order by id" --output json &> output2.json
  # store the time from 2nd query in `new_content`
  new_content=$(cat output2.json | jq '.[0].unique_col')

  echo $content
  echo $new_content

  # verify that `content` and `new_content` are the same
  assert_equal "$new_content" "$content"
}

@test "check cache functionality when the second query quals is not a subset of the first(operator1: '>'; operator2: '>'; cache miss)" {
  skip "TODO - verify behavior and re-enable - https://github.com/turbot/steampipe-plugin-sdk/issues/710"
  steampipe query "select int_col, a, b, unique_col from chaos_cache_check where int_col > 6 order by id" --output json &> output1.json
  # store the time from 1st query in `content`
  content=$(cat output1.json | jq '.[0].unique_col')

  steampipe query "select int_col, a, b, unique_col from chaos_cache_check where int_col > 5 order by id" --output json &> output2.json
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
}

@test "check cache functionality when the second query quals is a subset of the first(operator1: '>'; operator2: '>='; cache hit)" {
  skip "TODO - verify behavior and re-enable - https://github.com/turbot/steampipe-plugin-sdk/issues/710"
  steampipe query "select int_col, a, b, unique_col from chaos_cache_check where int_col > 6 order by id" --output json &> output1.json
  # store the time from 1st query in `content`
  content=$(cat output1.json | jq '.[0].unique_col')

  steampipe query "select int_col, a, b, unique_col from chaos_cache_check where int_col >= 7 order by id" --output json &> output2.json
  # store the time from 2nd query in `new_content`
  new_content=$(cat output2.json | jq '.[0].unique_col')

  echo $content
  echo $new_content

  # verify that `content` and `new_content` are the same
  assert_equal "$new_content" "$content"
}

@test "check cache functionality when the second query quals is not a subset of the first(operator1: '>'; operator2: '>='; cache miss)" {
  skip "TODO - verify behavior and re-enable - https://github.com/turbot/steampipe-plugin-sdk/issues/710"
  steampipe query "select int_col, a, b, unique_col from chaos_cache_check where int_col > 6 order by id" --output json &> output1.json
  # store the time from 1st query in `content`
  content=$(cat output1.json | jq '.[0].unique_col')

  steampipe query "select int_col, a, b, unique_col from chaos_cache_check where int_col >= 6 order by id" --output json &> output2.json
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
}

@test "check cache functionality when the second query quals is a subset of the first(operator1: '<'; operator2: '='; cache hit)" {
  skip "TODO - verify behavior and re-enable - https://github.com/turbot/steampipe-plugin-sdk/issues/710"
  steampipe query "select int_col, a, b, unique_col from chaos_cache_check where int_col < 6 order by id" --output json &> output1.json
  # store the time from 1st query in `content`
  content=$(cat output1.json | jq '.[5].unique_col')

  steampipe query "select int_col, a, b, unique_col from chaos_cache_check where int_col = 5 order by id" --output json &> output2.json
  # store the time from 2nd query in `new_content`
  new_content=$(cat output2.json | jq '.[0].unique_col')

  echo $content
  echo $new_content

  # verify that `content` and `new_content` are the same
  assert_equal "$new_content" "$content"
}

@test "check cache functionality when the second query quals is not a subset of the first(operator1: '<'; operator2: '='; cache miss)" {
  skip "TODO - verify behavior and re-enable - https://github.com/turbot/steampipe-plugin-sdk/issues/710"
  steampipe query "select int_col, a, b, unique_col from chaos_cache_check where int_col < 6 order by id" --output json &> output1.json
  # store the time from 1st query in `content`
  content=$(cat output1.json | jq '.[5].unique_col')

  steampipe query "select int_col, a, b, unique_col from chaos_cache_check where int_col = 6 order by id" --output json &> output2.json
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
}

@test "check cache functionality when the second query quals is a subset of the first(operator1: '<='; operator2: '='; cache hit)" {
  skip "TODO - verify behavior and re-enable - https://github.com/turbot/steampipe-plugin-sdk/issues/710"
  steampipe query "select int_col, a, b, unique_col from chaos_cache_check where int_col <= 6 order by id" --output json &> output1.json
  # store the time from 1st query in `content`
  content=$(cat output1.json | jq '.[5].unique_col')

  steampipe query "select int_col, a, b, unique_col from chaos_cache_check where int_col = 5 order by id" --output json &> output2.json
  # store the time from 2nd query in `new_content`
  new_content=$(cat output2.json | jq '.[0].unique_col')

  echo $content
  echo $new_content

  # verify that `content` and `new_content` are the same
  assert_equal "$new_content" "$content"
}

@test "check cache functionality when the second query quals is not a subset of the first(operator1: '<='; operator2: '='; cache miss)" {
  skip "TODO - verify behavior and re-enable - https://github.com/turbot/steampipe-plugin-sdk/issues/710"
  steampipe query "select int_col, a, b, unique_col from chaos_cache_check where int_col <= 6 order by id" --output json &> output1.json
  # store the time from 1st query in `content`
  content=$(cat output1.json | jq '.[5].unique_col')

  steampipe query "select int_col, a, b, unique_col from chaos_cache_check where int_col = 7 order by id" --output json &> output2.json
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
}

@test "check cache functionality when the second query quals is a subset of the first(operator1: '>'; operator2: '='; cache hit)" {
  skip "TODO - verify behavior and re-enable - https://github.com/turbot/steampipe-plugin-sdk/issues/710"
  steampipe query "select int_col, a, b, unique_col from chaos_cache_check where int_col > 6 order by id" --output json &> output1.json
  # store the time from 1st query in `content`
  content=$(cat output1.json | jq '.[0].unique_col')

  steampipe query "select int_col, a, b, unique_col from chaos_cache_check where int_col = 7 order by id" --output json &> output2.json
  # store the time from 2nd query in `new_content`
  new_content=$(cat output2.json | jq '.[0].unique_col')

  echo $content
  echo $new_content

  # verify that `content` and `new_content` are the same
  assert_equal "$new_content" "$content"
}

@test "check cache functionality when the second query quals is not a subset of the first(operator1: '>'; operator2: '='; cache miss)" {
  skip "TODO - verify behavior and re-enable - https://github.com/turbot/steampipe-plugin-sdk/issues/710"
  steampipe query "select int_col, a, b, unique_col from chaos_cache_check where int_col > 6 order by id" --output json &> output1.json
  # store the time from 1st query in `content`
  content=$(cat output1.json | jq '.[0].unique_col')

  steampipe query "select int_col, a, b, unique_col from chaos_cache_check where int_col = 6 order by id" --output json &> output2.json
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
}

@test "check cache functionality when the second query quals is a subset of the first(operator1: '>='; operator2: '='; cache hit)" {
  skip "TODO - verify behavior and re-enable - https://github.com/turbot/steampipe-plugin-sdk/issues/710"
  steampipe query "select int_col, a, b, unique_col from chaos_cache_check where int_col >= 6 order by id" --output json &> output1.json
  # store the time from 1st query in `content`
  content=$(cat output1.json | jq '.[0].unique_col')

  steampipe query "select int_col, a, b, unique_col from chaos_cache_check where int_col = 6 order by id" --output json &> output2.json
  # store the time from 2nd query in `new_content`
  new_content=$(cat output2.json | jq '.[0].unique_col')

  echo $content
  echo $new_content

  # verify that `content` and `new_content` are the same
  assert_equal "$new_content" "$content"
}

@test "check cache functionality when the second query quals is not a subset of the first(operator1: '>='; operator2: '='; cache miss)" {
  skip "TODO - verify behavior and re-enable - https://github.com/turbot/steampipe-plugin-sdk/issues/710"
  steampipe query "select int_col, a, b, unique_col from chaos_cache_check where int_col >= 6 order by id" --output json &> output1.json
  # store the time from 1st query in `content`
  content=$(cat output1.json | jq '.[0].unique_col')

  steampipe query "select int_col, a, b, unique_col from chaos_cache_check where int_col = 5 order by id" --output json &> output2.json
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
}

##### TIME #####

@test "check cache functionality when the second query quals is a subset of the first(operator1: '>='; operator2: '='; cache hit)[DISABLED]" {
  skip
  steampipe query "select time_col, a, b, unique_col from chaos_cache_check where time_col >= '2021-05-05 00:00:00 +0000 UTC' order by id" --output json &> output1.json
  # store the time from 1st query in `content`
  content=$(cat output1.json | jq '.[0].unique_col')

  steampipe query "select time_col, a, b, unique_col from chaos_cache_check where time_col = '2021-05-05 00:00:00 +0000 UTC' order by id" --output json &> output2.json
  # store the time from 2nd query in `new_content`
  new_content=$(cat output2.json | jq '.[0].unique_col')

  echo $content
  echo $new_content

  # verify that `content` and `new_content` are the same
  assert_equal "$new_content" "$content"
}

# @test "check cache functionality when the second query quals is not a subset of the first(operator1: '>='; operator2: '='; cache miss)" {
#   steampipe query "select time_col, a, b, unique_col from chaos_cache_check where time_col >= '2021-05-05 00:00:00 +0000 UTC' order by id" --output json &> output1.json
#   # store the time from 1st query in `content`
#   content=$(cat output1.json | jq '.[0].unique_col')

#   steampipe query "select time_col, a, b, unique_col from chaos_cache_check where time_col = '2021-04-04 00:00:00 +0000 UTC' order by id" --output json &> output2.json
#   # store the time from 2nd query in `new_content`
#   new_content=$(cat output2.json | jq '.[0].unique_col')

#   echo $content
#   echo $new_content

#   # verify that `content` and `new_content` are not the same
#   if [[ "$content" == "$new_content" ]]; then
#     flag=1
#   else
#     flag=0
#   fi
#   assert_equal "$flag" "0"

#   rm -f output?.json

# }

##### FLOAT #####

# @test "check cache functionality when the second query quals is a subset of the first(operator1: '<'; operator2: '<'; cache hit)" {
#   steampipe query "select float_col, a, b, unique_col from chaos_cache_check where float_col < 0.3 order by id" --output json &> output1.json
#   # store the time from 1st query in `content`
#   content=$(cat output1.json | jq '.[0].unique_col')

#   steampipe query "select float_col, a, b, unique_col from chaos_cache_check where float_col < 0.5 order by id" --output json &> output2.json
#   # store the time from 2nd query in `new_content`
#   new_content=$(cat output2.json | jq '.[0].unique_col')

#   echo $content
#   echo $new_content

#   # verify that `content` and `new_content` are the same
#   assert_equal "$new_content" "$content"

#   rm -f output?.json
# }

# @test "check cache functionality when the second query quals is not a subset of the first(operator1: '<'; operator2: '<'; cache miss)" {
#   steampipe query "select float_col, a, b, unique_col from chaos_cache_check where float_col < 0.5 order by id" --output json &> output1.json
#   # store the time from 1st query in `content`
#   content=$(cat output1.json | jq '.[0].unique_col')

#   steampipe query "select float_col, a, b, unique_col from chaos_cache_check where float_col < 0.7 order by id" --output json &> output2.json
#   # store the time from 2nd query in `new_content`
#   new_content=$(cat output2.json | jq '.[0].unique_col')

#   echo $content
#   echo $new_content

#   # verify that `content` and `new_content` are not the same
#   if [[ "$content" == "$new_content" ]]; then
#     flag=1
#   else
#     flag=0
#   fi
#   assert_equal "$flag" "0"

#   rm -f output?.json
# }