module github.com/monandkey/expedition3gpp/internal/pkg/expedition

go 1.16

replace (
	github.com/monandkey/expedition3gpp/internal/pkg/config => ../config
	github.com/monandkey/expedition3gpp/internal/pkg/fileutil => ../fileutil
)

require (
	github.com/PuerkitoBio/goquery v1.8.0
	github.com/monandkey/expedition3gpp/internal/pkg/config v0.0.0-00010101000000-000000000000
	github.com/monandkey/expedition3gpp/internal/pkg/fileutil v0.0.0-00010101000000-000000000000
	gopkg.in/yaml.v2 v2.4.0
)
