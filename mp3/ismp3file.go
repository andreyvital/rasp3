package mp3

import "strings"

func IsMp3File(file string) bool {
	return strings.HasSuffix(file, ".mp3") || strings.HasSuffix(file, ".MP3")
}
