// Copyright (c) 2016 - 2020 Sqreen. All Rights Reserved.
// Please refer to our terms for more information:
// https://www.sqreen.io/terms.html

package vulnerable_test

import (
	"io/ioutil"
	"testing"

	"github.com/sqreen/go-dvwa/vulnerable"
	"github.com/stretchr/testify/require"
)

func TestLocalFileInclusion(t *testing.T) {
	f, err := ioutil.TempFile("", "myfile")
	require.NoError(t, err)
	defer f.Close()

	// This isn't really a true LFi vulnerability - this is just opening a file
	// through a function accepting any filepath.
	opened, err := vulnerable.Open(f.Name())
	require.NoError(t, err)
	defer opened.Close()
}
