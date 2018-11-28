// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// InventoryAddNodeExporterAgentRequest inventory add node exporter agent request
// swagger:model inventoryAddNodeExporterAgentRequest
type InventoryAddNodeExporterAgentRequest struct {

	// Agent status: enabled or disabled.
	Disabled bool `json:"disabled,omitempty"`

	// Node identifier where Agent runs and for which insights are provided by that Agent.
	RunsOnNodeID int64 `json:"runs_on_node_id,omitempty"`
}

// Validate validates this inventory add node exporter agent request
func (m *InventoryAddNodeExporterAgentRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InventoryAddNodeExporterAgentRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InventoryAddNodeExporterAgentRequest) UnmarshalBinary(b []byte) error {
	var res InventoryAddNodeExporterAgentRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
