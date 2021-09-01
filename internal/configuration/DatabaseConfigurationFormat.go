package configuration

type DatabaseConfiguration struct {
	User     string `toml:"user"`
	Password string `toml:"password"`
	DbName   string `toml:"dbName"`
}
