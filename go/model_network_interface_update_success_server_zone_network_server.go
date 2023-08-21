/*
 * Morpheus API
 *
 * Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
 *
 * API version: 6.2.1
 * Contact: dev@morpheusdata.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// NetworkInterfaceUpdateSuccessServerZoneNetworkServer struct for NetworkInterfaceUpdateSuccessServerZoneNetworkServer
type NetworkInterfaceUpdateSuccessServerZoneNetworkServer struct {
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Type *NetworkInterfaceUpdateSuccessServerZoneNetworkServerType `json:"type,omitempty"`
	Status *string `json:"status,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	Visibility *string `json:"visibility,omitempty"`
}

// NewNetworkInterfaceUpdateSuccessServerZoneNetworkServer instantiates a new NetworkInterfaceUpdateSuccessServerZoneNetworkServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkInterfaceUpdateSuccessServerZoneNetworkServer() *NetworkInterfaceUpdateSuccessServerZoneNetworkServer {
	this := NetworkInterfaceUpdateSuccessServerZoneNetworkServer{}
	return &this
}

// NewNetworkInterfaceUpdateSuccessServerZoneNetworkServerWithDefaults instantiates a new NetworkInterfaceUpdateSuccessServerZoneNetworkServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkInterfaceUpdateSuccessServerZoneNetworkServerWithDefaults() *NetworkInterfaceUpdateSuccessServerZoneNetworkServer {
	this := NetworkInterfaceUpdateSuccessServerZoneNetworkServer{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NetworkInterfaceUpdateSuccessServerZoneNetworkServer) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkInterfaceUpdateSuccessServerZoneNetworkServer) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NetworkInterfaceUpdateSuccessServerZoneNetworkServer) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *NetworkInterfaceUpdateSuccessServerZoneNetworkServer) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworkInterfaceUpdateSuccessServerZoneNetworkServer) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkInterfaceUpdateSuccessServerZoneNetworkServer) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworkInterfaceUpdateSuccessServerZoneNetworkServer) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworkInterfaceUpdateSuccessServerZoneNetworkServer) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkInterfaceUpdateSuccessServerZoneNetworkServer) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkInterfaceUpdateSuccessServerZoneNetworkServer) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *NetworkInterfaceUpdateSuccessServerZoneNetworkServer) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *NetworkInterfaceUpdateSuccessServerZoneNetworkServer) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *NetworkInterfaceUpdateSuccessServerZoneNetworkServer) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *NetworkInterfaceUpdateSuccessServerZoneNetworkServer) UnsetDescription() {
	o.Description.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NetworkInterfaceUpdateSuccessServerZoneNetworkServer) GetType() NetworkInterfaceUpdateSuccessServerZoneNetworkServerType {
	if o == nil || o.Type == nil {
		var ret NetworkInterfaceUpdateSuccessServerZoneNetworkServerType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkInterfaceUpdateSuccessServerZoneNetworkServer) GetTypeOk() (*NetworkInterfaceUpdateSuccessServerZoneNetworkServerType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NetworkInterfaceUpdateSuccessServerZoneNetworkServer) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given NetworkInterfaceUpdateSuccessServerZoneNetworkServerType and assigns it to the Type field.
func (o *NetworkInterfaceUpdateSuccessServerZoneNetworkServer) SetType(v NetworkInterfaceUpdateSuccessServerZoneNetworkServerType) {
	o.Type = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *NetworkInterfaceUpdateSuccessServerZoneNetworkServer) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkInterfaceUpdateSuccessServerZoneNetworkServer) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *NetworkInterfaceUpdateSuccessServerZoneNetworkServer) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *NetworkInterfaceUpdateSuccessServerZoneNetworkServer) SetStatus(v string) {
	o.Status = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *NetworkInterfaceUpdateSuccessServerZoneNetworkServer) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkInterfaceUpdateSuccessServerZoneNetworkServer) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *NetworkInterfaceUpdateSuccessServerZoneNetworkServer) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *NetworkInterfaceUpdateSuccessServerZoneNetworkServer) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *NetworkInterfaceUpdateSuccessServerZoneNetworkServer) GetVisibility() string {
	if o == nil || o.Visibility == nil {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkInterfaceUpdateSuccessServerZoneNetworkServer) GetVisibilityOk() (*string, bool) {
	if o == nil || o.Visibility == nil {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *NetworkInterfaceUpdateSuccessServerZoneNetworkServer) HasVisibility() bool {
	if o != nil && o.Visibility != nil {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *NetworkInterfaceUpdateSuccessServerZoneNetworkServer) SetVisibility(v string) {
	o.Visibility = &v
}

func (o NetworkInterfaceUpdateSuccessServerZoneNetworkServer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Visibility != nil {
		toSerialize["visibility"] = o.Visibility
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkInterfaceUpdateSuccessServerZoneNetworkServer struct {
	value *NetworkInterfaceUpdateSuccessServerZoneNetworkServer
	isSet bool
}

func (v NullableNetworkInterfaceUpdateSuccessServerZoneNetworkServer) Get() *NetworkInterfaceUpdateSuccessServerZoneNetworkServer {
	return v.value
}

func (v *NullableNetworkInterfaceUpdateSuccessServerZoneNetworkServer) Set(val *NetworkInterfaceUpdateSuccessServerZoneNetworkServer) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkInterfaceUpdateSuccessServerZoneNetworkServer) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkInterfaceUpdateSuccessServerZoneNetworkServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkInterfaceUpdateSuccessServerZoneNetworkServer(val *NetworkInterfaceUpdateSuccessServerZoneNetworkServer) *NullableNetworkInterfaceUpdateSuccessServerZoneNetworkServer {
	return &NullableNetworkInterfaceUpdateSuccessServerZoneNetworkServer{value: val, isSet: true}
}

func (v NullableNetworkInterfaceUpdateSuccessServerZoneNetworkServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkInterfaceUpdateSuccessServerZoneNetworkServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


