package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var cfgFilenames = []string{".config.yml", "config.yml", "config.yaml"}

// LoadConfigFromDefaultLocations looks for a config file in the current directory, and all parent directories
// walking up the tree. The closest config file will be returned.
func LoadConfigFromDefaultLocations(dst interface{}) error {
	cfgFile, err := findCfg()
	if err != nil {
		return err
	}

	err = os.Chdir(filepath.Dir(cfgFile))
	if err != nil {
		return fmt.Errorf("unable to enter config dir %w", err)
	}

	if err := LoadConfig(cfgFile, dst); err != nil {
		return err
	}

	return nil
}

// LoadConfig reads the gqlgen.yml config file
func LoadConfig(filename string, dst interface{}) error {

	b, err := readFileEnvParsing(filename)
	if err != nil {
		return fmt.Errorf("unable to read config %w", err)
	}

	if err := Unmarshall(b, dst); err != nil {
		return err
	}
	return nil
}

func Unmarshall(data []byte, dst interface{}) error {
	if err := yaml.UnmarshalStrict(data, dst); err != nil {
		return fmt.Errorf("unable to parse config %w", err)
	}
	return nil
}

// findCfg searches for the config file in this directory and all parents up the tree
// looking for the closest match
func findCfg() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("unable to get working dir to findCfg %w", err)
	}

	cfg := findCfgInDir(dir)

	for cfg == "" && dir != filepath.Dir(dir) {
		dir = filepath.Dir(dir)
		cfg = findCfgInDir(dir)
	}

	if cfg == "" {
		return "", os.ErrNotExist
	}

	return cfg, nil
}

func findCfgInDir(dir string) string {
	for _, cfgName := range cfgFilenames {
		path := filepath.Join(dir, cfgName)
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}
	return ""
}

func readFileEnvParsing(path string) ([]byte, error) {
	confContent, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	// expand environment variables
	return []byte(os.Expand(string(confContent), ParseEnv)), nil
}

func ParseEnv(key string) string {
	parts := strings.Split(key, ":-")
	if env, ok := os.LookupEnv(parts[0]); ok {
		return env
	} else if len(parts) > 1 {
		return parts[1]
	}
	return ""
}
