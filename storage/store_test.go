package storage

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSaveShortURL_ShouldSuccessfullySaveShortURL(t *testing.T) {
	s := Init()
	err := s.SaveShortURL("short-url", "full-url")

	store := s.(*store)
	_, found := store.db["short-url"]

	assert.Nil(t, err)
	assert.True(t, found)
	assert.Equal(t, store.db["short-url"], "full-url")
}

func TestFindFullURLShouldReturnFullURLIfPresent(t *testing.T) {
	s := Init()
	store := s.(*store)
	store.db["short-url"] = "full-url"
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
	store.visitorMap["full-url"] = visitor{
		visitorCount: 0,
		shortUrl:     "short-url",
	}
	shortURL, found := s.IsShortURLExists("full-url")

	assert.True(t, found)
	assert.Equal(t, "short-url", shortURL)

}

func TestIsShortURLExistsShouldReturnFalseIfURLExists(t *testing.T) {
	s := Init()
	shortURL, found := s.IsShortURLExists("full-url")

	assert.False(t, found)
	assert.Equal(t, "", shortURL)

}
