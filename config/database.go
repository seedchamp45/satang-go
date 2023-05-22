package config

type DatabaseConfig struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Name     string `yaml:"name"`
		DBName   string `yaml:"dbname"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"database"`
}
