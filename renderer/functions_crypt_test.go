package renderer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithCryptFunctions(t *testing.T) {
	functions := New(WithCryptFunctions()).Configuration().ExtraFunctions

	for _, name := range []string{
		"encryptAWS",
		"decryptAWS",
		"encryptGCP",
		"decryptGCP",
		"encryptAzure",
		"decryptAzure",
		"encryptOCI",
		"decryptOCI",
	} {
		t.Run(name, func(t *testing.T) {
			function, ok := functions[name]
			assert.True(t, ok)
			assert.NotNil(t, function)
		})
	}
}
