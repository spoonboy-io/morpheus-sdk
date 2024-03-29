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

// StorageServerType struct for StorageServerType
type StorageServerType struct {
	Id *int64 `json:"id,omitempty"`
	Code *string `json:"code,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	Creatable *bool `json:"creatable,omitempty"`
	HasNamespaces *bool `json:"hasNamespaces,omitempty"`
	HasGroups *bool `json:"hasGroups,omitempty"`
	HasBlock *bool `json:"hasBlock,omitempty"`
	HasObject *bool `json:"hasObject,omitempty"`
	HasFile *bool `json:"hasFile,omitempty"`
	HasDatastore *bool `json:"hasDatastore,omitempty"`
	HasDisks *bool `json:"hasDisks,omitempty"`
	HasHosts *bool `json:"hasHosts,omitempty"`
	CreateNamespaces *bool `json:"createNamespaces,omitempty"`
	CreateGroup *bool `json:"createGroup,omitempty"`
	CreateBlock *bool `json:"createBlock,omitempty"`
	CreateObject *bool `json:"createObject,omitempty"`
	CreateFile *bool `json:"createFile,omitempty"`
	CreateDatastore *bool `json:"createDatastore,omitempty"`
	CreateDisk *bool `json:"createDisk,omitempty"`
	CreateHost *bool `json:"createHost,omitempty"`
	IconCode NullableString `json:"iconCode,omitempty"`
	HasFileBrowser *bool `json:"hasFileBrowser,omitempty"`
	OptionTypes *[]StorageServerTypeOptionTypes `json:"optionTypes,omitempty"`
	GroupOptionTypes *[]StorageServerTypeGroupOptionTypes `json:"groupOptionTypes,omitempty"`
	BucketOptionTypes []map[string]interface{} `json:"bucketOptionTypes,omitempty"`
	ShareOptionTypes []map[string]interface{} `json:"shareOptionTypes,omitempty"`
	ShareAccessOptionTypes []map[string]interface{} `json:"shareAccessOptionTypes,omitempty"`
	StorageVolumeTypes *[]StorageServerTypeStorageVolumeTypes `json:"storageVolumeTypes,omitempty"`
}

// NewStorageServerType instantiates a new StorageServerType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageServerType() *StorageServerType {
	this := StorageServerType{}
	return &this
}

// NewStorageServerTypeWithDefaults instantiates a new StorageServerType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageServerTypeWithDefaults() *StorageServerType {
	this := StorageServerType{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *StorageServerType) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageServerType) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *StorageServerType) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *StorageServerType) SetId(v int64) {
	o.Id = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *StorageServerType) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageServerType) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *StorageServerType) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *StorageServerType) SetCode(v string) {
	o.Code = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageServerType) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageServerType) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageServerType) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageServerType) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *StorageServerType) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageServerType) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *StorageServerType) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *StorageServerType) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *StorageServerType) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageServerType) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *StorageServerType) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *StorageServerType) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetCreatable returns the Creatable field value if set, zero value otherwise.
func (o *StorageServerType) GetCreatable() bool {
	if o == nil || o.Creatable == nil {
		var ret bool
		return ret
	}
	return *o.Creatable
}

// GetCreatableOk returns a tuple with the Creatable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageServerType) GetCreatableOk() (*bool, bool) {
	if o == nil || o.Creatable == nil {
		return nil, false
	}
	return o.Creatable, true
}

// HasCreatable returns a boolean if a field has been set.
func (o *StorageServerType) HasCreatable() bool {
	if o != nil && o.Creatable != nil {
		return true
	}

	return false
}

// SetCreatable gets a reference to the given bool and assigns it to the Creatable field.
func (o *StorageServerType) SetCreatable(v bool) {
	o.Creatable = &v
}

// GetHasNamespaces returns the HasNamespaces field value if set, zero value otherwise.
func (o *StorageServerType) GetHasNamespaces() bool {
	if o == nil || o.HasNamespaces == nil {
		var ret bool
		return ret
	}
	return *o.HasNamespaces
}

// GetHasNamespacesOk returns a tuple with the HasNamespaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageServerType) GetHasNamespacesOk() (*bool, bool) {
	if o == nil || o.HasNamespaces == nil {
		return nil, false
	}
	return o.HasNamespaces, true
}

// HasHasNamespaces returns a boolean if a field has been set.
func (o *StorageServerType) HasHasNamespaces() bool {
	if o != nil && o.HasNamespaces != nil {
		return true
	}

	return false
}

// SetHasNamespaces gets a reference to the given bool and assigns it to the HasNamespaces field.
func (o *StorageServerType) SetHasNamespaces(v bool) {
	o.HasNamespaces = &v
}

// GetHasGroups returns the HasGroups field value if set, zero value otherwise.
func (o *StorageServerType) GetHasGroups() bool {
	if o == nil || o.HasGroups == nil {
		var ret bool
		return ret
	}
	return *o.HasGroups
}

// GetHasGroupsOk returns a tuple with the HasGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageServerType) GetHasGroupsOk() (*bool, bool) {
	if o == nil || o.HasGroups == nil {
		return nil, false
	}
	return o.HasGroups, true
}

// HasHasGroups returns a boolean if a field has been set.
func (o *StorageServerType) HasHasGroups() bool {
	if o != nil && o.HasGroups != nil {
		return true
	}

	return false
}

// SetHasGroups gets a reference to the given bool and assigns it to the HasGroups field.
func (o *StorageServerType) SetHasGroups(v bool) {
	o.HasGroups = &v
}

// GetHasBlock returns the HasBlock field value if set, zero value otherwise.
func (o *StorageServerType) GetHasBlock() bool {
	if o == nil || o.HasBlock == nil {
		var ret bool
		return ret
	}
	return *o.HasBlock
}

// GetHasBlockOk returns a tuple with the HasBlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageServerType) GetHasBlockOk() (*bool, bool) {
	if o == nil || o.HasBlock == nil {
		return nil, false
	}
	return o.HasBlock, true
}

// HasHasBlock returns a boolean if a field has been set.
func (o *StorageServerType) HasHasBlock() bool {
	if o != nil && o.HasBlock != nil {
		return true
	}

	return false
}

// SetHasBlock gets a reference to the given bool and assigns it to the HasBlock field.
func (o *StorageServerType) SetHasBlock(v bool) {
	o.HasBlock = &v
}

// GetHasObject returns the HasObject field value if set, zero value otherwise.
func (o *StorageServerType) GetHasObject() bool {
	if o == nil || o.HasObject == nil {
		var ret bool
		return ret
	}
	return *o.HasObject
}

// GetHasObjectOk returns a tuple with the HasObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageServerType) GetHasObjectOk() (*bool, bool) {
	if o == nil || o.HasObject == nil {
		return nil, false
	}
	return o.HasObject, true
}

// HasHasObject returns a boolean if a field has been set.
func (o *StorageServerType) HasHasObject() bool {
	if o != nil && o.HasObject != nil {
		return true
	}

	return false
}

// SetHasObject gets a reference to the given bool and assigns it to the HasObject field.
func (o *StorageServerType) SetHasObject(v bool) {
	o.HasObject = &v
}

// GetHasFile returns the HasFile field value if set, zero value otherwise.
func (o *StorageServerType) GetHasFile() bool {
	if o == nil || o.HasFile == nil {
		var ret bool
		return ret
	}
	return *o.HasFile
}

// GetHasFileOk returns a tuple with the HasFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageServerType) GetHasFileOk() (*bool, bool) {
	if o == nil || o.HasFile == nil {
		return nil, false
	}
	return o.HasFile, true
}

// HasHasFile returns a boolean if a field has been set.
func (o *StorageServerType) HasHasFile() bool {
	if o != nil && o.HasFile != nil {
		return true
	}

	return false
}

// SetHasFile gets a reference to the given bool and assigns it to the HasFile field.
func (o *StorageServerType) SetHasFile(v bool) {
	o.HasFile = &v
}

// GetHasDatastore returns the HasDatastore field value if set, zero value otherwise.
func (o *StorageServerType) GetHasDatastore() bool {
	if o == nil || o.HasDatastore == nil {
		var ret bool
		return ret
	}
	return *o.HasDatastore
}

// GetHasDatastoreOk returns a tuple with the HasDatastore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageServerType) GetHasDatastoreOk() (*bool, bool) {
	if o == nil || o.HasDatastore == nil {
		return nil, false
	}
	return o.HasDatastore, true
}

// HasHasDatastore returns a boolean if a field has been set.
func (o *StorageServerType) HasHasDatastore() bool {
	if o != nil && o.HasDatastore != nil {
		return true
	}

	return false
}

// SetHasDatastore gets a reference to the given bool and assigns it to the HasDatastore field.
func (o *StorageServerType) SetHasDatastore(v bool) {
	o.HasDatastore = &v
}

// GetHasDisks returns the HasDisks field value if set, zero value otherwise.
func (o *StorageServerType) GetHasDisks() bool {
	if o == nil || o.HasDisks == nil {
		var ret bool
		return ret
	}
	return *o.HasDisks
}

// GetHasDisksOk returns a tuple with the HasDisks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageServerType) GetHasDisksOk() (*bool, bool) {
	if o == nil || o.HasDisks == nil {
		return nil, false
	}
	return o.HasDisks, true
}

// HasHasDisks returns a boolean if a field has been set.
func (o *StorageServerType) HasHasDisks() bool {
	if o != nil && o.HasDisks != nil {
		return true
	}

	return false
}

// SetHasDisks gets a reference to the given bool and assigns it to the HasDisks field.
func (o *StorageServerType) SetHasDisks(v bool) {
	o.HasDisks = &v
}

// GetHasHosts returns the HasHosts field value if set, zero value otherwise.
func (o *StorageServerType) GetHasHosts() bool {
	if o == nil || o.HasHosts == nil {
		var ret bool
		return ret
	}
	return *o.HasHosts
}

// GetHasHostsOk returns a tuple with the HasHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageServerType) GetHasHostsOk() (*bool, bool) {
	if o == nil || o.HasHosts == nil {
		return nil, false
	}
	return o.HasHosts, true
}

// HasHasHosts returns a boolean if a field has been set.
func (o *StorageServerType) HasHasHosts() bool {
	if o != nil && o.HasHosts != nil {
		return true
	}

	return false
}

// SetHasHosts gets a reference to the given bool and assigns it to the HasHosts field.
func (o *StorageServerType) SetHasHosts(v bool) {
	o.HasHosts = &v
}

// GetCreateNamespaces returns the CreateNamespaces field value if set, zero value otherwise.
func (o *StorageServerType) GetCreateNamespaces() bool {
	if o == nil || o.CreateNamespaces == nil {
		var ret bool
		return ret
	}
	return *o.CreateNamespaces
}

// GetCreateNamespacesOk returns a tuple with the CreateNamespaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageServerType) GetCreateNamespacesOk() (*bool, bool) {
	if o == nil || o.CreateNamespaces == nil {
		return nil, false
	}
	return o.CreateNamespaces, true
}

// HasCreateNamespaces returns a boolean if a field has been set.
func (o *StorageServerType) HasCreateNamespaces() bool {
	if o != nil && o.CreateNamespaces != nil {
		return true
	}

	return false
}

// SetCreateNamespaces gets a reference to the given bool and assigns it to the CreateNamespaces field.
func (o *StorageServerType) SetCreateNamespaces(v bool) {
	o.CreateNamespaces = &v
}

// GetCreateGroup returns the CreateGroup field value if set, zero value otherwise.
func (o *StorageServerType) GetCreateGroup() bool {
	if o == nil || o.CreateGroup == nil {
		var ret bool
		return ret
	}
	return *o.CreateGroup
}

// GetCreateGroupOk returns a tuple with the CreateGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageServerType) GetCreateGroupOk() (*bool, bool) {
	if o == nil || o.CreateGroup == nil {
		return nil, false
	}
	return o.CreateGroup, true
}

// HasCreateGroup returns a boolean if a field has been set.
func (o *StorageServerType) HasCreateGroup() bool {
	if o != nil && o.CreateGroup != nil {
		return true
	}

	return false
}

// SetCreateGroup gets a reference to the given bool and assigns it to the CreateGroup field.
func (o *StorageServerType) SetCreateGroup(v bool) {
	o.CreateGroup = &v
}

// GetCreateBlock returns the CreateBlock field value if set, zero value otherwise.
func (o *StorageServerType) GetCreateBlock() bool {
	if o == nil || o.CreateBlock == nil {
		var ret bool
		return ret
	}
	return *o.CreateBlock
}

// GetCreateBlockOk returns a tuple with the CreateBlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageServerType) GetCreateBlockOk() (*bool, bool) {
	if o == nil || o.CreateBlock == nil {
		return nil, false
	}
	return o.CreateBlock, true
}

// HasCreateBlock returns a boolean if a field has been set.
func (o *StorageServerType) HasCreateBlock() bool {
	if o != nil && o.CreateBlock != nil {
		return true
	}

	return false
}

// SetCreateBlock gets a reference to the given bool and assigns it to the CreateBlock field.
func (o *StorageServerType) SetCreateBlock(v bool) {
	o.CreateBlock = &v
}

// GetCreateObject returns the CreateObject field value if set, zero value otherwise.
func (o *StorageServerType) GetCreateObject() bool {
	if o == nil || o.CreateObject == nil {
		var ret bool
		return ret
	}
	return *o.CreateObject
}

// GetCreateObjectOk returns a tuple with the CreateObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageServerType) GetCreateObjectOk() (*bool, bool) {
	if o == nil || o.CreateObject == nil {
		return nil, false
	}
	return o.CreateObject, true
}

// HasCreateObject returns a boolean if a field has been set.
func (o *StorageServerType) HasCreateObject() bool {
	if o != nil && o.CreateObject != nil {
		return true
	}

	return false
}

// SetCreateObject gets a reference to the given bool and assigns it to the CreateObject field.
func (o *StorageServerType) SetCreateObject(v bool) {
	o.CreateObject = &v
}

// GetCreateFile returns the CreateFile field value if set, zero value otherwise.
func (o *StorageServerType) GetCreateFile() bool {
	if o == nil || o.CreateFile == nil {
		var ret bool
		return ret
	}
	return *o.CreateFile
}

// GetCreateFileOk returns a tuple with the CreateFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageServerType) GetCreateFileOk() (*bool, bool) {
	if o == nil || o.CreateFile == nil {
		return nil, false
	}
	return o.CreateFile, true
}

// HasCreateFile returns a boolean if a field has been set.
func (o *StorageServerType) HasCreateFile() bool {
	if o != nil && o.CreateFile != nil {
		return true
	}

	return false
}

// SetCreateFile gets a reference to the given bool and assigns it to the CreateFile field.
func (o *StorageServerType) SetCreateFile(v bool) {
	o.CreateFile = &v
}

// GetCreateDatastore returns the CreateDatastore field value if set, zero value otherwise.
func (o *StorageServerType) GetCreateDatastore() bool {
	if o == nil || o.CreateDatastore == nil {
		var ret bool
		return ret
	}
	return *o.CreateDatastore
}

// GetCreateDatastoreOk returns a tuple with the CreateDatastore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageServerType) GetCreateDatastoreOk() (*bool, bool) {
	if o == nil || o.CreateDatastore == nil {
		return nil, false
	}
	return o.CreateDatastore, true
}

// HasCreateDatastore returns a boolean if a field has been set.
func (o *StorageServerType) HasCreateDatastore() bool {
	if o != nil && o.CreateDatastore != nil {
		return true
	}

	return false
}

// SetCreateDatastore gets a reference to the given bool and assigns it to the CreateDatastore field.
func (o *StorageServerType) SetCreateDatastore(v bool) {
	o.CreateDatastore = &v
}

// GetCreateDisk returns the CreateDisk field value if set, zero value otherwise.
func (o *StorageServerType) GetCreateDisk() bool {
	if o == nil || o.CreateDisk == nil {
		var ret bool
		return ret
	}
	return *o.CreateDisk
}

// GetCreateDiskOk returns a tuple with the CreateDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageServerType) GetCreateDiskOk() (*bool, bool) {
	if o == nil || o.CreateDisk == nil {
		return nil, false
	}
	return o.CreateDisk, true
}

// HasCreateDisk returns a boolean if a field has been set.
func (o *StorageServerType) HasCreateDisk() bool {
	if o != nil && o.CreateDisk != nil {
		return true
	}

	return false
}

// SetCreateDisk gets a reference to the given bool and assigns it to the CreateDisk field.
func (o *StorageServerType) SetCreateDisk(v bool) {
	o.CreateDisk = &v
}

// GetCreateHost returns the CreateHost field value if set, zero value otherwise.
func (o *StorageServerType) GetCreateHost() bool {
	if o == nil || o.CreateHost == nil {
		var ret bool
		return ret
	}
	return *o.CreateHost
}

// GetCreateHostOk returns a tuple with the CreateHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageServerType) GetCreateHostOk() (*bool, bool) {
	if o == nil || o.CreateHost == nil {
		return nil, false
	}
	return o.CreateHost, true
}

// HasCreateHost returns a boolean if a field has been set.
func (o *StorageServerType) HasCreateHost() bool {
	if o != nil && o.CreateHost != nil {
		return true
	}

	return false
}

// SetCreateHost gets a reference to the given bool and assigns it to the CreateHost field.
func (o *StorageServerType) SetCreateHost(v bool) {
	o.CreateHost = &v
}

// GetIconCode returns the IconCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageServerType) GetIconCode() string {
	if o == nil || o.IconCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.IconCode.Get()
}

// GetIconCodeOk returns a tuple with the IconCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageServerType) GetIconCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IconCode.Get(), o.IconCode.IsSet()
}

// HasIconCode returns a boolean if a field has been set.
func (o *StorageServerType) HasIconCode() bool {
	if o != nil && o.IconCode.IsSet() {
		return true
	}

	return false
}

// SetIconCode gets a reference to the given NullableString and assigns it to the IconCode field.
func (o *StorageServerType) SetIconCode(v string) {
	o.IconCode.Set(&v)
}
// SetIconCodeNil sets the value for IconCode to be an explicit nil
func (o *StorageServerType) SetIconCodeNil() {
	o.IconCode.Set(nil)
}

// UnsetIconCode ensures that no value is present for IconCode, not even an explicit nil
func (o *StorageServerType) UnsetIconCode() {
	o.IconCode.Unset()
}

// GetHasFileBrowser returns the HasFileBrowser field value if set, zero value otherwise.
func (o *StorageServerType) GetHasFileBrowser() bool {
	if o == nil || o.HasFileBrowser == nil {
		var ret bool
		return ret
	}
	return *o.HasFileBrowser
}

// GetHasFileBrowserOk returns a tuple with the HasFileBrowser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageServerType) GetHasFileBrowserOk() (*bool, bool) {
	if o == nil || o.HasFileBrowser == nil {
		return nil, false
	}
	return o.HasFileBrowser, true
}

// HasHasFileBrowser returns a boolean if a field has been set.
func (o *StorageServerType) HasHasFileBrowser() bool {
	if o != nil && o.HasFileBrowser != nil {
		return true
	}

	return false
}

// SetHasFileBrowser gets a reference to the given bool and assigns it to the HasFileBrowser field.
func (o *StorageServerType) SetHasFileBrowser(v bool) {
	o.HasFileBrowser = &v
}

// GetOptionTypes returns the OptionTypes field value if set, zero value otherwise.
func (o *StorageServerType) GetOptionTypes() []StorageServerTypeOptionTypes {
	if o == nil || o.OptionTypes == nil {
		var ret []StorageServerTypeOptionTypes
		return ret
	}
	return *o.OptionTypes
}

// GetOptionTypesOk returns a tuple with the OptionTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageServerType) GetOptionTypesOk() (*[]StorageServerTypeOptionTypes, bool) {
	if o == nil || o.OptionTypes == nil {
		return nil, false
	}
	return o.OptionTypes, true
}

// HasOptionTypes returns a boolean if a field has been set.
func (o *StorageServerType) HasOptionTypes() bool {
	if o != nil && o.OptionTypes != nil {
		return true
	}

	return false
}

// SetOptionTypes gets a reference to the given []StorageServerTypeOptionTypes and assigns it to the OptionTypes field.
func (o *StorageServerType) SetOptionTypes(v []StorageServerTypeOptionTypes) {
	o.OptionTypes = &v
}

// GetGroupOptionTypes returns the GroupOptionTypes field value if set, zero value otherwise.
func (o *StorageServerType) GetGroupOptionTypes() []StorageServerTypeGroupOptionTypes {
	if o == nil || o.GroupOptionTypes == nil {
		var ret []StorageServerTypeGroupOptionTypes
		return ret
	}
	return *o.GroupOptionTypes
}

// GetGroupOptionTypesOk returns a tuple with the GroupOptionTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageServerType) GetGroupOptionTypesOk() (*[]StorageServerTypeGroupOptionTypes, bool) {
	if o == nil || o.GroupOptionTypes == nil {
		return nil, false
	}
	return o.GroupOptionTypes, true
}

// HasGroupOptionTypes returns a boolean if a field has been set.
func (o *StorageServerType) HasGroupOptionTypes() bool {
	if o != nil && o.GroupOptionTypes != nil {
		return true
	}

	return false
}

// SetGroupOptionTypes gets a reference to the given []StorageServerTypeGroupOptionTypes and assigns it to the GroupOptionTypes field.
func (o *StorageServerType) SetGroupOptionTypes(v []StorageServerTypeGroupOptionTypes) {
	o.GroupOptionTypes = &v
}

// GetBucketOptionTypes returns the BucketOptionTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageServerType) GetBucketOptionTypes() []map[string]interface{} {
	if o == nil  {
		var ret []map[string]interface{}
		return ret
	}
	return o.BucketOptionTypes
}

// GetBucketOptionTypesOk returns a tuple with the BucketOptionTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageServerType) GetBucketOptionTypesOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.BucketOptionTypes == nil {
		return nil, false
	}
	return &o.BucketOptionTypes, true
}

// HasBucketOptionTypes returns a boolean if a field has been set.
func (o *StorageServerType) HasBucketOptionTypes() bool {
	if o != nil && o.BucketOptionTypes != nil {
		return true
	}

	return false
}

// SetBucketOptionTypes gets a reference to the given []map[string]interface{} and assigns it to the BucketOptionTypes field.
func (o *StorageServerType) SetBucketOptionTypes(v []map[string]interface{}) {
	o.BucketOptionTypes = v
}

// GetShareOptionTypes returns the ShareOptionTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageServerType) GetShareOptionTypes() []map[string]interface{} {
	if o == nil  {
		var ret []map[string]interface{}
		return ret
	}
	return o.ShareOptionTypes
}

// GetShareOptionTypesOk returns a tuple with the ShareOptionTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageServerType) GetShareOptionTypesOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.ShareOptionTypes == nil {
		return nil, false
	}
	return &o.ShareOptionTypes, true
}

// HasShareOptionTypes returns a boolean if a field has been set.
func (o *StorageServerType) HasShareOptionTypes() bool {
	if o != nil && o.ShareOptionTypes != nil {
		return true
	}

	return false
}

// SetShareOptionTypes gets a reference to the given []map[string]interface{} and assigns it to the ShareOptionTypes field.
func (o *StorageServerType) SetShareOptionTypes(v []map[string]interface{}) {
	o.ShareOptionTypes = v
}

// GetShareAccessOptionTypes returns the ShareAccessOptionTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageServerType) GetShareAccessOptionTypes() []map[string]interface{} {
	if o == nil  {
		var ret []map[string]interface{}
		return ret
	}
	return o.ShareAccessOptionTypes
}

// GetShareAccessOptionTypesOk returns a tuple with the ShareAccessOptionTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageServerType) GetShareAccessOptionTypesOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.ShareAccessOptionTypes == nil {
		return nil, false
	}
	return &o.ShareAccessOptionTypes, true
}

// HasShareAccessOptionTypes returns a boolean if a field has been set.
func (o *StorageServerType) HasShareAccessOptionTypes() bool {
	if o != nil && o.ShareAccessOptionTypes != nil {
		return true
	}

	return false
}

// SetShareAccessOptionTypes gets a reference to the given []map[string]interface{} and assigns it to the ShareAccessOptionTypes field.
func (o *StorageServerType) SetShareAccessOptionTypes(v []map[string]interface{}) {
	o.ShareAccessOptionTypes = v
}

// GetStorageVolumeTypes returns the StorageVolumeTypes field value if set, zero value otherwise.
func (o *StorageServerType) GetStorageVolumeTypes() []StorageServerTypeStorageVolumeTypes {
	if o == nil || o.StorageVolumeTypes == nil {
		var ret []StorageServerTypeStorageVolumeTypes
		return ret
	}
	return *o.StorageVolumeTypes
}

// GetStorageVolumeTypesOk returns a tuple with the StorageVolumeTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageServerType) GetStorageVolumeTypesOk() (*[]StorageServerTypeStorageVolumeTypes, bool) {
	if o == nil || o.StorageVolumeTypes == nil {
		return nil, false
	}
	return o.StorageVolumeTypes, true
}

// HasStorageVolumeTypes returns a boolean if a field has been set.
func (o *StorageServerType) HasStorageVolumeTypes() bool {
	if o != nil && o.StorageVolumeTypes != nil {
		return true
	}

	return false
}

// SetStorageVolumeTypes gets a reference to the given []StorageServerTypeStorageVolumeTypes and assigns it to the StorageVolumeTypes field.
func (o *StorageServerType) SetStorageVolumeTypes(v []StorageServerTypeStorageVolumeTypes) {
	o.StorageVolumeTypes = &v
}

func (o StorageServerType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Creatable != nil {
		toSerialize["creatable"] = o.Creatable
	}
	if o.HasNamespaces != nil {
		toSerialize["hasNamespaces"] = o.HasNamespaces
	}
	if o.HasGroups != nil {
		toSerialize["hasGroups"] = o.HasGroups
	}
	if o.HasBlock != nil {
		toSerialize["hasBlock"] = o.HasBlock
	}
	if o.HasObject != nil {
		toSerialize["hasObject"] = o.HasObject
	}
	if o.HasFile != nil {
		toSerialize["hasFile"] = o.HasFile
	}
	if o.HasDatastore != nil {
		toSerialize["hasDatastore"] = o.HasDatastore
	}
	if o.HasDisks != nil {
		toSerialize["hasDisks"] = o.HasDisks
	}
	if o.HasHosts != nil {
		toSerialize["hasHosts"] = o.HasHosts
	}
	if o.CreateNamespaces != nil {
		toSerialize["createNamespaces"] = o.CreateNamespaces
	}
	if o.CreateGroup != nil {
		toSerialize["createGroup"] = o.CreateGroup
	}
	if o.CreateBlock != nil {
		toSerialize["createBlock"] = o.CreateBlock
	}
	if o.CreateObject != nil {
		toSerialize["createObject"] = o.CreateObject
	}
	if o.CreateFile != nil {
		toSerialize["createFile"] = o.CreateFile
	}
	if o.CreateDatastore != nil {
		toSerialize["createDatastore"] = o.CreateDatastore
	}
	if o.CreateDisk != nil {
		toSerialize["createDisk"] = o.CreateDisk
	}
	if o.CreateHost != nil {
		toSerialize["createHost"] = o.CreateHost
	}
	if o.IconCode.IsSet() {
		toSerialize["iconCode"] = o.IconCode.Get()
	}
	if o.HasFileBrowser != nil {
		toSerialize["hasFileBrowser"] = o.HasFileBrowser
	}
	if o.OptionTypes != nil {
		toSerialize["optionTypes"] = o.OptionTypes
	}
	if o.GroupOptionTypes != nil {
		toSerialize["groupOptionTypes"] = o.GroupOptionTypes
	}
	if o.BucketOptionTypes != nil {
		toSerialize["bucketOptionTypes"] = o.BucketOptionTypes
	}
	if o.ShareOptionTypes != nil {
		toSerialize["shareOptionTypes"] = o.ShareOptionTypes
	}
	if o.ShareAccessOptionTypes != nil {
		toSerialize["shareAccessOptionTypes"] = o.ShareAccessOptionTypes
	}
	if o.StorageVolumeTypes != nil {
		toSerialize["storageVolumeTypes"] = o.StorageVolumeTypes
	}
	return json.Marshal(toSerialize)
}

type NullableStorageServerType struct {
	value *StorageServerType
	isSet bool
}

func (v NullableStorageServerType) Get() *StorageServerType {
	return v.value
}

func (v *NullableStorageServerType) Set(val *StorageServerType) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageServerType) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageServerType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageServerType(val *StorageServerType) *NullableStorageServerType {
	return &NullableStorageServerType{value: val, isSet: true}
}

func (v NullableStorageServerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageServerType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


