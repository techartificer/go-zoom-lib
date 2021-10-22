package zoom

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	apiKey    = os.Getenv("ZOOM_API_KEY")
	apiSecret = os.Getenv("ZOOM_API_SECRET")
)

func TestSignatureGenerate(t *testing.T) {
	client := NewClient(apiKey, apiSecret)
	t.Run("hmac shoud be non empty", func(t *testing.T) {
		got := computeHmac256("12345", "secret")
		assert.NotEmpty(t, got, "Hmac can not be empty")
	})
	t.Run("signature shoud be non empty", func(t *testing.T) {
		got := client.GenerateSignature(12345, 1)
		assert.NotEmpty(t, got, "Signature can not be empty")
	})
}
