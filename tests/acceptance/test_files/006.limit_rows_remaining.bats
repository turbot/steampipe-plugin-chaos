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

###### rows remaining tests ######

# this test checks the SDK rows remaining functionality when a limit is passed to the query
@test "check rows remaining when a limit is passed in query" {
  run steampipe query "select limit_value, rows_streamed, sdk_rows_remaining from chaos_limit_verify_rows_remaining where total_row_count=2 limit 10" --output=json
  echo $output
  # Use jq command-line tool to parse the JSON string
  object_count=$(echo "$output" | jq length)

  for ((i=0; i<object_count; i++)); do
    # Extract each object using jq and print it
    object=$(echo "$output" | jq --raw-output ".[$i]")
    echo "$object"
    echo

    # Extract individual values using jq and assign them to variables
    limit_value=$(echo "$object" | jq --raw-output '.limit_value')
    rows_streamed=$(echo "$object" | jq --raw-output '.rows_streamed')
    sdk_rows_remaining=$(echo "$object" | jq --raw-output '.sdk_rows_remaining')
    sum=$((sdk_rows_remaining + rows_streamed))

    # Print them to help debugging in case of failures
    echo "limit: $limit_value"
    echo "rows streamed: $rows_streamed"
    echo "rows remaining: $sdk_rows_remaining"

    # Check if sdk_rows_remaining+rows_streamed=limit_value. When a limit is passed, sum of rows_streamed and sdk_rows_remaining should always be equal to limit_value.
    assert_equal $sum $limit_value
  done
}

# this test checks the SDK rows remaining functionality when no limit is passed to the query
@test "check rows remaining when no limit is passed in query" {
  run steampipe query "select limit_value, rows_streamed, sdk_rows_remaining, total_row_count from chaos_limit_verify_rows_remaining where total_row_count=10" --output=json
  echo $output
  # Use jq command-line tool to parse the JSON string
  object_count=$(echo "$output" | jq length)

  for ((i=0; i<object_count; i++)); do
    # Extract each object using jq and print it
    object=$(echo "$output" | jq --raw-output ".[$i]")
    echo "$object"
    echo

    # Extract individual values using jq and assign them to variables
    limit_value=$(echo "$object" | jq --raw-output '.limit_value')
    rows_streamed=$(echo "$object" | jq --raw-output '.rows_streamed')
    sdk_rows_remaining=$(echo "$object" | jq --raw-output '.sdk_rows_remaining')
    total_row_count=$(echo "$object" | jq --raw-output '.total_row_count')
    sum=$((sdk_rows_remaining + rows_streamed))

    # Print them to help debugging in case of failures
    echo "limit: $limit_value"
    echo "rows streamed: $rows_streamed"
    echo "rows remaining: $sdk_rows_remaining"
    echo "total rows: $total_row_count"

    # Check if sdk_rows_remaining+rows_streamed=2147483647(maximum value of a signed 32-bit integer is 2147483647). When no limit is passed, sum of rows_streamed and sdk_rows_remaining should always be equal to 2147483647 which is the maximum value of a signed 32-bit integer.
    assert_equal $sum 2147483647
    # Since no limit is passed limit_value should be null.
    assert_equal $limit_value null
  done
}