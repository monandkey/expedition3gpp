package expedition

type ExpeditionAction interface {
	SetParams(string, string, string, bool)
	Search() error
	Download() error
	Cache() error
}

type params struct {
	DocumentNumber  string
	DocumentVersion string
	OutputPath      string
	Cache           bool
}

type baseParams struct {
	params
	value []valueBody
}

type configParams struct {
	strageLocation     string
	cacheEnable        bool
	cacheRetentionTime int
	cacheLocation      string
}

type yamlStruct struct {
	YamlVersion int         `yaml:"version"`
	Title       string      `yaml:"title"`
	CreateDate  string      `yaml:"createdate"`
	Value       []valueBody `yaml:"value"`
}

type valueBody struct {
	Version string `yaml:"version"`
	Name    string `yaml:"name"`
	Url     string `yaml:"url"`
}
