// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Route route
// swagger:model Route
type Route struct {

	// Route annotations - this is a map of annotations attached to this route, keys must not exceed 128 bytes and must consist of non-whitespace printable ascii characters, and the seralized representation of individual values must not exeed 512 bytes
	Annotations map[string]interface{} `json:"annotations,omitempty"`

	// App ID
	AppID string `json:"app_id,omitempty"`

	// Route configuration - overrides application configuration
	Config map[string]string `json:"config,omitempty"`

	// Max usable CPU cores for this route. Value in MilliCPUs (eg. 500m) or as floating-point (eg. 0.5)
	Cpus string `json:"cpus,omitempty"`

	// Time when route was created. Always in UTC.
	// Read Only: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// Payload format sent into function.
	// Enum: [default http json]
	Format string `json:"format,omitempty"`

	// Map of http headers that will be sent with the response
	Headers map[string][]string `json:"headers,omitempty"`

	// Hot functions idle timeout before termination. Value in Seconds
	IDLETimeout *int32 `json:"idle_timeout,omitempty"`

	// Name of Docker image to use in this route. You should include the image tag, which should be a version number, to be more accurate. Can be overridden on a per route basis with route.image.
	Image string `json:"image,omitempty"`

	// Max usable memory for this route (MiB).
	Memory uint64 `json:"memory,omitempty"`

	// URL path that will be matched to this route
	// Read Only: true
	Path string `json:"path,omitempty"`

	// Timeout for executions of this route. Value in Seconds
	Timeout *int32 `json:"timeout,omitempty"`

	// Tmpfs size (MiB).
	TmpfsSize uint32 `json:"tmpfs_size,omitempty"`

	// Route type
	// Enum: [sync async]
	Type string `json:"type,omitempty"`

	// Most recent time that route was updated. Always in UTC.
	// Read Only: true
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`
}

// Validate validates this route
func (m *Route) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFormat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Route) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

var routeTypeFormatPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["default","http","json"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		routeTypeFormatPropEnum = append(routeTypeFormatPropEnum, v)
	}
}

const (

	// RouteFormatDefault captures enum value "default"
	RouteFormatDefault string = "default"

	// RouteFormatHTTP captures enum value "http"
	RouteFormatHTTP string = "http"

	// RouteFormatJSON captures enum value "json"
	RouteFormatJSON string = "json"
)

// prop value enum
func (m *Route) validateFormatEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, routeTypeFormatPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Route) validateFormat(formats strfmt.Registry) error {

	if swag.IsZero(m.Format) { // not required
		return nil
	}

	// value enum
	if err := m.validateFormatEnum("format", "body", m.Format); err != nil {
		return err
	}

	return nil
}

var routeTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["sync","async"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		routeTypeTypePropEnum = append(routeTypeTypePropEnum, v)
	}
}

const (

	// RouteTypeSync captures enum value "sync"
	RouteTypeSync string = "sync"

	// RouteTypeAsync captures enum value "async"
	RouteTypeAsync string = "async"
)

// prop value enum
func (m *Route) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, routeTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Route) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *Route) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Route) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Route) UnmarshalBinary(b []byte) error {
	var res Route
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
