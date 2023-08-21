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

// ZoneFolder struct for ZoneFolder
type ZoneFolder struct {
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Zone *InlineResponse20040AppDeployInstance `json:"zone,omitempty"`
	Parent NullableInlineResponse20082LoadBalancerInstanceSslCert `json:"parent,omitempty"`
	Type *string `json:"type,omitempty"`
	ExternalId *string `json:"externalId,omitempty"`
	Visibility *string `json:"visibility,omitempty"`
	ReadOnly *bool `json:"readOnly,omitempty"`
	DefaultFolder *bool `json:"defaultFolder,omitempty"`
	DefaultStore *bool `json:"defaultStore,omitempty"`
	Active *bool `json:"active,omitempty"`
	Tenants *[]ZoneDatastoreTenants `json:"tenants,omitempty"`
	ResourcePermissions *ResourcePermissions `json:"resourcePermissions,omitempty"`
	Depth *int64 `json:"depth,omitempty"`
}

// NewZoneFolder instantiates a new ZoneFolder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZoneFolder() *ZoneFolder {
	this := ZoneFolder{}
	return &this
}

// NewZoneFolderWithDefaults instantiates a new ZoneFolder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZoneFolderWithDefaults() *ZoneFolder {
	this := ZoneFolder{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ZoneFolder) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneFolder) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ZoneFolder) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ZoneFolder) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ZoneFolder) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneFolder) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ZoneFolder) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ZoneFolder) SetName(v string) {
	o.Name = &v
}

// GetZone returns the Zone field value if set, zero value otherwise.
func (o *ZoneFolder) GetZone() InlineResponse20040AppDeployInstance {
	if o == nil || o.Zone == nil {
		var ret InlineResponse20040AppDeployInstance
		return ret
	}
	return *o.Zone
}

// GetZoneOk returns a tuple with the Zone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneFolder) GetZoneOk() (*InlineResponse20040AppDeployInstance, bool) {
	if o == nil || o.Zone == nil {
		return nil, false
	}
	return o.Zone, true
}

// HasZone returns a boolean if a field has been set.
func (o *ZoneFolder) HasZone() bool {
	if o != nil && o.Zone != nil {
		return true
	}

	return false
}

// SetZone gets a reference to the given InlineResponse20040AppDeployInstance and assigns it to the Zone field.
func (o *ZoneFolder) SetZone(v InlineResponse20040AppDeployInstance) {
	o.Zone = &v
}

// GetParent returns the Parent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ZoneFolder) GetParent() InlineResponse20082LoadBalancerInstanceSslCert {
	if o == nil || o.Parent.Get() == nil {
		var ret InlineResponse20082LoadBalancerInstanceSslCert
		return ret
	}
	return *o.Parent.Get()
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ZoneFolder) GetParentOk() (*InlineResponse20082LoadBalancerInstanceSslCert, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Parent.Get(), o.Parent.IsSet()
}

// HasParent returns a boolean if a field has been set.
func (o *ZoneFolder) HasParent() bool {
	if o != nil && o.Parent.IsSet() {
		return true
	}

	return false
}

// SetParent gets a reference to the given NullableInlineResponse20082LoadBalancerInstanceSslCert and assigns it to the Parent field.
func (o *ZoneFolder) SetParent(v InlineResponse20082LoadBalancerInstanceSslCert) {
	o.Parent.Set(&v)
}
// SetParentNil sets the value for Parent to be an explicit nil
func (o *ZoneFolder) SetParentNil() {
	o.Parent.Set(nil)
}

// UnsetParent ensures that no value is present for Parent, not even an explicit nil
func (o *ZoneFolder) UnsetParent() {
	o.Parent.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ZoneFolder) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneFolder) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ZoneFolder) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ZoneFolder) SetType(v string) {
	o.Type = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *ZoneFolder) GetExternalId() string {
	if o == nil || o.ExternalId == nil {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneFolder) GetExternalIdOk() (*string, bool) {
	if o == nil || o.ExternalId == nil {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *ZoneFolder) HasExternalId() bool {
	if o != nil && o.ExternalId != nil {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *ZoneFolder) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *ZoneFolder) GetVisibility() string {
	if o == nil || o.Visibility == nil {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneFolder) GetVisibilityOk() (*string, bool) {
	if o == nil || o.Visibility == nil {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *ZoneFolder) HasVisibility() bool {
	if o != nil && o.Visibility != nil {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *ZoneFolder) SetVisibility(v string) {
	o.Visibility = &v
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *ZoneFolder) GetReadOnly() bool {
	if o == nil || o.ReadOnly == nil {
		var ret bool
		return ret
	}
	return *o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneFolder) GetReadOnlyOk() (*bool, bool) {
	if o == nil || o.ReadOnly == nil {
		return nil, false
	}
	return o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *ZoneFolder) HasReadOnly() bool {
	if o != nil && o.ReadOnly != nil {
		return true
	}

	return false
}

// SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.
func (o *ZoneFolder) SetReadOnly(v bool) {
	o.ReadOnly = &v
}

// GetDefaultFolder returns the DefaultFolder field value if set, zero value otherwise.
func (o *ZoneFolder) GetDefaultFolder() bool {
	if o == nil || o.DefaultFolder == nil {
		var ret bool
		return ret
	}
	return *o.DefaultFolder
}

// GetDefaultFolderOk returns a tuple with the DefaultFolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneFolder) GetDefaultFolderOk() (*bool, bool) {
	if o == nil || o.DefaultFolder == nil {
		return nil, false
	}
	return o.DefaultFolder, true
}

// HasDefaultFolder returns a boolean if a field has been set.
func (o *ZoneFolder) HasDefaultFolder() bool {
	if o != nil && o.DefaultFolder != nil {
		return true
	}

	return false
}

// SetDefaultFolder gets a reference to the given bool and assigns it to the DefaultFolder field.
func (o *ZoneFolder) SetDefaultFolder(v bool) {
	o.DefaultFolder = &v
}

// GetDefaultStore returns the DefaultStore field value if set, zero value otherwise.
func (o *ZoneFolder) GetDefaultStore() bool {
	if o == nil || o.DefaultStore == nil {
		var ret bool
		return ret
	}
	return *o.DefaultStore
}

// GetDefaultStoreOk returns a tuple with the DefaultStore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneFolder) GetDefaultStoreOk() (*bool, bool) {
	if o == nil || o.DefaultStore == nil {
		return nil, false
	}
	return o.DefaultStore, true
}

// HasDefaultStore returns a boolean if a field has been set.
func (o *ZoneFolder) HasDefaultStore() bool {
	if o != nil && o.DefaultStore != nil {
		return true
	}

	return false
}

// SetDefaultStore gets a reference to the given bool and assigns it to the DefaultStore field.
func (o *ZoneFolder) SetDefaultStore(v bool) {
	o.DefaultStore = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ZoneFolder) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneFolder) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *ZoneFolder) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *ZoneFolder) SetActive(v bool) {
	o.Active = &v
}

// GetTenants returns the Tenants field value if set, zero value otherwise.
func (o *ZoneFolder) GetTenants() []ZoneDatastoreTenants {
	if o == nil || o.Tenants == nil {
		var ret []ZoneDatastoreTenants
		return ret
	}
	return *o.Tenants
}

// GetTenantsOk returns a tuple with the Tenants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneFolder) GetTenantsOk() (*[]ZoneDatastoreTenants, bool) {
	if o == nil || o.Tenants == nil {
		return nil, false
	}
	return o.Tenants, true
}

// HasTenants returns a boolean if a field has been set.
func (o *ZoneFolder) HasTenants() bool {
	if o != nil && o.Tenants != nil {
		return true
	}

	return false
}

// SetTenants gets a reference to the given []ZoneDatastoreTenants and assigns it to the Tenants field.
func (o *ZoneFolder) SetTenants(v []ZoneDatastoreTenants) {
	o.Tenants = &v
}

// GetResourcePermissions returns the ResourcePermissions field value if set, zero value otherwise.
func (o *ZoneFolder) GetResourcePermissions() ResourcePermissions {
	if o == nil || o.ResourcePermissions == nil {
		var ret ResourcePermissions
		return ret
	}
	return *o.ResourcePermissions
}

// GetResourcePermissionsOk returns a tuple with the ResourcePermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneFolder) GetResourcePermissionsOk() (*ResourcePermissions, bool) {
	if o == nil || o.ResourcePermissions == nil {
		return nil, false
	}
	return o.ResourcePermissions, true
}

// HasResourcePermissions returns a boolean if a field has been set.
func (o *ZoneFolder) HasResourcePermissions() bool {
	if o != nil && o.ResourcePermissions != nil {
		return true
	}

	return false
}

// SetResourcePermissions gets a reference to the given ResourcePermissions and assigns it to the ResourcePermissions field.
func (o *ZoneFolder) SetResourcePermissions(v ResourcePermissions) {
	o.ResourcePermissions = &v
}

// GetDepth returns the Depth field value if set, zero value otherwise.
func (o *ZoneFolder) GetDepth() int64 {
	if o == nil || o.Depth == nil {
		var ret int64
		return ret
	}
	return *o.Depth
}

// GetDepthOk returns a tuple with the Depth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneFolder) GetDepthOk() (*int64, bool) {
	if o == nil || o.Depth == nil {
		return nil, false
	}
	return o.Depth, true
}

// HasDepth returns a boolean if a field has been set.
func (o *ZoneFolder) HasDepth() bool {
	if o != nil && o.Depth != nil {
		return true
	}

	return false
}

// SetDepth gets a reference to the given int64 and assigns it to the Depth field.
func (o *ZoneFolder) SetDepth(v int64) {
	o.Depth = &v
}

func (o ZoneFolder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Zone != nil {
		toSerialize["zone"] = o.Zone
	}
	if o.Parent.IsSet() {
		toSerialize["parent"] = o.Parent.Get()
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.ExternalId != nil {
		toSerialize["externalId"] = o.ExternalId
	}
	if o.Visibility != nil {
		toSerialize["visibility"] = o.Visibility
	}
	if o.ReadOnly != nil {
		toSerialize["readOnly"] = o.ReadOnly
	}
	if o.DefaultFolder != nil {
		toSerialize["defaultFolder"] = o.DefaultFolder
	}
	if o.DefaultStore != nil {
		toSerialize["defaultStore"] = o.DefaultStore
	}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.Tenants != nil {
		toSerialize["tenants"] = o.Tenants
	}
	if o.ResourcePermissions != nil {
		toSerialize["resourcePermissions"] = o.ResourcePermissions
	}
	if o.Depth != nil {
		toSerialize["depth"] = o.Depth
	}
	return json.Marshal(toSerialize)
}

type NullableZoneFolder struct {
	value *ZoneFolder
	isSet bool
}

func (v NullableZoneFolder) Get() *ZoneFolder {
	return v.value
}

func (v *NullableZoneFolder) Set(val *ZoneFolder) {
	v.value = val
	v.isSet = true
}

func (v NullableZoneFolder) IsSet() bool {
	return v.isSet
}

func (v *NullableZoneFolder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZoneFolder(val *ZoneFolder) *NullableZoneFolder {
	return &NullableZoneFolder{value: val, isSet: true}
}

func (v NullableZoneFolder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZoneFolder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

