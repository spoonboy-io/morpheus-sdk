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

// Cypher struct for Cypher
type Cypher struct {
	Id *int32 `json:"id,omitempty"`
	ItemKey *string `json:"itemKey,omitempty"`
	LeaseTimeout *int64 `json:"leaseTimeout,omitempty"`
	ExpireDate NullableTime `json:"expireDate,omitempty"`
	DateCreated NullableTime `json:"dateCreated,omitempty"`
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	LastAccessed *time.Time `json:"lastAccessed,omitempty"`
	CreatedBy *string `json:"createdBy,omitempty"`
}

// NewCypher instantiates a new Cypher object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCypher() *Cypher {
	this := Cypher{}
	return &this
}

// NewCypherWithDefaults instantiates a new Cypher object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCypherWithDefaults() *Cypher {
	this := Cypher{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Cypher) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cypher) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Cypher) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Cypher) SetId(v int32) {
	o.Id = &v
}

// GetItemKey returns the ItemKey field value if set, zero value otherwise.
func (o *Cypher) GetItemKey() string {
	if o == nil || o.ItemKey == nil {
		var ret string
		return ret
	}
	return *o.ItemKey
}

// GetItemKeyOk returns a tuple with the ItemKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cypher) GetItemKeyOk() (*string, bool) {
	if o == nil || o.ItemKey == nil {
		return nil, false
	}
	return o.ItemKey, true
}

// HasItemKey returns a boolean if a field has been set.
func (o *Cypher) HasItemKey() bool {
	if o != nil && o.ItemKey != nil {
		return true
	}

	return false
}

// SetItemKey gets a reference to the given string and assigns it to the ItemKey field.
func (o *Cypher) SetItemKey(v string) {
	o.ItemKey = &v
}

// GetLeaseTimeout returns the LeaseTimeout field value if set, zero value otherwise.
func (o *Cypher) GetLeaseTimeout() int64 {
	if o == nil || o.LeaseTimeout == nil {
		var ret int64
		return ret
	}
	return *o.LeaseTimeout
}

// GetLeaseTimeoutOk returns a tuple with the LeaseTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cypher) GetLeaseTimeoutOk() (*int64, bool) {
	if o == nil || o.LeaseTimeout == nil {
		return nil, false
	}
	return o.LeaseTimeout, true
}

// HasLeaseTimeout returns a boolean if a field has been set.
func (o *Cypher) HasLeaseTimeout() bool {
	if o != nil && o.LeaseTimeout != nil {
		return true
	}

	return false
}

// SetLeaseTimeout gets a reference to the given int64 and assigns it to the LeaseTimeout field.
func (o *Cypher) SetLeaseTimeout(v int64) {
	o.LeaseTimeout = &v
}

// GetExpireDate returns the ExpireDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Cypher) GetExpireDate() time.Time {
	if o == nil || o.ExpireDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpireDate.Get()
}

// GetExpireDateOk returns a tuple with the ExpireDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cypher) GetExpireDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExpireDate.Get(), o.ExpireDate.IsSet()
}

// HasExpireDate returns a boolean if a field has been set.
func (o *Cypher) HasExpireDate() bool {
	if o != nil && o.ExpireDate.IsSet() {
		return true
	}

	return false
}

// SetExpireDate gets a reference to the given NullableTime and assigns it to the ExpireDate field.
func (o *Cypher) SetExpireDate(v time.Time) {
	o.ExpireDate.Set(&v)
}
// SetExpireDateNil sets the value for ExpireDate to be an explicit nil
func (o *Cypher) SetExpireDateNil() {
	o.ExpireDate.Set(nil)
}

// UnsetExpireDate ensures that no value is present for ExpireDate, not even an explicit nil
func (o *Cypher) UnsetExpireDate() {
	o.ExpireDate.Unset()
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Cypher) GetDateCreated() time.Time {
	if o == nil || o.DateCreated.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DateCreated.Get()
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cypher) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DateCreated.Get(), o.DateCreated.IsSet()
}

// HasDateCreated returns a boolean if a field has been set.
func (o *Cypher) HasDateCreated() bool {
	if o != nil && o.DateCreated.IsSet() {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given NullableTime and assigns it to the DateCreated field.
func (o *Cypher) SetDateCreated(v time.Time) {
	o.DateCreated.Set(&v)
}
// SetDateCreatedNil sets the value for DateCreated to be an explicit nil
func (o *Cypher) SetDateCreatedNil() {
	o.DateCreated.Set(nil)
}

// UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
func (o *Cypher) UnsetDateCreated() {
	o.DateCreated.Unset()
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *Cypher) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cypher) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *Cypher) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *Cypher) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetLastAccessed returns the LastAccessed field value if set, zero value otherwise.
func (o *Cypher) GetLastAccessed() time.Time {
	if o == nil || o.LastAccessed == nil {
		var ret time.Time
		return ret
	}
	return *o.LastAccessed
}

// GetLastAccessedOk returns a tuple with the LastAccessed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cypher) GetLastAccessedOk() (*time.Time, bool) {
	if o == nil || o.LastAccessed == nil {
		return nil, false
	}
	return o.LastAccessed, true
}

// HasLastAccessed returns a boolean if a field has been set.
func (o *Cypher) HasLastAccessed() bool {
	if o != nil && o.LastAccessed != nil {
		return true
	}

	return false
}

// SetLastAccessed gets a reference to the given time.Time and assigns it to the LastAccessed field.
func (o *Cypher) SetLastAccessed(v time.Time) {
	o.LastAccessed = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *Cypher) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cypher) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *Cypher) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *Cypher) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

func (o Cypher) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ItemKey != nil {
		toSerialize["itemKey"] = o.ItemKey
	}
	if o.LeaseTimeout != nil {
		toSerialize["leaseTimeout"] = o.LeaseTimeout
	}
	if o.ExpireDate.IsSet() {
		toSerialize["expireDate"] = o.ExpireDate.Get()
	}
	if o.DateCreated.IsSet() {
		toSerialize["dateCreated"] = o.DateCreated.Get()
	}
	if o.LastUpdated != nil {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if o.LastAccessed != nil {
		toSerialize["lastAccessed"] = o.LastAccessed
	}
	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}
	return json.Marshal(toSerialize)
}

type NullableCypher struct {
	value *Cypher
	isSet bool
}

func (v NullableCypher) Get() *Cypher {
	return v.value
}

func (v *NullableCypher) Set(val *Cypher) {
	v.value = val
	v.isSet = true
}

func (v NullableCypher) IsSet() bool {
	return v.isSet
}

func (v *NullableCypher) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCypher(val *Cypher) *NullableCypher {
	return &NullableCypher{value: val, isSet: true}
}

func (v NullableCypher) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCypher) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


