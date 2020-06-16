package vulnerable_test

import (
	"context"
	"testing"

	"github.com/sqreen/go-dvwa/vulnerable"
	"github.com/stretchr/testify/require"
)

func TestGetProducts(t *testing.T) {
	// Create a database having 10 sneakers and users
	nbEntries := 10
	db, err := vulnerable.PrepareSQLDB(nbEntries)
	require.NoError(t, err)

	// Normal query getting every sneakers
	products, err := vulnerable.GetProducts(context.Background(), db, "sneaker")
	require.NoError(t, err)
	require.Len(t, products, nbEntries)

	// Now leverage the vulnerability to also retrieve users
	products, err = vulnerable.GetProducts(context.Background(), db, "sneaker' UNION SELECT * FROM user WHERE ''='")
	require.NoError(t, err)
	require.True(t, len(products) > nbEntries)
}
