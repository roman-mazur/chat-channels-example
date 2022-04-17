package sdk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	BaseUrl string
}

type Channel struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func (ch Channel) String() string {
	return ch.Name
}

func (c *Client) ListChannels(ctx context.Context) ([]Channel, error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodGet, c.BaseUrl+"/channels", nil)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	var res []Channel
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) CreateChannel(ctx context.Context, name string) error {
	var p Channel
	p.Name = name
	bb := new(bytes.Buffer)
	err := json.NewEncoder(bb).Encode(&p)
	if err != nil {
		panic(err)
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.BaseUrl+"/channels", bb)
	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code %d", resp.StatusCode)
	}
	return nil
}
