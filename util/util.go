package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/neosteamfriendgraphing/common"
	"go.uber.org/zap"
)

// IsValidFormatSteamID determines if a string is a valid
// format steam64ID (17 numerical digits)
func IsValidFormatSteamID(steamID string) bool {
	if len(steamID) != 17 {
		return false
	}
	match, _ := regexp.MatchString("([0-9]){17}", steamID)
	return match
}

// GetLocalIPAddress retrieves the local IP (port not included) for the current
// system as this is used in logs for quick access
func GetLocalIPAddress() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	addrWithNoPort := strings.Split(conn.LocalAddr().(*net.UDPAddr).String(), ":")
	return addrWithNoPort[0]
}

// LoadLoggingConfig loads required config from .env that are default logging fields
func LoadLoggingConfig() (common.LoggingFields, error) {
	logFieldsConfig := common.LoggingFields{
		NodeName: os.Getenv("NODE_NAME"),
		NodeDC:   os.Getenv("NODE_DC"),
		LogPaths: []string{"stdout", os.Getenv("LOG_PATH")},
		NodeIPV4: GetLocalIPAddress(),
		Service:  os.Getenv("SERVICE"),
	}
	if logFieldsConfig.NodeName == "" || logFieldsConfig.NodeDC == "" ||
		logFieldsConfig.LogPaths[1] == "" || logFieldsConfig.NodeIPV4 == "" {

		return common.LoggingFields{}, fmt.Errorf("one or more required environment variables are not set: %v", logFieldsConfig)
	}
	return logFieldsConfig, nil
}

// InitLogger Initialises the zap logger and returns a pointer to an instance of it,
// this also involves creating the logfile specified by LOG_PATH
func InitLogger(logFieldsConfig common.LoggingFields) *zap.Logger {
	os.OpenFile(logFieldsConfig.LogPaths[1], os.O_RDONLY|os.O_CREATE, 0666)
	c := zap.NewProductionConfig()
	c.OutputPaths = []string{"stdout", logFieldsConfig.LogPaths[1]}

	globalLogFields := make(map[string]interface{})
	globalLogFields["nodeName"] = logFieldsConfig.NodeName
	globalLogFields["nodeDC"] = logFieldsConfig.NodeDC
	globalLogFields["nodeIPV4"] = logFieldsConfig.NodeIPV4
	globalLogFields["service"] = logFieldsConfig.Service
	c.InitialFields = globalLogFields

	log, err := c.Build()
	if err != nil {
		panic(err)
	}
	return log
}

// GetCurrentTimeInMs returns the current timestamp in milliseconds
func GetCurrentTimeInMs() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// GetRequestStartTimeInTimeFormat returns the int64 timestamp (in milliseconds) of a string
// version of a timestamp
func GetRequestStartTimeInTimeFormat(requestStartTimeString string) int64 {
	requestStartTime, err := strconv.ParseInt(requestStartTimeString, 10, 64)
	if err != nil {
		panic(err)
	}
	return requestStartTime
}

// SendBasicInvalidResponse sends an invalid response back to the user with specified
// status code and error message. This is used for invalid user input
func SendBasicInvalidResponse(w http.ResponseWriter, req *http.Request, msg string, vars map[string]string, statusCode int) {
	w.WriteHeader(statusCode)
	response := struct {
		Error string `json:"error"`
	}{
		msg,
	}
	json.NewEncoder(w).Encode(response)
}

// SendBasicErrorResponse sends an error response back to the user with specified
// status code and error message. This is used for an error in the system
func SendBasicErrorResponse(w http.ResponseWriter, req *http.Request, err error, vars map[string]string, statusCode int) {
	w.WriteHeader(http.StatusInternalServerError)
	response := struct {
		Error string `json:"error"`
	}{
		fmt.Sprintf("Give the code monkeys this ID: '%s'", vars["requestID"]),
	}
	json.NewEncoder(w).Encode(response)
}

// GetAndRead executes a HTTP GET request and returns the body
// of the response in []byte format or an error if it's not nil
func GetAndRead(URL string, headers []http.Header) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return []byte{}, MakeErr(err)
	}
	for _, header := range headers {
		for key, val := range header {
			req.Header.Set(key, val[0])
		}
	}
	res, err := client.Do(req)
	if err != nil {
		return []byte{}, MakeErr(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return []byte{}, MakeErr(err)
	}

	return body, nil
}

func EnsureAllEnvVarsAreSet(serviceSpecificEnvVars ...string) error {
	resultString := ""

	if os.Getenv("AUTH_KEY") == "" {
		resultString += "AUTH_KEY\n"
	}
	if os.Getenv("API_PORT") == "" {
		resultString += "API_PORT\n"
	}
	if os.Getenv("LOG_PATH") == "" {
		resultString += "LOG_PATH\n"
	}
	if os.Getenv("NODE_NAME") == "" {
		resultString += "NODE_NAME\n"
	}
	if os.Getenv("SERVICE") == "" {
		resultString += "SERVICE\n"
	}
	if os.Getenv("SYSTEM_STATS_BUCKET") == "" {
		resultString += "SYSTEM_STATS_BUCKET\n"
	}
	if os.Getenv("SYSTEM_STATS_BUCKET_TOKEN") == "" {
		resultString += "SYSTEM_STATS_BUCKET_TOKEN\n"
	}
	if os.Getenv("ENDPOINT_LATENCIES_BUCKET") == "" {
		resultString += "ENDPOINT_LATENCIES_BUCKET\n"
	}
	if os.Getenv("ENDPOINT_LATENCIES_BUCKET_TOKEN") == "" {
		resultString += "ENDPOINT_LATENCIES_BUCKET_TOKEN\n"
	}
	if os.Getenv("ORG") == "" {
		resultString += "ORG\n"
	}
	if os.Getenv("INFLUXDB_URL") == "" {
		resultString += "INFLUXDB_URL\n"
	}

	// Service specific env vars
	for _, envVar := range serviceSpecificEnvVars {
		if os.Getenv(envVar) == "" {
			resultString += fmt.Sprintf("%s\n", envVar)
		}
	}

	if resultString != "" {
		return MakeErr(errors.New("all env vars were not set"), fmt.Sprintf("one or more env vars were not set: %s", resultString))
	}
	return nil
}

func SetupCORS(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

// MakeErr creates an error with a trace to where this function was
// called from
// 		errorWithLineTrace := MakeErr(err, "heres an error that was thrown because of x, y, z...")
func MakeErr(err error, msg ...string) error {
	if err == nil {
		return nil
	}
	_, file, line, _ := runtime.Caller(1)
	path, _ := os.Getwd()
	return fmt.Errorf("%s:%d %s: %s", strings.TrimPrefix(file, path), line, msg[0], err)
}

func GetBaseURLPath(r *http.Request) string {
	urlPath := strings.Split(r.URL.EscapedPath(), "/")

	if len(urlPath) > 2 {
		return strings.Join(urlPath[1:], "/")
	}
	if len(urlPath) == 2 && urlPath[1] != "" {
		return urlPath[1]
	}
	return "/"
}
