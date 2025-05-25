package config

import (
	"flag"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type DbConfig struct {
	DriverName string `yaml:"driver" env-default:"postgres"`
	Protocol   string `yaml:"protocol" env-default:"postgres"`
	User       string `yaml:"user" env-required:"true"`
	Password   string `yaml:"password" env-required:"true"`
	Host       string `yaml:"host" env-required:"true"`
}

type Config struct {
	DB            DbConfig `yaml:"db" env-required:"true"`
	ImagePath     string   `yaml:"img_path" env-required:"true"`
	BootstrapPath string   `yaml:"bootstrap_path" env-required:"true"`
	TemplatePath  string   `yaml:"tpl_path" env-required:"true"`
	Host          string   `yaml:"host" env-default:":8080"`
	Mode          string   `yaml:"mode" env-default:"release"`
}

func Load() (*Config, error) {
	configPath := fetchConfigPath()
	if configPath == "" {
		return nil, fmt.Errorf("config path is empty")
	}

	return LoadPath(configPath)
}

func LoadPath(configPath string) (*Config, error) {
	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return nil, fmt.Errorf("cannot read config: %w", err)
	}

	return &cfg, nil
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}

func (c *DbConfig) DataSourceName() string {
	return c.Protocol + "://" + c.User + ":" + c.Password + "@" + c.Host
}
