package config_test

import (
	"testing"

	"github.com/xiao-hub-create/vblog/config"
)

func TestLoadConfigFromYaml(t *testing.T) {
	if err := config.LoadConfigFromYaml("application.yaml"); err != nil {
		t.Fatal(err)
	}

	t.Log(config.Get().MySQL)
}
