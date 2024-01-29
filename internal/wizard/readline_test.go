package wizard

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestReadLine(t *testing.T) {
	testCases := []struct {
		input  string
		output string
	}{
		{input: "123", output: "123"},
		{input: "123\n456", output: "123"},
		{input: "", output: ""},
		{input: "\n123", output: ""},
		{input: " 1 2 3 \n 4 5 6", output: " 1 2 3 "},
		{input: "1号\n 4 5 6", output: "1号"},
	}

	for _, tc := range testCases {
		output, err := readLine(strings.NewReader(tc.input))
		require.NoError(t, err)
		require.Equal(t, tc.output, output)
	}
}
