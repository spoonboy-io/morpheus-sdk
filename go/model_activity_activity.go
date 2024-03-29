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

// ActivityActivity struct for ActivityActivity
type ActivityActivity struct {
	Id *string `json:"_id,omitempty"`
	Success *bool `json:"success,omitempty"`
	ActivityType *string `json:"activityType,omitempty"`
	Name *string `json:"name,omitempty"`
	Message *string `json:"message,omitempty"`
	ObjectType *string `json:"objectType,omitempty"`
	ObjectId *int64 `json:"objectId,omitempty"`
	User *InlineResponse200107NetworkPoolCreatedBy `json:"user,omitempty"`
	Ts *time.Time `json:"ts,omitempty"`
}

// NewActivityActivity instantiates a new ActivityActivity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivityActivity() *ActivityActivity {
	this := ActivityActivity{}
	return &this
}

// NewActivityActivityWithDefaults instantiates a new ActivityActivity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivityActivityWithDefaults() *ActivityActivity {
	this := ActivityActivity{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ActivityActivity) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityActivity) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ActivityActivity) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ActivityActivity) SetId(v string) {
	o.Id = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *ActivityActivity) GetSuccess() bool {
	if o == nil || o.Success == nil {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityActivity) GetSuccessOk() (*bool, bool) {
	if o == nil || o.Success == nil {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *ActivityActivity) HasSuccess() bool {
	if o != nil && o.Success != nil {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *ActivityActivity) SetSuccess(v bool) {
	o.Success = &v
}

// GetActivityType returns the ActivityType field value if set, zero value otherwise.
func (o *ActivityActivity) GetActivityType() string {
	if o == nil || o.ActivityType == nil {
		var ret string
		return ret
	}
	return *o.ActivityType
}

// GetActivityTypeOk returns a tuple with the ActivityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityActivity) GetActivityTypeOk() (*string, bool) {
	if o == nil || o.ActivityType == nil {
		return nil, false
	}
	return o.ActivityType, true
}

// HasActivityType returns a boolean if a field has been set.
func (o *ActivityActivity) HasActivityType() bool {
	if o != nil && o.ActivityType != nil {
		return true
	}

	return false
}

// SetActivityType gets a reference to the given string and assigns it to the ActivityType field.
func (o *ActivityActivity) SetActivityType(v string) {
	o.ActivityType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ActivityActivity) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityActivity) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ActivityActivity) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ActivityActivity) SetName(v string) {
	o.Name = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ActivityActivity) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityActivity) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ActivityActivity) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ActivityActivity) SetMessage(v string) {
	o.Message = &v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *ActivityActivity) GetObjectType() string {
	if o == nil || o.ObjectType == nil {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityActivity) GetObjectTypeOk() (*string, bool) {
	if o == nil || o.ObjectType == nil {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *ActivityActivity) HasObjectType() bool {
	if o != nil && o.ObjectType != nil {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *ActivityActivity) SetObjectType(v string) {
	o.ObjectType = &v
}

// GetObjectId returns the ObjectId field value if set, zero value otherwise.
func (o *ActivityActivity) GetObjectId() int64 {
	if o == nil || o.ObjectId == nil {
		var ret int64
		return ret
	}
	return *o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityActivity) GetObjectIdOk() (*int64, bool) {
	if o == nil || o.ObjectId == nil {
		return nil, false
	}
	return o.ObjectId, true
}

// HasObjectId returns a boolean if a field has been set.
func (o *ActivityActivity) HasObjectId() bool {
	if o != nil && o.ObjectId != nil {
		return true
	}

	return false
}

// SetObjectId gets a reference to the given int64 and assigns it to the ObjectId field.
func (o *ActivityActivity) SetObjectId(v int64) {
	o.ObjectId = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *ActivityActivity) GetUser() InlineResponse200107NetworkPoolCreatedBy {
	if o == nil || o.User == nil {
		var ret InlineResponse200107NetworkPoolCreatedBy
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityActivity) GetUserOk() (*InlineResponse200107NetworkPoolCreatedBy, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *ActivityActivity) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given InlineResponse200107NetworkPoolCreatedBy and assigns it to the User field.
func (o *ActivityActivity) SetUser(v InlineResponse200107NetworkPoolCreatedBy) {
	o.User = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *ActivityActivity) GetTs() time.Time {
	if o == nil || o.Ts == nil {
		var ret time.Time
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityActivity) GetTsOk() (*time.Time, bool) {
	if o == nil || o.Ts == nil {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *ActivityActivity) HasTs() bool {
	if o != nil && o.Ts != nil {
		return true
	}

	return false
}

// SetTs gets a reference to the given time.Time and assigns it to the Ts field.
func (o *ActivityActivity) SetTs(v time.Time) {
	o.Ts = &v
}

func (o ActivityActivity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["_id"] = o.Id
	}
	if o.Success != nil {
		toSerialize["success"] = o.Success
	}
	if o.ActivityType != nil {
		toSerialize["activityType"] = o.ActivityType
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.ObjectType != nil {
		toSerialize["objectType"] = o.ObjectType
	}
	if o.ObjectId != nil {
		toSerialize["objectId"] = o.ObjectId
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.Ts != nil {
		toSerialize["ts"] = o.Ts
	}
	return json.Marshal(toSerialize)
}

type NullableActivityActivity struct {
	value *ActivityActivity
	isSet bool
}

func (v NullableActivityActivity) Get() *ActivityActivity {
	return v.value
}

func (v *NullableActivityActivity) Set(val *ActivityActivity) {
	v.value = val
	v.isSet = true
}

func (v NullableActivityActivity) IsSet() bool {
	return v.isSet
}

func (v *NullableActivityActivity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivityActivity(val *ActivityActivity) *NullableActivityActivity {
	return &NullableActivityActivity{value: val, isSet: true}
}

func (v NullableActivityActivity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivityActivity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


