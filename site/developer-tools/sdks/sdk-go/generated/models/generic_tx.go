// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/validate"
)

// GenericTx generic tx
// swagger:discriminator GenericTx type
type GenericTx interface {
	runtime.Validatable

	// type
	// Required: true
	Type() string
	SetType(string)

	// version
	// Required: true
	Version() *uint64
	SetVersion(*uint64)
}

type genericTx struct {
	typeField string

	versionField *uint64
}

// Type gets the type of this polymorphic type
func (m *genericTx) Type() string {
	return "GenericTx"
}

// SetType sets the type of this polymorphic type
func (m *genericTx) SetType(val string) {

}

// Version gets the version of this polymorphic type
func (m *genericTx) Version() *uint64 {
	return m.versionField
}

// SetVersion sets the version of this polymorphic type
func (m *genericTx) SetVersion(val *uint64) {
	m.versionField = val
}

// UnmarshalGenericTxSlice unmarshals polymorphic slices of GenericTx
func UnmarshalGenericTxSlice(reader io.Reader, consumer runtime.Consumer) ([]GenericTx, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []GenericTx
	for _, element := range elements {
		obj, err := unmarshalGenericTx(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalGenericTx unmarshals polymorphic GenericTx
func UnmarshalGenericTx(reader io.Reader, consumer runtime.Consumer) (GenericTx, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalGenericTx(data, consumer)
}

func unmarshalGenericTx(data []byte, consumer runtime.Consumer) (GenericTx, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the type property.
	var getType struct {
		Type string `json:"type"`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("type", "body", getType.Type); err != nil {
		return nil, err
	}

	// The value of type is used to determine which type to create and unmarshal the data into
	switch getType.Type {
	case "ChannelCloseMutualTxJSON":
		var result ChannelCloseMutualTxJSON
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "ChannelCloseSoloTxJSON":
		var result ChannelCloseSoloTxJSON
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "ChannelCreateTxJSON":
		var result ChannelCreateTxJSON
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "ChannelDepositTxJSON":
		var result ChannelDepositTxJSON
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "ChannelForceProgressTxJSON":
		var result ChannelForceProgressTxJSON
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "ChannelSettleTxJSON":
		var result ChannelSettleTxJSON
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "ChannelSlashTxJSON":
		var result ChannelSLASHTxJSON
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "ChannelSnapshotSoloTxJSON":
		var result ChannelSnapshotSoloTxJSON
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "ChannelWithdrawalTxJSON":
		var result ChannelWithdrawalTxJSON
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "ContractCallTxJSON":
		var result ContractCallTxJSON
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "ContractCreateTxJSON":
		var result ContractCreateTxJSON
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "GenericTx":
		var result genericTx
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "NameClaimTxJSON":
		var result NameClaimTxJSON
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "NamePreclaimTxJSON":
		var result NamePreclaimTxJSON
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "NameRevokeTxJSON":
		var result NameRevokeTxJSON
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "NameTransferTxJSON":
		var result NameTransferTxJSON
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "NameUpdateTxJSON":
		var result NameUpdateTxJSON
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "OracleExtendTxJSON":
		var result OracleExtendTxJSON
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "OracleQueryTxJSON":
		var result OracleQueryTxJSON
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "OracleRegisterTxJSON":
		var result OracleRegisterTxJSON
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "OracleResponseTxJSON":
		var result OracleResponseTxJSON
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "SpendTxJSON":
		var result SpendTxJSON
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	}
	return nil, errors.New(422, "invalid type value: %q", getType.Type)

}

// Validate validates this generic tx
func (m *genericTx) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *genericTx) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version()); err != nil {
		return err
	}

	return nil
}
