// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"io"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GenericSignedTx generic signed tx
// swagger:model GenericSignedTx
type GenericSignedTx struct {

	// Value "none" means no block in the chain includes the transaction
	// Required: true
	BlockHash EncodedHash `json:"block_hash"`

	// block height
	// Required: true
	BlockHeight *uint64 `json:"block_height"`

	// hash
	// Required: true
	Hash EncodedHash `json:"hash"`

	// signatures
	// Required: true
	// Min Items: 1
	Signatures []string `json:"signatures"`

	txField GenericTx
}

// Tx gets the tx of this base type
func (m *GenericSignedTx) Tx() GenericTx {
	return m.txField
}

// SetTx sets the tx of this base type
func (m *GenericSignedTx) SetTx(val GenericTx) {
	m.txField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *GenericSignedTx) UnmarshalJSON(raw []byte) error {
	var data struct {
		BlockHash EncodedHash `json:"block_hash"`

		BlockHeight *uint64 `json:"block_height"`

		Hash EncodedHash `json:"hash"`

		Signatures []string `json:"signatures"`

		Tx json.RawMessage `json:"tx"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	propTx, err := UnmarshalGenericTx(bytes.NewBuffer(data.Tx), runtime.JSONConsumer())
	if err != nil && err != io.EOF {
		return err
	}

	var result GenericSignedTx

	// block_hash
	result.BlockHash = data.BlockHash

	// block_height
	result.BlockHeight = data.BlockHeight

	// hash
	result.Hash = data.Hash

	// signatures
	result.Signatures = data.Signatures

	// tx
	result.txField = propTx

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m GenericSignedTx) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		BlockHash EncodedHash `json:"block_hash"`

		BlockHeight *uint64 `json:"block_height"`

		Hash EncodedHash `json:"hash"`

		Signatures []string `json:"signatures"`
	}{

		BlockHash: m.BlockHash,

		BlockHeight: m.BlockHeight,

		Hash: m.Hash,

		Signatures: m.Signatures,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Tx GenericTx `json:"tx"`
	}{

		Tx: m.txField,
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this generic signed tx
func (m *GenericSignedTx) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBlockHash(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBlockHeight(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHash(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSignatures(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTx(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GenericSignedTx) validateBlockHash(formats strfmt.Registry) error {

	if err := m.BlockHash.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("block_hash")
		}
		return err
	}

	return nil
}

func (m *GenericSignedTx) validateBlockHeight(formats strfmt.Registry) error {

	if err := validate.Required("block_height", "body", m.BlockHeight); err != nil {
		return err
	}

	return nil
}

func (m *GenericSignedTx) validateHash(formats strfmt.Registry) error {

	if err := m.Hash.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("hash")
		}
		return err
	}

	return nil
}

func (m *GenericSignedTx) validateSignatures(formats strfmt.Registry) error {

	if err := validate.Required("signatures", "body", m.Signatures); err != nil {
		return err
	}

	iSignaturesSize := int64(len(m.Signatures))

	if err := validate.MinItems("signatures", "body", iSignaturesSize, 1); err != nil {
		return err
	}

	return nil
}

func (m *GenericSignedTx) validateTx(formats strfmt.Registry) error {

	if err := m.Tx().Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("tx")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GenericSignedTx) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GenericSignedTx) UnmarshalBinary(b []byte) error {
	var res GenericSignedTx
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
