package model

type ShortURLRequest struct {
	URL string `json:"url" binding:"required"`
}
