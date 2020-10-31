package config

import (
	"testing"
)

func TestParse(t *testing.T) {
	config, err := Parse()
	if err != nil {
		t.Errorf("Couldn't load configuration %v", err)
	}

	if config.EntryPoint != "hightower.go" {
		t.Error("Entrypoint is not correct")
	}
}
