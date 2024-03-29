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

// SnapshotSnapshot struct for SnapshotSnapshot
type SnapshotSnapshot struct {
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Description NullableString `json:"description,omitempty"`
	ExternalId *string `json:"externalId,omitempty"`
	Status *string `json:"status,omitempty"`
	State NullableString `json:"state,omitempty"`
	SnapshotType *string `json:"snapshotType,omitempty"`
	SnapshotCreated *time.Time `json:"snapshotCreated,omitempty"`
	Zone *SnapshotSnapshotZone `json:"zone,omitempty"`
	Datastore NullableString `json:"datastore,omitempty"`
	ParentSnapshot NullableString `json:"parentSnapshot,omitempty"`
	CurrentlyActive *bool `json:"currentlyActive,omitempty"`
	DateCreated *time.Time `json:"dateCreated,omitempty"`
}

// NewSnapshotSnapshot instantiates a new SnapshotSnapshot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshotSnapshot() *SnapshotSnapshot {
	this := SnapshotSnapshot{}
	return &this
}

// NewSnapshotSnapshotWithDefaults instantiates a new SnapshotSnapshot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotSnapshotWithDefaults() *SnapshotSnapshot {
	this := SnapshotSnapshot{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SnapshotSnapshot) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotSnapshot) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SnapshotSnapshot) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *SnapshotSnapshot) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SnapshotSnapshot) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotSnapshot) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SnapshotSnapshot) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SnapshotSnapshot) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SnapshotSnapshot) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnapshotSnapshot) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *SnapshotSnapshot) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *SnapshotSnapshot) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *SnapshotSnapshot) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *SnapshotSnapshot) UnsetDescription() {
	o.Description.Unset()
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *SnapshotSnapshot) GetExternalId() string {
	if o == nil || o.ExternalId == nil {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotSnapshot) GetExternalIdOk() (*string, bool) {
	if o == nil || o.ExternalId == nil {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *SnapshotSnapshot) HasExternalId() bool {
	if o != nil && o.ExternalId != nil {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *SnapshotSnapshot) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SnapshotSnapshot) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotSnapshot) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SnapshotSnapshot) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SnapshotSnapshot) SetStatus(v string) {
	o.Status = &v
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SnapshotSnapshot) GetState() string {
	if o == nil || o.State.Get() == nil {
		var ret string
		return ret
	}
	return *o.State.Get()
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnapshotSnapshot) GetStateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.State.Get(), o.State.IsSet()
}

// HasState returns a boolean if a field has been set.
func (o *SnapshotSnapshot) HasState() bool {
	if o != nil && o.State.IsSet() {
		return true
	}

	return false
}

// SetState gets a reference to the given NullableString and assigns it to the State field.
func (o *SnapshotSnapshot) SetState(v string) {
	o.State.Set(&v)
}
// SetStateNil sets the value for State to be an explicit nil
func (o *SnapshotSnapshot) SetStateNil() {
	o.State.Set(nil)
}

// UnsetState ensures that no value is present for State, not even an explicit nil
func (o *SnapshotSnapshot) UnsetState() {
	o.State.Unset()
}

// GetSnapshotType returns the SnapshotType field value if set, zero value otherwise.
func (o *SnapshotSnapshot) GetSnapshotType() string {
	if o == nil || o.SnapshotType == nil {
		var ret string
		return ret
	}
	return *o.SnapshotType
}

// GetSnapshotTypeOk returns a tuple with the SnapshotType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotSnapshot) GetSnapshotTypeOk() (*string, bool) {
	if o == nil || o.SnapshotType == nil {
		return nil, false
	}
	return o.SnapshotType, true
}

// HasSnapshotType returns a boolean if a field has been set.
func (o *SnapshotSnapshot) HasSnapshotType() bool {
	if o != nil && o.SnapshotType != nil {
		return true
	}

	return false
}

// SetSnapshotType gets a reference to the given string and assigns it to the SnapshotType field.
func (o *SnapshotSnapshot) SetSnapshotType(v string) {
	o.SnapshotType = &v
}

// GetSnapshotCreated returns the SnapshotCreated field value if set, zero value otherwise.
func (o *SnapshotSnapshot) GetSnapshotCreated() time.Time {
	if o == nil || o.SnapshotCreated == nil {
		var ret time.Time
		return ret
	}
	return *o.SnapshotCreated
}

// GetSnapshotCreatedOk returns a tuple with the SnapshotCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotSnapshot) GetSnapshotCreatedOk() (*time.Time, bool) {
	if o == nil || o.SnapshotCreated == nil {
		return nil, false
	}
	return o.SnapshotCreated, true
}

// HasSnapshotCreated returns a boolean if a field has been set.
func (o *SnapshotSnapshot) HasSnapshotCreated() bool {
	if o != nil && o.SnapshotCreated != nil {
		return true
	}

	return false
}

// SetSnapshotCreated gets a reference to the given time.Time and assigns it to the SnapshotCreated field.
func (o *SnapshotSnapshot) SetSnapshotCreated(v time.Time) {
	o.SnapshotCreated = &v
}

// GetZone returns the Zone field value if set, zero value otherwise.
func (o *SnapshotSnapshot) GetZone() SnapshotSnapshotZone {
	if o == nil || o.Zone == nil {
		var ret SnapshotSnapshotZone
		return ret
	}
	return *o.Zone
}

// GetZoneOk returns a tuple with the Zone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotSnapshot) GetZoneOk() (*SnapshotSnapshotZone, bool) {
	if o == nil || o.Zone == nil {
		return nil, false
	}
	return o.Zone, true
}

// HasZone returns a boolean if a field has been set.
func (o *SnapshotSnapshot) HasZone() bool {
	if o != nil && o.Zone != nil {
		return true
	}

	return false
}

// SetZone gets a reference to the given SnapshotSnapshotZone and assigns it to the Zone field.
func (o *SnapshotSnapshot) SetZone(v SnapshotSnapshotZone) {
	o.Zone = &v
}

// GetDatastore returns the Datastore field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SnapshotSnapshot) GetDatastore() string {
	if o == nil || o.Datastore.Get() == nil {
		var ret string
		return ret
	}
	return *o.Datastore.Get()
}

// GetDatastoreOk returns a tuple with the Datastore field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnapshotSnapshot) GetDatastoreOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Datastore.Get(), o.Datastore.IsSet()
}

// HasDatastore returns a boolean if a field has been set.
func (o *SnapshotSnapshot) HasDatastore() bool {
	if o != nil && o.Datastore.IsSet() {
		return true
	}

	return false
}

// SetDatastore gets a reference to the given NullableString and assigns it to the Datastore field.
func (o *SnapshotSnapshot) SetDatastore(v string) {
	o.Datastore.Set(&v)
}
// SetDatastoreNil sets the value for Datastore to be an explicit nil
func (o *SnapshotSnapshot) SetDatastoreNil() {
	o.Datastore.Set(nil)
}

// UnsetDatastore ensures that no value is present for Datastore, not even an explicit nil
func (o *SnapshotSnapshot) UnsetDatastore() {
	o.Datastore.Unset()
}

// GetParentSnapshot returns the ParentSnapshot field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SnapshotSnapshot) GetParentSnapshot() string {
	if o == nil || o.ParentSnapshot.Get() == nil {
		var ret string
		return ret
	}
	return *o.ParentSnapshot.Get()
}

// GetParentSnapshotOk returns a tuple with the ParentSnapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SnapshotSnapshot) GetParentSnapshotOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ParentSnapshot.Get(), o.ParentSnapshot.IsSet()
}

// HasParentSnapshot returns a boolean if a field has been set.
func (o *SnapshotSnapshot) HasParentSnapshot() bool {
	if o != nil && o.ParentSnapshot.IsSet() {
		return true
	}

	return false
}

// SetParentSnapshot gets a reference to the given NullableString and assigns it to the ParentSnapshot field.
func (o *SnapshotSnapshot) SetParentSnapshot(v string) {
	o.ParentSnapshot.Set(&v)
}
// SetParentSnapshotNil sets the value for ParentSnapshot to be an explicit nil
func (o *SnapshotSnapshot) SetParentSnapshotNil() {
	o.ParentSnapshot.Set(nil)
}

// UnsetParentSnapshot ensures that no value is present for ParentSnapshot, not even an explicit nil
func (o *SnapshotSnapshot) UnsetParentSnapshot() {
	o.ParentSnapshot.Unset()
}

// GetCurrentlyActive returns the CurrentlyActive field value if set, zero value otherwise.
func (o *SnapshotSnapshot) GetCurrentlyActive() bool {
	if o == nil || o.CurrentlyActive == nil {
		var ret bool
		return ret
	}
	return *o.CurrentlyActive
}

// GetCurrentlyActiveOk returns a tuple with the CurrentlyActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotSnapshot) GetCurrentlyActiveOk() (*bool, bool) {
	if o == nil || o.CurrentlyActive == nil {
		return nil, false
	}
	return o.CurrentlyActive, true
}

// HasCurrentlyActive returns a boolean if a field has been set.
func (o *SnapshotSnapshot) HasCurrentlyActive() bool {
	if o != nil && o.CurrentlyActive != nil {
		return true
	}

	return false
}

// SetCurrentlyActive gets a reference to the given bool and assigns it to the CurrentlyActive field.
func (o *SnapshotSnapshot) SetCurrentlyActive(v bool) {
	o.CurrentlyActive = &v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *SnapshotSnapshot) GetDateCreated() time.Time {
	if o == nil || o.DateCreated == nil {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotSnapshot) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || o.DateCreated == nil {
		return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *SnapshotSnapshot) HasDateCreated() bool {
	if o != nil && o.DateCreated != nil {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *SnapshotSnapshot) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

func (o SnapshotSnapshot) MarshalJSON() ([]byte, error) {
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
	if o.ExternalId != nil {
		toSerialize["externalId"] = o.ExternalId
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.State.IsSet() {
		toSerialize["state"] = o.State.Get()
	}
	if o.SnapshotType != nil {
		toSerialize["snapshotType"] = o.SnapshotType
	}
	if o.SnapshotCreated != nil {
		toSerialize["snapshotCreated"] = o.SnapshotCreated
	}
	if o.Zone != nil {
		toSerialize["zone"] = o.Zone
	}
	if o.Datastore.IsSet() {
		toSerialize["datastore"] = o.Datastore.Get()
	}
	if o.ParentSnapshot.IsSet() {
		toSerialize["parentSnapshot"] = o.ParentSnapshot.Get()
	}
	if o.CurrentlyActive != nil {
		toSerialize["currentlyActive"] = o.CurrentlyActive
	}
	if o.DateCreated != nil {
		toSerialize["dateCreated"] = o.DateCreated
	}
	return json.Marshal(toSerialize)
}

type NullableSnapshotSnapshot struct {
	value *SnapshotSnapshot
	isSet bool
}

func (v NullableSnapshotSnapshot) Get() *SnapshotSnapshot {
	return v.value
}

func (v *NullableSnapshotSnapshot) Set(val *SnapshotSnapshot) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotSnapshot) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotSnapshot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotSnapshot(val *SnapshotSnapshot) *NullableSnapshotSnapshot {
	return &NullableSnapshotSnapshot{value: val, isSet: true}
}

func (v NullableSnapshotSnapshot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotSnapshot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


