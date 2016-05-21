package titanclient

import (
	"bytes"
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/t11e/go-discoveryclient"
	pc "github.com/t11e/go-pebbleclient"
)

//go:generate go run vendor/github.com/vektra/mockery/cmd/mockery/mockery.go -name=Client -case=underscore

type Client interface {
	Query(dataset string, query *discoveryclient.Query) (*discoveryclient.Results, error)
}

type client struct {
	c pc.Client
}

// Register registers us in a connector.
func Register(connector *pc.Connector) {
	connector.Register((*Client)(nil), func(client pc.Client) (pc.Service, error) {
		return New(client)
	})
}

func New(pebbleClient pc.Client) (Client, error) {
	return &client{pebbleClient.WithOptions(pc.Options{
		ServiceName: "titan",
		ApiVersion:  1,
	})}, nil
}

func (c *client) Query(dataset string, query *discoveryclient.Query) (*discoveryclient.Results, error) {
	b, err := json.Marshal(query)
	if err != nil {
		return nil, errors.Wrap(err, "Could not marshal query")
	}

	var out discoveryclient.Results
	err = c.c.Post("/query/:dataset", &pc.RequestOptions{
		Params: pc.Params{
			"dataset": dataset,
		},
	}, bytes.NewReader(b), &out)
	if err != nil {
		return nil, err
	}
	return &out, err
}
