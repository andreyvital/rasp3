package gql

type ArtworkReference struct {
	URL    string `json:"url"`
	MIME   string `json:"mime"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}
