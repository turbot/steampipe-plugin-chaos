#!/bin/bash -e

set -e

ps -ef | grep "steampipe"
# export STEAMPIPE_CACHE_PENDING_QUERY_TIMEOUT=20
# export STEAMPIPE_LOG=trace
echo "starting test"
steampipe plugin install chaos
# steampipe service start

# steampipe query "select unique_col, a, b, c, error_after_delay from chaos_cache_check" --output json &> output4.json &
# sleep .5s
# steampipe query "select unique_col, a, b from chaos_cache_check" --output json &> output5.json &
# sleep .5s
# steampipe query "select unique_col, a, b from chaos_cache_check" --output json &> output6.json


# # query
# steampipe query --output json  "select * from chaos.chaos_all_column_types order by id limit 5"
# steampipe query --output json "select * from chaos.chaos_hydrate_columns_dependency order by id"
# steampipe query --output json "select * from chaos.chaos_parallel_hydrate_columns order by id"
# steampipe query --output json "select * from chaos.chaos_all_numeric_column order by id"
  
# ###
# steampipe query "create table all_columns (booleancolumn boolean, textcolumn1 CHAR(20), textcolumn2 VARCHAR(20),  textcolumn3 text, integercolumn1 smallint, integercolumn2 int, integercolumn3 SERIAL, integercolumn4 bigint,  integercolumn5 bigserial, numericColumn numeric(6,4), realColumn real, floatcolumn float,  date1 DATE,  time1 TIME,  timestamp1 TIMESTAMP, interval1 TIMESTAMPTZ, timestamp2 INTERVAL, array1 text[], jsondata jsonb, jsondata2 json, uuidcolumn UUID, ipAddress inet, macAddress macaddr, cidrRange cidr, xmlData xml, currency money)"
# steampipe query "INSERT INTO all_columns (booleancolumn, textcolumn1, textcolumn2, textcolumn3, integercolumn1, integercolumn2, integercolumn3, integercolumn4, integercolumn5, numericColumn, realColumn, floatcolumn, date1, time1, timestamp1, interval1, timestamp2, array1, jsondata, jsondata2, uuidcolumn, ipAddress, macAddress, cidrRange, xmlData, currency) VALUES (TRUE, 'Yes', 'test for varchar', 'This is a very long text for the PostgreSQL text column', 3278, 21445454, 2147483645, 92233720368547758, 922337203685477580, 23.5141543, 4660.33777, 4.6816421254887534, '1978-02-05', '08:00:00', '2016-06-22 19:10:25-07', '2016-06-22 19:10:25-07', '1 year 2 months 3 days', '{\"(408)-589-5841\"}','{ \"customer\": \"John Doe\", \"items\": {\"product\": \"Beer\",\"qty\": 6}}', '{ \"customer\": \"John Doe\", \"items\": {\"product\": \"Beer\",\"qty\": 6}}', '6948DF80-14BD-4E04-8842-7668D9C001F5', '192.168.0.0', '08:00:2b:01:02:03', '10.1.2.3/32', '<?xml version=\"1.0\"?><book><title>Manual</title><chapter>...</chapter></book>', 922337203685477.57)"
# steampipe query --output json "select * from all_columns"
# steampipe query "drop table all_columns"
# ###

# list_tables
# echo "LIST TABLES ############################"
# steampipe query "select fatal_error from chaos.chaos_list_errors" || true
# steampipe query "select fatal_error_after_streaming from chaos.chaos_list_errors" || true
# steampipe query --output json "select retryable_error from chaos.chaos_list_errors order by id limit 5" || true
# steampipe query "select retryable_error_after_streaming from chaos.chaos_list_errors" || true
# steampipe query "select should_ignore_error from chaos.chaos_list_errors" || true
# steampipe query --output json "select should_ignore_error_after_streaming from chaos.chaos_list_errors order by id" || true
# steampipe query --output json "select delay from chaos.chaos_list_errors order by id limit 5" || true
# steampipe query "select panic from chaos.chaos_list_errors" || true
# steampipe query "select panic_after_streaming from chaos.chaos_list_errors" || true
# steampipe query "select parent_fatal_error from chaos.chaos_list_parent_child" || true
# steampipe query "select parent_fatal_error_after_streaming from chaos.chaos_list_parent_child" || true
# steampipe query --output json "select parent_retryable_error from chaos.chaos_list_parent_child order by id limit 5" || true
# steampipe query "select parent_retryable_error_after_streaming from chaos.chaos_list_parent_child" || true
# steampipe query "select parent_should_ignore_error from chaos.chaos_list_parent_child" || true
# steampipe query --output json "select parent_should_ignore_error_after_streaming from chaos.chaos_list_parent_child order by id" || true
# steampipe query --output json "select parent_delay from chaos.chaos_list_parent_child order by id limit 5" || true
# steampipe query "select parent_panic from chaos.chaos_list_parent_child" || true
# steampipe query "select child_fatal_error from chaos.chaos_list_parent_child" || true
# steampipe query "select child_fatal_error_after_streaming from chaos.chaos_list_parent_child" || true
# steampipe query "select child_retryable_error from chaos.chaos_list_parent_child" || true
# steampipe query --output json "select child_retryable_error_after_streaming from chaos.chaos_list_parent_child order by id" || true
# steampipe query "select child_should_ignore_error from chaos.chaos_list_parent_child" || true
# steampipe query --output json "select child_should_ignore_error_after_streaming from chaos.chaos_list_parent_child order by id" || true
# steampipe query --output json "select child_delay from chaos.chaos_list_parent_child order by id" || true
# steampipe query "select child_panic from chaos.chaos_list_parent_child" || true
# echo "LIST TABLES ############################"

# get_tables
# echo "GET TABLES ############################"
# steampipe query "select fatal_error from chaos.chaos_get_errors where id=0" || true
# steampipe query "select retryable_error from chaos.chaos_get_errors where id=0" || true
# steampipe query "select ignorable_error from chaos.chaos_get_errors where id=0" || true
# steampipe query --output json "select delay from chaos.chaos_get_errors where id=0 order by id" || true
# steampipe query --output json "select panic from chaos.chaos_get_errors where id=0" || true
# steampipe query "select retryable_error_default_config from chaos.chaos_get_errors_default_config where id=0" || true
# steampipe query "select ignorable_error_default_config from chaos.chaos_get_errors_default_config where id=0" || true
# echo "GET TABLES ############################"

echo "HYDRATE TABLES ############################"
steampipe query "select fatal_error from chaos.chaos_hydrate_errors" || true
STEAMPIPE_CACHE=FALSE && steampipe query "select retryable_error from chaos.chaos_hydrate_errors" || true
steampipe query "select ignorable_error from chaos.chaos_hydrate_errors" || true
steampipe query --output json "select delay from chaos.chaos_hydrate_errors order by id" || true
steampipe query --output json "select panic from chaos.chaos_hydrate_errors" || true
echo "HYDRATE TABLES ############################"

# steampipe service stop --force
echo "fin"