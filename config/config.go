package config

type Config struct {
	App      *AppConfig      `yaml:"app"`
	Log      *LogConfig      `yaml:"log"`
	Database *DatabaseConfig `yaml:"database"`
}

type AppConfig struct {
	Env        string `yaml:"env"`
	Name       string `yaml:"name"`
	ServerAddr string `yaml:"service_addr"`
	Timeout    int    `yaml:"timeout"`
}

type LogConfig struct {
	Level      string `yaml:"level" toml:"level"`
	File       string `yaml:"file" toml:"file"`
	Format     string `yaml:"format" toml:"format"`
	MaxSize    int    `yaml:"max-size" toml:"max-size"`
	MaxAge     int    `yaml:"max-age" toml:"max-age"`
	MaxBackups int    `yaml:"max-backups" toml:"max-backups"`
	Compress   bool   `yaml:"compress" toml:"compress"`
}

type DatabaseConfig struct {
	Type             string `yaml:"type"`
	Address          string `yaml:"address"`
	Port             int    `yaml:"port"`
	Username         string `yaml:"username"`
	Password         string `yaml:"password"`
	DBName           string `yaml:"db-name"`
	ConnConfig       string `yaml:"conn-config"`
	EnableForeignKey bool   `yaml:"enable-foreign-key"`
	MaxIdleConns     int    `yaml:"max-idle-conns"`
	MaxOpenConns     int    `yaml:"max-open-conns"`
	ConnMaxLifetime  int    `yaml:"conn-max-lifetime"`
	ConnMaxIdleTime  int    `yaml:"conn-max-idle-time"`
}
