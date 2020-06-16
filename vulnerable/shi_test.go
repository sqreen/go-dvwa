package vulnerable_test

import (
	"context"
	"testing"

	"github.com/sqreen/go-dvwa/vulnerable"
	"github.com/stretchr/testify/require"
)

func TestShellInjection(t *testing.T) {

	// Use the vulnerable System() to add our attack
	output, err := vulnerable.System(context.Background(), "echo ok"+"; echo vulnerable")
	require.NoError(t, err)
	require.Equal(t, "ok\nvulnerable\n", string(output))
}
