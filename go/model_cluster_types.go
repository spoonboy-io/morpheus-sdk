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

// ClusterTypes struct for ClusterTypes
type ClusterTypes struct {
	Id *int64 `json:"id,omitempty"`
	DeployTargetService *string `json:"deployTargetService,omitempty"`
	ShortName *string `json:"shortName,omitempty"`
	ProviderType *string `json:"providerType,omitempty"`
	Code *string `json:"code,omitempty"`
	HostService *string `json:"hostService,omitempty"`
	Managed *bool `json:"managed,omitempty"`
	HasMasters *bool `json:"hasMasters,omitempty"`
	HasWorkers *bool `json:"hasWorkers,omitempty"`
	ViewSet *string `json:"viewSet,omitempty"`
	ImageCode *string `json:"imageCode,omitempty"`
	KubeCtlLocal *bool `json:"kubeCtlLocal,omitempty"`
	HasDatastore *bool `json:"hasDatastore,omitempty"`
	SupportsCloudScaling *bool `json:"supportsCloudScaling,omitempty"`
	Name *string `json:"name,omitempty"`
	HasDefaultDataDisk *bool `json:"hasDefaultDataDisk,omitempty"`
	CanManage *bool `json:"canManage,omitempty"`
	HasCluster *bool `json:"hasCluster,omitempty"`
	Description *string `json:"description,omitempty"`
	OptionTypes *[]OptionType `json:"optionTypes,omitempty"`
	ControllerTypes *[]ClusterTypesControllerTypes `json:"controllerTypes,omitempty"`
	WorkerTypes *[]ClusterTypesControllerTypes `json:"workerTypes,omitempty"`
}

// NewClusterTypes instantiates a new ClusterTypes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterTypes() *ClusterTypes {
	this := ClusterTypes{}
	return &this
}

// NewClusterTypesWithDefaults instantiates a new ClusterTypes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterTypesWithDefaults() *ClusterTypes {
	this := ClusterTypes{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ClusterTypes) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterTypes) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ClusterTypes) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ClusterTypes) SetId(v int64) {
	o.Id = &v
}

// GetDeployTargetService returns the DeployTargetService field value if set, zero value otherwise.
func (o *ClusterTypes) GetDeployTargetService() string {
	if o == nil || o.DeployTargetService == nil {
		var ret string
		return ret
	}
	return *o.DeployTargetService
}

// GetDeployTargetServiceOk returns a tuple with the DeployTargetService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterTypes) GetDeployTargetServiceOk() (*string, bool) {
	if o == nil || o.DeployTargetService == nil {
		return nil, false
	}
	return o.DeployTargetService, true
}

// HasDeployTargetService returns a boolean if a field has been set.
func (o *ClusterTypes) HasDeployTargetService() bool {
	if o != nil && o.DeployTargetService != nil {
		return true
	}

	return false
}

// SetDeployTargetService gets a reference to the given string and assigns it to the DeployTargetService field.
func (o *ClusterTypes) SetDeployTargetService(v string) {
	o.DeployTargetService = &v
}

// GetShortName returns the ShortName field value if set, zero value otherwise.
func (o *ClusterTypes) GetShortName() string {
	if o == nil || o.ShortName == nil {
		var ret string
		return ret
	}
	return *o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterTypes) GetShortNameOk() (*string, bool) {
	if o == nil || o.ShortName == nil {
		return nil, false
	}
	return o.ShortName, true
}

// HasShortName returns a boolean if a field has been set.
func (o *ClusterTypes) HasShortName() bool {
	if o != nil && o.ShortName != nil {
		return true
	}

	return false
}

// SetShortName gets a reference to the given string and assigns it to the ShortName field.
func (o *ClusterTypes) SetShortName(v string) {
	o.ShortName = &v
}

// GetProviderType returns the ProviderType field value if set, zero value otherwise.
func (o *ClusterTypes) GetProviderType() string {
	if o == nil || o.ProviderType == nil {
		var ret string
		return ret
	}
	return *o.ProviderType
}

// GetProviderTypeOk returns a tuple with the ProviderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterTypes) GetProviderTypeOk() (*string, bool) {
	if o == nil || o.ProviderType == nil {
		return nil, false
	}
	return o.ProviderType, true
}

// HasProviderType returns a boolean if a field has been set.
func (o *ClusterTypes) HasProviderType() bool {
	if o != nil && o.ProviderType != nil {
		return true
	}

	return false
}

// SetProviderType gets a reference to the given string and assigns it to the ProviderType field.
func (o *ClusterTypes) SetProviderType(v string) {
	o.ProviderType = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ClusterTypes) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterTypes) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ClusterTypes) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ClusterTypes) SetCode(v string) {
	o.Code = &v
}

// GetHostService returns the HostService field value if set, zero value otherwise.
func (o *ClusterTypes) GetHostService() string {
	if o == nil || o.HostService == nil {
		var ret string
		return ret
	}
	return *o.HostService
}

// GetHostServiceOk returns a tuple with the HostService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterTypes) GetHostServiceOk() (*string, bool) {
	if o == nil || o.HostService == nil {
		return nil, false
	}
	return o.HostService, true
}

// HasHostService returns a boolean if a field has been set.
func (o *ClusterTypes) HasHostService() bool {
	if o != nil && o.HostService != nil {
		return true
	}

	return false
}

// SetHostService gets a reference to the given string and assigns it to the HostService field.
func (o *ClusterTypes) SetHostService(v string) {
	o.HostService = &v
}

// GetManaged returns the Managed field value if set, zero value otherwise.
func (o *ClusterTypes) GetManaged() bool {
	if o == nil || o.Managed == nil {
		var ret bool
		return ret
	}
	return *o.Managed
}

// GetManagedOk returns a tuple with the Managed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterTypes) GetManagedOk() (*bool, bool) {
	if o == nil || o.Managed == nil {
		return nil, false
	}
	return o.Managed, true
}

// HasManaged returns a boolean if a field has been set.
func (o *ClusterTypes) HasManaged() bool {
	if o != nil && o.Managed != nil {
		return true
	}

	return false
}

// SetManaged gets a reference to the given bool and assigns it to the Managed field.
func (o *ClusterTypes) SetManaged(v bool) {
	o.Managed = &v
}

// GetHasMasters returns the HasMasters field value if set, zero value otherwise.
func (o *ClusterTypes) GetHasMasters() bool {
	if o == nil || o.HasMasters == nil {
		var ret bool
		return ret
	}
	return *o.HasMasters
}

// GetHasMastersOk returns a tuple with the HasMasters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterTypes) GetHasMastersOk() (*bool, bool) {
	if o == nil || o.HasMasters == nil {
		return nil, false
	}
	return o.HasMasters, true
}

// HasHasMasters returns a boolean if a field has been set.
func (o *ClusterTypes) HasHasMasters() bool {
	if o != nil && o.HasMasters != nil {
		return true
	}

	return false
}

// SetHasMasters gets a reference to the given bool and assigns it to the HasMasters field.
func (o *ClusterTypes) SetHasMasters(v bool) {
	o.HasMasters = &v
}

// GetHasWorkers returns the HasWorkers field value if set, zero value otherwise.
func (o *ClusterTypes) GetHasWorkers() bool {
	if o == nil || o.HasWorkers == nil {
		var ret bool
		return ret
	}
	return *o.HasWorkers
}

// GetHasWorkersOk returns a tuple with the HasWorkers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterTypes) GetHasWorkersOk() (*bool, bool) {
	if o == nil || o.HasWorkers == nil {
		return nil, false
	}
	return o.HasWorkers, true
}

// HasHasWorkers returns a boolean if a field has been set.
func (o *ClusterTypes) HasHasWorkers() bool {
	if o != nil && o.HasWorkers != nil {
		return true
	}

	return false
}

// SetHasWorkers gets a reference to the given bool and assigns it to the HasWorkers field.
func (o *ClusterTypes) SetHasWorkers(v bool) {
	o.HasWorkers = &v
}

// GetViewSet returns the ViewSet field value if set, zero value otherwise.
func (o *ClusterTypes) GetViewSet() string {
	if o == nil || o.ViewSet == nil {
		var ret string
		return ret
	}
	return *o.ViewSet
}

// GetViewSetOk returns a tuple with the ViewSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterTypes) GetViewSetOk() (*string, bool) {
	if o == nil || o.ViewSet == nil {
		return nil, false
	}
	return o.ViewSet, true
}

// HasViewSet returns a boolean if a field has been set.
func (o *ClusterTypes) HasViewSet() bool {
	if o != nil && o.ViewSet != nil {
		return true
	}

	return false
}

// SetViewSet gets a reference to the given string and assigns it to the ViewSet field.
func (o *ClusterTypes) SetViewSet(v string) {
	o.ViewSet = &v
}

// GetImageCode returns the ImageCode field value if set, zero value otherwise.
func (o *ClusterTypes) GetImageCode() string {
	if o == nil || o.ImageCode == nil {
		var ret string
		return ret
	}
	return *o.ImageCode
}

// GetImageCodeOk returns a tuple with the ImageCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterTypes) GetImageCodeOk() (*string, bool) {
	if o == nil || o.ImageCode == nil {
		return nil, false
	}
	return o.ImageCode, true
}

// HasImageCode returns a boolean if a field has been set.
func (o *ClusterTypes) HasImageCode() bool {
	if o != nil && o.ImageCode != nil {
		return true
	}

	return false
}

// SetImageCode gets a reference to the given string and assigns it to the ImageCode field.
func (o *ClusterTypes) SetImageCode(v string) {
	o.ImageCode = &v
}

// GetKubeCtlLocal returns the KubeCtlLocal field value if set, zero value otherwise.
func (o *ClusterTypes) GetKubeCtlLocal() bool {
	if o == nil || o.KubeCtlLocal == nil {
		var ret bool
		return ret
	}
	return *o.KubeCtlLocal
}

// GetKubeCtlLocalOk returns a tuple with the KubeCtlLocal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterTypes) GetKubeCtlLocalOk() (*bool, bool) {
	if o == nil || o.KubeCtlLocal == nil {
		return nil, false
	}
	return o.KubeCtlLocal, true
}

// HasKubeCtlLocal returns a boolean if a field has been set.
func (o *ClusterTypes) HasKubeCtlLocal() bool {
	if o != nil && o.KubeCtlLocal != nil {
		return true
	}

	return false
}

// SetKubeCtlLocal gets a reference to the given bool and assigns it to the KubeCtlLocal field.
func (o *ClusterTypes) SetKubeCtlLocal(v bool) {
	o.KubeCtlLocal = &v
}

// GetHasDatastore returns the HasDatastore field value if set, zero value otherwise.
func (o *ClusterTypes) GetHasDatastore() bool {
	if o == nil || o.HasDatastore == nil {
		var ret bool
		return ret
	}
	return *o.HasDatastore
}

// GetHasDatastoreOk returns a tuple with the HasDatastore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterTypes) GetHasDatastoreOk() (*bool, bool) {
	if o == nil || o.HasDatastore == nil {
		return nil, false
	}
	return o.HasDatastore, true
}

// HasHasDatastore returns a boolean if a field has been set.
func (o *ClusterTypes) HasHasDatastore() bool {
	if o != nil && o.HasDatastore != nil {
		return true
	}

	return false
}

// SetHasDatastore gets a reference to the given bool and assigns it to the HasDatastore field.
func (o *ClusterTypes) SetHasDatastore(v bool) {
	o.HasDatastore = &v
}

// GetSupportsCloudScaling returns the SupportsCloudScaling field value if set, zero value otherwise.
func (o *ClusterTypes) GetSupportsCloudScaling() bool {
	if o == nil || o.SupportsCloudScaling == nil {
		var ret bool
		return ret
	}
	return *o.SupportsCloudScaling
}

// GetSupportsCloudScalingOk returns a tuple with the SupportsCloudScaling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterTypes) GetSupportsCloudScalingOk() (*bool, bool) {
	if o == nil || o.SupportsCloudScaling == nil {
		return nil, false
	}
	return o.SupportsCloudScaling, true
}

// HasSupportsCloudScaling returns a boolean if a field has been set.
func (o *ClusterTypes) HasSupportsCloudScaling() bool {
	if o != nil && o.SupportsCloudScaling != nil {
		return true
	}

	return false
}

// SetSupportsCloudScaling gets a reference to the given bool and assigns it to the SupportsCloudScaling field.
func (o *ClusterTypes) SetSupportsCloudScaling(v bool) {
	o.SupportsCloudScaling = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ClusterTypes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterTypes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ClusterTypes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ClusterTypes) SetName(v string) {
	o.Name = &v
}

// GetHasDefaultDataDisk returns the HasDefaultDataDisk field value if set, zero value otherwise.
func (o *ClusterTypes) GetHasDefaultDataDisk() bool {
	if o == nil || o.HasDefaultDataDisk == nil {
		var ret bool
		return ret
	}
	return *o.HasDefaultDataDisk
}

// GetHasDefaultDataDiskOk returns a tuple with the HasDefaultDataDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterTypes) GetHasDefaultDataDiskOk() (*bool, bool) {
	if o == nil || o.HasDefaultDataDisk == nil {
		return nil, false
	}
	return o.HasDefaultDataDisk, true
}

// HasHasDefaultDataDisk returns a boolean if a field has been set.
func (o *ClusterTypes) HasHasDefaultDataDisk() bool {
	if o != nil && o.HasDefaultDataDisk != nil {
		return true
	}

	return false
}

// SetHasDefaultDataDisk gets a reference to the given bool and assigns it to the HasDefaultDataDisk field.
func (o *ClusterTypes) SetHasDefaultDataDisk(v bool) {
	o.HasDefaultDataDisk = &v
}

// GetCanManage returns the CanManage field value if set, zero value otherwise.
func (o *ClusterTypes) GetCanManage() bool {
	if o == nil || o.CanManage == nil {
		var ret bool
		return ret
	}
	return *o.CanManage
}

// GetCanManageOk returns a tuple with the CanManage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterTypes) GetCanManageOk() (*bool, bool) {
	if o == nil || o.CanManage == nil {
		return nil, false
	}
	return o.CanManage, true
}

// HasCanManage returns a boolean if a field has been set.
func (o *ClusterTypes) HasCanManage() bool {
	if o != nil && o.CanManage != nil {
		return true
	}

	return false
}

// SetCanManage gets a reference to the given bool and assigns it to the CanManage field.
func (o *ClusterTypes) SetCanManage(v bool) {
	o.CanManage = &v
}

// GetHasCluster returns the HasCluster field value if set, zero value otherwise.
func (o *ClusterTypes) GetHasCluster() bool {
	if o == nil || o.HasCluster == nil {
		var ret bool
		return ret
	}
	return *o.HasCluster
}

// GetHasClusterOk returns a tuple with the HasCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterTypes) GetHasClusterOk() (*bool, bool) {
	if o == nil || o.HasCluster == nil {
		return nil, false
	}
	return o.HasCluster, true
}

// HasHasCluster returns a boolean if a field has been set.
func (o *ClusterTypes) HasHasCluster() bool {
	if o != nil && o.HasCluster != nil {
		return true
	}

	return false
}

// SetHasCluster gets a reference to the given bool and assigns it to the HasCluster field.
func (o *ClusterTypes) SetHasCluster(v bool) {
	o.HasCluster = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ClusterTypes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterTypes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ClusterTypes) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ClusterTypes) SetDescription(v string) {
	o.Description = &v
}

// GetOptionTypes returns the OptionTypes field value if set, zero value otherwise.
func (o *ClusterTypes) GetOptionTypes() []OptionType {
	if o == nil || o.OptionTypes == nil {
		var ret []OptionType
		return ret
	}
	return *o.OptionTypes
}

// GetOptionTypesOk returns a tuple with the OptionTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterTypes) GetOptionTypesOk() (*[]OptionType, bool) {
	if o == nil || o.OptionTypes == nil {
		return nil, false
	}
	return o.OptionTypes, true
}

// HasOptionTypes returns a boolean if a field has been set.
func (o *ClusterTypes) HasOptionTypes() bool {
	if o != nil && o.OptionTypes != nil {
		return true
	}

	return false
}

// SetOptionTypes gets a reference to the given []OptionType and assigns it to the OptionTypes field.
func (o *ClusterTypes) SetOptionTypes(v []OptionType) {
	o.OptionTypes = &v
}

// GetControllerTypes returns the ControllerTypes field value if set, zero value otherwise.
func (o *ClusterTypes) GetControllerTypes() []ClusterTypesControllerTypes {
	if o == nil || o.ControllerTypes == nil {
		var ret []ClusterTypesControllerTypes
		return ret
	}
	return *o.ControllerTypes
}

// GetControllerTypesOk returns a tuple with the ControllerTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterTypes) GetControllerTypesOk() (*[]ClusterTypesControllerTypes, bool) {
	if o == nil || o.ControllerTypes == nil {
		return nil, false
	}
	return o.ControllerTypes, true
}

// HasControllerTypes returns a boolean if a field has been set.
func (o *ClusterTypes) HasControllerTypes() bool {
	if o != nil && o.ControllerTypes != nil {
		return true
	}

	return false
}

// SetControllerTypes gets a reference to the given []ClusterTypesControllerTypes and assigns it to the ControllerTypes field.
func (o *ClusterTypes) SetControllerTypes(v []ClusterTypesControllerTypes) {
	o.ControllerTypes = &v
}

// GetWorkerTypes returns the WorkerTypes field value if set, zero value otherwise.
func (o *ClusterTypes) GetWorkerTypes() []ClusterTypesControllerTypes {
	if o == nil || o.WorkerTypes == nil {
		var ret []ClusterTypesControllerTypes
		return ret
	}
	return *o.WorkerTypes
}

// GetWorkerTypesOk returns a tuple with the WorkerTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterTypes) GetWorkerTypesOk() (*[]ClusterTypesControllerTypes, bool) {
	if o == nil || o.WorkerTypes == nil {
		return nil, false
	}
	return o.WorkerTypes, true
}

// HasWorkerTypes returns a boolean if a field has been set.
func (o *ClusterTypes) HasWorkerTypes() bool {
	if o != nil && o.WorkerTypes != nil {
		return true
	}

	return false
}

// SetWorkerTypes gets a reference to the given []ClusterTypesControllerTypes and assigns it to the WorkerTypes field.
func (o *ClusterTypes) SetWorkerTypes(v []ClusterTypesControllerTypes) {
	o.WorkerTypes = &v
}

func (o ClusterTypes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.DeployTargetService != nil {
		toSerialize["deployTargetService"] = o.DeployTargetService
	}
	if o.ShortName != nil {
		toSerialize["shortName"] = o.ShortName
	}
	if o.ProviderType != nil {
		toSerialize["providerType"] = o.ProviderType
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.HostService != nil {
		toSerialize["hostService"] = o.HostService
	}
	if o.Managed != nil {
		toSerialize["managed"] = o.Managed
	}
	if o.HasMasters != nil {
		toSerialize["hasMasters"] = o.HasMasters
	}
	if o.HasWorkers != nil {
		toSerialize["hasWorkers"] = o.HasWorkers
	}
	if o.ViewSet != nil {
		toSerialize["viewSet"] = o.ViewSet
	}
	if o.ImageCode != nil {
		toSerialize["imageCode"] = o.ImageCode
	}
	if o.KubeCtlLocal != nil {
		toSerialize["kubeCtlLocal"] = o.KubeCtlLocal
	}
	if o.HasDatastore != nil {
		toSerialize["hasDatastore"] = o.HasDatastore
	}
	if o.SupportsCloudScaling != nil {
		toSerialize["supportsCloudScaling"] = o.SupportsCloudScaling
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.HasDefaultDataDisk != nil {
		toSerialize["hasDefaultDataDisk"] = o.HasDefaultDataDisk
	}
	if o.CanManage != nil {
		toSerialize["canManage"] = o.CanManage
	}
	if o.HasCluster != nil {
		toSerialize["hasCluster"] = o.HasCluster
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.OptionTypes != nil {
		toSerialize["optionTypes"] = o.OptionTypes
	}
	if o.ControllerTypes != nil {
		toSerialize["controllerTypes"] = o.ControllerTypes
	}
	if o.WorkerTypes != nil {
		toSerialize["workerTypes"] = o.WorkerTypes
	}
	return json.Marshal(toSerialize)
}

type NullableClusterTypes struct {
	value *ClusterTypes
	isSet bool
}

func (v NullableClusterTypes) Get() *ClusterTypes {
	return v.value
}

func (v *NullableClusterTypes) Set(val *ClusterTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterTypes(val *ClusterTypes) *NullableClusterTypes {
	return &NullableClusterTypes{value: val, isSet: true}
}

func (v NullableClusterTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


