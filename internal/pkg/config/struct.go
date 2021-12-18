package config

// ConfigAction is an interface that defines methods for manipulating the configuration.
type ConfigAction interface {
	SetParams(string, bool, int, string)
	Load() params
	Write() error
}

// params is a structure that manages parameters for configuring
type params struct {
	StrageLocation     string `yaml:"strageLocation"`
	CacheEnable        bool   `yaml:"cacheEnable"`
	CacheRetentionTime int    `yaml:"cacheRetentionTime"`
	CacheLocation      string `yaml:"cacheLocation"`
}

// baseParams is a structure that inherits from params
type baseParams struct {
	params
}
