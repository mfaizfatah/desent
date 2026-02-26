package repository

import (
	"sync"
	"time"
)

type MemoryTokenRepository struct {
	mu     sync.RWMutex
	tokens map[string]time.Time
	ttl    time.Duration
}

func NewMemoryTokenRepository(ttl time.Duration) *MemoryTokenRepository {
	return &MemoryTokenRepository{
		tokens: make(map[string]time.Time),
		ttl:    ttl,
	}
}

func (r *MemoryTokenRepository) Store(token string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.tokens[token] = time.Now().Add(r.ttl)
}

func (r *MemoryTokenRepository) Generate() string {
	return ""
}

func (r *MemoryTokenRepository) Exists(token string) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()
	exp, ok := r.tokens[token]
	if !ok {
		return false
	}
	return time.Now().Before(exp)
}
