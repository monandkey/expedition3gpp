package expedition

// ExpeditionAction is an interface that defines methods to manipulate this package.
type ExpeditionAction interface {
	SetParams(string, string, string, bool, string)
	Search() error
	Download() error
	Cache() error
}

// params is a structure that stores the information needed to download the 3GPP document
type params struct {
	DocumentNumber  string
	DocumentVersion string
	OutputPath      string
	Cache           bool
	releaseNumber   string
}

// baseParams is a structure that inherits from params
type baseParams struct {
	params
	value []valueBody
}

// configParams is a structure to store configuration information
type configParams struct {
	strageLocation     string
	cacheEnable        bool
	cacheRetentionTime int
	cacheLocation      string
}

// yamlStruct is a structure for loading yaml
type yamlStruct struct {
	YamlVersion int         `yaml:"version"`
	Title       string      `yaml:"title"`
	CreateDate  string      `yaml:"createdate"`
	Value       []valueBody `yaml:"value"`
}

// valueBody is a structure for storing information when creating a cache
type valueBody struct {
	Version string `yaml:"version"`
	Name    string `yaml:"name"`
	Url     string `yaml:"url"`
}
