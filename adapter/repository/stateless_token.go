package repository

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"strconv"
	"strings"
	"time"
)

// StatelessTokenRepository validates tokens without server-side storage.
// Tokens are self-contained (expiry + HMAC) so they work across serverless instances.
type StatelessTokenRepository struct {
	secret []byte
	ttl    time.Duration
}

func NewStatelessTokenRepository(ttl time.Duration) *StatelessTokenRepository {
	// Fixed secret so tokens work across serverless instances
	secret := []byte("desent-challenge-secret-key-v1")
	return &StatelessTokenRepository{secret: secret, ttl: ttl}
}

func (r *StatelessTokenRepository) Store(token string) {
	// No-op: tokens are self-contained, no storage needed
}

func (r *StatelessTokenRepository) Exists(token string) bool {
	parts := strings.Split(token, ".")
	if len(parts) != 2 {
		return false
	}
	expiryB64, sigHex := parts[0], parts[1]
	expiryBytes, err := base64.RawURLEncoding.DecodeString(expiryB64)
	if err != nil {
		return false
	}
	expiry, err := strconv.ParseInt(string(expiryBytes), 10, 64)
	if err != nil {
		return false
	}
	if time.Now().Unix() > expiry {
		return false
	}
	mac := hmac.New(sha256.New, r.secret)
	mac.Write(expiryBytes)
	expected := hex.EncodeToString(mac.Sum(nil))
	return hmac.Equal([]byte(expected), []byte(sigHex))
}

func (r *StatelessTokenRepository) Generate() string {
	expiry := time.Now().Add(r.ttl).Unix()
	expiryStr := strconv.FormatInt(expiry, 10)
	mac := hmac.New(sha256.New, r.secret)
	mac.Write([]byte(expiryStr))
	sig := hex.EncodeToString(mac.Sum(nil))
	return base64.RawURLEncoding.EncodeToString([]byte(expiryStr)) + "." + sig
}
