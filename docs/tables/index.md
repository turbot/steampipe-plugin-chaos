---
organization: Turbot
category: ["software development"]
icon_url: "/images/plugins/turbot/steampipe.svg"
brand_color: "#a42a2d"
display_name: "Chaos (Steampipe)"
name: "chaos"
description: "Steampipe plugin for testing Steampipe in weird and wonderful ways."
---

# Chaos

The Chaos plugin is used internally to test Steampipe features and functionality in weird and wonderful ways  .

## Installation

Download and install the latest Chaos plugin:

```bash
$ steampipe plugin install chaos
Installing plugin chaos...
$
```

Run a query:

```bash
$ steampipe query
Welcome to Steampipe v0.0.11
Type ".inspect" for more information.
> select
  id,
  string_column,
  boolean_column,
  date_time_column,
  double_column,
  cidr_column,
  json_column
from
  all_column_types;

```
+----+------------------------+----------------+---------------------+---------------------+----------------+------------------------------------------------------------------------------------------------------------------------+
| id |     string_column      | boolean_column |  date_time_column   |    double_column    |  cidr_column   |                                                      json_column                                                       |
+----+------------------------+----------------+---------------------+---------------------+----------------+------------------------------------------------------------------------------------------------------------------------+
|  4 | stringValuesomething-4 | true           | 2001-07-14 18:00:34 | 0.23529411764705882 | 192.168.0.0/22 | {"Id":4,"Name":"stringValuesomething-4","Statement":{"Action":"iam:GetContextKeysForCustomPolicy","Effect":"Allow"}}   |
|  5 | stringValuesomething-5 | false          | 2001-09-11 07:49:00 | 0.29411764705882354 | 172.16.0.0/16  | {"Id":5,"Name":"stringValuesomething-5","Statement":{"Action":"iam:GetContextKeysForPrincipalPolicy","Effect":"Deny"}} |
|  9 | stringValuesomething-9 | false          | 2001-08-27 05:00:00 |  0.5294117647058824 | 10.0.0.1/32    | {"Id":9,"Name":"stringValuesomething-9","Statement":{"Action":"iam:GetContextKeysForPrincipalPolicy","Effect":"Deny"}} |
|  6 | stringValuesomething-6 | true           | 2001-06-04 22:00:32 | 0.35294117647058826 | 10.1.0.0/16    | {"Id":6,"Name":"stringValuesomething-6","Statement":{"Action":"iam:SimulateCustomPolicy","Effect":"Allow"}}            |
|  3 | stringValuesomething-3 | false          | 2001-12-12 06:19:00 | 0.17647058823529413 | 172.31.0.0/16  | {"Id":3,"Name":"stringValuesomething-3","Statement":{"Action":"iam:SimulatePrincipalPolicy","Effect":"Deny"}}          |
+----+------------------------+----------------+---------------------+---------------------+----------------+------------------------------------------------------------------------------------------------------------------------+
```