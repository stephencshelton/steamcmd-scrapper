package main

import (
  "bytes"
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "os"

  "gopkg.in/yaml.v2"
)

func main() {
  // Read the contents of the file "applications.yaml"
  fileContents, err := ioutil.ReadFile("applications.yaml")
  if err != nil {
    log.Fatalf("Error reading file: %v", err)
  }

  // Unmarshal the file contents into a map
  var applications map[string]map[string]int
  if err := yaml.Unmarshal(fileContents, &applications); err != nil {
    log.Fatalf("Error unmarshalling file contents: %v", err)
  }

  // Iterate over the map
  for _, appInfo := range applications {
    appId := appInfo["appId"]

    // Make an HTTP request to the specified URL
    response, err := http.Get(fmt.Sprintf("https://api.steamcmd.net/v1/info/%d", appId))
    if err != nil {
      log.Fatalf("Error making HTTP request: %v", err)
    }
    defer response.Body.Close()

    // Read the response body
    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
      log.Fatalf("Error reading response body: %v", err)
    }

    // Get the WEBHOOK_URL environment variable
    webhookURL := os.Getenv("WEBHOOK_URL")
    if webhookURL == "" {
      log.Fatalf("WEBHOOK_URL environment variable is not set")
    }

    // Get the WEBHOOK_TOKEN environment variable
    webhookToken := os.Getenv("WEBHOOK_TOKEN")
    if webhookToken == "" {
      log.Fatalf("WEBHOOK_TOKEN environment variable is not set")
    }

    // Make an HTTP POST request to the webhook URL
    req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(body))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", webhookToken))
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
      log.Fatalf("Error posting to webhook: %v", err)
    }

    // Read the response from the POST request
    postResponse, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      log.Fatalf("Error reading response from POST request: %v", err)
    }

    // Print the POST request and the response
    log.Printf("POST Request: %v\nResponse: %s\n", req, postResponse)
  }
}
