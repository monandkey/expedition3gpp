module github.com/monandkey/expedition3gpp/internal/pkg/config

go 1.16

require (
	github.com/monandkey/expedition3gpp/internal/pkg/fileutil v0.0.0-00010101000000-000000000000
	gopkg.in/yaml.v2 v2.4.0
)

replace github.com/monandkey/expedition3gpp/internal/pkg/fileutil => ../fileutil
