package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/thegrandpackard/PackardQuest/models"
)

type Client interface {
	GetPlayer(id int) (*models.Player, error)
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

func (c *client) GetPlayer(id int) (*models.Player, error) {
	resp, err := c.httpClient.Get(fmt.Sprintf("%s/api/%s/player/%d", c.host, c.apiVersion, id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	player := &models.Player{}
	err = json.Unmarshal(bodyBytes, &player)
	if err != nil {
		return nil, err
	}

	return player, nil
}
