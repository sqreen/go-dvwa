// Copyright (c) 2016 - 2020 Sqreen. All Rights Reserved.
// Please refer to our terms for more information:
// https://www.sqreen.io/terms.html

package vulnerable

import "os"

func Open(filepath string) (*os.File, error) {
	// Nothing special is needed to make Open vulnerable to local file inclusion
	// (LFi). LFi is actually possible when the filepath is not properly
	// restricted.
	return os.Open(filepath)
}
