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

// InlineResponse20095NetworkRouterNetworkServer struct for InlineResponse20095NetworkRouterNetworkServer
type InlineResponse20095NetworkRouterNetworkServer struct {
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Integration *InlineResponse20095NetworkRouterNetworkServerIntegration `json:"integration,omitempty"`
}

// NewInlineResponse20095NetworkRouterNetworkServer instantiates a new InlineResponse20095NetworkRouterNetworkServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20095NetworkRouterNetworkServer() *InlineResponse20095NetworkRouterNetworkServer {
	this := InlineResponse20095NetworkRouterNetworkServer{}
	return &this
}

// NewInlineResponse20095NetworkRouterNetworkServerWithDefaults instantiates a new InlineResponse20095NetworkRouterNetworkServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20095NetworkRouterNetworkServerWithDefaults() *InlineResponse20095NetworkRouterNetworkServer {
	this := InlineResponse20095NetworkRouterNetworkServer{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse20095NetworkRouterNetworkServer) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095NetworkRouterNetworkServer) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse20095NetworkRouterNetworkServer) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *InlineResponse20095NetworkRouterNetworkServer) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20095NetworkRouterNetworkServer) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095NetworkRouterNetworkServer) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20095NetworkRouterNetworkServer) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20095NetworkRouterNetworkServer) SetName(v string) {
	o.Name = &v
}

// GetIntegration returns the Integration field value if set, zero value otherwise.
func (o *InlineResponse20095NetworkRouterNetworkServer) GetIntegration() InlineResponse20095NetworkRouterNetworkServerIntegration {
	if o == nil || o.Integration == nil {
		var ret InlineResponse20095NetworkRouterNetworkServerIntegration
		return ret
	}
	return *o.Integration
}

// GetIntegrationOk returns a tuple with the Integration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20095NetworkRouterNetworkServer) GetIntegrationOk() (*InlineResponse20095NetworkRouterNetworkServerIntegration, bool) {
	if o == nil || o.Integration == nil {
		return nil, false
	}
	return o.Integration, true
}

// HasIntegration returns a boolean if a field has been set.
func (o *InlineResponse20095NetworkRouterNetworkServer) HasIntegration() bool {
	if o != nil && o.Integration != nil {
		return true
	}

	return false
}

// SetIntegration gets a reference to the given InlineResponse20095NetworkRouterNetworkServerIntegration and assigns it to the Integration field.
func (o *InlineResponse20095NetworkRouterNetworkServer) SetIntegration(v InlineResponse20095NetworkRouterNetworkServerIntegration) {
	o.Integration = &v
}

func (o InlineResponse20095NetworkRouterNetworkServer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Integration != nil {
		toSerialize["integration"] = o.Integration
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20095NetworkRouterNetworkServer struct {
	value *InlineResponse20095NetworkRouterNetworkServer
	isSet bool
}

func (v NullableInlineResponse20095NetworkRouterNetworkServer) Get() *InlineResponse20095NetworkRouterNetworkServer {
	return v.value
}

func (v *NullableInlineResponse20095NetworkRouterNetworkServer) Set(val *InlineResponse20095NetworkRouterNetworkServer) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20095NetworkRouterNetworkServer) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20095NetworkRouterNetworkServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20095NetworkRouterNetworkServer(val *InlineResponse20095NetworkRouterNetworkServer) *NullableInlineResponse20095NetworkRouterNetworkServer {
	return &NullableInlineResponse20095NetworkRouterNetworkServer{value: val, isSet: true}
}

func (v NullableInlineResponse20095NetworkRouterNetworkServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20095NetworkRouterNetworkServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


