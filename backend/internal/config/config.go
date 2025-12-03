package config

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Secret string
}

type ServerConfig struct {
	Address string
}

type DatabaseConfig struct {
	Address      string
	MaxOpenConns int
	MaxIdleConns int
	MaxIdleTime  string
}
