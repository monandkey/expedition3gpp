module github.com/monandkey/expedition3gpp

go 1.16

require (
	github.com/PuerkitoBio/goquery v1.7.1 // indirect
	github.com/monandkey/expedition3gpp/cmd v0.0.0-00010101000000-000000000000 // indirect
	github.com/monandkey/expedition3gpp/pkg/config v0.0.0-00010101000000-000000000000 // indirect
	github.com/spf13/cobra v1.2.1 // indirect
	local.packages/expedition3gpp v0.0.0-00010101000000-000000000000 // indirect
)

replace (
	github.com/monandkey/expedition3gpp/cmd => ./cmd
	github.com/monandkey/expedition3gpp/pkg/config => ./pkg/config
	local.packages/expedition3gpp => ./pkg/expedition3gpp
)
