package titanclient

import (
	"github.com/stretchr/testify/mock"
	"github.com/t11e/go-discoveryclient"
)

type MockClient struct {
	mock.Mock
}

func (c *MockClient) Query(dataset string, q *discoveryclient.Query) (*discoveryclient.Results, error) {
	args := c.Called(dataset, q)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*discoveryclient.Results), nil
}
