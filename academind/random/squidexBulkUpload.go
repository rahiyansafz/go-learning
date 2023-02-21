package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func SquidexBulkUpload() error {
	// Set the Squidex API credentials
	appName := "konnectingdots"
	clientId := "konnectingdots:default"
	clientSecret := "5w5d9w0byuo4q8lxjidhyhyksyc0fugbqjy3cz0rzrsx"
	contentTypeName := "company"

	// Set the API endpoint
	endpoint := "https://cms.pdcl-online.com/api/content/" + appName + "/" + contentTypeName + "?publish=true"

	// Read the JSON data from a file or any other source
	data := []byte(`[{"name":"a","age":2},{"name":"b","age":3}]`)

	// Unmarshal the JSON data into a Go slice of maps
	var items []map[string]interface{}
	if err := json.Unmarshal(data, &items); err != nil {
		return fmt.Errorf("unmarshal json data: %w", err)
	}

	// Encode the data as a JSON array
	payload, err := json.Marshal(items)
	if err != nil {
		return fmt.Errorf("marshal items: %w", err)
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	// Concatenate the client ID and secret to form the access token
	accessToken := clientId + ":" + clientSecret

	// Encode the access token in base64 format
	encodedAccessToken := base64.StdEncoding.EncodeToString([]byte(accessToken))

	// Set the necessary headers for the request
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+encodedAccessToken)

	// Send the request to the Squidex API
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("send request: %w", err)
	}
	defer res.Body.Close()

	// Read the response from the Squidex API
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, res.Body); err != nil {
		return fmt.Errorf("read response: %w", err)
	}

	log.Println(buf.String())
	return nil
}
