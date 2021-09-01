package configuration

type GeneralConfiguration struct {
	Server          *Configuration         `toml:"server"`
	Logging         *LoggerConfiguration   `toml:"logging"`
	DbConfiguration *DatabaseConfiguration `toml:"database"`
}
