package mp3

type NameValueTag struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func NewNamedTag(name, value string) NameValueTag {
	return NameValueTag{
		Name:  name,
		Value: value,
	}
}
