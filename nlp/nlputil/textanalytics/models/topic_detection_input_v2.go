package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// TopicDetectionInputV2 topic detection input v2
// swagger:model TopicDetectionInputV2
type TopicDetectionInputV2 struct {

	// documents
	Documents []*InputV2 `json:"documents"`

	// List of words to ignore from all documents during pre-processing.
	StopWords []string `json:"stopWords"`

	// List of topics to omit from the response.
	TopicsToExclude []string `json:"topicsToExclude"`
}

// Validate validates this topic detection input v2
func (m *TopicDetectionInputV2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDocuments(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStopWords(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTopicsToExclude(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TopicDetectionInputV2) validateDocuments(formats strfmt.Registry) error {

	if swag.IsZero(m.Documents) { // not required
		return nil
	}

	for i := 0; i < len(m.Documents); i++ {

		if swag.IsZero(m.Documents[i]) { // not required
			continue
		}

		if m.Documents[i] != nil {

			if err := m.Documents[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *TopicDetectionInputV2) validateStopWords(formats strfmt.Registry) error {

	if swag.IsZero(m.StopWords) { // not required
		return nil
	}

	return nil
}

func (m *TopicDetectionInputV2) validateTopicsToExclude(formats strfmt.Registry) error {

	if swag.IsZero(m.TopicsToExclude) { // not required
		return nil
	}

	return nil
}
