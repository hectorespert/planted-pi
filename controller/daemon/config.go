package daemon

import (
	"os"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Database string `json:"database" yaml:"database"`
	PAM      bool   `json:"pam" yaml:"pam"`
}

var DefaultConfig = Config{
	Database: "reef-pi.db",
	PAM:      false,
}

func ParseConfig(filename string) (Config, error) {
	c := DefaultConfig
	content, err := os.ReadFile(filename)
	if err != nil {
		return c, err
	}
	return c, yaml.Unmarshal(content, &c)
}
