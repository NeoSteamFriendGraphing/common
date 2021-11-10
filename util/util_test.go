package util

import (
	"os"
	"testing"

	"github.com/neosteamfriendgraphing/common"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	code := m.Run()

	os.Exit(code)
}

func TestIsValidFormatSteamIDWithValidSteamID(t *testing.T) {
	assert.True(t, IsValidFormatSteamID("76561197969081524"))
}

func TestIsValidFormatSteamIDWithInValidSteamID(t *testing.T) {
	assert.False(t, IsValidFormatSteamID("76561197969081524123456"))
}

func TestGetLocalIPAddress(t *testing.T) {
	IPAddress := GetLocalIPAddress()
	assert.Greater(t, len(IPAddress), 0, "returned IP address must not be empty")
}

func TestLoadLoggingConfigWithSpecifiedVariables(t *testing.T) {
	expectedConfig := common.LoggingFields{
		NodeName: "expectedName",
		NodeDC:   "expectedDC",
		LogPaths: []string{"stdout", "expectedLogPath"},
		NodeIPV4: GetLocalIPAddress(),
	}

	os.Setenv("NODE_NAME", "expectedName")
	os.Setenv("NODE_DC", "expectedDC")
	os.Setenv("LOG_PATH", "expectedLogPath")
	realLoggingConfig, err := LoadLoggingConfig()

	assert.Nil(t, err, "err should be nil for a valid logging config")
	assert.EqualValues(t, realLoggingConfig, expectedConfig)
}

func TestLoadLoggingConfigWithoutAllVariablesSetReturnsAnError(t *testing.T) {
	os.Setenv("NODE_NAME", "")
	os.Setenv("NODE_DC", "expectedDC")
	os.Setenv("LOG_PATH", "expectedLogPath")
	_, err := LoadLoggingConfig()

	assert.Contains(t, err.Error(), "one or more required environment variables are not set:")
}

func TestIsValidFormatGraphID(t *testing.T) {
	isValid, err := IsValidFormatGraphID("helloWorld1234234")
	assert.True(t, isValid)
	assert.Nil(t, err)
}
func TestIsValidFormatGraphIDWithPathTraversalInput(t *testing.T) {
	isValid, err := IsValidFormatGraphID("../../../../.env")
	assert.False(t, isValid)
	assert.Nil(t, err)
}

func TestEnsureAllEnvvarsAreSet(t *testing.T) {
	os.Setenv("API_PORT", "techno")
	os.Setenv("LOG_PATH", "techno")
	os.Setenv("NODE_NAME", "techno")
	os.Setenv("SYSTEM_STATS_BUCKET", "techno")
	os.Setenv("SYSTEM_STATS_BUCKET_TOKEN", "techno")
	os.Setenv("ORG", "techno")
	os.Setenv("INFLUXDB_URL", "techno")

	assert.NoError(t, EnsureAllEnvVarsAreSet())

	os.Setenv("API_PORT", "")
	os.Setenv("LOG_PATH", "")
	os.Setenv("NODE_NAME", "")
	os.Setenv("SYSTEM_STATS_BUCKET", "")
	os.Setenv("SYSTEM_STATS_BUCKET_TOKEN", "")
	os.Setenv("ORG", "")
	os.Setenv("INFLUXDB_URL", "")
}

func TestEnsureAllEnvvarsCatchesAnUnsetDefaultVariable(t *testing.T) {
	os.Setenv("API_PORT", "techno")
	os.Setenv("LOG_PATH", "techno")
	os.Setenv("NODE_NAME", "techno")
	os.Setenv("SYSTEM_STATS_BUCKET_TOKEN", "techno")
	os.Setenv("ORG", "techno")
	os.Setenv("INFLUXDB_URL", "techno")

	expectedErrorMsg := "One or more env vars were not set: SYSTEM_STATS_BUCKET\n"
	assert.EqualError(t, EnsureAllEnvVarsAreSet(), expectedErrorMsg)

	os.Setenv("API_PORT", "")
	os.Setenv("LOG_PATH", "")
	os.Setenv("NODE_NAME", "")
	os.Setenv("SYSTEM_STATS_BUCKET", "")
	os.Setenv("SYSTEM_STATS_BUCKET_TOKEN", "")
	os.Setenv("ORG", "")
	os.Setenv("INFLUXDB_URL", "")
}

func TestEnsureAllEnvvarsCatchesAnUnsetServiceSpecificVariable(t *testing.T) {
	os.Setenv("API_PORT", "techno")
	os.Setenv("LOG_PATH", "techno")
	os.Setenv("NODE_NAME", "techno")
	os.Setenv("SYSTEM_STATS_BUCKET", "techno")
	os.Setenv("SYSTEM_STATS_BUCKET_TOKEN", "techno")
	os.Setenv("ORG", "techno")
	os.Setenv("INFLUXDB_URL", "techno")

	expectedErrorMsg := "One or more env vars were not set: ANYCANS\n"
	assert.EqualError(t, EnsureAllEnvVarsAreSet("ANYCANS"), expectedErrorMsg)
}
