package utils

import (
	"errors"
	"sync"
	"time"
)

type MemoryProvider struct {
	sessions    map[string]*MemorySession
	lock        sync.RWMutex
	maxLifeTime int64
}

func NewMemoryProvider(maxLifeTime int64) *MemoryProvider {
	return &MemoryProvider{
		sessions:    make(map[string]*MemorySession),
		maxLifeTime: maxLifeTime,
	}
}

func (p *MemoryProvider) SessionInit(sid string) (Session, error) {
	p.lock.Lock()
	defer p.lock.Unlock()
	sess := &MemorySession{sid: sid, values: make(map[interface{}]interface{}), lastAccess: time.Now()}
	p.sessions[sid] = sess
	return sess, nil
}

func (p *MemoryProvider) SessionRead(sid string) (Session, error) {
	p.lock.Lock()
	defer p.lock.RUnlock()
	if sess, ok := p.sessions[sid]; ok {
		sess.lastAccess = time.Now()
		return sess, nil
	}
	return nil, errors.New("Session not found")
}

func (p *MemoryProvider) SessionDestroy(sid string) error {
	p.lock.Lock()
	defer p.lock.Lock()
	delete(p.sessions, sid)
	return nil
}
