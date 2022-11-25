package api

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"testing"
)

func TestPing(t *testing.T) {
	const expectedMsg = "10.9.3-MariaDB-1:10.9.3+maria~ubu2204"
	os.Setenv("ENVIROMENT", "test")
	defer os.Unsetenv("ENVIROMENT")

	res, err := http.Get("http://localhost:8080/ping")
	if err != nil {
		t.Error("Error when making request to /ping. Error: ", err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Error("Error when reading response body. Error:", err)
	}

	response := make(map[string]string)
	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Error("Invalid response")
	}

	if response["message"] != expectedMsg {
		t.Error("Expected message: ", expectedMsg, ", Got", response["message"])
	}

}
