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

// LogData struct for LogData
type LogData struct {
	TypeCode *string `json:"typeCode,omitempty"`
	Message *string `json:"message,omitempty"`
	Level *string `json:"level,omitempty"`
	Ts *time.Time `json:"ts,omitempty"`
	SourceType *string `json:"sourceType,omitempty"`
	Title *string `json:"title,omitempty"`
	LogSignature *string `json:"logSignature,omitempty"`
	ObjectId *string `json:"objectId,omitempty"`
	Seq *int64 `json:"seq,omitempty"`
	Id *string `json:"_id,omitempty"`
	SignatureVerified *bool `json:"signatureVerified,omitempty"`
}

// NewLogData instantiates a new LogData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogData() *LogData {
	this := LogData{}
	return &this
}

// NewLogDataWithDefaults instantiates a new LogData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogDataWithDefaults() *LogData {
	this := LogData{}
	return &this
}

// GetTypeCode returns the TypeCode field value if set, zero value otherwise.
func (o *LogData) GetTypeCode() string {
	if o == nil || o.TypeCode == nil {
		var ret string
		return ret
	}
	return *o.TypeCode
}

// GetTypeCodeOk returns a tuple with the TypeCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogData) GetTypeCodeOk() (*string, bool) {
	if o == nil || o.TypeCode == nil {
		return nil, false
	}
	return o.TypeCode, true
}

// HasTypeCode returns a boolean if a field has been set.
func (o *LogData) HasTypeCode() bool {
	if o != nil && o.TypeCode != nil {
		return true
	}

	return false
}

// SetTypeCode gets a reference to the given string and assigns it to the TypeCode field.
func (o *LogData) SetTypeCode(v string) {
	o.TypeCode = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *LogData) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogData) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *LogData) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *LogData) SetMessage(v string) {
	o.Message = &v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *LogData) GetLevel() string {
	if o == nil || o.Level == nil {
		var ret string
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogData) GetLevelOk() (*string, bool) {
	if o == nil || o.Level == nil {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *LogData) HasLevel() bool {
	if o != nil && o.Level != nil {
		return true
	}

	return false
}

// SetLevel gets a reference to the given string and assigns it to the Level field.
func (o *LogData) SetLevel(v string) {
	o.Level = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *LogData) GetTs() time.Time {
	if o == nil || o.Ts == nil {
		var ret time.Time
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogData) GetTsOk() (*time.Time, bool) {
	if o == nil || o.Ts == nil {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *LogData) HasTs() bool {
	if o != nil && o.Ts != nil {
		return true
	}

	return false
}

// SetTs gets a reference to the given time.Time and assigns it to the Ts field.
func (o *LogData) SetTs(v time.Time) {
	o.Ts = &v
}

// GetSourceType returns the SourceType field value if set, zero value otherwise.
func (o *LogData) GetSourceType() string {
	if o == nil || o.SourceType == nil {
		var ret string
		return ret
	}
	return *o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogData) GetSourceTypeOk() (*string, bool) {
	if o == nil || o.SourceType == nil {
		return nil, false
	}
	return o.SourceType, true
}

// HasSourceType returns a boolean if a field has been set.
func (o *LogData) HasSourceType() bool {
	if o != nil && o.SourceType != nil {
		return true
	}

	return false
}

// SetSourceType gets a reference to the given string and assigns it to the SourceType field.
func (o *LogData) SetSourceType(v string) {
	o.SourceType = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *LogData) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogData) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *LogData) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *LogData) SetTitle(v string) {
	o.Title = &v
}

// GetLogSignature returns the LogSignature field value if set, zero value otherwise.
func (o *LogData) GetLogSignature() string {
	if o == nil || o.LogSignature == nil {
		var ret string
		return ret
	}
	return *o.LogSignature
}

// GetLogSignatureOk returns a tuple with the LogSignature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogData) GetLogSignatureOk() (*string, bool) {
	if o == nil || o.LogSignature == nil {
		return nil, false
	}
	return o.LogSignature, true
}

// HasLogSignature returns a boolean if a field has been set.
func (o *LogData) HasLogSignature() bool {
	if o != nil && o.LogSignature != nil {
		return true
	}

	return false
}

// SetLogSignature gets a reference to the given string and assigns it to the LogSignature field.
func (o *LogData) SetLogSignature(v string) {
	o.LogSignature = &v
}

// GetObjectId returns the ObjectId field value if set, zero value otherwise.
func (o *LogData) GetObjectId() string {
	if o == nil || o.ObjectId == nil {
		var ret string
		return ret
	}
	return *o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogData) GetObjectIdOk() (*string, bool) {
	if o == nil || o.ObjectId == nil {
		return nil, false
	}
	return o.ObjectId, true
}

// HasObjectId returns a boolean if a field has been set.
func (o *LogData) HasObjectId() bool {
	if o != nil && o.ObjectId != nil {
		return true
	}

	return false
}

// SetObjectId gets a reference to the given string and assigns it to the ObjectId field.
func (o *LogData) SetObjectId(v string) {
	o.ObjectId = &v
}

// GetSeq returns the Seq field value if set, zero value otherwise.
func (o *LogData) GetSeq() int64 {
	if o == nil || o.Seq == nil {
		var ret int64
		return ret
	}
	return *o.Seq
}

// GetSeqOk returns a tuple with the Seq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogData) GetSeqOk() (*int64, bool) {
	if o == nil || o.Seq == nil {
		return nil, false
	}
	return o.Seq, true
}

// HasSeq returns a boolean if a field has been set.
func (o *LogData) HasSeq() bool {
	if o != nil && o.Seq != nil {
		return true
	}

	return false
}

// SetSeq gets a reference to the given int64 and assigns it to the Seq field.
func (o *LogData) SetSeq(v int64) {
	o.Seq = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LogData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LogData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LogData) SetId(v string) {
	o.Id = &v
}

// GetSignatureVerified returns the SignatureVerified field value if set, zero value otherwise.
func (o *LogData) GetSignatureVerified() bool {
	if o == nil || o.SignatureVerified == nil {
		var ret bool
		return ret
	}
	return *o.SignatureVerified
}

// GetSignatureVerifiedOk returns a tuple with the SignatureVerified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogData) GetSignatureVerifiedOk() (*bool, bool) {
	if o == nil || o.SignatureVerified == nil {
		return nil, false
	}
	return o.SignatureVerified, true
}

// HasSignatureVerified returns a boolean if a field has been set.
func (o *LogData) HasSignatureVerified() bool {
	if o != nil && o.SignatureVerified != nil {
		return true
	}

	return false
}

// SetSignatureVerified gets a reference to the given bool and assigns it to the SignatureVerified field.
func (o *LogData) SetSignatureVerified(v bool) {
	o.SignatureVerified = &v
}

func (o LogData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TypeCode != nil {
		toSerialize["typeCode"] = o.TypeCode
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Level != nil {
		toSerialize["level"] = o.Level
	}
	if o.Ts != nil {
		toSerialize["ts"] = o.Ts
	}
	if o.SourceType != nil {
		toSerialize["sourceType"] = o.SourceType
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.LogSignature != nil {
		toSerialize["logSignature"] = o.LogSignature
	}
	if o.ObjectId != nil {
		toSerialize["objectId"] = o.ObjectId
	}
	if o.Seq != nil {
		toSerialize["seq"] = o.Seq
	}
	if o.Id != nil {
		toSerialize["_id"] = o.Id
	}
	if o.SignatureVerified != nil {
		toSerialize["signatureVerified"] = o.SignatureVerified
	}
	return json.Marshal(toSerialize)
}

type NullableLogData struct {
	value *LogData
	isSet bool
}

func (v NullableLogData) Get() *LogData {
	return v.value
}

func (v *NullableLogData) Set(val *LogData) {
	v.value = val
	v.isSet = true
}

func (v NullableLogData) IsSet() bool {
	return v.isSet
}

func (v *NullableLogData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogData(val *LogData) *NullableLogData {
	return &NullableLogData{value: val, isSet: true}
}

func (v NullableLogData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


