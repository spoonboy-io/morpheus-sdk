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

// ImageBuildCreate struct for ImageBuildCreate
type ImageBuildCreate struct {
	// A name for the image build
	Name *string `json:"name,omitempty"`
	// A description for the image build
	Description NullableString `json:"description,omitempty"`
	// The image builder type.
	Type *string `json:"type,omitempty"`
	Site *ImageBuildCreateSite `json:"site,omitempty"`
	Zone *ImageBuildCreateZone `json:"zone,omitempty"`
	// A map of config values. This is the instance config that is used for provisioning. See Provisioning Types.
	Config *map[string]interface{} `json:"config,omitempty"`
	BootScript *ImageBuildCreateBootScript `json:"bootScript,omitempty"`
	PreseedScript *ImageBuildCreatePreseedScript `json:"preseedScript,omitempty"`
	// SSH Username
	SshUsername *string `json:"sshUsername,omitempty"`
	// SSH Password
	SshPassword *string `json:"sshPassword,omitempty"`
	// Storage Provider
	StorageProvider NullableString `json:"storageProvider,omitempty"`
	// Cloud Init
	IsCloudInit *string `json:"isCloudInit,omitempty"`
	// Build Output Name
	BuildOutputName NullableString `json:"buildOutputName,omitempty"`
	// Conversion Formats
	ConversionFormats NullableString `json:"conversionFormats,omitempty"`
	// Keep Results - Keep only the most recent builds. Older executions will be deleted along with their associated Virtual Images. The value 0 disables this functionality.
	KeepResults *int64 `json:"keepResults,omitempty"`
}

// NewImageBuildCreate instantiates a new ImageBuildCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageBuildCreate() *ImageBuildCreate {
	this := ImageBuildCreate{}
	var keepResults int64 = 0
	this.KeepResults = &keepResults
	return &this
}

// NewImageBuildCreateWithDefaults instantiates a new ImageBuildCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageBuildCreateWithDefaults() *ImageBuildCreate {
	this := ImageBuildCreate{}
	var keepResults int64 = 0
	this.KeepResults = &keepResults
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ImageBuildCreate) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageBuildCreate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ImageBuildCreate) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ImageBuildCreate) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImageBuildCreate) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageBuildCreate) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ImageBuildCreate) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ImageBuildCreate) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ImageBuildCreate) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ImageBuildCreate) UnsetDescription() {
	o.Description.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ImageBuildCreate) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageBuildCreate) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ImageBuildCreate) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ImageBuildCreate) SetType(v string) {
	o.Type = &v
}

// GetSite returns the Site field value if set, zero value otherwise.
func (o *ImageBuildCreate) GetSite() ImageBuildCreateSite {
	if o == nil || o.Site == nil {
		var ret ImageBuildCreateSite
		return ret
	}
	return *o.Site
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageBuildCreate) GetSiteOk() (*ImageBuildCreateSite, bool) {
	if o == nil || o.Site == nil {
		return nil, false
	}
	return o.Site, true
}

// HasSite returns a boolean if a field has been set.
func (o *ImageBuildCreate) HasSite() bool {
	if o != nil && o.Site != nil {
		return true
	}

	return false
}

// SetSite gets a reference to the given ImageBuildCreateSite and assigns it to the Site field.
func (o *ImageBuildCreate) SetSite(v ImageBuildCreateSite) {
	o.Site = &v
}

// GetZone returns the Zone field value if set, zero value otherwise.
func (o *ImageBuildCreate) GetZone() ImageBuildCreateZone {
	if o == nil || o.Zone == nil {
		var ret ImageBuildCreateZone
		return ret
	}
	return *o.Zone
}

// GetZoneOk returns a tuple with the Zone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageBuildCreate) GetZoneOk() (*ImageBuildCreateZone, bool) {
	if o == nil || o.Zone == nil {
		return nil, false
	}
	return o.Zone, true
}

// HasZone returns a boolean if a field has been set.
func (o *ImageBuildCreate) HasZone() bool {
	if o != nil && o.Zone != nil {
		return true
	}

	return false
}

// SetZone gets a reference to the given ImageBuildCreateZone and assigns it to the Zone field.
func (o *ImageBuildCreate) SetZone(v ImageBuildCreateZone) {
	o.Zone = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *ImageBuildCreate) GetConfig() map[string]interface{} {
	if o == nil || o.Config == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageBuildCreate) GetConfigOk() (*map[string]interface{}, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *ImageBuildCreate) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given map[string]interface{} and assigns it to the Config field.
func (o *ImageBuildCreate) SetConfig(v map[string]interface{}) {
	o.Config = &v
}

// GetBootScript returns the BootScript field value if set, zero value otherwise.
func (o *ImageBuildCreate) GetBootScript() ImageBuildCreateBootScript {
	if o == nil || o.BootScript == nil {
		var ret ImageBuildCreateBootScript
		return ret
	}
	return *o.BootScript
}

// GetBootScriptOk returns a tuple with the BootScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageBuildCreate) GetBootScriptOk() (*ImageBuildCreateBootScript, bool) {
	if o == nil || o.BootScript == nil {
		return nil, false
	}
	return o.BootScript, true
}

// HasBootScript returns a boolean if a field has been set.
func (o *ImageBuildCreate) HasBootScript() bool {
	if o != nil && o.BootScript != nil {
		return true
	}

	return false
}

// SetBootScript gets a reference to the given ImageBuildCreateBootScript and assigns it to the BootScript field.
func (o *ImageBuildCreate) SetBootScript(v ImageBuildCreateBootScript) {
	o.BootScript = &v
}

// GetPreseedScript returns the PreseedScript field value if set, zero value otherwise.
func (o *ImageBuildCreate) GetPreseedScript() ImageBuildCreatePreseedScript {
	if o == nil || o.PreseedScript == nil {
		var ret ImageBuildCreatePreseedScript
		return ret
	}
	return *o.PreseedScript
}

// GetPreseedScriptOk returns a tuple with the PreseedScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageBuildCreate) GetPreseedScriptOk() (*ImageBuildCreatePreseedScript, bool) {
	if o == nil || o.PreseedScript == nil {
		return nil, false
	}
	return o.PreseedScript, true
}

// HasPreseedScript returns a boolean if a field has been set.
func (o *ImageBuildCreate) HasPreseedScript() bool {
	if o != nil && o.PreseedScript != nil {
		return true
	}

	return false
}

// SetPreseedScript gets a reference to the given ImageBuildCreatePreseedScript and assigns it to the PreseedScript field.
func (o *ImageBuildCreate) SetPreseedScript(v ImageBuildCreatePreseedScript) {
	o.PreseedScript = &v
}

// GetSshUsername returns the SshUsername field value if set, zero value otherwise.
func (o *ImageBuildCreate) GetSshUsername() string {
	if o == nil || o.SshUsername == nil {
		var ret string
		return ret
	}
	return *o.SshUsername
}

// GetSshUsernameOk returns a tuple with the SshUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageBuildCreate) GetSshUsernameOk() (*string, bool) {
	if o == nil || o.SshUsername == nil {
		return nil, false
	}
	return o.SshUsername, true
}

// HasSshUsername returns a boolean if a field has been set.
func (o *ImageBuildCreate) HasSshUsername() bool {
	if o != nil && o.SshUsername != nil {
		return true
	}

	return false
}

// SetSshUsername gets a reference to the given string and assigns it to the SshUsername field.
func (o *ImageBuildCreate) SetSshUsername(v string) {
	o.SshUsername = &v
}

// GetSshPassword returns the SshPassword field value if set, zero value otherwise.
func (o *ImageBuildCreate) GetSshPassword() string {
	if o == nil || o.SshPassword == nil {
		var ret string
		return ret
	}
	return *o.SshPassword
}

// GetSshPasswordOk returns a tuple with the SshPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageBuildCreate) GetSshPasswordOk() (*string, bool) {
	if o == nil || o.SshPassword == nil {
		return nil, false
	}
	return o.SshPassword, true
}

// HasSshPassword returns a boolean if a field has been set.
func (o *ImageBuildCreate) HasSshPassword() bool {
	if o != nil && o.SshPassword != nil {
		return true
	}

	return false
}

// SetSshPassword gets a reference to the given string and assigns it to the SshPassword field.
func (o *ImageBuildCreate) SetSshPassword(v string) {
	o.SshPassword = &v
}

// GetStorageProvider returns the StorageProvider field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImageBuildCreate) GetStorageProvider() string {
	if o == nil || o.StorageProvider.Get() == nil {
		var ret string
		return ret
	}
	return *o.StorageProvider.Get()
}

// GetStorageProviderOk returns a tuple with the StorageProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageBuildCreate) GetStorageProviderOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StorageProvider.Get(), o.StorageProvider.IsSet()
}

// HasStorageProvider returns a boolean if a field has been set.
func (o *ImageBuildCreate) HasStorageProvider() bool {
	if o != nil && o.StorageProvider.IsSet() {
		return true
	}

	return false
}

// SetStorageProvider gets a reference to the given NullableString and assigns it to the StorageProvider field.
func (o *ImageBuildCreate) SetStorageProvider(v string) {
	o.StorageProvider.Set(&v)
}
// SetStorageProviderNil sets the value for StorageProvider to be an explicit nil
func (o *ImageBuildCreate) SetStorageProviderNil() {
	o.StorageProvider.Set(nil)
}

// UnsetStorageProvider ensures that no value is present for StorageProvider, not even an explicit nil
func (o *ImageBuildCreate) UnsetStorageProvider() {
	o.StorageProvider.Unset()
}

// GetIsCloudInit returns the IsCloudInit field value if set, zero value otherwise.
func (o *ImageBuildCreate) GetIsCloudInit() string {
	if o == nil || o.IsCloudInit == nil {
		var ret string
		return ret
	}
	return *o.IsCloudInit
}

// GetIsCloudInitOk returns a tuple with the IsCloudInit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageBuildCreate) GetIsCloudInitOk() (*string, bool) {
	if o == nil || o.IsCloudInit == nil {
		return nil, false
	}
	return o.IsCloudInit, true
}

// HasIsCloudInit returns a boolean if a field has been set.
func (o *ImageBuildCreate) HasIsCloudInit() bool {
	if o != nil && o.IsCloudInit != nil {
		return true
	}

	return false
}

// SetIsCloudInit gets a reference to the given string and assigns it to the IsCloudInit field.
func (o *ImageBuildCreate) SetIsCloudInit(v string) {
	o.IsCloudInit = &v
}

// GetBuildOutputName returns the BuildOutputName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImageBuildCreate) GetBuildOutputName() string {
	if o == nil || o.BuildOutputName.Get() == nil {
		var ret string
		return ret
	}
	return *o.BuildOutputName.Get()
}

// GetBuildOutputNameOk returns a tuple with the BuildOutputName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageBuildCreate) GetBuildOutputNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BuildOutputName.Get(), o.BuildOutputName.IsSet()
}

// HasBuildOutputName returns a boolean if a field has been set.
func (o *ImageBuildCreate) HasBuildOutputName() bool {
	if o != nil && o.BuildOutputName.IsSet() {
		return true
	}

	return false
}

// SetBuildOutputName gets a reference to the given NullableString and assigns it to the BuildOutputName field.
func (o *ImageBuildCreate) SetBuildOutputName(v string) {
	o.BuildOutputName.Set(&v)
}
// SetBuildOutputNameNil sets the value for BuildOutputName to be an explicit nil
func (o *ImageBuildCreate) SetBuildOutputNameNil() {
	o.BuildOutputName.Set(nil)
}

// UnsetBuildOutputName ensures that no value is present for BuildOutputName, not even an explicit nil
func (o *ImageBuildCreate) UnsetBuildOutputName() {
	o.BuildOutputName.Unset()
}

// GetConversionFormats returns the ConversionFormats field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImageBuildCreate) GetConversionFormats() string {
	if o == nil || o.ConversionFormats.Get() == nil {
		var ret string
		return ret
	}
	return *o.ConversionFormats.Get()
}

// GetConversionFormatsOk returns a tuple with the ConversionFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImageBuildCreate) GetConversionFormatsOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ConversionFormats.Get(), o.ConversionFormats.IsSet()
}

// HasConversionFormats returns a boolean if a field has been set.
func (o *ImageBuildCreate) HasConversionFormats() bool {
	if o != nil && o.ConversionFormats.IsSet() {
		return true
	}

	return false
}

// SetConversionFormats gets a reference to the given NullableString and assigns it to the ConversionFormats field.
func (o *ImageBuildCreate) SetConversionFormats(v string) {
	o.ConversionFormats.Set(&v)
}
// SetConversionFormatsNil sets the value for ConversionFormats to be an explicit nil
func (o *ImageBuildCreate) SetConversionFormatsNil() {
	o.ConversionFormats.Set(nil)
}

// UnsetConversionFormats ensures that no value is present for ConversionFormats, not even an explicit nil
func (o *ImageBuildCreate) UnsetConversionFormats() {
	o.ConversionFormats.Unset()
}

// GetKeepResults returns the KeepResults field value if set, zero value otherwise.
func (o *ImageBuildCreate) GetKeepResults() int64 {
	if o == nil || o.KeepResults == nil {
		var ret int64
		return ret
	}
	return *o.KeepResults
}

// GetKeepResultsOk returns a tuple with the KeepResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageBuildCreate) GetKeepResultsOk() (*int64, bool) {
	if o == nil || o.KeepResults == nil {
		return nil, false
	}
	return o.KeepResults, true
}

// HasKeepResults returns a boolean if a field has been set.
func (o *ImageBuildCreate) HasKeepResults() bool {
	if o != nil && o.KeepResults != nil {
		return true
	}

	return false
}

// SetKeepResults gets a reference to the given int64 and assigns it to the KeepResults field.
func (o *ImageBuildCreate) SetKeepResults(v int64) {
	o.KeepResults = &v
}

func (o ImageBuildCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Site != nil {
		toSerialize["site"] = o.Site
	}
	if o.Zone != nil {
		toSerialize["zone"] = o.Zone
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if o.BootScript != nil {
		toSerialize["bootScript"] = o.BootScript
	}
	if o.PreseedScript != nil {
		toSerialize["preseedScript"] = o.PreseedScript
	}
	if o.SshUsername != nil {
		toSerialize["sshUsername"] = o.SshUsername
	}
	if o.SshPassword != nil {
		toSerialize["sshPassword"] = o.SshPassword
	}
	if o.StorageProvider.IsSet() {
		toSerialize["storageProvider"] = o.StorageProvider.Get()
	}
	if o.IsCloudInit != nil {
		toSerialize["isCloudInit"] = o.IsCloudInit
	}
	if o.BuildOutputName.IsSet() {
		toSerialize["buildOutputName"] = o.BuildOutputName.Get()
	}
	if o.ConversionFormats.IsSet() {
		toSerialize["conversionFormats"] = o.ConversionFormats.Get()
	}
	if o.KeepResults != nil {
		toSerialize["keepResults"] = o.KeepResults
	}
	return json.Marshal(toSerialize)
}

type NullableImageBuildCreate struct {
	value *ImageBuildCreate
	isSet bool
}

func (v NullableImageBuildCreate) Get() *ImageBuildCreate {
	return v.value
}

func (v *NullableImageBuildCreate) Set(val *ImageBuildCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableImageBuildCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableImageBuildCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageBuildCreate(val *ImageBuildCreate) *NullableImageBuildCreate {
	return &NullableImageBuildCreate{value: val, isSet: true}
}

func (v NullableImageBuildCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageBuildCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


