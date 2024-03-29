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

// NetworkConfig struct for NetworkConfig
type NetworkConfig struct {
	VlanIDs NullableString `json:"vlanIDs,omitempty"`
	ConnectedGateway *string `json:"connectedGateway,omitempty"`
	SubnetIpManagementType *string `json:"subnetIpManagementType,omitempty"`
	SubnetIpServerId *string `json:"subnetIpServerId,omitempty"`
	DhcpRange *string `json:"dhcpRange,omitempty"`
	SubnetDhcpServerAddress *string `json:"subnetDhcpServerAddress,omitempty"`
	SubnetDhcpLeaseTime *string `json:"subnetDhcpLeaseTime,omitempty"`
}

// NewNetworkConfig instantiates a new NetworkConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkConfig() *NetworkConfig {
	this := NetworkConfig{}
	return &this
}

// NewNetworkConfigWithDefaults instantiates a new NetworkConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkConfigWithDefaults() *NetworkConfig {
	this := NetworkConfig{}
	return &this
}

// GetVlanIDs returns the VlanIDs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkConfig) GetVlanIDs() string {
	if o == nil || o.VlanIDs.Get() == nil {
		var ret string
		return ret
	}
	return *o.VlanIDs.Get()
}

// GetVlanIDsOk returns a tuple with the VlanIDs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkConfig) GetVlanIDsOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.VlanIDs.Get(), o.VlanIDs.IsSet()
}

// HasVlanIDs returns a boolean if a field has been set.
func (o *NetworkConfig) HasVlanIDs() bool {
	if o != nil && o.VlanIDs.IsSet() {
		return true
	}

	return false
}

// SetVlanIDs gets a reference to the given NullableString and assigns it to the VlanIDs field.
func (o *NetworkConfig) SetVlanIDs(v string) {
	o.VlanIDs.Set(&v)
}
// SetVlanIDsNil sets the value for VlanIDs to be an explicit nil
func (o *NetworkConfig) SetVlanIDsNil() {
	o.VlanIDs.Set(nil)
}

// UnsetVlanIDs ensures that no value is present for VlanIDs, not even an explicit nil
func (o *NetworkConfig) UnsetVlanIDs() {
	o.VlanIDs.Unset()
}

// GetConnectedGateway returns the ConnectedGateway field value if set, zero value otherwise.
func (o *NetworkConfig) GetConnectedGateway() string {
	if o == nil || o.ConnectedGateway == nil {
		var ret string
		return ret
	}
	return *o.ConnectedGateway
}

// GetConnectedGatewayOk returns a tuple with the ConnectedGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConfig) GetConnectedGatewayOk() (*string, bool) {
	if o == nil || o.ConnectedGateway == nil {
		return nil, false
	}
	return o.ConnectedGateway, true
}

// HasConnectedGateway returns a boolean if a field has been set.
func (o *NetworkConfig) HasConnectedGateway() bool {
	if o != nil && o.ConnectedGateway != nil {
		return true
	}

	return false
}

// SetConnectedGateway gets a reference to the given string and assigns it to the ConnectedGateway field.
func (o *NetworkConfig) SetConnectedGateway(v string) {
	o.ConnectedGateway = &v
}

// GetSubnetIpManagementType returns the SubnetIpManagementType field value if set, zero value otherwise.
func (o *NetworkConfig) GetSubnetIpManagementType() string {
	if o == nil || o.SubnetIpManagementType == nil {
		var ret string
		return ret
	}
	return *o.SubnetIpManagementType
}

// GetSubnetIpManagementTypeOk returns a tuple with the SubnetIpManagementType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConfig) GetSubnetIpManagementTypeOk() (*string, bool) {
	if o == nil || o.SubnetIpManagementType == nil {
		return nil, false
	}
	return o.SubnetIpManagementType, true
}

// HasSubnetIpManagementType returns a boolean if a field has been set.
func (o *NetworkConfig) HasSubnetIpManagementType() bool {
	if o != nil && o.SubnetIpManagementType != nil {
		return true
	}

	return false
}

// SetSubnetIpManagementType gets a reference to the given string and assigns it to the SubnetIpManagementType field.
func (o *NetworkConfig) SetSubnetIpManagementType(v string) {
	o.SubnetIpManagementType = &v
}

// GetSubnetIpServerId returns the SubnetIpServerId field value if set, zero value otherwise.
func (o *NetworkConfig) GetSubnetIpServerId() string {
	if o == nil || o.SubnetIpServerId == nil {
		var ret string
		return ret
	}
	return *o.SubnetIpServerId
}

// GetSubnetIpServerIdOk returns a tuple with the SubnetIpServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConfig) GetSubnetIpServerIdOk() (*string, bool) {
	if o == nil || o.SubnetIpServerId == nil {
		return nil, false
	}
	return o.SubnetIpServerId, true
}

// HasSubnetIpServerId returns a boolean if a field has been set.
func (o *NetworkConfig) HasSubnetIpServerId() bool {
	if o != nil && o.SubnetIpServerId != nil {
		return true
	}

	return false
}

// SetSubnetIpServerId gets a reference to the given string and assigns it to the SubnetIpServerId field.
func (o *NetworkConfig) SetSubnetIpServerId(v string) {
	o.SubnetIpServerId = &v
}

// GetDhcpRange returns the DhcpRange field value if set, zero value otherwise.
func (o *NetworkConfig) GetDhcpRange() string {
	if o == nil || o.DhcpRange == nil {
		var ret string
		return ret
	}
	return *o.DhcpRange
}

// GetDhcpRangeOk returns a tuple with the DhcpRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConfig) GetDhcpRangeOk() (*string, bool) {
	if o == nil || o.DhcpRange == nil {
		return nil, false
	}
	return o.DhcpRange, true
}

// HasDhcpRange returns a boolean if a field has been set.
func (o *NetworkConfig) HasDhcpRange() bool {
	if o != nil && o.DhcpRange != nil {
		return true
	}

	return false
}

// SetDhcpRange gets a reference to the given string and assigns it to the DhcpRange field.
func (o *NetworkConfig) SetDhcpRange(v string) {
	o.DhcpRange = &v
}

// GetSubnetDhcpServerAddress returns the SubnetDhcpServerAddress field value if set, zero value otherwise.
func (o *NetworkConfig) GetSubnetDhcpServerAddress() string {
	if o == nil || o.SubnetDhcpServerAddress == nil {
		var ret string
		return ret
	}
	return *o.SubnetDhcpServerAddress
}

// GetSubnetDhcpServerAddressOk returns a tuple with the SubnetDhcpServerAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConfig) GetSubnetDhcpServerAddressOk() (*string, bool) {
	if o == nil || o.SubnetDhcpServerAddress == nil {
		return nil, false
	}
	return o.SubnetDhcpServerAddress, true
}

// HasSubnetDhcpServerAddress returns a boolean if a field has been set.
func (o *NetworkConfig) HasSubnetDhcpServerAddress() bool {
	if o != nil && o.SubnetDhcpServerAddress != nil {
		return true
	}

	return false
}

// SetSubnetDhcpServerAddress gets a reference to the given string and assigns it to the SubnetDhcpServerAddress field.
func (o *NetworkConfig) SetSubnetDhcpServerAddress(v string) {
	o.SubnetDhcpServerAddress = &v
}

// GetSubnetDhcpLeaseTime returns the SubnetDhcpLeaseTime field value if set, zero value otherwise.
func (o *NetworkConfig) GetSubnetDhcpLeaseTime() string {
	if o == nil || o.SubnetDhcpLeaseTime == nil {
		var ret string
		return ret
	}
	return *o.SubnetDhcpLeaseTime
}

// GetSubnetDhcpLeaseTimeOk returns a tuple with the SubnetDhcpLeaseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkConfig) GetSubnetDhcpLeaseTimeOk() (*string, bool) {
	if o == nil || o.SubnetDhcpLeaseTime == nil {
		return nil, false
	}
	return o.SubnetDhcpLeaseTime, true
}

// HasSubnetDhcpLeaseTime returns a boolean if a field has been set.
func (o *NetworkConfig) HasSubnetDhcpLeaseTime() bool {
	if o != nil && o.SubnetDhcpLeaseTime != nil {
		return true
	}

	return false
}

// SetSubnetDhcpLeaseTime gets a reference to the given string and assigns it to the SubnetDhcpLeaseTime field.
func (o *NetworkConfig) SetSubnetDhcpLeaseTime(v string) {
	o.SubnetDhcpLeaseTime = &v
}

func (o NetworkConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.VlanIDs.IsSet() {
		toSerialize["vlanIDs"] = o.VlanIDs.Get()
	}
	if o.ConnectedGateway != nil {
		toSerialize["connectedGateway"] = o.ConnectedGateway
	}
	if o.SubnetIpManagementType != nil {
		toSerialize["subnetIpManagementType"] = o.SubnetIpManagementType
	}
	if o.SubnetIpServerId != nil {
		toSerialize["subnetIpServerId"] = o.SubnetIpServerId
	}
	if o.DhcpRange != nil {
		toSerialize["dhcpRange"] = o.DhcpRange
	}
	if o.SubnetDhcpServerAddress != nil {
		toSerialize["subnetDhcpServerAddress"] = o.SubnetDhcpServerAddress
	}
	if o.SubnetDhcpLeaseTime != nil {
		toSerialize["subnetDhcpLeaseTime"] = o.SubnetDhcpLeaseTime
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkConfig struct {
	value *NetworkConfig
	isSet bool
}

func (v NullableNetworkConfig) Get() *NetworkConfig {
	return v.value
}

func (v *NullableNetworkConfig) Set(val *NetworkConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkConfig(val *NetworkConfig) *NullableNetworkConfig {
	return &NullableNetworkConfig{value: val, isSet: true}
}

func (v NullableNetworkConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


