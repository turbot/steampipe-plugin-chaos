#!/bin/bash -e

ps -ef | grep "steampipe"
# export STEAMPIPE_CACHE_PENDING_QUERY_TIMEOUT=20
# export STEAMPIPE_LOG=trace
echo "starting test"
steampipe plugin install chaos
steampipe service start
steampipe query "select unique_col, a, b, c, error_after_delay from chaos_cache_check" --output json &> output4.json &
sleep .5s
steampipe query "select unique_col, a, b from chaos_cache_check" --output json &> output5.json &
sleep .5s
steampipe query "select unique_col, a, b from chaos_cache_check" --output json &> output6.json
steampipe service stop --force
echo "fin"