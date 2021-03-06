package cmd

type params struct {
	documentNumber  string
	documentVersion string
	outputPath      string
	cache           bool
	releaseNumber   string
}

type configParams struct {
	strageLocation     string
	cacheEnable        bool
	cacheRetentionTime int
	cacheLocation      string
}
