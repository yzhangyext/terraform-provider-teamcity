// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TriggeredBy triggered by
// swagger:model TriggeredBy
type TriggeredBy struct {

	// build
	Build *Build `json:"build,omitempty"`

	// build type
	BuildType *BuildType `json:"buildType,omitempty"`

	// date
	Date string `json:"date,omitempty" xml:"date"`

	// details
	Details string `json:"details,omitempty" xml:"details"`

	// properties
	Properties *Properties `json:"properties,omitempty"`

	// raw value
	RawValue string `json:"rawValue,omitempty" xml:"rawValue"`

	// type
	Type string `json:"type,omitempty" xml:"type"`

	// user
	User *User `json:"user,omitempty"`
}

// Validate validates this triggered by
func (m *TriggeredBy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuild(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateBuildType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProperties(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TriggeredBy) validateBuild(formats strfmt.Registry) error {

	if swag.IsZero(m.Build) { // not required
		return nil
	}

	if m.Build != nil {

		if err := m.Build.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("build")
			}
			return err
		}
	}

	return nil
}

func (m *TriggeredBy) validateBuildType(formats strfmt.Registry) error {

	if swag.IsZero(m.BuildType) { // not required
		return nil
	}

	if m.BuildType != nil {

		if err := m.BuildType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("buildType")
			}
			return err
		}
	}

	return nil
}

func (m *TriggeredBy) validateProperties(formats strfmt.Registry) error {

	if swag.IsZero(m.Properties) { // not required
		return nil
	}

	if m.Properties != nil {

		if err := m.Properties.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("properties")
			}
			return err
		}
	}

	return nil
}

func (m *TriggeredBy) validateUser(formats strfmt.Registry) error {

	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {

		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TriggeredBy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TriggeredBy) UnmarshalBinary(b []byte) error {
	var res TriggeredBy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
