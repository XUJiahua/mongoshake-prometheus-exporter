package config

import (
	"fmt"
	"testing"

	"github.com/BurntSushi/toml"
)

func TestLoadConfig(t *testing.T) {
	newCfg := &Config{
		Probes: []ProbeConfig{
			{BaseURL: "http://localhost:9100"},
			{BaseURL: "http://localhost:9100"},
		},
	}
	if err := newCfg.Validate(); err != nil {
		t.Fatalf("validate config error: %v", err)
	}
	cfgBytes, err := toml.Marshal(newCfg)
	if err != nil {
		t.Fatalf("marshal config error: %v", err)
	}
	fmt.Println(string(cfgBytes))
}
