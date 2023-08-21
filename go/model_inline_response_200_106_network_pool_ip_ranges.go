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
	"time"
)

// InlineResponse200106NetworkPoolIpRanges struct for InlineResponse200106NetworkPoolIpRanges
type InlineResponse200106NetworkPoolIpRanges struct {
	Id *int64 `json:"id,omitempty"`
	StartAddress NullableString `json:"startAddress,omitempty"`
	EndAddress NullableString `json:"endAddress,omitempty"`
	InternalId NullableString `json:"internalId,omitempty"`
	ExternalId NullableString `json:"externalId,omitempty"`
	Description NullableString `json:"description,omitempty"`
	AddressCount *int64 `json:"addressCount,omitempty"`
	Active *bool `json:"active,omitempty"`
	DateCreated *time.Time `json:"dateCreated,omitempty"`
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	Cidr NullableString `json:"cidr,omitempty"`
	CidrIPv6 NullableString `json:"cidrIPv6,omitempty"`
}

// NewInlineResponse200106NetworkPoolIpRanges instantiates a new InlineResponse200106NetworkPoolIpRanges object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200106NetworkPoolIpRanges() *InlineResponse200106NetworkPoolIpRanges {
	this := InlineResponse200106NetworkPoolIpRanges{}
	return &this
}

// NewInlineResponse200106NetworkPoolIpRangesWithDefaults instantiates a new InlineResponse200106NetworkPoolIpRanges object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200106NetworkPoolIpRangesWithDefaults() *InlineResponse200106NetworkPoolIpRanges {
	this := InlineResponse200106NetworkPoolIpRanges{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse200106NetworkPoolIpRanges) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200106NetworkPoolIpRanges) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse200106NetworkPoolIpRanges) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *InlineResponse200106NetworkPoolIpRanges) SetId(v int64) {
	o.Id = &v
}

// GetStartAddress returns the StartAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineResponse200106NetworkPoolIpRanges) GetStartAddress() string {
	if o == nil || o.StartAddress.Get() == nil {
		var ret string
		return ret
	}
	return *o.StartAddress.Get()
}

// GetStartAddressOk returns a tuple with the StartAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineResponse200106NetworkPoolIpRanges) GetStartAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StartAddress.Get(), o.StartAddress.IsSet()
}

// HasStartAddress returns a boolean if a field has been set.
func (o *InlineResponse200106NetworkPoolIpRanges) HasStartAddress() bool {
	if o != nil && o.StartAddress.IsSet() {
		return true
	}

	return false
}

// SetStartAddress gets a reference to the given NullableString and assigns it to the StartAddress field.
func (o *InlineResponse200106NetworkPoolIpRanges) SetStartAddress(v string) {
	o.StartAddress.Set(&v)
}
// SetStartAddressNil sets the value for StartAddress to be an explicit nil
func (o *InlineResponse200106NetworkPoolIpRanges) SetStartAddressNil() {
	o.StartAddress.Set(nil)
}

// UnsetStartAddress ensures that no value is present for StartAddress, not even an explicit nil
func (o *InlineResponse200106NetworkPoolIpRanges) UnsetStartAddress() {
	o.StartAddress.Unset()
}

// GetEndAddress returns the EndAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineResponse200106NetworkPoolIpRanges) GetEndAddress() string {
	if o == nil || o.EndAddress.Get() == nil {
		var ret string
		return ret
	}
	return *o.EndAddress.Get()
}

// GetEndAddressOk returns a tuple with the EndAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineResponse200106NetworkPoolIpRanges) GetEndAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndAddress.Get(), o.EndAddress.IsSet()
}

// HasEndAddress returns a boolean if a field has been set.
func (o *InlineResponse200106NetworkPoolIpRanges) HasEndAddress() bool {
	if o != nil && o.EndAddress.IsSet() {
		return true
	}

	return false
}

// SetEndAddress gets a reference to the given NullableString and assigns it to the EndAddress field.
func (o *InlineResponse200106NetworkPoolIpRanges) SetEndAddress(v string) {
	o.EndAddress.Set(&v)
}
// SetEndAddressNil sets the value for EndAddress to be an explicit nil
func (o *InlineResponse200106NetworkPoolIpRanges) SetEndAddressNil() {
	o.EndAddress.Set(nil)
}

// UnsetEndAddress ensures that no value is present for EndAddress, not even an explicit nil
func (o *InlineResponse200106NetworkPoolIpRanges) UnsetEndAddress() {
	o.EndAddress.Unset()
}

// GetInternalId returns the InternalId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineResponse200106NetworkPoolIpRanges) GetInternalId() string {
	if o == nil || o.InternalId.Get() == nil {
		var ret string
		return ret
	}
	return *o.InternalId.Get()
}

// GetInternalIdOk returns a tuple with the InternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineResponse200106NetworkPoolIpRanges) GetInternalIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.InternalId.Get(), o.InternalId.IsSet()
}

// HasInternalId returns a boolean if a field has been set.
func (o *InlineResponse200106NetworkPoolIpRanges) HasInternalId() bool {
	if o != nil && o.InternalId.IsSet() {
		return true
	}

	return false
}

// SetInternalId gets a reference to the given NullableString and assigns it to the InternalId field.
func (o *InlineResponse200106NetworkPoolIpRanges) SetInternalId(v string) {
	o.InternalId.Set(&v)
}
// SetInternalIdNil sets the value for InternalId to be an explicit nil
func (o *InlineResponse200106NetworkPoolIpRanges) SetInternalIdNil() {
	o.InternalId.Set(nil)
}

// UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
func (o *InlineResponse200106NetworkPoolIpRanges) UnsetInternalId() {
	o.InternalId.Unset()
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineResponse200106NetworkPoolIpRanges) GetExternalId() string {
	if o == nil || o.ExternalId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExternalId.Get()
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineResponse200106NetworkPoolIpRanges) GetExternalIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExternalId.Get(), o.ExternalId.IsSet()
}

// HasExternalId returns a boolean if a field has been set.
func (o *InlineResponse200106NetworkPoolIpRanges) HasExternalId() bool {
	if o != nil && o.ExternalId.IsSet() {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given NullableString and assigns it to the ExternalId field.
func (o *InlineResponse200106NetworkPoolIpRanges) SetExternalId(v string) {
	o.ExternalId.Set(&v)
}
// SetExternalIdNil sets the value for ExternalId to be an explicit nil
func (o *InlineResponse200106NetworkPoolIpRanges) SetExternalIdNil() {
	o.ExternalId.Set(nil)
}

// UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
func (o *InlineResponse200106NetworkPoolIpRanges) UnsetExternalId() {
	o.ExternalId.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineResponse200106NetworkPoolIpRanges) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineResponse200106NetworkPoolIpRanges) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *InlineResponse200106NetworkPoolIpRanges) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *InlineResponse200106NetworkPoolIpRanges) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *InlineResponse200106NetworkPoolIpRanges) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *InlineResponse200106NetworkPoolIpRanges) UnsetDescription() {
	o.Description.Unset()
}

// GetAddressCount returns the AddressCount field value if set, zero value otherwise.
func (o *InlineResponse200106NetworkPoolIpRanges) GetAddressCount() int64 {
	if o == nil || o.AddressCount == nil {
		var ret int64
		return ret
	}
	return *o.AddressCount
}

// GetAddressCountOk returns a tuple with the AddressCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200106NetworkPoolIpRanges) GetAddressCountOk() (*int64, bool) {
	if o == nil || o.AddressCount == nil {
		return nil, false
	}
	return o.AddressCount, true
}

// HasAddressCount returns a boolean if a field has been set.
func (o *InlineResponse200106NetworkPoolIpRanges) HasAddressCount() bool {
	if o != nil && o.AddressCount != nil {
		return true
	}

	return false
}

// SetAddressCount gets a reference to the given int64 and assigns it to the AddressCount field.
func (o *InlineResponse200106NetworkPoolIpRanges) SetAddressCount(v int64) {
	o.AddressCount = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *InlineResponse200106NetworkPoolIpRanges) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200106NetworkPoolIpRanges) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *InlineResponse200106NetworkPoolIpRanges) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *InlineResponse200106NetworkPoolIpRanges) SetActive(v bool) {
	o.Active = &v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *InlineResponse200106NetworkPoolIpRanges) GetDateCreated() time.Time {
	if o == nil || o.DateCreated == nil {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200106NetworkPoolIpRanges) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || o.DateCreated == nil {
		return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *InlineResponse200106NetworkPoolIpRanges) HasDateCreated() bool {
	if o != nil && o.DateCreated != nil {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *InlineResponse200106NetworkPoolIpRanges) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *InlineResponse200106NetworkPoolIpRanges) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200106NetworkPoolIpRanges) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *InlineResponse200106NetworkPoolIpRanges) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *InlineResponse200106NetworkPoolIpRanges) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetCidr returns the Cidr field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineResponse200106NetworkPoolIpRanges) GetCidr() string {
	if o == nil || o.Cidr.Get() == nil {
		var ret string
		return ret
	}
	return *o.Cidr.Get()
}

// GetCidrOk returns a tuple with the Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineResponse200106NetworkPoolIpRanges) GetCidrOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Cidr.Get(), o.Cidr.IsSet()
}

// HasCidr returns a boolean if a field has been set.
func (o *InlineResponse200106NetworkPoolIpRanges) HasCidr() bool {
	if o != nil && o.Cidr.IsSet() {
		return true
	}

	return false
}

// SetCidr gets a reference to the given NullableString and assigns it to the Cidr field.
func (o *InlineResponse200106NetworkPoolIpRanges) SetCidr(v string) {
	o.Cidr.Set(&v)
}
// SetCidrNil sets the value for Cidr to be an explicit nil
func (o *InlineResponse200106NetworkPoolIpRanges) SetCidrNil() {
	o.Cidr.Set(nil)
}

// UnsetCidr ensures that no value is present for Cidr, not even an explicit nil
func (o *InlineResponse200106NetworkPoolIpRanges) UnsetCidr() {
	o.Cidr.Unset()
}

// GetCidrIPv6 returns the CidrIPv6 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineResponse200106NetworkPoolIpRanges) GetCidrIPv6() string {
	if o == nil || o.CidrIPv6.Get() == nil {
		var ret string
		return ret
	}
	return *o.CidrIPv6.Get()
}

// GetCidrIPv6Ok returns a tuple with the CidrIPv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineResponse200106NetworkPoolIpRanges) GetCidrIPv6Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CidrIPv6.Get(), o.CidrIPv6.IsSet()
}

// HasCidrIPv6 returns a boolean if a field has been set.
func (o *InlineResponse200106NetworkPoolIpRanges) HasCidrIPv6() bool {
	if o != nil && o.CidrIPv6.IsSet() {
		return true
	}

	return false
}

// SetCidrIPv6 gets a reference to the given NullableString and assigns it to the CidrIPv6 field.
func (o *InlineResponse200106NetworkPoolIpRanges) SetCidrIPv6(v string) {
	o.CidrIPv6.Set(&v)
}
// SetCidrIPv6Nil sets the value for CidrIPv6 to be an explicit nil
func (o *InlineResponse200106NetworkPoolIpRanges) SetCidrIPv6Nil() {
	o.CidrIPv6.Set(nil)
}

// UnsetCidrIPv6 ensures that no value is present for CidrIPv6, not even an explicit nil
func (o *InlineResponse200106NetworkPoolIpRanges) UnsetCidrIPv6() {
	o.CidrIPv6.Unset()
}

func (o InlineResponse200106NetworkPoolIpRanges) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.StartAddress.IsSet() {
		toSerialize["startAddress"] = o.StartAddress.Get()
	}
	if o.EndAddress.IsSet() {
		toSerialize["endAddress"] = o.EndAddress.Get()
	}
	if o.InternalId.IsSet() {
		toSerialize["internalId"] = o.InternalId.Get()
	}
	if o.ExternalId.IsSet() {
		toSerialize["externalId"] = o.ExternalId.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.AddressCount != nil {
		toSerialize["addressCount"] = o.AddressCount
	}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.DateCreated != nil {
		toSerialize["dateCreated"] = o.DateCreated
	}
	if o.LastUpdated != nil {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if o.Cidr.IsSet() {
		toSerialize["cidr"] = o.Cidr.Get()
	}
	if o.CidrIPv6.IsSet() {
		toSerialize["cidrIPv6"] = o.CidrIPv6.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200106NetworkPoolIpRanges struct {
	value *InlineResponse200106NetworkPoolIpRanges
	isSet bool
}

func (v NullableInlineResponse200106NetworkPoolIpRanges) Get() *InlineResponse200106NetworkPoolIpRanges {
	return v.value
}

func (v *NullableInlineResponse200106NetworkPoolIpRanges) Set(val *InlineResponse200106NetworkPoolIpRanges) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200106NetworkPoolIpRanges) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200106NetworkPoolIpRanges) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200106NetworkPoolIpRanges(val *InlineResponse200106NetworkPoolIpRanges) *NullableInlineResponse200106NetworkPoolIpRanges {
	return &NullableInlineResponse200106NetworkPoolIpRanges{value: val, isSet: true}
}

func (v NullableInlineResponse200106NetworkPoolIpRanges) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200106NetworkPoolIpRanges) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


