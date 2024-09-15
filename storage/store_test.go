package storage

import (
	"github.com/stretchr/testify/assert"
	"shorted/model"
	"testing"
)

func TestSaveShortURL_ShouldSuccessfullySaveShortURL(t *testing.T) {
	s := Init()
	s.SaveShortURL("short-url", "full-url")

	store := s.(*store)
	_, found := store.shortToFullMap["short-url"]

	assert.True(t, found)
	assert.Equal(t, store.shortToFullMap["short-url"], "full-url")
}

func TestSaveShortURL_ShouldNotStoreSameURLAgain(t *testing.T) {
	s := Init()
	s.SaveShortURL("short-url", "full-url")
	s.SaveShortURL("short-url", "full-url")

	store := s.(*store)
	_, found := store.shortToFullMap["short-url"]

	assert.True(t, found)
	assert.Equal(t, store.shortToFullMap["short-url"], "full-url")
	assert.Equal(t, 1, len(store.shortToFullMap))
}

func TestFindFullURLShouldReturnFullURLIfPresent(t *testing.T) {
	s := Init()
	store := s.(*store)
	store.shortToFullMap["short-url"] = "full-url"
	fullURL, found := s.FindFullURL("short-url")
	assert.True(t, found)
	assert.Equal(t, "full-url", fullURL)
}

func TestFindFullURLShouldReturnFalseWhenURLNotPresent(t *testing.T) {
	s := Init()
	fullURL, found := s.FindFullURL("short-url")
	assert.False(t, found)
	assert.Equal(t, "", fullURL)
}

func TestIsShortURLExistsShouldReturnTrueIfURLExists(t *testing.T) {
	s := Init()
	store := s.(*store)
	store.fullToShortMap["full-url"] = "short-url"
	shortURL, found := s.IsShortURLExistsForFullURL("full-url")

	assert.True(t, found)
	assert.Equal(t, "short-url", shortURL)

}

func TestIsShortURLExistsShouldReturnFalseIfURLExists(t *testing.T) {
	s := Init()
	shortURL, found := s.IsShortURLExistsForFullURL("full-url")

	assert.False(t, found)
	assert.Equal(t, "", shortURL)

}

func TestUpdateMetricsForDomainShouldUpdateDomainMetricsToOneIfCalledFirstTime(t *testing.T) {
	s := Init()
	s.UpdateMetricsForDomain("domain.com")
	store := s.(*store)
	assert.Equal(t, store.metricsMap["domain.com"], 1)

}

func TestUpdateMetricsForDomainShouldUpdateDomainMetricsToTwoIfCalledSecondTimeTime(t *testing.T) {
	s := Init()
	s.UpdateMetricsForDomain("domain.com")
	s.UpdateMetricsForDomain("domain.com")
	store := s.(*store)
	assert.Equal(t, store.metricsMap["domain.com"], 2)

}

func TestGetMetricsForTopDomainShouldReturnMetrics(t *testing.T) {
	s := Init()
	expectedResponse := model.MetricsResponse{TopHits: []model.TopHit{{
		URL:  "domain.com",
		Hits: 2,
	}}}
	s.UpdateMetricsForDomain("domain.com")
	s.UpdateMetricsForDomain("domain.com")
	response := s.GetMetricsForTopDomain(1)
	assert.Equal(t, expectedResponse, response)
}

func TestGetMetricsForTopDomainShouldReturnMaxNMetricsIfNIsPassed(t *testing.T) {
	s := Init()
	expectedResponse := model.MetricsResponse{TopHits: []model.TopHit{{
		URL:  "domain2.com",
		Hits: 3,
	}, {
		URL:  "domain1.com",
		Hits: 1,
	}}}
	s.UpdateMetricsForDomain("domain1.com")
	s.UpdateMetricsForDomain("domain2.com")
	s.UpdateMetricsForDomain("domain2.com")
	s.UpdateMetricsForDomain("domain2.com")

	response := s.GetMetricsForTopDomain(5)
	assert.Equal(t, expectedResponse, response)
}

func TestGetMetricsForTopDomainShouldReturnAllMetricsIfRecordsAreLessThanN(t *testing.T) {
	s := Init()
	expectedResponse := model.MetricsResponse{TopHits: []model.TopHit{{
		URL:  "domain4.com",
		Hits: 4,
	}, {
		URL:  "domain2.com",
		Hits: 3,
	}}}
	s.UpdateMetricsForDomain("domain1.com")
	s.UpdateMetricsForDomain("domain2.com")
	s.UpdateMetricsForDomain("domain2.com")
	s.UpdateMetricsForDomain("domain2.com")
	s.UpdateMetricsForDomain("domain3.com")
	s.UpdateMetricsForDomain("domain3.com")
	s.UpdateMetricsForDomain("domain4.com")
	s.UpdateMetricsForDomain("domain4.com")
	s.UpdateMetricsForDomain("domain4.com")
	s.UpdateMetricsForDomain("domain4.com")

	response := s.GetMetricsForTopDomain(2)
	assert.Equal(t, expectedResponse, response)
}
