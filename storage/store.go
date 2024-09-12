package storage

import (
	"sync"
)

type Store interface {
	SaveShortURL(key string, value string) error
	FindFullURL(key string) (string, bool)
	IsShortURLExists(key string) (string, bool)
}

type store struct {
	db           map[string]string
	visitorMap   map[string]visitor
	dbMutex      sync.RWMutex
	visitorMutex sync.RWMutex
}

type visitor struct {
	visitorCount int
	shortUrl     string
}

func Init() Store {
	return &store{db: map[string]string{}, visitorMap: map[string]visitor{}}
}

func (s *store) SaveShortURL(shortURL string, fullURL string) error {
	s.dbMutex.Lock()
	s.db[shortURL] = fullURL
	s.dbMutex.Unlock()

	s.addVisitor(fullURL, shortURL)
	return nil
}

func (s *store) FindFullURL(shortURL string) (string, bool) {
	s.dbMutex.RLock()
	fullURL, found := s.db[shortURL]
	s.dbMutex.RUnlock()
	if !found {
		return "", false
	}
	s.updateVisit(fullURL)
	return fullURL, true
}

func (s *store) updateVisit(fullURL string) {
	s.visitorMutex.Lock()
	defer s.visitorMutex.Unlock()
	v, _ := s.visitorMap[fullURL]
	v.visitorCount++
	s.visitorMap[fullURL] = v
}

func (s *store) IsShortURLExists(fullURL string) (string, bool) {
	s.visitorMutex.RLock()
	defer s.visitorMutex.RUnlock()
	v, found := s.visitorMap[fullURL]
	return v.shortUrl, found
}

func (s *store) addVisitor(fullURL string, shortURL string) {
	s.visitorMutex.Lock()
	defer s.visitorMutex.Unlock()
	s.visitorMap[fullURL] = visitor{0, shortURL}
}
