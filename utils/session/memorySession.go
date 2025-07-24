package utils

import "time"

type MemorySession struct {
	sid        string
	values     map[interface{}]interface{}
	lastAccess time.Time
}

func (s *MemorySession) Set(key, value interface{}) {
	(*s).values[key] = value
	(*s).lastAccess = time.Now()
}

func (s *MemorySession) Get(key interface{}) interface{} {
	(*s).lastAccess = time.Now()
	return (*s).values[key]
}

func (s *MemorySession) Delete(key interface{}) {
	delete((*s).values, key)
}

func (s *MemorySession) SessionID() string {
	return (*s).sid
}
