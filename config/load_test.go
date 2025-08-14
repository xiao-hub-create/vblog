package config

import "testing"

func TestLoadConfigFromYaml(t *testing.T) {
	if err := config.LoadConfigFromYaml("application.yaml"); err != nil {
		t.Fatal(err)
	}

	t.Log(config.Get().MySQL)
}
