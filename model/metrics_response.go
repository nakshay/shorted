package model

type MetricsResponse struct {
	TopHits []TopHit `json:"top_hits"`
}

type TopHit struct {
	URL  string `json:"url"`
	Hits int    `json:"hits"`
}
