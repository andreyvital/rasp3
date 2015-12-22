package mp3

type ID3 struct {
	Title   string `json:"title"`
	Artist  string `json:"artist"`
	Album   string `json:"album"`
	Year    int    `json:"year"`
	Genre   string `json:"genre"`
	Track   int    `json:"track"`
	Picture struct {
		Encoded string
		Mime    string
	} `json:"-"`
}
