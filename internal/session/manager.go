package session

import (
	"math/rand/v2"
	"net"
	"sync"
)

type Manager struct {
	mu       sync.RWMutex
	sessions map[uint32]*Session
}

func NewManager() *Manager {
	return &Manager{
		sessions: make(map[uint32]*Session),
	}
}

func (m *Manager) Create(id uint32, udpAddr *net.UDPAddr) *Session {

	m.mu.Lock()
	defer m.mu.Unlock()

	s := Session{
		ID:         id,
		RemoteAddr: udpAddr,
	}

	m.sessions[id] = &s
	return &s

}

func (m *Manager) Get(id uint32) (*Session, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	s, ok := m.sessions[id]
	return s, ok
}

func GenerateSessionID() uint32 {
	return rand.Uint32()
}
