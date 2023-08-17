/*
Morpheus API

Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.

API version: 6.1.1
Contact: dev@morpheusdata.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore{}

// BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore struct for BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore
type BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore struct {
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
	ExternalId *string `json:"externalId,omitempty"`
	InternalId NullableString `json:"internalId,omitempty"`
	ExternalPath NullableString `json:"externalPath,omitempty"`
}

// NewBillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore instantiates a new BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore() *BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore {
	this := BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore{}
	return &this
}

// NewBillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastoreWithDefaults instantiates a new BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastoreWithDefaults() *BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore {
	this := BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore) SetType(v string) {
	o.Type = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetInternalId returns the InternalId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore) GetInternalId() string {
	if o == nil || IsNil(o.InternalId.Get()) {
		var ret string
		return ret
	}
	return *o.InternalId.Get()
}

// GetInternalIdOk returns a tuple with the InternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore) GetInternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InternalId.Get(), o.InternalId.IsSet()
}

// HasInternalId returns a boolean if a field has been set.
func (o *BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore) HasInternalId() bool {
	if o != nil && o.InternalId.IsSet() {
		return true
	}

	return false
}

// SetInternalId gets a reference to the given NullableString and assigns it to the InternalId field.
func (o *BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore) SetInternalId(v string) {
	o.InternalId.Set(&v)
}
// SetInternalIdNil sets the value for InternalId to be an explicit nil
func (o *BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore) SetInternalIdNil() {
	o.InternalId.Set(nil)
}

// UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
func (o *BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore) UnsetInternalId() {
	o.InternalId.Unset()
}

// GetExternalPath returns the ExternalPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore) GetExternalPath() string {
	if o == nil || IsNil(o.ExternalPath.Get()) {
		var ret string
		return ret
	}
	return *o.ExternalPath.Get()
}

// GetExternalPathOk returns a tuple with the ExternalPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore) GetExternalPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalPath.Get(), o.ExternalPath.IsSet()
}

// HasExternalPath returns a boolean if a field has been set.
func (o *BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore) HasExternalPath() bool {
	if o != nil && o.ExternalPath.IsSet() {
		return true
	}

	return false
}

// SetExternalPath gets a reference to the given NullableString and assigns it to the ExternalPath field.
func (o *BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore) SetExternalPath(v string) {
	o.ExternalPath.Set(&v)
}
// SetExternalPathNil sets the value for ExternalPath to be an explicit nil
func (o *BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore) SetExternalPathNil() {
	o.ExternalPath.Set(nil)
}

// UnsetExternalPath ensures that no value is present for ExternalPath, not even an explicit nil
func (o *BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore) UnsetExternalPath() {
	o.ExternalPath.Unset()
}

func (o BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.ExternalId) {
		toSerialize["externalId"] = o.ExternalId
	}
	if o.InternalId.IsSet() {
		toSerialize["internalId"] = o.InternalId.Get()
	}
	if o.ExternalPath.IsSet() {
		toSerialize["externalPath"] = o.ExternalPath.Get()
	}
	return toSerialize, nil
}

type NullableBillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore struct {
	value *BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore
	isSet bool
}

func (v NullableBillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore) Get() *BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore {
	return v.value
}

func (v *NullableBillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore) Set(val *BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore(val *BillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore) *NullableBillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore {
	return &NullableBillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore{value: val, isSet: true}
}

func (v NullableBillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingInstancesInstancesInnerContainersInnerUsagesInnerVolumesInnerDatastore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


