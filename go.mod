module github.com/monandkey/expedition3gpp

go 1.16

require (
	github.com/monandkey/expedition3gpp/cmd v0.0.0-00010101000000-000000000000
	github.com/monandkey/expedition3gpp/internal/pkg/config v0.0.0-00010101000000-000000000000 // indirect
	github.com/spf13/cobra v1.2.1 // indirect
)

replace (
	github.com/monandkey/expedition3gpp/cmd => ./cmd
	github.com/monandkey/expedition3gpp/internal/pkg/config => ./internal/pkg/config
	github.com/monandkey/expedition3gpp/internal/pkg/expedition => ./internal/pkg/expedition
)
