module github.com/monandkey/expedition3gpp/cmd

go 1.16

replace (
	github.com/monandkey/expedition3gpp/internal/pkg/config => ../internal/pkg/config
	github.com/monandkey/expedition3gpp/internal/pkg/expedition => ../internal/pkg/expedition
	github.com/monandkey/expedition3gpp/internal/pkg/fileutil => ../internal/pkg/fileutil
)

require (
	github.com/monandkey/expedition3gpp/internal/pkg/config v0.0.0-00010101000000-000000000000
	github.com/monandkey/expedition3gpp/internal/pkg/expedition v0.0.0-00010101000000-000000000000
	github.com/spf13/cobra v1.2.1
)
