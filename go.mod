module github.com/monandkey/expedition3gpp

go 1.16

require github.com/monandkey/expedition3gpp/cmd v0.0.0-00010101000000-000000000000

replace (
	github.com/monandkey/expedition3gpp/cmd => ./cmd
	github.com/monandkey/expedition3gpp/internal/pkg/config => ./internal/pkg/config
	github.com/monandkey/expedition3gpp/internal/pkg/expedition => ./internal/pkg/expedition
	github.com/monandkey/expedition3gpp/internal/pkg/fileutil => ./internal/pkg/fileutil
)
