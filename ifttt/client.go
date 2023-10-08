package ifttt

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const apiURL = "https://maker.ifttt.com"

type IFTTT interface {
	TriggerJSONWithKey(eventName string, json io.Reader) error
}

type client struct {
	apiKey     string
	httpClient *http.Client
}

func NewClient(apiKey string) (IFTTT, error) {
	if apiKey == "" {
		return nil, errors.New("empty apiKey")
	}

	return &client{
		apiKey:     apiKey,
		httpClient: &http.Client{},
	}, nil
}

func extractErrorMessages(resp *http.Response) string {
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err.Error()
	}

	respBody := &responseObject{}
	err = json.Unmarshal(bodyBytes, &respBody)
	if err != nil {
		return err.Error()
	}

	var errorStrings []string
	for _, err := range respBody.Errors {
		errorStrings = append(errorStrings, err.Message)
	}

	return strings.Join(errorStrings, ",")
}

func (c *client) TriggerJSONWithKey(eventName string, json io.Reader) error {
	resp, err := c.httpClient.Post(fmt.Sprintf("%s/trigger/wand_interaction/json/with/key/%s", apiURL, c.apiKey), "application/json", json)
	if err != nil {
		return err
	} else if resp.StatusCode != 200 {
		errorString := extractErrorMessages(resp)
		return fmt.Errorf("unexpected status code: %d with errors: %s", resp.StatusCode, errorString)
	}

	return nil
}
