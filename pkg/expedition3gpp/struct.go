package expedition3gpp

type params struct {
	StrageLocation     string `yaml:"strageLocation"`
	CacheEnable        bool   `yaml:"cacheEnable"`
	CacheRetentionTime int    `yaml:"cacheRetentionTime"`
	CacheLocation      string `yaml:"cacheLocation"`
}

type configPath struct {
	path string
}

type disassembledCharacter struct {
	homedir  string
	separate string
	filename string
}

type InitConfig struct {
	StrageLocation     string
	CacheEnable        bool
	CacheRetentionTime int
	CacheLocation      string
}
