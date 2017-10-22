// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ArchiveRequest archive request
// swagger:model ArchiveRequest

type ArchiveRequest struct {

	// url
	// Required: true
	// Min Length: 1
	URL *string `json:"url"`
}

/* polymorph ArchiveRequest url false */

// Validate validates this archive request
func (m *ArchiveRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateURL(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ArchiveRequest) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	if err := validate.MinLength("url", "body", string(*m.URL), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ArchiveRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ArchiveRequest) UnmarshalBinary(b []byte) error {
	var res ArchiveRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}