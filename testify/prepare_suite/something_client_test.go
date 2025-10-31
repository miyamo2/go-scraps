package prepare_suite

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type SomethingClientSuite struct {
	suite.Suite
	client *SomethingClient
}

func NewSomethingClientSuite(ctx context.Context) (*SomethingClientSuite, error) {
	client, err := NewSomethingClient(ctx)
	if err != nil {
		return nil, err
	}
	return &SomethingClientSuite{
		client: client,
	}, nil
}

// Here, PrepareSuite is used instead of SetupSuite.
//func (s *SomethingClientSuite) SetupSuite() {
//	client, err := NewSomethingClient(s.T().Context())
//	s.Require().NoError(err)
//	s.client = client
//}

func Test_SomethingClient(t *testing.T) {
	// instead of suite.Run(t, new(SomethingClientSuite))
	suite.Run(t, PrepareSuite[*SomethingClientSuite](t, t.Context(), NewSomethingClientSuite))
}

func Test_SomethingClient_With_Error_In_PrepareSuite(t *testing.T) {
	ctx := context.WithValue(context.Background(), ErrKey{}, fmt.Errorf("this error is expected"))
	suite.Run(t, PrepareSuite[*SomethingClientSuite](t, ctx, NewSomethingClientSuite))
	// Output: === RUN   Test_SomethingClient_With_Error_In_PrepareSuite
	//    something_client_test.go:40: failed to prepare *testifyrelated.SomethingClientSuite: this error is expected
	//--- FAIL: Test_SomethingClient_With_Error_In_PrepareSuite (0.00s)
}
