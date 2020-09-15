package pigo_test

import (
	"os"
	"testing"

	"github.com/regeda/faced/internal/pigo"
	"github.com/regeda/faced/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDetector(t *testing.T) {
	d, err := pigo.New("cascade/facefinder", "cascade/puploc", "cascade/lps")

	require.NoError(t, err)

	f, err := os.Open("testdata/fixtures/katy-perry.jpg")

	require.NoError(t, err)

	faces, err := d.Detect(f)

	require.NoError(t, err)
	require.Len(t, faces, 1)

	assert.Equal(t, models.Rect{
		X:      81,
		Y:      84,
		Height: 265,
		Width:  265,
	}, *faces[0].Bounds)

	assert.Equal(t, models.Point{X: 160, Y: 197}, *faces[0].LeftEye)
	assert.Equal(t, models.Point{X: 263, Y: 199}, *faces[0].RightEye)
	assert.Equal(t, models.Point{X: 214, Y: 290}, *faces[0].Mouth)
}
