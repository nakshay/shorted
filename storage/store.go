package storage

import (
	"shorted/model"
	"sort"
	"sync"
)

//go:generate mockgen -source=./store.go -destination=../mocks/mock_store.go -package=mocks

type Store interface {
	SaveShortURL(shortURL string, fullURL string)
	FindFullURL(shortURL string) (string, bool)
	IsShortURLExistsForFullURL(fullURL string) (string, bool)
	UpdateMetricsForDomain(domain string)
	GetMetricsForTopDomain(topNDomains int) model.MetricsResponse
}

type store struct {
	shortToFullMap     map[string]string
	shortToFullRWMutex sync.RWMutex
	fullToShortMap     map[string]string
	fullToShortRMutex  sync.RWMutex
	metricsMap         map[string]int
	metricsRWMutex     sync.RWMutex
}

func Init() Store {
	return &store{
		shortToFullMap: map[string]string{},
		fullToShortMap: map[string]string{},
		metricsMap:     make(map[string]int),
	}
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

func (s *store) UpdateMetricsForDomain(domain string) {
	s.metricsRWMutex.Lock()
	defer s.metricsRWMutex.Unlock()
	s.metricsMap[domain]++
}

func (s *store) GetMetricsForTopDomain(topNDomains int) model.MetricsResponse {
	var topHits []model.TopHit

	s.metricsRWMutex.RLock()
	defer s.metricsRWMutex.RUnlock()

	if len(s.metricsMap) < 1 {
		return model.MetricsResponse{TopHits: make([]model.TopHit, 0)}
	}

	for domain, count := range s.metricsMap {
		topHits = append(topHits, model.TopHit{URL: domain, Hits: count})
	}
	sort.Slice(topHits, func(i, j int) bool {
		return topHits[i].Hits > topHits[j].Hits
	})

	if len(topHits) < topNDomains {
		return model.MetricsResponse{
			TopHits: topHits,
		}
	}
	topHits = topHits[:topNDomains]
	return model.MetricsResponse{
		TopHits: topHits,
	}
}
