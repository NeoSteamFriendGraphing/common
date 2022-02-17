package util

import (
	"errors"
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
		Service:  "common",
	}

	os.Setenv("NODE_NAME", "expectedName")
	os.Setenv("NODE_DC", "expectedDC")
	os.Setenv("LOG_PATH", "expectedLogPath")
	os.Setenv("SERVICE", "common")
	realLoggingConfig, err := LoadLoggingConfig()

	assert.Nil(t, err, "err should be nil for a valid logging config")
	assert.EqualValues(t, realLoggingConfig, expectedConfig)
}

func TestLoadLoggingConfigWithoutAllVariablesSetReturnsAnError(t *testing.T) {
	os.Setenv("NODE_NAME", "")
	os.Setenv("NODE_DC", "expectedDC")
	os.Setenv("LOG_PATH", "expectedLogPath")
	os.Setenv("SERVICE", "common")
	_, err := LoadLoggingConfig()

	assert.Contains(t, err.Error(), "one or more required environment variables are not set:")
}

func TestEnsureAllEnvvarsAreSet(t *testing.T) {
	os.Setenv("AUTH_KEY", "dnb")
	os.Setenv("API_PORT", "techno")
	os.Setenv("LOG_PATH", "techno")
	os.Setenv("NODE_NAME", "techno")
	os.Setenv("SERVICE", "breakbeat")
	os.Setenv("SYSTEM_STATS_BUCKET", "techno")
	os.Setenv("SYSTEM_STATS_BUCKET_TOKEN", "groove")
	os.Setenv("ENDPOINT_LATENCIES_BUCKET", "industrial")
	os.Setenv("ENDPOINT_LATENCIES_BUCKET_TOKEN", "dnb")
	os.Setenv("ORG", "techno")
	os.Setenv("INFLUXDB_URL", "techno")

	assert.NoError(t, EnsureAllEnvVarsAreSet())

	os.Setenv("AUTH_KEY", "")
	os.Setenv("API_PORT", "")
	os.Setenv("LOG_PATH", "")
	os.Setenv("NODE_NAME", "")
	os.Setenv("SERVICE", "")
	os.Setenv("SYSTEM_STATS_BUCKET", "")
	os.Setenv("SYSTEM_STATS_BUCKET_TOKEN", "")
	os.Setenv("ENDPOINT_LATENCIES_BUCKET", "")
	os.Setenv("ENDPOINT_LATENCIES_BUCKET_TOKEN", "")
	os.Setenv("ORG", "")
	os.Setenv("INFLUXDB_URL", "")
}

func TestEnsureAllEnvvarsCatchesAnUnsetDefaultVariable(t *testing.T) {
	os.Setenv("AUTH_KEY", "dnb")
	os.Setenv("API_PORT", "techno")
	os.Setenv("LOG_PATH", "techno")
	os.Setenv("NODE_NAME", "techno")
	os.Setenv("SERVICE", "breakbeat")
	os.Setenv("SYSTEM_STATS_BUCKET_TOKEN", "techno")
	os.Setenv("ORG", "techno")
	os.Setenv("INFLUXDB_URL", "techno")
	os.Setenv("ENDPOINT_LATENCIES_BUCKET", "industrial")
	os.Setenv("ENDPOINT_LATENCIES_BUCKET_TOKEN", "dnb")

	expectedErrorMsg := "one or more env vars were not set: SYSTEM_STATS_BUCKET\n"
	assert.Contains(t, EnsureAllEnvVarsAreSet().Error(), expectedErrorMsg)

	os.Setenv("AUTH_KEY", "")
	os.Setenv("API_PORT", "")
	os.Setenv("LOG_PATH", "")
	os.Setenv("NODE_NAME", "")
	os.Setenv("SERVICE", "")
	os.Setenv("SYSTEM_STATS_BUCKET", "")
	os.Setenv("SYSTEM_STATS_BUCKET_TOKEN", "")
	os.Setenv("ENDPOINT_LATENCIES_BUCKET", "")
	os.Setenv("ENDPOINT_LATENCIES_BUCKET_TOKEN", "")
	os.Setenv("ORG", "")
	os.Setenv("INFLUXDB_URL", "")
}

func TestEnsureAllEnvvarsCatchesAnUnsetServiceSpecificVariable(t *testing.T) {
	os.Setenv("AUTH_KEY", "dnb")
	os.Setenv("API_PORT", "techno")
	os.Setenv("LOG_PATH", "techno")
	os.Setenv("NODE_NAME", "techno")
	os.Setenv("SERVICE", "breakbeat")
	os.Setenv("SYSTEM_STATS_BUCKET", "techno")
	os.Setenv("SYSTEM_STATS_BUCKET_TOKEN", "techno")
	os.Setenv("ENDPOINT_LATENCIES_BUCKET", "industrial")
	os.Setenv("ENDPOINT_LATENCIES_BUCKET_TOKEN", "dnb")
	os.Setenv("ORG", "techno")
	os.Setenv("INFLUXDB_URL", "techno")

	expectedErrorMsg := "one or more env vars were not set: ANYCANS\n"
	assert.Contains(t, EnsureAllEnvVarsAreSet("ANYCANS").Error(), expectedErrorMsg)

	os.Setenv("AUTH_KEY", "")
	os.Setenv("API_PORT", "")
	os.Setenv("LOG_PATH", "")
	os.Setenv("NODE_NAME", "")
	os.Setenv("SERVICE", "")
	os.Setenv("SYSTEM_STATS_BUCKET", "")
	os.Setenv("SYSTEM_STATS_BUCKET_TOKEN", "")
	os.Setenv("ENDPOINT_LATENCIES_BUCKET", "")
	os.Setenv("ENDPOINT_LATENCIES_BUCKET_TOKEN", "")
	os.Setenv("ORG", "")
	os.Setenv("INFLUXDB_URL", "")
}

func TestMakeErr(t *testing.T) {
	expectedSubstring := "detailed explanation: could not do it"
	randomErr := errors.New("could not do it")
	actualErr := MakeErr(randomErr, "detailed explanation")

	assert.Contains(t, actualErr.Error(), expectedSubstring)
}
