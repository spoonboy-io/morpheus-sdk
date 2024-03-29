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

// GuidanceVmwareSizingResourceInterfaces struct for GuidanceVmwareSizingResourceInterfaces
type GuidanceVmwareSizingResourceInterfaces struct {
	Id *int64 `json:"id,omitempty"`
	RefType NullableString `json:"refType,omitempty"`
	RefId NullableString `json:"refId,omitempty"`
	Name *string `json:"name,omitempty"`
	InternalId *string `json:"internalId,omitempty"`
	ExternalId *string `json:"externalId,omitempty"`
	UniqueId NullableString `json:"uniqueId,omitempty"`
	PublicIpAddress *string `json:"publicIpAddress,omitempty"`
	PublicIpv6Address NullableString `json:"publicIpv6Address,omitempty"`
	IpAddress *string `json:"ipAddress,omitempty"`
	Ipv6Address *string `json:"ipv6Address,omitempty"`
	IpSubnet NullableString `json:"ipSubnet,omitempty"`
	Ipv6Subnet NullableString `json:"ipv6Subnet,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Dhcp *bool `json:"dhcp,omitempty"`
	Active *bool `json:"active,omitempty"`
	PoolAssigned *bool `json:"poolAssigned,omitempty"`
	PrimaryInterface *bool `json:"primaryInterface,omitempty"`
	Network *InlineResponse20040AppDeployInstance `json:"network,omitempty"`
	Subnet NullableString `json:"subnet,omitempty"`
	NetworkGroup NullableString `json:"networkGroup,omitempty"`
	NetworkPosition NullableString `json:"networkPosition,omitempty"`
	NetworkPool *InlineResponse20040AppDeployInstance `json:"networkPool,omitempty"`
	NetworkDomain NullableString `json:"networkDomain,omitempty"`
	Type *InlineResponse20079LoadBalancerMonitorLoadBalancerType `json:"type,omitempty"`
	IpMode *string `json:"ipMode,omitempty"`
	MacAddress *string `json:"macAddress,omitempty"`
}

// NewGuidanceVmwareSizingResourceInterfaces instantiates a new GuidanceVmwareSizingResourceInterfaces object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGuidanceVmwareSizingResourceInterfaces() *GuidanceVmwareSizingResourceInterfaces {
	this := GuidanceVmwareSizingResourceInterfaces{}
	return &this
}

// NewGuidanceVmwareSizingResourceInterfacesWithDefaults instantiates a new GuidanceVmwareSizingResourceInterfaces object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGuidanceVmwareSizingResourceInterfacesWithDefaults() *GuidanceVmwareSizingResourceInterfaces {
	this := GuidanceVmwareSizingResourceInterfaces{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GuidanceVmwareSizingResourceInterfaces) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceVmwareSizingResourceInterfaces) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GuidanceVmwareSizingResourceInterfaces) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *GuidanceVmwareSizingResourceInterfaces) SetId(v int64) {
	o.Id = &v
}

// GetRefType returns the RefType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuidanceVmwareSizingResourceInterfaces) GetRefType() string {
	if o == nil || o.RefType.Get() == nil {
		var ret string
		return ret
	}
	return *o.RefType.Get()
}

// GetRefTypeOk returns a tuple with the RefType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuidanceVmwareSizingResourceInterfaces) GetRefTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RefType.Get(), o.RefType.IsSet()
}

// HasRefType returns a boolean if a field has been set.
func (o *GuidanceVmwareSizingResourceInterfaces) HasRefType() bool {
	if o != nil && o.RefType.IsSet() {
		return true
	}

	return false
}

// SetRefType gets a reference to the given NullableString and assigns it to the RefType field.
func (o *GuidanceVmwareSizingResourceInterfaces) SetRefType(v string) {
	o.RefType.Set(&v)
}
// SetRefTypeNil sets the value for RefType to be an explicit nil
func (o *GuidanceVmwareSizingResourceInterfaces) SetRefTypeNil() {
	o.RefType.Set(nil)
}

// UnsetRefType ensures that no value is present for RefType, not even an explicit nil
func (o *GuidanceVmwareSizingResourceInterfaces) UnsetRefType() {
	o.RefType.Unset()
}

// GetRefId returns the RefId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuidanceVmwareSizingResourceInterfaces) GetRefId() string {
	if o == nil || o.RefId.Get() == nil {
		var ret string
		return ret
	}
	return *o.RefId.Get()
}

// GetRefIdOk returns a tuple with the RefId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuidanceVmwareSizingResourceInterfaces) GetRefIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RefId.Get(), o.RefId.IsSet()
}

// HasRefId returns a boolean if a field has been set.
func (o *GuidanceVmwareSizingResourceInterfaces) HasRefId() bool {
	if o != nil && o.RefId.IsSet() {
		return true
	}

	return false
}

// SetRefId gets a reference to the given NullableString and assigns it to the RefId field.
func (o *GuidanceVmwareSizingResourceInterfaces) SetRefId(v string) {
	o.RefId.Set(&v)
}
// SetRefIdNil sets the value for RefId to be an explicit nil
func (o *GuidanceVmwareSizingResourceInterfaces) SetRefIdNil() {
	o.RefId.Set(nil)
}

// UnsetRefId ensures that no value is present for RefId, not even an explicit nil
func (o *GuidanceVmwareSizingResourceInterfaces) UnsetRefId() {
	o.RefId.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GuidanceVmwareSizingResourceInterfaces) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceVmwareSizingResourceInterfaces) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GuidanceVmwareSizingResourceInterfaces) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GuidanceVmwareSizingResourceInterfaces) SetName(v string) {
	o.Name = &v
}

// GetInternalId returns the InternalId field value if set, zero value otherwise.
func (o *GuidanceVmwareSizingResourceInterfaces) GetInternalId() string {
	if o == nil || o.InternalId == nil {
		var ret string
		return ret
	}
	return *o.InternalId
}

// GetInternalIdOk returns a tuple with the InternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceVmwareSizingResourceInterfaces) GetInternalIdOk() (*string, bool) {
	if o == nil || o.InternalId == nil {
		return nil, false
	}
	return o.InternalId, true
}

// HasInternalId returns a boolean if a field has been set.
func (o *GuidanceVmwareSizingResourceInterfaces) HasInternalId() bool {
	if o != nil && o.InternalId != nil {
		return true
	}

	return false
}

// SetInternalId gets a reference to the given string and assigns it to the InternalId field.
func (o *GuidanceVmwareSizingResourceInterfaces) SetInternalId(v string) {
	o.InternalId = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *GuidanceVmwareSizingResourceInterfaces) GetExternalId() string {
	if o == nil || o.ExternalId == nil {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceVmwareSizingResourceInterfaces) GetExternalIdOk() (*string, bool) {
	if o == nil || o.ExternalId == nil {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *GuidanceVmwareSizingResourceInterfaces) HasExternalId() bool {
	if o != nil && o.ExternalId != nil {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *GuidanceVmwareSizingResourceInterfaces) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetUniqueId returns the UniqueId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuidanceVmwareSizingResourceInterfaces) GetUniqueId() string {
	if o == nil || o.UniqueId.Get() == nil {
		var ret string
		return ret
	}
	return *o.UniqueId.Get()
}

// GetUniqueIdOk returns a tuple with the UniqueId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuidanceVmwareSizingResourceInterfaces) GetUniqueIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UniqueId.Get(), o.UniqueId.IsSet()
}

// HasUniqueId returns a boolean if a field has been set.
func (o *GuidanceVmwareSizingResourceInterfaces) HasUniqueId() bool {
	if o != nil && o.UniqueId.IsSet() {
		return true
	}

	return false
}

// SetUniqueId gets a reference to the given NullableString and assigns it to the UniqueId field.
func (o *GuidanceVmwareSizingResourceInterfaces) SetUniqueId(v string) {
	o.UniqueId.Set(&v)
}
// SetUniqueIdNil sets the value for UniqueId to be an explicit nil
func (o *GuidanceVmwareSizingResourceInterfaces) SetUniqueIdNil() {
	o.UniqueId.Set(nil)
}

// UnsetUniqueId ensures that no value is present for UniqueId, not even an explicit nil
func (o *GuidanceVmwareSizingResourceInterfaces) UnsetUniqueId() {
	o.UniqueId.Unset()
}

// GetPublicIpAddress returns the PublicIpAddress field value if set, zero value otherwise.
func (o *GuidanceVmwareSizingResourceInterfaces) GetPublicIpAddress() string {
	if o == nil || o.PublicIpAddress == nil {
		var ret string
		return ret
	}
	return *o.PublicIpAddress
}

// GetPublicIpAddressOk returns a tuple with the PublicIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceVmwareSizingResourceInterfaces) GetPublicIpAddressOk() (*string, bool) {
	if o == nil || o.PublicIpAddress == nil {
		return nil, false
	}
	return o.PublicIpAddress, true
}

// HasPublicIpAddress returns a boolean if a field has been set.
func (o *GuidanceVmwareSizingResourceInterfaces) HasPublicIpAddress() bool {
	if o != nil && o.PublicIpAddress != nil {
		return true
	}

	return false
}

// SetPublicIpAddress gets a reference to the given string and assigns it to the PublicIpAddress field.
func (o *GuidanceVmwareSizingResourceInterfaces) SetPublicIpAddress(v string) {
	o.PublicIpAddress = &v
}

// GetPublicIpv6Address returns the PublicIpv6Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuidanceVmwareSizingResourceInterfaces) GetPublicIpv6Address() string {
	if o == nil || o.PublicIpv6Address.Get() == nil {
		var ret string
		return ret
	}
	return *o.PublicIpv6Address.Get()
}

// GetPublicIpv6AddressOk returns a tuple with the PublicIpv6Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuidanceVmwareSizingResourceInterfaces) GetPublicIpv6AddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PublicIpv6Address.Get(), o.PublicIpv6Address.IsSet()
}

// HasPublicIpv6Address returns a boolean if a field has been set.
func (o *GuidanceVmwareSizingResourceInterfaces) HasPublicIpv6Address() bool {
	if o != nil && o.PublicIpv6Address.IsSet() {
		return true
	}

	return false
}

// SetPublicIpv6Address gets a reference to the given NullableString and assigns it to the PublicIpv6Address field.
func (o *GuidanceVmwareSizingResourceInterfaces) SetPublicIpv6Address(v string) {
	o.PublicIpv6Address.Set(&v)
}
// SetPublicIpv6AddressNil sets the value for PublicIpv6Address to be an explicit nil
func (o *GuidanceVmwareSizingResourceInterfaces) SetPublicIpv6AddressNil() {
	o.PublicIpv6Address.Set(nil)
}

// UnsetPublicIpv6Address ensures that no value is present for PublicIpv6Address, not even an explicit nil
func (o *GuidanceVmwareSizingResourceInterfaces) UnsetPublicIpv6Address() {
	o.PublicIpv6Address.Unset()
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *GuidanceVmwareSizingResourceInterfaces) GetIpAddress() string {
	if o == nil || o.IpAddress == nil {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceVmwareSizingResourceInterfaces) GetIpAddressOk() (*string, bool) {
	if o == nil || o.IpAddress == nil {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *GuidanceVmwareSizingResourceInterfaces) HasIpAddress() bool {
	if o != nil && o.IpAddress != nil {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *GuidanceVmwareSizingResourceInterfaces) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetIpv6Address returns the Ipv6Address field value if set, zero value otherwise.
func (o *GuidanceVmwareSizingResourceInterfaces) GetIpv6Address() string {
	if o == nil || o.Ipv6Address == nil {
		var ret string
		return ret
	}
	return *o.Ipv6Address
}

// GetIpv6AddressOk returns a tuple with the Ipv6Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceVmwareSizingResourceInterfaces) GetIpv6AddressOk() (*string, bool) {
	if o == nil || o.Ipv6Address == nil {
		return nil, false
	}
	return o.Ipv6Address, true
}

// HasIpv6Address returns a boolean if a field has been set.
func (o *GuidanceVmwareSizingResourceInterfaces) HasIpv6Address() bool {
	if o != nil && o.Ipv6Address != nil {
		return true
	}

	return false
}

// SetIpv6Address gets a reference to the given string and assigns it to the Ipv6Address field.
func (o *GuidanceVmwareSizingResourceInterfaces) SetIpv6Address(v string) {
	o.Ipv6Address = &v
}

// GetIpSubnet returns the IpSubnet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuidanceVmwareSizingResourceInterfaces) GetIpSubnet() string {
	if o == nil || o.IpSubnet.Get() == nil {
		var ret string
		return ret
	}
	return *o.IpSubnet.Get()
}

// GetIpSubnetOk returns a tuple with the IpSubnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuidanceVmwareSizingResourceInterfaces) GetIpSubnetOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IpSubnet.Get(), o.IpSubnet.IsSet()
}

// HasIpSubnet returns a boolean if a field has been set.
func (o *GuidanceVmwareSizingResourceInterfaces) HasIpSubnet() bool {
	if o != nil && o.IpSubnet.IsSet() {
		return true
	}

	return false
}

// SetIpSubnet gets a reference to the given NullableString and assigns it to the IpSubnet field.
func (o *GuidanceVmwareSizingResourceInterfaces) SetIpSubnet(v string) {
	o.IpSubnet.Set(&v)
}
// SetIpSubnetNil sets the value for IpSubnet to be an explicit nil
func (o *GuidanceVmwareSizingResourceInterfaces) SetIpSubnetNil() {
	o.IpSubnet.Set(nil)
}

// UnsetIpSubnet ensures that no value is present for IpSubnet, not even an explicit nil
func (o *GuidanceVmwareSizingResourceInterfaces) UnsetIpSubnet() {
	o.IpSubnet.Unset()
}

// GetIpv6Subnet returns the Ipv6Subnet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuidanceVmwareSizingResourceInterfaces) GetIpv6Subnet() string {
	if o == nil || o.Ipv6Subnet.Get() == nil {
		var ret string
		return ret
	}
	return *o.Ipv6Subnet.Get()
}

// GetIpv6SubnetOk returns a tuple with the Ipv6Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuidanceVmwareSizingResourceInterfaces) GetIpv6SubnetOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Ipv6Subnet.Get(), o.Ipv6Subnet.IsSet()
}

// HasIpv6Subnet returns a boolean if a field has been set.
func (o *GuidanceVmwareSizingResourceInterfaces) HasIpv6Subnet() bool {
	if o != nil && o.Ipv6Subnet.IsSet() {
		return true
	}

	return false
}

// SetIpv6Subnet gets a reference to the given NullableString and assigns it to the Ipv6Subnet field.
func (o *GuidanceVmwareSizingResourceInterfaces) SetIpv6Subnet(v string) {
	o.Ipv6Subnet.Set(&v)
}
// SetIpv6SubnetNil sets the value for Ipv6Subnet to be an explicit nil
func (o *GuidanceVmwareSizingResourceInterfaces) SetIpv6SubnetNil() {
	o.Ipv6Subnet.Set(nil)
}

// UnsetIpv6Subnet ensures that no value is present for Ipv6Subnet, not even an explicit nil
func (o *GuidanceVmwareSizingResourceInterfaces) UnsetIpv6Subnet() {
	o.Ipv6Subnet.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuidanceVmwareSizingResourceInterfaces) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuidanceVmwareSizingResourceInterfaces) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *GuidanceVmwareSizingResourceInterfaces) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *GuidanceVmwareSizingResourceInterfaces) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *GuidanceVmwareSizingResourceInterfaces) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *GuidanceVmwareSizingResourceInterfaces) UnsetDescription() {
	o.Description.Unset()
}

// GetDhcp returns the Dhcp field value if set, zero value otherwise.
func (o *GuidanceVmwareSizingResourceInterfaces) GetDhcp() bool {
	if o == nil || o.Dhcp == nil {
		var ret bool
		return ret
	}
	return *o.Dhcp
}

// GetDhcpOk returns a tuple with the Dhcp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceVmwareSizingResourceInterfaces) GetDhcpOk() (*bool, bool) {
	if o == nil || o.Dhcp == nil {
		return nil, false
	}
	return o.Dhcp, true
}

// HasDhcp returns a boolean if a field has been set.
func (o *GuidanceVmwareSizingResourceInterfaces) HasDhcp() bool {
	if o != nil && o.Dhcp != nil {
		return true
	}

	return false
}

// SetDhcp gets a reference to the given bool and assigns it to the Dhcp field.
func (o *GuidanceVmwareSizingResourceInterfaces) SetDhcp(v bool) {
	o.Dhcp = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *GuidanceVmwareSizingResourceInterfaces) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceVmwareSizingResourceInterfaces) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *GuidanceVmwareSizingResourceInterfaces) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *GuidanceVmwareSizingResourceInterfaces) SetActive(v bool) {
	o.Active = &v
}

// GetPoolAssigned returns the PoolAssigned field value if set, zero value otherwise.
func (o *GuidanceVmwareSizingResourceInterfaces) GetPoolAssigned() bool {
	if o == nil || o.PoolAssigned == nil {
		var ret bool
		return ret
	}
	return *o.PoolAssigned
}

// GetPoolAssignedOk returns a tuple with the PoolAssigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceVmwareSizingResourceInterfaces) GetPoolAssignedOk() (*bool, bool) {
	if o == nil || o.PoolAssigned == nil {
		return nil, false
	}
	return o.PoolAssigned, true
}

// HasPoolAssigned returns a boolean if a field has been set.
func (o *GuidanceVmwareSizingResourceInterfaces) HasPoolAssigned() bool {
	if o != nil && o.PoolAssigned != nil {
		return true
	}

	return false
}

// SetPoolAssigned gets a reference to the given bool and assigns it to the PoolAssigned field.
func (o *GuidanceVmwareSizingResourceInterfaces) SetPoolAssigned(v bool) {
	o.PoolAssigned = &v
}

// GetPrimaryInterface returns the PrimaryInterface field value if set, zero value otherwise.
func (o *GuidanceVmwareSizingResourceInterfaces) GetPrimaryInterface() bool {
	if o == nil || o.PrimaryInterface == nil {
		var ret bool
		return ret
	}
	return *o.PrimaryInterface
}

// GetPrimaryInterfaceOk returns a tuple with the PrimaryInterface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceVmwareSizingResourceInterfaces) GetPrimaryInterfaceOk() (*bool, bool) {
	if o == nil || o.PrimaryInterface == nil {
		return nil, false
	}
	return o.PrimaryInterface, true
}

// HasPrimaryInterface returns a boolean if a field has been set.
func (o *GuidanceVmwareSizingResourceInterfaces) HasPrimaryInterface() bool {
	if o != nil && o.PrimaryInterface != nil {
		return true
	}

	return false
}

// SetPrimaryInterface gets a reference to the given bool and assigns it to the PrimaryInterface field.
func (o *GuidanceVmwareSizingResourceInterfaces) SetPrimaryInterface(v bool) {
	o.PrimaryInterface = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *GuidanceVmwareSizingResourceInterfaces) GetNetwork() InlineResponse20040AppDeployInstance {
	if o == nil || o.Network == nil {
		var ret InlineResponse20040AppDeployInstance
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceVmwareSizingResourceInterfaces) GetNetworkOk() (*InlineResponse20040AppDeployInstance, bool) {
	if o == nil || o.Network == nil {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *GuidanceVmwareSizingResourceInterfaces) HasNetwork() bool {
	if o != nil && o.Network != nil {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given InlineResponse20040AppDeployInstance and assigns it to the Network field.
func (o *GuidanceVmwareSizingResourceInterfaces) SetNetwork(v InlineResponse20040AppDeployInstance) {
	o.Network = &v
}

// GetSubnet returns the Subnet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuidanceVmwareSizingResourceInterfaces) GetSubnet() string {
	if o == nil || o.Subnet.Get() == nil {
		var ret string
		return ret
	}
	return *o.Subnet.Get()
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuidanceVmwareSizingResourceInterfaces) GetSubnetOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Subnet.Get(), o.Subnet.IsSet()
}

// HasSubnet returns a boolean if a field has been set.
func (o *GuidanceVmwareSizingResourceInterfaces) HasSubnet() bool {
	if o != nil && o.Subnet.IsSet() {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given NullableString and assigns it to the Subnet field.
func (o *GuidanceVmwareSizingResourceInterfaces) SetSubnet(v string) {
	o.Subnet.Set(&v)
}
// SetSubnetNil sets the value for Subnet to be an explicit nil
func (o *GuidanceVmwareSizingResourceInterfaces) SetSubnetNil() {
	o.Subnet.Set(nil)
}

// UnsetSubnet ensures that no value is present for Subnet, not even an explicit nil
func (o *GuidanceVmwareSizingResourceInterfaces) UnsetSubnet() {
	o.Subnet.Unset()
}

// GetNetworkGroup returns the NetworkGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuidanceVmwareSizingResourceInterfaces) GetNetworkGroup() string {
	if o == nil || o.NetworkGroup.Get() == nil {
		var ret string
		return ret
	}
	return *o.NetworkGroup.Get()
}

// GetNetworkGroupOk returns a tuple with the NetworkGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuidanceVmwareSizingResourceInterfaces) GetNetworkGroupOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NetworkGroup.Get(), o.NetworkGroup.IsSet()
}

// HasNetworkGroup returns a boolean if a field has been set.
func (o *GuidanceVmwareSizingResourceInterfaces) HasNetworkGroup() bool {
	if o != nil && o.NetworkGroup.IsSet() {
		return true
	}

	return false
}

// SetNetworkGroup gets a reference to the given NullableString and assigns it to the NetworkGroup field.
func (o *GuidanceVmwareSizingResourceInterfaces) SetNetworkGroup(v string) {
	o.NetworkGroup.Set(&v)
}
// SetNetworkGroupNil sets the value for NetworkGroup to be an explicit nil
func (o *GuidanceVmwareSizingResourceInterfaces) SetNetworkGroupNil() {
	o.NetworkGroup.Set(nil)
}

// UnsetNetworkGroup ensures that no value is present for NetworkGroup, not even an explicit nil
func (o *GuidanceVmwareSizingResourceInterfaces) UnsetNetworkGroup() {
	o.NetworkGroup.Unset()
}

// GetNetworkPosition returns the NetworkPosition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuidanceVmwareSizingResourceInterfaces) GetNetworkPosition() string {
	if o == nil || o.NetworkPosition.Get() == nil {
		var ret string
		return ret
	}
	return *o.NetworkPosition.Get()
}

// GetNetworkPositionOk returns a tuple with the NetworkPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuidanceVmwareSizingResourceInterfaces) GetNetworkPositionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NetworkPosition.Get(), o.NetworkPosition.IsSet()
}

// HasNetworkPosition returns a boolean if a field has been set.
func (o *GuidanceVmwareSizingResourceInterfaces) HasNetworkPosition() bool {
	if o != nil && o.NetworkPosition.IsSet() {
		return true
	}

	return false
}

// SetNetworkPosition gets a reference to the given NullableString and assigns it to the NetworkPosition field.
func (o *GuidanceVmwareSizingResourceInterfaces) SetNetworkPosition(v string) {
	o.NetworkPosition.Set(&v)
}
// SetNetworkPositionNil sets the value for NetworkPosition to be an explicit nil
func (o *GuidanceVmwareSizingResourceInterfaces) SetNetworkPositionNil() {
	o.NetworkPosition.Set(nil)
}

// UnsetNetworkPosition ensures that no value is present for NetworkPosition, not even an explicit nil
func (o *GuidanceVmwareSizingResourceInterfaces) UnsetNetworkPosition() {
	o.NetworkPosition.Unset()
}

// GetNetworkPool returns the NetworkPool field value if set, zero value otherwise.
func (o *GuidanceVmwareSizingResourceInterfaces) GetNetworkPool() InlineResponse20040AppDeployInstance {
	if o == nil || o.NetworkPool == nil {
		var ret InlineResponse20040AppDeployInstance
		return ret
	}
	return *o.NetworkPool
}

// GetNetworkPoolOk returns a tuple with the NetworkPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceVmwareSizingResourceInterfaces) GetNetworkPoolOk() (*InlineResponse20040AppDeployInstance, bool) {
	if o == nil || o.NetworkPool == nil {
		return nil, false
	}
	return o.NetworkPool, true
}

// HasNetworkPool returns a boolean if a field has been set.
func (o *GuidanceVmwareSizingResourceInterfaces) HasNetworkPool() bool {
	if o != nil && o.NetworkPool != nil {
		return true
	}

	return false
}

// SetNetworkPool gets a reference to the given InlineResponse20040AppDeployInstance and assigns it to the NetworkPool field.
func (o *GuidanceVmwareSizingResourceInterfaces) SetNetworkPool(v InlineResponse20040AppDeployInstance) {
	o.NetworkPool = &v
}

// GetNetworkDomain returns the NetworkDomain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GuidanceVmwareSizingResourceInterfaces) GetNetworkDomain() string {
	if o == nil || o.NetworkDomain.Get() == nil {
		var ret string
		return ret
	}
	return *o.NetworkDomain.Get()
}

// GetNetworkDomainOk returns a tuple with the NetworkDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GuidanceVmwareSizingResourceInterfaces) GetNetworkDomainOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NetworkDomain.Get(), o.NetworkDomain.IsSet()
}

// HasNetworkDomain returns a boolean if a field has been set.
func (o *GuidanceVmwareSizingResourceInterfaces) HasNetworkDomain() bool {
	if o != nil && o.NetworkDomain.IsSet() {
		return true
	}

	return false
}

// SetNetworkDomain gets a reference to the given NullableString and assigns it to the NetworkDomain field.
func (o *GuidanceVmwareSizingResourceInterfaces) SetNetworkDomain(v string) {
	o.NetworkDomain.Set(&v)
}
// SetNetworkDomainNil sets the value for NetworkDomain to be an explicit nil
func (o *GuidanceVmwareSizingResourceInterfaces) SetNetworkDomainNil() {
	o.NetworkDomain.Set(nil)
}

// UnsetNetworkDomain ensures that no value is present for NetworkDomain, not even an explicit nil
func (o *GuidanceVmwareSizingResourceInterfaces) UnsetNetworkDomain() {
	o.NetworkDomain.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GuidanceVmwareSizingResourceInterfaces) GetType() InlineResponse20079LoadBalancerMonitorLoadBalancerType {
	if o == nil || o.Type == nil {
		var ret InlineResponse20079LoadBalancerMonitorLoadBalancerType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceVmwareSizingResourceInterfaces) GetTypeOk() (*InlineResponse20079LoadBalancerMonitorLoadBalancerType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GuidanceVmwareSizingResourceInterfaces) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given InlineResponse20079LoadBalancerMonitorLoadBalancerType and assigns it to the Type field.
func (o *GuidanceVmwareSizingResourceInterfaces) SetType(v InlineResponse20079LoadBalancerMonitorLoadBalancerType) {
	o.Type = &v
}

// GetIpMode returns the IpMode field value if set, zero value otherwise.
func (o *GuidanceVmwareSizingResourceInterfaces) GetIpMode() string {
	if o == nil || o.IpMode == nil {
		var ret string
		return ret
	}
	return *o.IpMode
}

// GetIpModeOk returns a tuple with the IpMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceVmwareSizingResourceInterfaces) GetIpModeOk() (*string, bool) {
	if o == nil || o.IpMode == nil {
		return nil, false
	}
	return o.IpMode, true
}

// HasIpMode returns a boolean if a field has been set.
func (o *GuidanceVmwareSizingResourceInterfaces) HasIpMode() bool {
	if o != nil && o.IpMode != nil {
		return true
	}

	return false
}

// SetIpMode gets a reference to the given string and assigns it to the IpMode field.
func (o *GuidanceVmwareSizingResourceInterfaces) SetIpMode(v string) {
	o.IpMode = &v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *GuidanceVmwareSizingResourceInterfaces) GetMacAddress() string {
	if o == nil || o.MacAddress == nil {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GuidanceVmwareSizingResourceInterfaces) GetMacAddressOk() (*string, bool) {
	if o == nil || o.MacAddress == nil {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *GuidanceVmwareSizingResourceInterfaces) HasMacAddress() bool {
	if o != nil && o.MacAddress != nil {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *GuidanceVmwareSizingResourceInterfaces) SetMacAddress(v string) {
	o.MacAddress = &v
}

func (o GuidanceVmwareSizingResourceInterfaces) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.RefType.IsSet() {
		toSerialize["refType"] = o.RefType.Get()
	}
	if o.RefId.IsSet() {
		toSerialize["refId"] = o.RefId.Get()
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.InternalId != nil {
		toSerialize["internalId"] = o.InternalId
	}
	if o.ExternalId != nil {
		toSerialize["externalId"] = o.ExternalId
	}
	if o.UniqueId.IsSet() {
		toSerialize["uniqueId"] = o.UniqueId.Get()
	}
	if o.PublicIpAddress != nil {
		toSerialize["publicIpAddress"] = o.PublicIpAddress
	}
	if o.PublicIpv6Address.IsSet() {
		toSerialize["publicIpv6Address"] = o.PublicIpv6Address.Get()
	}
	if o.IpAddress != nil {
		toSerialize["ipAddress"] = o.IpAddress
	}
	if o.Ipv6Address != nil {
		toSerialize["ipv6Address"] = o.Ipv6Address
	}
	if o.IpSubnet.IsSet() {
		toSerialize["ipSubnet"] = o.IpSubnet.Get()
	}
	if o.Ipv6Subnet.IsSet() {
		toSerialize["ipv6Subnet"] = o.Ipv6Subnet.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Dhcp != nil {
		toSerialize["dhcp"] = o.Dhcp
	}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.PoolAssigned != nil {
		toSerialize["poolAssigned"] = o.PoolAssigned
	}
	if o.PrimaryInterface != nil {
		toSerialize["primaryInterface"] = o.PrimaryInterface
	}
	if o.Network != nil {
		toSerialize["network"] = o.Network
	}
	if o.Subnet.IsSet() {
		toSerialize["subnet"] = o.Subnet.Get()
	}
	if o.NetworkGroup.IsSet() {
		toSerialize["networkGroup"] = o.NetworkGroup.Get()
	}
	if o.NetworkPosition.IsSet() {
		toSerialize["networkPosition"] = o.NetworkPosition.Get()
	}
	if o.NetworkPool != nil {
		toSerialize["networkPool"] = o.NetworkPool
	}
	if o.NetworkDomain.IsSet() {
		toSerialize["networkDomain"] = o.NetworkDomain.Get()
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.IpMode != nil {
		toSerialize["ipMode"] = o.IpMode
	}
	if o.MacAddress != nil {
		toSerialize["macAddress"] = o.MacAddress
	}
	return json.Marshal(toSerialize)
}

type NullableGuidanceVmwareSizingResourceInterfaces struct {
	value *GuidanceVmwareSizingResourceInterfaces
	isSet bool
}

func (v NullableGuidanceVmwareSizingResourceInterfaces) Get() *GuidanceVmwareSizingResourceInterfaces {
	return v.value
}

func (v *NullableGuidanceVmwareSizingResourceInterfaces) Set(val *GuidanceVmwareSizingResourceInterfaces) {
	v.value = val
	v.isSet = true
}

func (v NullableGuidanceVmwareSizingResourceInterfaces) IsSet() bool {
	return v.isSet
}

func (v *NullableGuidanceVmwareSizingResourceInterfaces) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGuidanceVmwareSizingResourceInterfaces(val *GuidanceVmwareSizingResourceInterfaces) *NullableGuidanceVmwareSizingResourceInterfaces {
	return &NullableGuidanceVmwareSizingResourceInterfaces{value: val, isSet: true}
}

func (v NullableGuidanceVmwareSizingResourceInterfaces) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGuidanceVmwareSizingResourceInterfaces) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


