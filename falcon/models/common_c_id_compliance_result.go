// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CommonCIDComplianceResult common c ID compliance result
//
// swagger:model common.CIDComplianceResult
type CommonCIDComplianceResult struct {

	// average overall score
	// Required: true
	AverageOverallScore *float64 `json:"average_overall_score"`

	// cid
	// Required: true
	Cid *string `json:"cid"`

	// num aids
	// Required: true
	NumAids *int64 `json:"num_aids"`

	// platforms
	// Required: true
	Platforms []*CommonOSCompliance `json:"platforms"`
}

// Validate validates this common c ID compliance result
func (m *CommonCIDComplianceResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAverageOverallScore(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumAids(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatforms(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CommonCIDComplianceResult) validateAverageOverallScore(formats strfmt.Registry) error {

	if err := validate.Required("average_overall_score", "body", m.AverageOverallScore); err != nil {
		return err
	}

	return nil
}

func (m *CommonCIDComplianceResult) validateCid(formats strfmt.Registry) error {

	if err := validate.Required("cid", "body", m.Cid); err != nil {
		return err
	}

	return nil
}

func (m *CommonCIDComplianceResult) validateNumAids(formats strfmt.Registry) error {

	if err := validate.Required("num_aids", "body", m.NumAids); err != nil {
		return err
	}

	return nil
}

func (m *CommonCIDComplianceResult) validatePlatforms(formats strfmt.Registry) error {

	if err := validate.Required("platforms", "body", m.Platforms); err != nil {
		return err
	}

	for i := 0; i < len(m.Platforms); i++ {
		if swag.IsZero(m.Platforms[i]) { // not required
			continue
		}

		if m.Platforms[i] != nil {
			if err := m.Platforms[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("platforms" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("platforms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this common c ID compliance result based on the context it is used
func (m *CommonCIDComplianceResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePlatforms(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CommonCIDComplianceResult) contextValidatePlatforms(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Platforms); i++ {

		if m.Platforms[i] != nil {
			if err := m.Platforms[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("platforms" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("platforms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CommonCIDComplianceResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommonCIDComplianceResult) UnmarshalBinary(b []byte) error {
	var res CommonCIDComplianceResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
