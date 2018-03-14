// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ProblemOccurrence problem occurrence
// swagger:model problemOccurrence
type ProblemOccurrence struct {

	// additional data
	AdditionalData string `json:"additionalData,omitempty"`

	// build
	Build *Build `json:"build,omitempty"`

	// currently investigated
	CurrentlyInvestigated *bool `json:"currentlyInvestigated,omitempty" xml:"currentlyInvestigated"`

	// currently muted
	CurrentlyMuted *bool `json:"currentlyMuted,omitempty" xml:"currentlyMuted"`

	// details
	Details string `json:"details,omitempty"`

	// href
	Href string `json:"href,omitempty" xml:"href"`

	// id
	ID string `json:"id,omitempty" xml:"id"`

	// identity
	Identity string `json:"identity,omitempty" xml:"identity"`

	// mute
	Mute *Mute `json:"mute,omitempty"`

	// muted
	Muted *bool `json:"muted,omitempty" xml:"muted"`

	// problem
	Problem *Problem `json:"problem,omitempty"`

	// type
	Type string `json:"type,omitempty" xml:"type"`
}

// Validate validates this problem occurrence
func (m *ProblemOccurrence) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuild(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMute(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProblem(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProblemOccurrence) validateBuild(formats strfmt.Registry) error {

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

func (m *ProblemOccurrence) validateMute(formats strfmt.Registry) error {

	if swag.IsZero(m.Mute) { // not required
		return nil
	}

	if m.Mute != nil {

		if err := m.Mute.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mute")
			}
			return err
		}
	}

	return nil
}

func (m *ProblemOccurrence) validateProblem(formats strfmt.Registry) error {

	if swag.IsZero(m.Problem) { // not required
		return nil
	}

	if m.Problem != nil {

		if err := m.Problem.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("problem")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProblemOccurrence) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProblemOccurrence) UnmarshalBinary(b []byte) error {
	var res ProblemOccurrence
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
