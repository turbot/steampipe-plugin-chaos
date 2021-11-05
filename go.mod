module github.com/turbot/steampipe-plugin-chaos

go 1.15

require (
	github.com/agext/levenshtein v1.2.2 // indirect
	github.com/apparentlymart/go-dump v0.0.0-20190214190832-042adf3cf4a0 // indirect
	github.com/fatih/color v1.10.0 // indirect
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/hashicorp/yamux v0.0.0-20200609203250-aecfd211c9ce // indirect
	github.com/iancoleman/strcase v0.1.3 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/mattn/go-runewidth v0.0.10 // indirect
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/mitchellh/go-wordwrap v1.0.0 // indirect
	github.com/mitchellh/mapstructure v1.4.2 // indirect
	github.com/oklog/run v1.1.0 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/turbot/go-kit v0.3.0
	github.com/turbot/steampipe-plugin-sdk v1.7.3-0.20211105154226-b6d5a7bbac8d
	golang.org/x/net v0.0.0-20211005215030-d2e5035098b3 // indirect
	golang.org/x/sys v0.0.0-20211004093028-2c5d950f24ef // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20211005153810-c76a74d43a8e // indirect
	google.golang.org/grpc v1.41.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

//replace github.com/turbot/steampipe-plugin-sdk => github.com/turbot/steampipe-plugin-sdk pm_test_acc
replace github.com/turbot/steampipe-plugin-sdk => github.com/turbot/steampipe-plugin-sdk v1.7.3-0.20211105182102-9ed5a7d29686

//replace github.com/turbot/steampipe-plugin-sdk => /Users/kai/Dev/github/turbot/steampipe-plugin-sdk
