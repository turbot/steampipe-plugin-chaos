module github.com/turbot/steampipe-plugin-chaos

go 1.15

require (
	github.com/fatih/color v1.10.0 // indirect
	github.com/hashicorp/yamux v0.0.0-20200609203250-aecfd211c9ce // indirect
	github.com/iancoleman/strcase v0.1.3 // indirect
	github.com/mattn/go-runewidth v0.0.10 // indirect
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/oklog/run v1.1.0 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/turbot/go-kit v0.3.0
	github.com/turbot/steampipe-plugin-sdk v1.7.0
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

// main
//replace github.com/turbot/steampipe-plugin-sdk => github.com/turbot/steampipe-plugin-sdk v1.7.3-0.20211108132648-aad452166788
replace github.com/turbot/steampipe-plugin-sdk => /Users/kai/Dev/github/turbot/steampipe-plugin-sdk
