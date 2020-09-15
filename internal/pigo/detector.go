package pigo

import (
	"fmt"
	"io"
	"io/ioutil"

	// Supported image formats for the image decoder
	_ "image/jpeg"
	_ "image/png"

	pigo "github.com/esimov/pigo/core"
	"github.com/regeda/faced/models"
)

// Magic constants given from the pigo library.
const (
	angle        = 0.0
	iouThreshold = 0.2
	minSize      = 20
	maxSize      = 1000
	minScale     = 50
	shiftFactor  = 0.1
	scaleFactor  = 1.1
	qThresh      = 5.0
	perturb      = 63
)

var mouthCascade = [...]string{"lp93", "lp84", "lp82", "lp81"}

// Detector gives a simple API on top of pigo library.
type Detector struct {
	classifier *pigo.Pigo
	plc        *pigo.PuplocCascade
	flpcs      map[string][]*pigo.FlpCascade
}

// New creates a new Detector.
func New(cascadeFile, puplocFile, flplocDir string) (*Detector, error) {
	cascadeBytes, err := ioutil.ReadFile(cascadeFile)
	if err != nil {
		return nil, fmt.Errorf("[pigo] failed to read the cascade file: %w", err)
	}
	p := pigo.NewPigo()
	classifier, err := p.Unpack(cascadeBytes)
	if err != nil {
		return nil, fmt.Errorf("[pigo] failed to unpack the cascade file: %w", err)
	}

	pl := pigo.NewPuplocCascade()
	puplocBytes, err := ioutil.ReadFile(puplocFile)
	if err != nil {
		return nil, fmt.Errorf("[pigo] failed to read the puploc file: %w", err)
	}
	plc, err := pl.UnpackCascade(puplocBytes)
	if err != nil {
		return nil, fmt.Errorf("[pigo] failed to unpack the puploc file: %w", err)
	}
	flpcs, err := pl.ReadCascadeDir(flplocDir)
	if err != nil {
		return nil, fmt.Errorf("[pigo] failed to read the cascade dir: %w", err)
	}

	return &Detector{
		classifier: classifier,
		plc:        plc,
		flpcs:      flpcs,
	}, nil
}

// Detect predicts faces from a reader.
func (d *Detector) Detect(r io.Reader) ([]*models.Face, error) {
	src, err := pigo.DecodeImage(r)
	if err != nil {
		return nil, fmt.Errorf("[pigo] failed to decode an image: %w", err)
	}

	pixels := pigo.RgbToGrayscale(src)
	cols, rows := src.Bounds().Max.X, src.Bounds().Max.Y

	params := pigo.CascadeParams{
		MinSize:     minSize,
		MaxSize:     maxSize,
		ShiftFactor: shiftFactor,
		ScaleFactor: scaleFactor,
		ImageParams: pigo.ImageParams{
			Pixels: pixels,
			Rows:   rows,
			Cols:   cols,
			Dim:    cols,
		},
	}

	dets := d.classifier.ClusterDetections(d.classifier.RunCascade(params, angle), iouThreshold)

	//@TODO make a pool for memory usage optimization
	faces := make([]*models.Face, 0, len(dets))

	for _, det := range dets {
		if det.Q < qThresh {
			continue
		}
		face := models.Face{
			Bounds: &models.Rect{
				X:      int64(det.Col - det.Scale/2),
				Y:      int64(det.Row - det.Scale/2),
				Width:  int64(det.Scale),
				Height: int64(det.Scale),
			},
		}
		if det.Scale > minScale {
			leftPuploc := pigo.Puploc{
				Row:      det.Row - int(0.075*float32(det.Scale)),
				Col:      det.Col - int(0.175*float32(det.Scale)),
				Scale:    float32(det.Scale) * 0.25,
				Perturbs: perturb,
			}
			leftEye := d.plc.RunDetector(leftPuploc, params.ImageParams, angle, false)
			face.LeftEye = &models.Point{
				X: int64(leftEye.Col),
				Y: int64(leftEye.Row),
			}
			rightPuploc := pigo.Puploc{
				Row:      det.Row - int(0.075*float32(det.Scale)),
				Col:      det.Col + int(0.185*float32(det.Scale)),
				Scale:    float32(det.Scale) * 0.25,
				Perturbs: perturb,
			}
			rightEye := d.plc.RunDetector(rightPuploc, params.ImageParams, angle, false)
			face.RightEye = &models.Point{
				X: int64(rightEye.Col),
				Y: int64(rightEye.Row),
			}
			for _, mouth := range mouthCascade {
				for _, flpc := range d.flpcs[mouth] {
					flp := flpc.FindLandmarkPoints(leftEye, rightEye, params.ImageParams, perturb, false)
					face.Mouth = &models.Point{
						X: int64(flp.Col),
						Y: int64(flp.Row),
					}
				}
			}
		}
		faces = append(faces, &face)
	}

	return faces, nil
}
