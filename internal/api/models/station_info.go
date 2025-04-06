// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StationInfo station info
//
// swagger:model StationInfo
type StationInfo struct {

	// bikes available
	BikesAvailable int64 `json:"bikesAvailable,omitempty"`

	// docks available
	DocksAvailable int64 `json:"docksAvailable,omitempty"`

	// latitude
	Latitude float32 `json:"latitude,omitempty"`

	// longitude
	Longitude float32 `json:"longitude,omitempty"`

	// station Id
	StationID string `json:"stationId,omitempty"`

	// station name
	StationName string `json:"stationName,omitempty"`
}

// Validate validates this station info
func (m *StationInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this station info based on context it is used
func (m *StationInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StationInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StationInfo) UnmarshalBinary(b []byte) error {
	var res StationInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
