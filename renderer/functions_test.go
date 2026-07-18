package renderer

import (
	"bytes"
	"compress/gzip"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGzipProducesCompleteStream(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected []byte
	}{
		{name: "empty string", input: "", expected: []byte{}},
		{name: "text", input: "OCI KMS test", expected: []byte("OCI KMS test")},
		{name: "binary", input: []byte{0x00, 0x01, 0x02, 0xfe, 0xff}, expected: []byte{0x00, 0x01, 0x02, 0xfe, 0xff}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			compressed, err := Gzip(tt.input)
			require.NoError(t, err)

			reader, err := gzip.NewReader(bytes.NewReader([]byte(compressed)))
			require.NoError(t, err)
			actual, err := io.ReadAll(reader)
			require.NoError(t, err)
			require.NoError(t, reader.Close())

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestGzipUngzipRoundTrip(t *testing.T) {
	expected := "OCI KMS render gzip round trip"

	compressed, err := Gzip(expected)
	require.NoError(t, err)
	actual, err := Ungzip(compressed)

	require.NoError(t, err)
	assert.Equal(t, expected, actual)
}
