package storage

import (
	"github.com/stretchr/testify/assert"
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
