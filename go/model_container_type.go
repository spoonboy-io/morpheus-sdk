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

// ContainerType struct for ContainerType
type ContainerType struct {
	Id *int32 `json:"id,omitempty"`
	Account NullableContainerTypeAccount `json:"account,omitempty"`
	Name *string `json:"name,omitempty"`
	Labels *[]string `json:"labels,omitempty"`
	ShortName *string `json:"shortName,omitempty"`
	Code *string `json:"code,omitempty"`
	ContainerVersion *string `json:"containerVersion,omitempty"`
	ProvisionType *ContainerTypeProvisionType `json:"provisionType,omitempty"`
	VirtualImage NullableContainerTypeAccount `json:"virtualImage,omitempty"`
	Category NullableString `json:"category,omitempty"`
	Config map[string]interface{} `json:"config,omitempty"`
	ContainerPorts []ContainerTypeContainerPorts `json:"containerPorts,omitempty"`
	ContainerScripts []map[string]interface{} `json:"containerScripts,omitempty"`
	ContainerTemplates []map[string]interface{} `json:"containerTemplates,omitempty"`
	EnvironmentVariables []map[string]interface{} `json:"environmentVariables,omitempty"`
}

// NewContainerType instantiates a new ContainerType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerType() *ContainerType {
	this := ContainerType{}
	return &this
}

// NewContainerTypeWithDefaults instantiates a new ContainerType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerTypeWithDefaults() *ContainerType {
	this := ContainerType{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ContainerType) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerType) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ContainerType) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ContainerType) SetId(v int32) {
	o.Id = &v
}

// GetAccount returns the Account field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContainerType) GetAccount() ContainerTypeAccount {
	if o == nil || o.Account.Get() == nil {
		var ret ContainerTypeAccount
		return ret
	}
	return *o.Account.Get()
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContainerType) GetAccountOk() (*ContainerTypeAccount, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Account.Get(), o.Account.IsSet()
}

// HasAccount returns a boolean if a field has been set.
func (o *ContainerType) HasAccount() bool {
	if o != nil && o.Account.IsSet() {
		return true
	}

	return false
}

// SetAccount gets a reference to the given NullableContainerTypeAccount and assigns it to the Account field.
func (o *ContainerType) SetAccount(v ContainerTypeAccount) {
	o.Account.Set(&v)
}
// SetAccountNil sets the value for Account to be an explicit nil
func (o *ContainerType) SetAccountNil() {
	o.Account.Set(nil)
}

// UnsetAccount ensures that no value is present for Account, not even an explicit nil
func (o *ContainerType) UnsetAccount() {
	o.Account.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ContainerType) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerType) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ContainerType) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ContainerType) SetName(v string) {
	o.Name = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *ContainerType) GetLabels() []string {
	if o == nil || o.Labels == nil {
		var ret []string
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerType) GetLabelsOk() (*[]string, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *ContainerType) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *ContainerType) SetLabels(v []string) {
	o.Labels = &v
}

// GetShortName returns the ShortName field value if set, zero value otherwise.
func (o *ContainerType) GetShortName() string {
	if o == nil || o.ShortName == nil {
		var ret string
		return ret
	}
	return *o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerType) GetShortNameOk() (*string, bool) {
	if o == nil || o.ShortName == nil {
		return nil, false
	}
	return o.ShortName, true
}

// HasShortName returns a boolean if a field has been set.
func (o *ContainerType) HasShortName() bool {
	if o != nil && o.ShortName != nil {
		return true
	}

	return false
}

// SetShortName gets a reference to the given string and assigns it to the ShortName field.
func (o *ContainerType) SetShortName(v string) {
	o.ShortName = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ContainerType) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerType) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ContainerType) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ContainerType) SetCode(v string) {
	o.Code = &v
}

// GetContainerVersion returns the ContainerVersion field value if set, zero value otherwise.
func (o *ContainerType) GetContainerVersion() string {
	if o == nil || o.ContainerVersion == nil {
		var ret string
		return ret
	}
	return *o.ContainerVersion
}

// GetContainerVersionOk returns a tuple with the ContainerVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerType) GetContainerVersionOk() (*string, bool) {
	if o == nil || o.ContainerVersion == nil {
		return nil, false
	}
	return o.ContainerVersion, true
}

// HasContainerVersion returns a boolean if a field has been set.
func (o *ContainerType) HasContainerVersion() bool {
	if o != nil && o.ContainerVersion != nil {
		return true
	}

	return false
}

// SetContainerVersion gets a reference to the given string and assigns it to the ContainerVersion field.
func (o *ContainerType) SetContainerVersion(v string) {
	o.ContainerVersion = &v
}

// GetProvisionType returns the ProvisionType field value if set, zero value otherwise.
func (o *ContainerType) GetProvisionType() ContainerTypeProvisionType {
	if o == nil || o.ProvisionType == nil {
		var ret ContainerTypeProvisionType
		return ret
	}
	return *o.ProvisionType
}

// GetProvisionTypeOk returns a tuple with the ProvisionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerType) GetProvisionTypeOk() (*ContainerTypeProvisionType, bool) {
	if o == nil || o.ProvisionType == nil {
		return nil, false
	}
	return o.ProvisionType, true
}

// HasProvisionType returns a boolean if a field has been set.
func (o *ContainerType) HasProvisionType() bool {
	if o != nil && o.ProvisionType != nil {
		return true
	}

	return false
}

// SetProvisionType gets a reference to the given ContainerTypeProvisionType and assigns it to the ProvisionType field.
func (o *ContainerType) SetProvisionType(v ContainerTypeProvisionType) {
	o.ProvisionType = &v
}

// GetVirtualImage returns the VirtualImage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContainerType) GetVirtualImage() ContainerTypeAccount {
	if o == nil || o.VirtualImage.Get() == nil {
		var ret ContainerTypeAccount
		return ret
	}
	return *o.VirtualImage.Get()
}

// GetVirtualImageOk returns a tuple with the VirtualImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContainerType) GetVirtualImageOk() (*ContainerTypeAccount, bool) {
	if o == nil  {
		return nil, false
	}
	return o.VirtualImage.Get(), o.VirtualImage.IsSet()
}

// HasVirtualImage returns a boolean if a field has been set.
func (o *ContainerType) HasVirtualImage() bool {
	if o != nil && o.VirtualImage.IsSet() {
		return true
	}

	return false
}

// SetVirtualImage gets a reference to the given NullableContainerTypeAccount and assigns it to the VirtualImage field.
func (o *ContainerType) SetVirtualImage(v ContainerTypeAccount) {
	o.VirtualImage.Set(&v)
}
// SetVirtualImageNil sets the value for VirtualImage to be an explicit nil
func (o *ContainerType) SetVirtualImageNil() {
	o.VirtualImage.Set(nil)
}

// UnsetVirtualImage ensures that no value is present for VirtualImage, not even an explicit nil
func (o *ContainerType) UnsetVirtualImage() {
	o.VirtualImage.Unset()
}

// GetCategory returns the Category field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContainerType) GetCategory() string {
	if o == nil || o.Category.Get() == nil {
		var ret string
		return ret
	}
	return *o.Category.Get()
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContainerType) GetCategoryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Category.Get(), o.Category.IsSet()
}

// HasCategory returns a boolean if a field has been set.
func (o *ContainerType) HasCategory() bool {
	if o != nil && o.Category.IsSet() {
		return true
	}

	return false
}

// SetCategory gets a reference to the given NullableString and assigns it to the Category field.
func (o *ContainerType) SetCategory(v string) {
	o.Category.Set(&v)
}
// SetCategoryNil sets the value for Category to be an explicit nil
func (o *ContainerType) SetCategoryNil() {
	o.Category.Set(nil)
}

// UnsetCategory ensures that no value is present for Category, not even an explicit nil
func (o *ContainerType) UnsetCategory() {
	o.Category.Unset()
}

// GetConfig returns the Config field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContainerType) GetConfig() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}
	return o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContainerType) GetConfigOk() (*map[string]interface{}, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return &o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *ContainerType) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given map[string]interface{} and assigns it to the Config field.
func (o *ContainerType) SetConfig(v map[string]interface{}) {
	o.Config = v
}

// GetContainerPorts returns the ContainerPorts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContainerType) GetContainerPorts() []ContainerTypeContainerPorts {
	if o == nil  {
		var ret []ContainerTypeContainerPorts
		return ret
	}
	return o.ContainerPorts
}

// GetContainerPortsOk returns a tuple with the ContainerPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContainerType) GetContainerPortsOk() (*[]ContainerTypeContainerPorts, bool) {
	if o == nil || o.ContainerPorts == nil {
		return nil, false
	}
	return &o.ContainerPorts, true
}

// HasContainerPorts returns a boolean if a field has been set.
func (o *ContainerType) HasContainerPorts() bool {
	if o != nil && o.ContainerPorts != nil {
		return true
	}

	return false
}

// SetContainerPorts gets a reference to the given []ContainerTypeContainerPorts and assigns it to the ContainerPorts field.
func (o *ContainerType) SetContainerPorts(v []ContainerTypeContainerPorts) {
	o.ContainerPorts = v
}

// GetContainerScripts returns the ContainerScripts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContainerType) GetContainerScripts() []map[string]interface{} {
	if o == nil  {
		var ret []map[string]interface{}
		return ret
	}
	return o.ContainerScripts
}

// GetContainerScriptsOk returns a tuple with the ContainerScripts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContainerType) GetContainerScriptsOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.ContainerScripts == nil {
		return nil, false
	}
	return &o.ContainerScripts, true
}

// HasContainerScripts returns a boolean if a field has been set.
func (o *ContainerType) HasContainerScripts() bool {
	if o != nil && o.ContainerScripts != nil {
		return true
	}

	return false
}

// SetContainerScripts gets a reference to the given []map[string]interface{} and assigns it to the ContainerScripts field.
func (o *ContainerType) SetContainerScripts(v []map[string]interface{}) {
	o.ContainerScripts = v
}

// GetContainerTemplates returns the ContainerTemplates field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContainerType) GetContainerTemplates() []map[string]interface{} {
	if o == nil  {
		var ret []map[string]interface{}
		return ret
	}
	return o.ContainerTemplates
}

// GetContainerTemplatesOk returns a tuple with the ContainerTemplates field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContainerType) GetContainerTemplatesOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.ContainerTemplates == nil {
		return nil, false
	}
	return &o.ContainerTemplates, true
}

// HasContainerTemplates returns a boolean if a field has been set.
func (o *ContainerType) HasContainerTemplates() bool {
	if o != nil && o.ContainerTemplates != nil {
		return true
	}

	return false
}

// SetContainerTemplates gets a reference to the given []map[string]interface{} and assigns it to the ContainerTemplates field.
func (o *ContainerType) SetContainerTemplates(v []map[string]interface{}) {
	o.ContainerTemplates = v
}

// GetEnvironmentVariables returns the EnvironmentVariables field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ContainerType) GetEnvironmentVariables() []map[string]interface{} {
	if o == nil  {
		var ret []map[string]interface{}
		return ret
	}
	return o.EnvironmentVariables
}

// GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ContainerType) GetEnvironmentVariablesOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.EnvironmentVariables == nil {
		return nil, false
	}
	return &o.EnvironmentVariables, true
}

// HasEnvironmentVariables returns a boolean if a field has been set.
func (o *ContainerType) HasEnvironmentVariables() bool {
	if o != nil && o.EnvironmentVariables != nil {
		return true
	}

	return false
}

// SetEnvironmentVariables gets a reference to the given []map[string]interface{} and assigns it to the EnvironmentVariables field.
func (o *ContainerType) SetEnvironmentVariables(v []map[string]interface{}) {
	o.EnvironmentVariables = v
}

func (o ContainerType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Account.IsSet() {
		toSerialize["account"] = o.Account.Get()
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if o.ShortName != nil {
		toSerialize["shortName"] = o.ShortName
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.ContainerVersion != nil {
		toSerialize["containerVersion"] = o.ContainerVersion
	}
	if o.ProvisionType != nil {
		toSerialize["provisionType"] = o.ProvisionType
	}
	if o.VirtualImage.IsSet() {
		toSerialize["virtualImage"] = o.VirtualImage.Get()
	}
	if o.Category.IsSet() {
		toSerialize["category"] = o.Category.Get()
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if o.ContainerPorts != nil {
		toSerialize["containerPorts"] = o.ContainerPorts
	}
	if o.ContainerScripts != nil {
		toSerialize["containerScripts"] = o.ContainerScripts
	}
	if o.ContainerTemplates != nil {
		toSerialize["containerTemplates"] = o.ContainerTemplates
	}
	if o.EnvironmentVariables != nil {
		toSerialize["environmentVariables"] = o.EnvironmentVariables
	}
	return json.Marshal(toSerialize)
}

type NullableContainerType struct {
	value *ContainerType
	isSet bool
}

func (v NullableContainerType) Get() *ContainerType {
	return v.value
}

func (v *NullableContainerType) Set(val *ContainerType) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerType) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerType(val *ContainerType) *NullableContainerType {
	return &NullableContainerType{value: val, isSet: true}
}

func (v NullableContainerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

