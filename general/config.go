package general

type MysqlConfig struct {
	DNS string `yaml:"dns"`
}

type LogConfig struct {
	Output    uint32 `yaml:"output"`
	Formatter uint32 `yaml:"formatter"`
	Level     uint32 `yaml:"level"`
}

type Terminal struct {
	ConfigFilePath string
	LogDirectory   string
}
