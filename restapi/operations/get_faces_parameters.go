// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewGetFacesParams creates a new GetFacesParams object
// no default values defined in spec.
func NewGetFacesParams() GetFacesParams {

	return GetFacesParams{}
}

// GetFacesParams contains all the bound params for the get faces operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetFaces
type GetFacesParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: query
	*/
	URL strfmt.URI
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetFacesParams() beforehand.
func (o *GetFacesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qURL, qhkURL, _ := qs.GetOK("url")
	if err := o.bindURL(qURL, qhkURL, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindURL binds and validates parameter URL from query.
func (o *GetFacesParams) bindURL(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("url", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("url", "query", raw); err != nil {
		return err
	}

	// Format: uri
	value, err := formats.Parse("uri", raw)
	if err != nil {
		return errors.InvalidType("url", "query", "strfmt.URI", raw)
	}
	o.URL = *(value.(*strfmt.URI))

	if err := o.validateURL(formats); err != nil {
		return err
	}

	return nil
}

// validateURL carries on validations for parameter URL
func (o *GetFacesParams) validateURL(formats strfmt.Registry) error {

	if err := validate.FormatOf("url", "query", "uri", o.URL.String(), formats); err != nil {
		return err
	}
	return nil
}