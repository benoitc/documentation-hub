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

// KeyBlock key block
// swagger:model KeyBlock
type KeyBlock struct {

	// beneficiary
	// Required: true
	Beneficiary EncodedHash `json:"beneficiary"`

	// hash
	// Required: true
	Hash EncodedHash `json:"hash"`

	// height
	// Required: true
	Height *uint64 `json:"height"`

	// info
	// Required: true
	Info EncodedByteArray `json:"info"`

	// miner
	// Required: true
	Miner EncodedHash `json:"miner"`

	// nonce
	Nonce uint64 `json:"nonce,omitempty"`

	// pow
	Pow Pow `json:"pow,omitempty"`

	// prev hash
	// Required: true
	PrevHash EncodedHash `json:"prev_hash"`

	// prev key hash
	// Required: true
	PrevKeyHash EncodedHash `json:"prev_key_hash"`

	// state hash
	// Required: true
	StateHash EncodedHash `json:"state_hash"`

	// target
	// Required: true
	Target *uint64 `json:"target"`

	// time
	// Required: true
	Time *int64 `json:"time"`

	// version
	// Required: true
	Version *uint64 `json:"version"`
}

// Validate validates this key block
func (m *KeyBlock) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBeneficiary(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHash(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHeight(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMiner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePow(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrevHash(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrevKeyHash(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStateHash(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTarget(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KeyBlock) validateBeneficiary(formats strfmt.Registry) error {

	if err := m.Beneficiary.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("beneficiary")
		}
		return err
	}

	return nil
}

func (m *KeyBlock) validateHash(formats strfmt.Registry) error {

	if err := m.Hash.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("hash")
		}
		return err
	}

	return nil
}

func (m *KeyBlock) validateHeight(formats strfmt.Registry) error {

	if err := validate.Required("height", "body", m.Height); err != nil {
		return err
	}

	return nil
}

func (m *KeyBlock) validateInfo(formats strfmt.Registry) error {

	if err := m.Info.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("info")
		}
		return err
	}

	return nil
}

func (m *KeyBlock) validateMiner(formats strfmt.Registry) error {

	if err := m.Miner.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("miner")
		}
		return err
	}

	return nil
}

func (m *KeyBlock) validatePow(formats strfmt.Registry) error {

	if swag.IsZero(m.Pow) { // not required
		return nil
	}

	if err := m.Pow.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("pow")
		}
		return err
	}

	return nil
}

func (m *KeyBlock) validatePrevHash(formats strfmt.Registry) error {

	if err := m.PrevHash.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("prev_hash")
		}
		return err
	}

	return nil
}

func (m *KeyBlock) validatePrevKeyHash(formats strfmt.Registry) error {

	if err := m.PrevKeyHash.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("prev_key_hash")
		}
		return err
	}

	return nil
}

func (m *KeyBlock) validateStateHash(formats strfmt.Registry) error {

	if err := m.StateHash.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("state_hash")
		}
		return err
	}

	return nil
}

func (m *KeyBlock) validateTarget(formats strfmt.Registry) error {

	if err := validate.Required("target", "body", m.Target); err != nil {
		return err
	}

	return nil
}

func (m *KeyBlock) validateTime(formats strfmt.Registry) error {

	if err := validate.Required("time", "body", m.Time); err != nil {
		return err
	}

	return nil
}

func (m *KeyBlock) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KeyBlock) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KeyBlock) UnmarshalBinary(b []byte) error {
	var res KeyBlock
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
