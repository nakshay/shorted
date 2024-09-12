package store

import "sync"

type Store interface {
	Save(key string, value string) error
	Read(key string) (string, bool)
	isExist(key string) bool
}

type store struct {
	db      map[string]string
	visitor map[string]int
	mutex   sync.RWMutex
}

func Init() Store {
	return &store{db: map[string]string{}, visitor: map[string]int{}}
}

func (s *store) Save(key string, value string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.db[key] = value
	s.visitor[value] = 0
	return nil
}

func (s *store) Read(key string) (string, bool) {
	s.mutex.RLock()
	s.mutex.RUnlock()
	value, found := s.db[key]
	if !found {
		return "", false
	}
	return value, true
}
func (s *store) isExist(key string) bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	_, found := s.visitor[key]
	return found
}
