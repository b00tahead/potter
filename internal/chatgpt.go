package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

var apiKeyFile = "./config/config.json"

type Config struct {
	APIKey string `json:"api_key"`
}

func SaveAPIKey(apiKey string) error {
	config := Config{
		APIKey: apiKey,
	}

	data, err := json.MarshalIndent(config, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile(apiKeyFile, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func GetAPIKey() (string, error) {
	data, err := os.ReadFile(apiKeyFile)
	if err != nil {
		return "", err
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return "", err
	}

	return config.APIKey, nil
}

func GetChatGPTResponse(input string) (string, error) {
	apiKey, err := GetAPIKey()
	if err != nil {
		return "", err
	}

	apiURL := "https://api.openai.com/v1/engines/davinci-codex/completions"
	payload := fmt.Sprintf(`{"prompt": "%s", "max_tokens": 50}`, input)

	req, err := http.NewRequest("POST",  apiURL, strings.NewReader(payload))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}

	if choices, ok := result["choices"].([]interface{}); ok && len(choices) > 0 {
		if choice, ok := choices[0].(map[string]interface{})["text"].(string); ok {
			return choice, nil
		}
	}

	return "", fmt.Errorf("Failed to get response from ChatGPT")
}
