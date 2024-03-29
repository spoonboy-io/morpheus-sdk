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

// WikiPage struct for WikiPage
type WikiPage struct {
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	UrlName *string `json:"urlName,omitempty"`
	Category *string `json:"category,omitempty"`
	RefId NullableString `json:"refId,omitempty"`
	RefType NullableString `json:"refType,omitempty"`
	Format *string `json:"format,omitempty"`
	Content *string `json:"content,omitempty"`
	CreatedBy *InlineResponse200107NetworkPoolCreatedBy `json:"createdBy,omitempty"`
	UpdatedBy *InlineResponse200107NetworkPoolCreatedBy `json:"updatedBy,omitempty"`
	DateCreated *time.Time `json:"dateCreated,omitempty"`
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
}

// NewWikiPage instantiates a new WikiPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWikiPage() *WikiPage {
	this := WikiPage{}
	return &this
}

// NewWikiPageWithDefaults instantiates a new WikiPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWikiPageWithDefaults() *WikiPage {
	this := WikiPage{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WikiPage) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WikiPage) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WikiPage) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *WikiPage) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WikiPage) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WikiPage) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WikiPage) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WikiPage) SetName(v string) {
	o.Name = &v
}

// GetUrlName returns the UrlName field value if set, zero value otherwise.
func (o *WikiPage) GetUrlName() string {
	if o == nil || o.UrlName == nil {
		var ret string
		return ret
	}
	return *o.UrlName
}

// GetUrlNameOk returns a tuple with the UrlName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WikiPage) GetUrlNameOk() (*string, bool) {
	if o == nil || o.UrlName == nil {
		return nil, false
	}
	return o.UrlName, true
}

// HasUrlName returns a boolean if a field has been set.
func (o *WikiPage) HasUrlName() bool {
	if o != nil && o.UrlName != nil {
		return true
	}

	return false
}

// SetUrlName gets a reference to the given string and assigns it to the UrlName field.
func (o *WikiPage) SetUrlName(v string) {
	o.UrlName = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *WikiPage) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WikiPage) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *WikiPage) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *WikiPage) SetCategory(v string) {
	o.Category = &v
}

// GetRefId returns the RefId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WikiPage) GetRefId() string {
	if o == nil || o.RefId.Get() == nil {
		var ret string
		return ret
	}
	return *o.RefId.Get()
}

// GetRefIdOk returns a tuple with the RefId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WikiPage) GetRefIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RefId.Get(), o.RefId.IsSet()
}

// HasRefId returns a boolean if a field has been set.
func (o *WikiPage) HasRefId() bool {
	if o != nil && o.RefId.IsSet() {
		return true
	}

	return false
}

// SetRefId gets a reference to the given NullableString and assigns it to the RefId field.
func (o *WikiPage) SetRefId(v string) {
	o.RefId.Set(&v)
}
// SetRefIdNil sets the value for RefId to be an explicit nil
func (o *WikiPage) SetRefIdNil() {
	o.RefId.Set(nil)
}

// UnsetRefId ensures that no value is present for RefId, not even an explicit nil
func (o *WikiPage) UnsetRefId() {
	o.RefId.Unset()
}

// GetRefType returns the RefType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WikiPage) GetRefType() string {
	if o == nil || o.RefType.Get() == nil {
		var ret string
		return ret
	}
	return *o.RefType.Get()
}

// GetRefTypeOk returns a tuple with the RefType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WikiPage) GetRefTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RefType.Get(), o.RefType.IsSet()
}

// HasRefType returns a boolean if a field has been set.
func (o *WikiPage) HasRefType() bool {
	if o != nil && o.RefType.IsSet() {
		return true
	}

	return false
}

// SetRefType gets a reference to the given NullableString and assigns it to the RefType field.
func (o *WikiPage) SetRefType(v string) {
	o.RefType.Set(&v)
}
// SetRefTypeNil sets the value for RefType to be an explicit nil
func (o *WikiPage) SetRefTypeNil() {
	o.RefType.Set(nil)
}

// UnsetRefType ensures that no value is present for RefType, not even an explicit nil
func (o *WikiPage) UnsetRefType() {
	o.RefType.Unset()
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *WikiPage) GetFormat() string {
	if o == nil || o.Format == nil {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WikiPage) GetFormatOk() (*string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *WikiPage) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *WikiPage) SetFormat(v string) {
	o.Format = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *WikiPage) GetContent() string {
	if o == nil || o.Content == nil {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WikiPage) GetContentOk() (*string, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *WikiPage) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *WikiPage) SetContent(v string) {
	o.Content = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *WikiPage) GetCreatedBy() InlineResponse200107NetworkPoolCreatedBy {
	if o == nil || o.CreatedBy == nil {
		var ret InlineResponse200107NetworkPoolCreatedBy
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WikiPage) GetCreatedByOk() (*InlineResponse200107NetworkPoolCreatedBy, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *WikiPage) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given InlineResponse200107NetworkPoolCreatedBy and assigns it to the CreatedBy field.
func (o *WikiPage) SetCreatedBy(v InlineResponse200107NetworkPoolCreatedBy) {
	o.CreatedBy = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *WikiPage) GetUpdatedBy() InlineResponse200107NetworkPoolCreatedBy {
	if o == nil || o.UpdatedBy == nil {
		var ret InlineResponse200107NetworkPoolCreatedBy
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WikiPage) GetUpdatedByOk() (*InlineResponse200107NetworkPoolCreatedBy, bool) {
	if o == nil || o.UpdatedBy == nil {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *WikiPage) HasUpdatedBy() bool {
	if o != nil && o.UpdatedBy != nil {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given InlineResponse200107NetworkPoolCreatedBy and assigns it to the UpdatedBy field.
func (o *WikiPage) SetUpdatedBy(v InlineResponse200107NetworkPoolCreatedBy) {
	o.UpdatedBy = &v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *WikiPage) GetDateCreated() time.Time {
	if o == nil || o.DateCreated == nil {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WikiPage) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || o.DateCreated == nil {
		return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *WikiPage) HasDateCreated() bool {
	if o != nil && o.DateCreated != nil {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *WikiPage) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *WikiPage) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WikiPage) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *WikiPage) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *WikiPage) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

func (o WikiPage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.UrlName != nil {
		toSerialize["urlName"] = o.UrlName
	}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.RefId.IsSet() {
		toSerialize["refId"] = o.RefId.Get()
	}
	if o.RefType.IsSet() {
		toSerialize["refType"] = o.RefType.Get()
	}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.UpdatedBy != nil {
		toSerialize["updatedBy"] = o.UpdatedBy
	}
	if o.DateCreated != nil {
		toSerialize["dateCreated"] = o.DateCreated
	}
	if o.LastUpdated != nil {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	return json.Marshal(toSerialize)
}

type NullableWikiPage struct {
	value *WikiPage
	isSet bool
}

func (v NullableWikiPage) Get() *WikiPage {
	return v.value
}

func (v *NullableWikiPage) Set(val *WikiPage) {
	v.value = val
	v.isSet = true
}

func (v NullableWikiPage) IsSet() bool {
	return v.isSet
}

func (v *NullableWikiPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWikiPage(val *WikiPage) *NullableWikiPage {
	return &NullableWikiPage{value: val, isSet: true}
}

func (v NullableWikiPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWikiPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


