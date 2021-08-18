package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	baseURL = "http://localhost:8761/settings?appID=%d"
)

type settingsResponse struct {
	AppID int    `json:"appID"`
	Name  string `json:"name"`
	Host  string `json:"host"`
	Port  string `json:"port"`
	Group string `json:"group,omitempty"`
}

func main() {

	client := http.Client{Timeout: 1 * time.Second}

	appIds := []int{100, 200, 300, 400, 500, 600, 700, 800, 900, 1000}

	allSettings := fetchSettings(appIds, client)
	log.Printf("[INFO] fetched settings amont: %d\n", len(allSettings))

}

func fetchSettings(appIds []int, client http.Client) []settingsResponse {
	var allSettings []settingsResponse
	for _, id := range appIds {
		response, err := client.Get(fmt.Sprintf(baseURL, id))
		if err != nil {
			log.Printf("[WARN] cant fetch apps settings for appID: %d. Error: %v", id, err)
		}
		var settings []settingsResponse
		err = json.NewDecoder(response.Body).Decode(&settings)
		if err != nil {
			log.Printf("[ERROR] cant fetch apps settings for appID: %d. Error: %v", id, err)
		}
		allSettings = append(allSettings, settings...)
	}
	return allSettings
}
