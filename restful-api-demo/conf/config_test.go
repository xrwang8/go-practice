package conf

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestLoadConfigFromToml(t *testing.T) {
	should := assert.New(t)
	err := LoadConfigFromToml("../etc/demo.toml")
	if should.NoError(err) {
		should.Equal("restful-api-demo", C().App.Name)
	}
}

func TestLoadConfigFromEnv(t *testing.T) {
	should := assert.New(t)
	os.Setenv("MYSQL_DATABASE", "unit_test")

	err := LoadConfigFromEnv()
	if should.NoError(err) {
		should.Equal("unit_test", C().MySQL.Database)
	}
}
