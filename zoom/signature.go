package zoom

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"time"
)

func computeHmac256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func (c *Client) GenerateSignature(meetingNumber, role int64) string {
	timestamp := time.Now().UnixNano()/1e6 - 30000
	data := fmt.Sprintf("%s%d%d%d", client.apiKey, meetingNumber, timestamp, role)
	msg := base64.StdEncoding.EncodeToString([]byte(data))
	hash := computeHmac256(msg, client.secretKey)
	data = fmt.Sprintf("%s.%d.%d.%d.%s", client.apiKey, meetingNumber, timestamp, role, hash)
	return base64.StdEncoding.EncodeToString([]byte(data))
}
