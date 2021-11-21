package config

type ConfigAction interface {
	SetParams(string, bool, int, string)
	Load() params
	Write() error
}

type params struct {
	StrageLocation     string `yaml:"strageLocation"`
	CacheEnable        bool   `yaml:"cacheEnable"`
	CacheRetentionTime int    `yaml:"cacheRetentionTime"`
	CacheLocation      string `yaml:"cacheLocation"`
}

type baseParams struct {
	params
}
