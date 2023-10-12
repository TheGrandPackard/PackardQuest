package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/thegrandpackard/PackardQuest/models"
)

type Client interface {
	GetPlayerByID(playerID int) (*models.Player, error)
	GetPlayerByWandID(wandID int) (*models.Player, error)
	UpdatePlayer(playerID int, request models.UpdatePlayerRequest) (*models.Player, error)
}

type client struct {
	host       string
	apiVersion string
	httpClient http.Client
}

func NewClient(host string) Client {
	return &client{
		host:       host,
		apiVersion: "latest",
		httpClient: *http.DefaultClient,
	}
}

func (c *client) GetPlayerByID(id int) (*models.Player, error) {
	resp, err := c.httpClient.Get(fmt.Sprintf("%s/api/%s/player/%d", c.host, c.apiVersion, id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	playerResponse := &models.PlayerResponse{}
	err = json.Unmarshal(bodyBytes, &playerResponse)
	if err != nil {
		return nil, err
	}

	return playerResponse.Player, nil
}

func (c *client) GetPlayerByWandID(wandID int) (*models.Player, error) {
	resp, err := c.httpClient.Get(fmt.Sprintf("%s/api/%s/player/wand/%d", c.host, c.apiVersion, wandID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	playerResponse := &models.PlayerResponse{}
	err = json.Unmarshal(bodyBytes, &playerResponse)
	if err != nil {
		return nil, err
	}

	return playerResponse.Player, nil
}

func (c *client) UpdatePlayer(playerID int, request models.UpdatePlayerRequest) (*models.Player, error) {
	reqBytes, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/api/%s/player/%d", c.host, c.apiVersion, playerID), bytes.NewBuffer(reqBytes))
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	playerResponse := &models.PlayerResponse{}
	err = json.Unmarshal(bodyBytes, &playerResponse)
	if err != nil {
		return nil, err
	}

	return playerResponse.Player, nil
}
