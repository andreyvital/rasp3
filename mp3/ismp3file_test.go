package mp3_test

import (
	"testing"

	"github.com/CentaurWarchief/rasp3/mp3"
	"github.com/stretchr/testify/assert"
)

func TestIsMp3File(t *testing.T) {
	assert.True(t, mp3.IsMp3File("In Any Tongue.mp3"))
	assert.True(t, mp3.IsMp3File("In Any Tongue.MP3"))
	assert.False(t, mp3.IsMp3File("In Any Tongue.wav"))
	assert.False(t, mp3.IsMp3File("In Any Tongue.flac"))
}
