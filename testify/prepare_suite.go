package testifyrelated

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

func PrepareSuite[T suite.TestingSuite](t *testing.T, ctx context.Context, fn func(ctx context.Context) (T, error)) T {
	t.Helper()
	v, err := fn(ctx)
	if err != nil {
		var zero T
		t.Fatal(fmt.Errorf("failed to prepare %T: %w", zero, err))
		return zero
	}
	return v
}
