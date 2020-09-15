package handlers

import (
	"errors"
	"io"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/regeda/faced/models"
	"github.com/regeda/faced/restapi/operations"
)

var (
	errNoFacesFound = errors.New("No faces found. Please load a big picture.")
)

type Detector interface {
	Detect(io.Reader) ([]*models.Face, error)
}

type Faces struct {
	d Detector
	c *http.Client
}

type FacesOpt func(*Faces)

func FacesWithHTTPClient(c *http.Client) FacesOpt {
	return func(f *Faces) {
		f.c = c
	}
}

func NewFaces(d Detector, opts ...FacesOpt) *Faces {
	f := Faces{
		d: d,
	}
	for _, opt := range opts {
		opt(&f)
	}
	if f.c == nil {
		f.c = http.DefaultClient
	}
	return &f
}

func (f *Faces) Handle(params operations.GetFacesParams) middleware.Responder {
	resp, err := f.c.Get(params.URL.String())
	if err != nil {
		return errFacesToResponse(http.StatusBadRequest, err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	faces, err := f.d.Detect(resp.Body)
	if err != nil {
		return errFacesToResponse(http.StatusBadRequest, err)
	}
	if len(faces) == 0 {
		return errFacesToResponse(http.StatusNotFound, errNoFacesFound)
	}
	return operations.NewGetFacesOK().WithPayload(&operations.GetFacesOKBody{
		Faces: faces,
	})
}

func errFacesToResponse(code int, err error) middleware.Responder {
	msg := err.Error()
	return operations.NewGetFacesDefault(code).WithPayload(&models.Error{
		Code:    int64(code),
		Message: &msg,
	})
}
