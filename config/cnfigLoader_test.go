package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLoading(t *testing.T) {
	var confFile ConfGoPath
	confFile.FromFile("resources/pha_test.yml")
	assert.Equal(t, confFile.PhaConf.Endpoint, "https://domain.example", "The params should be the same.")
	assert.Equal(t, confFile.PhaConf.Token, "cli-34567890", "The params should be the same.")
	assert.Equal(t, confFile.PhaConf.Repo, "repo01", "The params should be the same.")
	assert.Equal(t, confFile.PhaJenkins.Endpoint, "https://ci-jenkins", "The params should be the same.")
	assert.Equal(t, confFile.PhaJenkins.Username, "example", "The params should be the same.")
	assert.Equal(t, confFile.PhaJenkins.Token, "34567890", "The params should be the same.")
	assert.Equal(t, confFile.PhaJenkins.Pipeline, "testingJenkins", "The params should be the same.")
	assert.Equal(t, confFile.PhaImap.Username, "my-test-email@gmail.com", "The params should be the same.")
	assert.Equal(t, confFile.PhaImap.Password, "my-dummy-password", "The params should be the same.")
	assert.Equal(t, confFile.PhaImap.Address, "imap.gmail.com:993", "The params should be the same.")
}
