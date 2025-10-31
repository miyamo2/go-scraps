package testifyrelated

import "context"

type SomethingClient struct{}

func (c *SomethingClient) DoSomething(ctx context.Context) error {
	return nil
}

func NewSomethingClient(ctx context.Context) (*SomethingClient, error) {
	err, ok := ctx.Value(ErrKey{}).(error)
	if ok {
		return nil, err
	}
	return &SomethingClient{}, nil
}

type ErrKey struct{}
