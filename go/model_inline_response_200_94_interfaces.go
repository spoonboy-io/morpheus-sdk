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

// InlineResponse20094Interfaces struct for InlineResponse20094Interfaces
type InlineResponse20094Interfaces struct {
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Code NullableString `json:"code,omitempty"`
	InterfaceType NullableString `json:"interfaceType,omitempty"`
	NetworkPosition NullableString `json:"networkPosition,omitempty"`
	IpAddress *string `json:"ipAddress,omitempty"`
	Cidr *string `json:"cidr,omitempty"`
	ExternalLink NullableString `json:"externalLink,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	Network *InlineResponse20094Network `json:"network,omitempty"`
}

// NewInlineResponse20094Interfaces instantiates a new InlineResponse20094Interfaces object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20094Interfaces() *InlineResponse20094Interfaces {
	this := InlineResponse20094Interfaces{}
	return &this
}

// NewInlineResponse20094InterfacesWithDefaults instantiates a new InlineResponse20094Interfaces object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20094InterfacesWithDefaults() *InlineResponse20094Interfaces {
	this := InlineResponse20094Interfaces{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse20094Interfaces) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20094Interfaces) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse20094Interfaces) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *InlineResponse20094Interfaces) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20094Interfaces) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20094Interfaces) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20094Interfaces) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20094Interfaces) SetName(v string) {
	o.Name = &v
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineResponse20094Interfaces) GetCode() string {
	if o == nil || o.Code.Get() == nil {
		var ret string
		return ret
	}
	return *o.Code.Get()
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineResponse20094Interfaces) GetCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Code.Get(), o.Code.IsSet()
}

// HasCode returns a boolean if a field has been set.
func (o *InlineResponse20094Interfaces) HasCode() bool {
	if o != nil && o.Code.IsSet() {
		return true
	}

	return false
}

// SetCode gets a reference to the given NullableString and assigns it to the Code field.
func (o *InlineResponse20094Interfaces) SetCode(v string) {
	o.Code.Set(&v)
}
// SetCodeNil sets the value for Code to be an explicit nil
func (o *InlineResponse20094Interfaces) SetCodeNil() {
	o.Code.Set(nil)
}

// UnsetCode ensures that no value is present for Code, not even an explicit nil
func (o *InlineResponse20094Interfaces) UnsetCode() {
	o.Code.Unset()
}

// GetInterfaceType returns the InterfaceType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineResponse20094Interfaces) GetInterfaceType() string {
	if o == nil || o.InterfaceType.Get() == nil {
		var ret string
		return ret
	}
	return *o.InterfaceType.Get()
}

// GetInterfaceTypeOk returns a tuple with the InterfaceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineResponse20094Interfaces) GetInterfaceTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.InterfaceType.Get(), o.InterfaceType.IsSet()
}

// HasInterfaceType returns a boolean if a field has been set.
func (o *InlineResponse20094Interfaces) HasInterfaceType() bool {
	if o != nil && o.InterfaceType.IsSet() {
		return true
	}

	return false
}

// SetInterfaceType gets a reference to the given NullableString and assigns it to the InterfaceType field.
func (o *InlineResponse20094Interfaces) SetInterfaceType(v string) {
	o.InterfaceType.Set(&v)
}
// SetInterfaceTypeNil sets the value for InterfaceType to be an explicit nil
func (o *InlineResponse20094Interfaces) SetInterfaceTypeNil() {
	o.InterfaceType.Set(nil)
}

// UnsetInterfaceType ensures that no value is present for InterfaceType, not even an explicit nil
func (o *InlineResponse20094Interfaces) UnsetInterfaceType() {
	o.InterfaceType.Unset()
}

// GetNetworkPosition returns the NetworkPosition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineResponse20094Interfaces) GetNetworkPosition() string {
	if o == nil || o.NetworkPosition.Get() == nil {
		var ret string
		return ret
	}
	return *o.NetworkPosition.Get()
}

// GetNetworkPositionOk returns a tuple with the NetworkPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineResponse20094Interfaces) GetNetworkPositionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NetworkPosition.Get(), o.NetworkPosition.IsSet()
}

// HasNetworkPosition returns a boolean if a field has been set.
func (o *InlineResponse20094Interfaces) HasNetworkPosition() bool {
	if o != nil && o.NetworkPosition.IsSet() {
		return true
	}

	return false
}

// SetNetworkPosition gets a reference to the given NullableString and assigns it to the NetworkPosition field.
func (o *InlineResponse20094Interfaces) SetNetworkPosition(v string) {
	o.NetworkPosition.Set(&v)
}
// SetNetworkPositionNil sets the value for NetworkPosition to be an explicit nil
func (o *InlineResponse20094Interfaces) SetNetworkPositionNil() {
	o.NetworkPosition.Set(nil)
}

// UnsetNetworkPosition ensures that no value is present for NetworkPosition, not even an explicit nil
func (o *InlineResponse20094Interfaces) UnsetNetworkPosition() {
	o.NetworkPosition.Unset()
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *InlineResponse20094Interfaces) GetIpAddress() string {
	if o == nil || o.IpAddress == nil {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20094Interfaces) GetIpAddressOk() (*string, bool) {
	if o == nil || o.IpAddress == nil {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *InlineResponse20094Interfaces) HasIpAddress() bool {
	if o != nil && o.IpAddress != nil {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *InlineResponse20094Interfaces) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetCidr returns the Cidr field value if set, zero value otherwise.
func (o *InlineResponse20094Interfaces) GetCidr() string {
	if o == nil || o.Cidr == nil {
		var ret string
		return ret
	}
	return *o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20094Interfaces) GetCidrOk() (*string, bool) {
	if o == nil || o.Cidr == nil {
		return nil, false
	}
	return o.Cidr, true
}

// HasCidr returns a boolean if a field has been set.
func (o *InlineResponse20094Interfaces) HasCidr() bool {
	if o != nil && o.Cidr != nil {
		return true
	}

	return false
}

// SetCidr gets a reference to the given string and assigns it to the Cidr field.
func (o *InlineResponse20094Interfaces) SetCidr(v string) {
	o.Cidr = &v
}

// GetExternalLink returns the ExternalLink field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineResponse20094Interfaces) GetExternalLink() string {
	if o == nil || o.ExternalLink.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExternalLink.Get()
}

// GetExternalLinkOk returns a tuple with the ExternalLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineResponse20094Interfaces) GetExternalLinkOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExternalLink.Get(), o.ExternalLink.IsSet()
}

// HasExternalLink returns a boolean if a field has been set.
func (o *InlineResponse20094Interfaces) HasExternalLink() bool {
	if o != nil && o.ExternalLink.IsSet() {
		return true
	}

	return false
}

// SetExternalLink gets a reference to the given NullableString and assigns it to the ExternalLink field.
func (o *InlineResponse20094Interfaces) SetExternalLink(v string) {
	o.ExternalLink.Set(&v)
}
// SetExternalLinkNil sets the value for ExternalLink to be an explicit nil
func (o *InlineResponse20094Interfaces) SetExternalLinkNil() {
	o.ExternalLink.Set(nil)
}

// UnsetExternalLink ensures that no value is present for ExternalLink, not even an explicit nil
func (o *InlineResponse20094Interfaces) UnsetExternalLink() {
	o.ExternalLink.Unset()
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse20094Interfaces) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20094Interfaces) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse20094Interfaces) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse20094Interfaces) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *InlineResponse20094Interfaces) GetNetwork() InlineResponse20094Network {
	if o == nil || o.Network == nil {
		var ret InlineResponse20094Network
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20094Interfaces) GetNetworkOk() (*InlineResponse20094Network, bool) {
	if o == nil || o.Network == nil {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *InlineResponse20094Interfaces) HasNetwork() bool {
	if o != nil && o.Network != nil {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given InlineResponse20094Network and assigns it to the Network field.
func (o *InlineResponse20094Interfaces) SetNetwork(v InlineResponse20094Network) {
	o.Network = &v
}

func (o InlineResponse20094Interfaces) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Code.IsSet() {
		toSerialize["code"] = o.Code.Get()
	}
	if o.InterfaceType.IsSet() {
		toSerialize["interfaceType"] = o.InterfaceType.Get()
	}
	if o.NetworkPosition.IsSet() {
		toSerialize["networkPosition"] = o.NetworkPosition.Get()
	}
	if o.IpAddress != nil {
		toSerialize["ipAddress"] = o.IpAddress
	}
	if o.Cidr != nil {
		toSerialize["cidr"] = o.Cidr
	}
	if o.ExternalLink.IsSet() {
		toSerialize["externalLink"] = o.ExternalLink.Get()
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Network != nil {
		toSerialize["network"] = o.Network
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20094Interfaces struct {
	value *InlineResponse20094Interfaces
	isSet bool
}

func (v NullableInlineResponse20094Interfaces) Get() *InlineResponse20094Interfaces {
	return v.value
}

func (v *NullableInlineResponse20094Interfaces) Set(val *InlineResponse20094Interfaces) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20094Interfaces) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20094Interfaces) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20094Interfaces(val *InlineResponse20094Interfaces) *NullableInlineResponse20094Interfaces {
	return &NullableInlineResponse20094Interfaces{value: val, isSet: true}
}

func (v NullableInlineResponse20094Interfaces) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20094Interfaces) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


