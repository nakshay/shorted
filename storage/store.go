package storage

import (
	"sync"
)

//go:generate mockgen -source=./store.go -destination=../mocks/mock_store.go -package=mocks

type Store interface {
	SaveShortURL(shortURL string, fullURL string)
	FindFullURL(shortURL string) (string, bool)
	IsShortURLExistsForFullURL(fullURL string) (string, bool)
}

type store struct {
	shortToFullMap     map[string]string
	shortToFullRWMutex sync.RWMutex
	fullToShortMap     map[string]string
	fullToShortRMutex  sync.RWMutex
}

func Init() Store {
	return &store{shortToFullMap: map[string]string{}, fullToShortMap: map[string]string{}}
}

func (s *store) SaveShortURL(shortURL string, fullURL string) {
	s.mapShortToFullURL(shortURL, fullURL)
	s.mapFullToShortURL(fullURL, shortURL)
}

func (s *store) mapShortToFullURL(shortURL string, fullURL string) {
	s.shortToFullRWMutex.Lock()
	defer s.shortToFullRWMutex.Unlock()
	s.shortToFullMap[shortURL] = fullURL
}
func (s *store) mapFullToShortURL(fullURL string, shortURL string) {
	s.fullToShortRMutex.Lock()
	defer s.fullToShortRMutex.Unlock()
	s.fullToShortMap[fullURL] = shortURL
}

func (s *store) FindFullURL(shortURL string) (string, bool) {
	s.shortToFullRWMutex.RLock()
	defer s.shortToFullRWMutex.RUnlock()
	fullURL, found := s.shortToFullMap[shortURL]
	return fullURL, found
}

func (s *store) IsShortURLExistsForFullURL(fullURL string) (string, bool) {
	s.fullToShortRMutex.RLock()
	defer s.fullToShortRMutex.RUnlock()
	shortURL, found := s.fullToShortMap[fullURL]
	return shortURL, found
}
