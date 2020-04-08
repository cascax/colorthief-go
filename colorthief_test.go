package colorthief

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"image"
	"image/color"
	"os"
	"testing"
)

func printColor(t *testing.T, colors []color.Color) {
	err := PrintColor(colors, "img/color.png")
	require.Nil(t, err)
}

func TestGetPaletteFromFile(t *testing.T) {
	colors, err := GetPaletteFromFile("img/image-1.png", 6)
	require.Nil(t, err)
	assert.Equal(t, 6, len(colors))
	t.Log(colors)
	printColor(t, colors)
}

func BenchmarkGetPaletteFromFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := GetPaletteFromFile("img/image-4.jpg", 5)
		assert.Nil(b, err)
	}
}

func BenchmarkGetColor(b *testing.B) {
	f, err := os.Open("img/image-4.jpg")
	require.Nil(b, err)

	defer f.Close()
	img, _, err := image.Decode(f)
	require.Nil(b, err)

	for i := 0; i < b.N; i++ {
		_, err = GetColor(img)
		assert.Nil(b, err)
	}
}
