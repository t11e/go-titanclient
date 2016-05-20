package titanclient

import (
	"bytes"
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/t11e/go-discoveryclient"
	pc "github.com/t11e/go-pebbleclient"
)

type Client struct {
	c pc.Client
}

func New(client pc.Client) (*Client, error) {
	return &Client{client.Options(pc.Options{
		ServiceName: "titan",
		ApiVersion:  1,
	})}, nil
}

func (client *Client) Query(dataset string, query *discoveryclient.Query) (*discoveryclient.Results, error) {
	b, err := json.Marshal(query)
	if err != nil {
		return nil, errors.Wrap(err, "Could not marshal query")
	}

	var out discoveryclient.Results
	err = client.c.Post("/query/:dataset", &pc.RequestOptions{
		Params: pc.Params{
			"dataset": dataset,
		},
	}, bytes.NewReader(b), &out)
	if err != nil {
		return nil, err
	}
	return &out, err
}
