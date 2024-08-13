// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetMessages get messages
//
// swagger:model GetMessages
type GetMessages struct {

	// id
	// Example: 1
	ID int64 `json:"id,omitempty"`

	// message
	// Example: some text
	Message string `json:"message,omitempty"`

	// user name
	// Example: user_123
	UserName string `json:"user_name,omitempty"`
}

// Validate validates this get messages
func (m *GetMessages) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get messages based on context it is used
func (m *GetMessages) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetMessages) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetMessages) UnmarshalBinary(b []byte) error {
	var res GetMessages
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}