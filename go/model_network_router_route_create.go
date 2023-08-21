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

// NetworkRouterRouteCreate struct for NetworkRouterRouteCreate
type NetworkRouterRouteCreate struct {
	// Route name
	Name *string `json:"name,omitempty"`
	// Route description
	Description *string `json:"description,omitempty"`
	// Can be used to enable / disable the route (true, false). Default is off
	Enabled *bool `json:"enabled,omitempty"`
	// Can be used to set as default route (true, false). Default is off
	DefaultRoute *bool `json:"defaultRoute,omitempty"`
	// Source address or range
	Source string `json:"source"`
	// Destination address or range
	Destination string `json:"destination"`
	// MTU
	NetworkMtu string `json:"networkMtu"`
}

// NewNetworkRouterRouteCreate instantiates a new NetworkRouterRouteCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkRouterRouteCreate(source string, destination string, networkMtu string, ) *NetworkRouterRouteCreate {
	this := NetworkRouterRouteCreate{}
	var enabled bool = false
	this.Enabled = &enabled
	var defaultRoute bool = false
	this.DefaultRoute = &defaultRoute
	this.Source = source
	this.Destination = destination
	this.NetworkMtu = networkMtu
	return &this
}

// NewNetworkRouterRouteCreateWithDefaults instantiates a new NetworkRouterRouteCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkRouterRouteCreateWithDefaults() *NetworkRouterRouteCreate {
	this := NetworkRouterRouteCreate{}
	var enabled bool = false
	this.Enabled = &enabled
	var defaultRoute bool = false
	this.DefaultRoute = &defaultRoute
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworkRouterRouteCreate) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkRouterRouteCreate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworkRouterRouteCreate) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworkRouterRouteCreate) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NetworkRouterRouteCreate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkRouterRouteCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NetworkRouterRouteCreate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NetworkRouterRouteCreate) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *NetworkRouterRouteCreate) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkRouterRouteCreate) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *NetworkRouterRouteCreate) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *NetworkRouterRouteCreate) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetDefaultRoute returns the DefaultRoute field value if set, zero value otherwise.
func (o *NetworkRouterRouteCreate) GetDefaultRoute() bool {
	if o == nil || o.DefaultRoute == nil {
		var ret bool
		return ret
	}
	return *o.DefaultRoute
}

// GetDefaultRouteOk returns a tuple with the DefaultRoute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkRouterRouteCreate) GetDefaultRouteOk() (*bool, bool) {
	if o == nil || o.DefaultRoute == nil {
		return nil, false
	}
	return o.DefaultRoute, true
}

// HasDefaultRoute returns a boolean if a field has been set.
func (o *NetworkRouterRouteCreate) HasDefaultRoute() bool {
	if o != nil && o.DefaultRoute != nil {
		return true
	}

	return false
}

// SetDefaultRoute gets a reference to the given bool and assigns it to the DefaultRoute field.
func (o *NetworkRouterRouteCreate) SetDefaultRoute(v bool) {
	o.DefaultRoute = &v
}

// GetSource returns the Source field value
func (o *NetworkRouterRouteCreate) GetSource() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *NetworkRouterRouteCreate) GetSourceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *NetworkRouterRouteCreate) SetSource(v string) {
	o.Source = v
}

// GetDestination returns the Destination field value
func (o *NetworkRouterRouteCreate) GetDestination() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *NetworkRouterRouteCreate) GetDestinationOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *NetworkRouterRouteCreate) SetDestination(v string) {
	o.Destination = v
}

// GetNetworkMtu returns the NetworkMtu field value
func (o *NetworkRouterRouteCreate) GetNetworkMtu() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.NetworkMtu
}

// GetNetworkMtuOk returns a tuple with the NetworkMtu field value
// and a boolean to check if the value has been set.
func (o *NetworkRouterRouteCreate) GetNetworkMtuOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NetworkMtu, true
}

// SetNetworkMtu sets field value
func (o *NetworkRouterRouteCreate) SetNetworkMtu(v string) {
	o.NetworkMtu = v
}

func (o NetworkRouterRouteCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.DefaultRoute != nil {
		toSerialize["defaultRoute"] = o.DefaultRoute
	}
	if true {
		toSerialize["source"] = o.Source
	}
	if true {
		toSerialize["destination"] = o.Destination
	}
	if true {
		toSerialize["networkMtu"] = o.NetworkMtu
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkRouterRouteCreate struct {
	value *NetworkRouterRouteCreate
	isSet bool
}

func (v NullableNetworkRouterRouteCreate) Get() *NetworkRouterRouteCreate {
	return v.value
}

func (v *NullableNetworkRouterRouteCreate) Set(val *NetworkRouterRouteCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkRouterRouteCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkRouterRouteCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkRouterRouteCreate(val *NetworkRouterRouteCreate) *NullableNetworkRouterRouteCreate {
	return &NullableNetworkRouterRouteCreate{value: val, isSet: true}
}

func (v NullableNetworkRouterRouteCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkRouterRouteCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


