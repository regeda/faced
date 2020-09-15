package handlers_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/golang/mock/gomock"
	"github.com/regeda/faced/handlers"
	"github.com/regeda/faced/models"
	"github.com/regeda/faced/restapi/operations"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFacesHandle(t *testing.T) {
	fakeURL := "http://localhost"
	fakeReq, err := http.NewRequest("GET", fakeURL, nil)
	require.NoError(t, err)
	params := operations.GetFacesParams{URL: strfmt.URI(fakeURL)}

	t.Run("should_detect_successfully", func(t *testing.T) {
		faces := []*models.Face{{
			Bounds: &models.Rect{X: 1, Y: 1},
		}}

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		rc := NewMockReadCloser(ctrl)
		rc.
			EXPECT().
			Close().
			Return(nil)

		rt := NewMockRoundTripper(ctrl)
		rt.
			EXPECT().
			RoundTrip(gomock.Eq(fakeReq)).
			Return(&http.Response{
				StatusCode: http.StatusOK,
				Body:       rc,
			}, nil)

		d := NewMockDetector(ctrl)
		d.
			EXPECT().
			Detect(gomock.Eq(rc)).
			Return(faces, nil)

		h := handlers.NewFaces(d, handlers.FacesWithHTTPClient(&http.Client{
			Transport: rt,
		}))

		resp := h.Handle(params)
		require.IsType(t, (*operations.GetFacesOK)(nil), resp)
		require.Equal(t, faces, resp.(*operations.GetFacesOK).Payload.Faces)
	})

	t.Run("empty_result", func(t *testing.T) {
		faces := []*models.Face{}

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		rc := NewMockReadCloser(ctrl)
		rc.
			EXPECT().
			Close().
			Return(nil)

		rt := NewMockRoundTripper(ctrl)
		rt.
			EXPECT().
			RoundTrip(gomock.Eq(fakeReq)).
			Return(&http.Response{
				StatusCode: http.StatusOK,
				Body:       rc,
			}, nil)

		d := NewMockDetector(ctrl)
		d.
			EXPECT().
			Detect(gomock.Eq(rc)).
			Return(faces, nil)

		h := handlers.NewFaces(d, handlers.FacesWithHTTPClient(&http.Client{
			Transport: rt,
		}))

		resp := h.Handle(params)
		require.IsType(t, (*operations.GetFacesDefault)(nil), resp)
		assert.EqualValues(t, http.StatusNotFound, resp.(*operations.GetFacesDefault).Payload.Code)
		assert.EqualValues(t, "No faces found. Please load a big picture.", *resp.(*operations.GetFacesDefault).Payload.Message)
	})

	t.Run("http_client_fails", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		rt := NewMockRoundTripper(ctrl)
		rt.
			EXPECT().
			RoundTrip(gomock.Eq(fakeReq)).
			Return(nil, http.ErrHijacked)

		d := NewMockDetector(ctrl)

		h := handlers.NewFaces(d, handlers.FacesWithHTTPClient(&http.Client{
			Transport: rt,
		}))

		resp := h.Handle(params)
		require.IsType(t, (*operations.GetFacesDefault)(nil), resp)
		assert.EqualValues(t, http.StatusBadRequest, resp.(*operations.GetFacesDefault).Payload.Code)
		assert.EqualValues(t, `Get "http://localhost": http: connection has been hijacked`, *resp.(*operations.GetFacesDefault).Payload.Message)
	})

	t.Run("detector_fails", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		rc := NewMockReadCloser(ctrl)
		rc.
			EXPECT().
			Close().
			Return(nil)

		rt := NewMockRoundTripper(ctrl)
		rt.
			EXPECT().
			RoundTrip(gomock.Eq(fakeReq)).
			Return(&http.Response{
				StatusCode: http.StatusOK,
				Body:       rc,
			}, nil)

		err := errors.New("oops")

		d := NewMockDetector(ctrl)
		d.
			EXPECT().
			Detect(gomock.Eq(rc)).
			Return(nil, err)

		h := handlers.NewFaces(d, handlers.FacesWithHTTPClient(&http.Client{
			Transport: rt,
		}))

		resp := h.Handle(params)
		require.IsType(t, (*operations.GetFacesDefault)(nil), resp)
		assert.EqualValues(t, http.StatusBadRequest, resp.(*operations.GetFacesDefault).Payload.Code)
		assert.EqualValues(t, err.Error(), *resp.(*operations.GetFacesDefault).Payload.Message)
	})
}
