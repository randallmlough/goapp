package command

import (
	"context"
	"fmt"
	"github.com/randallmlough/goapp/config"
	"github.com/randallmlough/goapp/internal/utils"
	"github.com/urfave/cli/v2"
	"path"
)

func DefaultConfig() *Config {
	return &Config{}
}

type Config struct {
	Project *Project
}

var CFG = &cli.StringFlag{
	Name:  "config, c",
	Usage: "Load configuration from `FILE`",
}

func LoadConfig(c *cli.Context) error {
	workingPath, err := utils.GetWorkingPath()
	if err != nil {
		return fmt.Errorf("failed to get working path %w", err)
	}
	configPath := path.Join(workingPath, "config.yml")
	if !utils.FileExists(configPath) {
		configToContext(c, "config", DefaultConfig())
	} else {
		var cfg Config
		if err := config.LoadConfig(configPath, &cfg); err != nil {
			return fmt.Errorf("failed to load config %w", err)
		}
		configToContext(c, "config", cfg)
	}
	return nil
}

func configFromContext(c *cli.Context) *Config {
	if tmp := c.Context.Value("config"); tmp != nil {
		if cfg, ok := tmp.(*Config); ok {
			return cfg
		}
	}
	return nil
}

func configToContext(c *cli.Context, key, value interface{}) {
	c.Context = context.WithValue(c.Context, key, value)
}
