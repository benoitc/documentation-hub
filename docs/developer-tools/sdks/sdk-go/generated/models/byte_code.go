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

// ByteCode byte code
// swagger:model ByteCode
type ByteCode struct {

	// bytecode
	// Required: true
	Bytecode *string `json:"bytecode"`
}

// Validate validates this byte code
func (m *ByteCode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBytecode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ByteCode) validateBytecode(formats strfmt.Registry) error {

	if err := validate.Required("bytecode", "body", m.Bytecode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ByteCode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ByteCode) UnmarshalBinary(b []byte) error {
	var res ByteCode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
