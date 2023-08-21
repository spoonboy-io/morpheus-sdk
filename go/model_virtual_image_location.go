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

// VirtualImageLocation struct for VirtualImageLocation
type VirtualImageLocation struct {
	Id *int64 `json:"id,omitempty"`
	Cloud *InlineResponse20079LoadBalancerMonitorLoadBalancerType `json:"cloud,omitempty"`
	Code *string `json:"code,omitempty"`
	InternalId NullableString `json:"internalId,omitempty"`
	ExternalId *string `json:"externalId,omitempty"`
	ExternalDiskId *string `json:"externalDiskId,omitempty"`
	RemotePath NullableString `json:"remotePath,omitempty"`
	ImagePath NullableString `json:"imagePath,omitempty"`
	ImageName *string `json:"imageName,omitempty"`
	ImageRegion *string `json:"imageRegion,omitempty"`
	ImageFolder NullableString `json:"imageFolder,omitempty"`
	RefType *string `json:"refType,omitempty"`
	RefId *int64 `json:"refId,omitempty"`
	NodeRefType NullableString `json:"nodeRefType,omitempty"`
	NodeRefId NullableString `json:"nodeRefId,omitempty"`
	SubRefType NullableString `json:"subRefType,omitempty"`
	SubRefId NullableString `json:"subRefId,omitempty"`
	IsPublic *bool `json:"isPublic,omitempty"`
	SystemImage *bool `json:"systemImage,omitempty"`
	DiskIndex *int64 `json:"diskIndex,omitempty"`
	PricePlan NullableString `json:"pricePlan,omitempty"`
	Volumes *[]map[string]interface{} `json:"volumes,omitempty"`
	StorageControllers *[]map[string]interface{} `json:"storageControllers,omitempty"`
	NetworkInterfaces *[]map[string]interface{} `json:"networkInterfaces,omitempty"`
	VirtualImage *VirtualImageLocationVirtualImage `json:"virtualImage,omitempty"`
}

// NewVirtualImageLocation instantiates a new VirtualImageLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualImageLocation() *VirtualImageLocation {
	this := VirtualImageLocation{}
	return &this
}

// NewVirtualImageLocationWithDefaults instantiates a new VirtualImageLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualImageLocationWithDefaults() *VirtualImageLocation {
	this := VirtualImageLocation{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VirtualImageLocation) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualImageLocation) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VirtualImageLocation) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *VirtualImageLocation) SetId(v int64) {
	o.Id = &v
}

// GetCloud returns the Cloud field value if set, zero value otherwise.
func (o *VirtualImageLocation) GetCloud() InlineResponse20079LoadBalancerMonitorLoadBalancerType {
	if o == nil || o.Cloud == nil {
		var ret InlineResponse20079LoadBalancerMonitorLoadBalancerType
		return ret
	}
	return *o.Cloud
}

// GetCloudOk returns a tuple with the Cloud field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualImageLocation) GetCloudOk() (*InlineResponse20079LoadBalancerMonitorLoadBalancerType, bool) {
	if o == nil || o.Cloud == nil {
		return nil, false
	}
	return o.Cloud, true
}

// HasCloud returns a boolean if a field has been set.
func (o *VirtualImageLocation) HasCloud() bool {
	if o != nil && o.Cloud != nil {
		return true
	}

	return false
}

// SetCloud gets a reference to the given InlineResponse20079LoadBalancerMonitorLoadBalancerType and assigns it to the Cloud field.
func (o *VirtualImageLocation) SetCloud(v InlineResponse20079LoadBalancerMonitorLoadBalancerType) {
	o.Cloud = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *VirtualImageLocation) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualImageLocation) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *VirtualImageLocation) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *VirtualImageLocation) SetCode(v string) {
	o.Code = &v
}

// GetInternalId returns the InternalId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualImageLocation) GetInternalId() string {
	if o == nil || o.InternalId.Get() == nil {
		var ret string
		return ret
	}
	return *o.InternalId.Get()
}

// GetInternalIdOk returns a tuple with the InternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualImageLocation) GetInternalIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.InternalId.Get(), o.InternalId.IsSet()
}

// HasInternalId returns a boolean if a field has been set.
func (o *VirtualImageLocation) HasInternalId() bool {
	if o != nil && o.InternalId.IsSet() {
		return true
	}

	return false
}

// SetInternalId gets a reference to the given NullableString and assigns it to the InternalId field.
func (o *VirtualImageLocation) SetInternalId(v string) {
	o.InternalId.Set(&v)
}
// SetInternalIdNil sets the value for InternalId to be an explicit nil
func (o *VirtualImageLocation) SetInternalIdNil() {
	o.InternalId.Set(nil)
}

// UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
func (o *VirtualImageLocation) UnsetInternalId() {
	o.InternalId.Unset()
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *VirtualImageLocation) GetExternalId() string {
	if o == nil || o.ExternalId == nil {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualImageLocation) GetExternalIdOk() (*string, bool) {
	if o == nil || o.ExternalId == nil {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *VirtualImageLocation) HasExternalId() bool {
	if o != nil && o.ExternalId != nil {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *VirtualImageLocation) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetExternalDiskId returns the ExternalDiskId field value if set, zero value otherwise.
func (o *VirtualImageLocation) GetExternalDiskId() string {
	if o == nil || o.ExternalDiskId == nil {
		var ret string
		return ret
	}
	return *o.ExternalDiskId
}

// GetExternalDiskIdOk returns a tuple with the ExternalDiskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualImageLocation) GetExternalDiskIdOk() (*string, bool) {
	if o == nil || o.ExternalDiskId == nil {
		return nil, false
	}
	return o.ExternalDiskId, true
}

// HasExternalDiskId returns a boolean if a field has been set.
func (o *VirtualImageLocation) HasExternalDiskId() bool {
	if o != nil && o.ExternalDiskId != nil {
		return true
	}

	return false
}

// SetExternalDiskId gets a reference to the given string and assigns it to the ExternalDiskId field.
func (o *VirtualImageLocation) SetExternalDiskId(v string) {
	o.ExternalDiskId = &v
}

// GetRemotePath returns the RemotePath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualImageLocation) GetRemotePath() string {
	if o == nil || o.RemotePath.Get() == nil {
		var ret string
		return ret
	}
	return *o.RemotePath.Get()
}

// GetRemotePathOk returns a tuple with the RemotePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualImageLocation) GetRemotePathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RemotePath.Get(), o.RemotePath.IsSet()
}

// HasRemotePath returns a boolean if a field has been set.
func (o *VirtualImageLocation) HasRemotePath() bool {
	if o != nil && o.RemotePath.IsSet() {
		return true
	}

	return false
}

// SetRemotePath gets a reference to the given NullableString and assigns it to the RemotePath field.
func (o *VirtualImageLocation) SetRemotePath(v string) {
	o.RemotePath.Set(&v)
}
// SetRemotePathNil sets the value for RemotePath to be an explicit nil
func (o *VirtualImageLocation) SetRemotePathNil() {
	o.RemotePath.Set(nil)
}

// UnsetRemotePath ensures that no value is present for RemotePath, not even an explicit nil
func (o *VirtualImageLocation) UnsetRemotePath() {
	o.RemotePath.Unset()
}

// GetImagePath returns the ImagePath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualImageLocation) GetImagePath() string {
	if o == nil || o.ImagePath.Get() == nil {
		var ret string
		return ret
	}
	return *o.ImagePath.Get()
}

// GetImagePathOk returns a tuple with the ImagePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualImageLocation) GetImagePathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ImagePath.Get(), o.ImagePath.IsSet()
}

// HasImagePath returns a boolean if a field has been set.
func (o *VirtualImageLocation) HasImagePath() bool {
	if o != nil && o.ImagePath.IsSet() {
		return true
	}

	return false
}

// SetImagePath gets a reference to the given NullableString and assigns it to the ImagePath field.
func (o *VirtualImageLocation) SetImagePath(v string) {
	o.ImagePath.Set(&v)
}
// SetImagePathNil sets the value for ImagePath to be an explicit nil
func (o *VirtualImageLocation) SetImagePathNil() {
	o.ImagePath.Set(nil)
}

// UnsetImagePath ensures that no value is present for ImagePath, not even an explicit nil
func (o *VirtualImageLocation) UnsetImagePath() {
	o.ImagePath.Unset()
}

// GetImageName returns the ImageName field value if set, zero value otherwise.
func (o *VirtualImageLocation) GetImageName() string {
	if o == nil || o.ImageName == nil {
		var ret string
		return ret
	}
	return *o.ImageName
}

// GetImageNameOk returns a tuple with the ImageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualImageLocation) GetImageNameOk() (*string, bool) {
	if o == nil || o.ImageName == nil {
		return nil, false
	}
	return o.ImageName, true
}

// HasImageName returns a boolean if a field has been set.
func (o *VirtualImageLocation) HasImageName() bool {
	if o != nil && o.ImageName != nil {
		return true
	}

	return false
}

// SetImageName gets a reference to the given string and assigns it to the ImageName field.
func (o *VirtualImageLocation) SetImageName(v string) {
	o.ImageName = &v
}

// GetImageRegion returns the ImageRegion field value if set, zero value otherwise.
func (o *VirtualImageLocation) GetImageRegion() string {
	if o == nil || o.ImageRegion == nil {
		var ret string
		return ret
	}
	return *o.ImageRegion
}

// GetImageRegionOk returns a tuple with the ImageRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualImageLocation) GetImageRegionOk() (*string, bool) {
	if o == nil || o.ImageRegion == nil {
		return nil, false
	}
	return o.ImageRegion, true
}

// HasImageRegion returns a boolean if a field has been set.
func (o *VirtualImageLocation) HasImageRegion() bool {
	if o != nil && o.ImageRegion != nil {
		return true
	}

	return false
}

// SetImageRegion gets a reference to the given string and assigns it to the ImageRegion field.
func (o *VirtualImageLocation) SetImageRegion(v string) {
	o.ImageRegion = &v
}

// GetImageFolder returns the ImageFolder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualImageLocation) GetImageFolder() string {
	if o == nil || o.ImageFolder.Get() == nil {
		var ret string
		return ret
	}
	return *o.ImageFolder.Get()
}

// GetImageFolderOk returns a tuple with the ImageFolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualImageLocation) GetImageFolderOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ImageFolder.Get(), o.ImageFolder.IsSet()
}

// HasImageFolder returns a boolean if a field has been set.
func (o *VirtualImageLocation) HasImageFolder() bool {
	if o != nil && o.ImageFolder.IsSet() {
		return true
	}

	return false
}

// SetImageFolder gets a reference to the given NullableString and assigns it to the ImageFolder field.
func (o *VirtualImageLocation) SetImageFolder(v string) {
	o.ImageFolder.Set(&v)
}
// SetImageFolderNil sets the value for ImageFolder to be an explicit nil
func (o *VirtualImageLocation) SetImageFolderNil() {
	o.ImageFolder.Set(nil)
}

// UnsetImageFolder ensures that no value is present for ImageFolder, not even an explicit nil
func (o *VirtualImageLocation) UnsetImageFolder() {
	o.ImageFolder.Unset()
}

// GetRefType returns the RefType field value if set, zero value otherwise.
func (o *VirtualImageLocation) GetRefType() string {
	if o == nil || o.RefType == nil {
		var ret string
		return ret
	}
	return *o.RefType
}

// GetRefTypeOk returns a tuple with the RefType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualImageLocation) GetRefTypeOk() (*string, bool) {
	if o == nil || o.RefType == nil {
		return nil, false
	}
	return o.RefType, true
}

// HasRefType returns a boolean if a field has been set.
func (o *VirtualImageLocation) HasRefType() bool {
	if o != nil && o.RefType != nil {
		return true
	}

	return false
}

// SetRefType gets a reference to the given string and assigns it to the RefType field.
func (o *VirtualImageLocation) SetRefType(v string) {
	o.RefType = &v
}

// GetRefId returns the RefId field value if set, zero value otherwise.
func (o *VirtualImageLocation) GetRefId() int64 {
	if o == nil || o.RefId == nil {
		var ret int64
		return ret
	}
	return *o.RefId
}

// GetRefIdOk returns a tuple with the RefId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualImageLocation) GetRefIdOk() (*int64, bool) {
	if o == nil || o.RefId == nil {
		return nil, false
	}
	return o.RefId, true
}

// HasRefId returns a boolean if a field has been set.
func (o *VirtualImageLocation) HasRefId() bool {
	if o != nil && o.RefId != nil {
		return true
	}

	return false
}

// SetRefId gets a reference to the given int64 and assigns it to the RefId field.
func (o *VirtualImageLocation) SetRefId(v int64) {
	o.RefId = &v
}

// GetNodeRefType returns the NodeRefType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualImageLocation) GetNodeRefType() string {
	if o == nil || o.NodeRefType.Get() == nil {
		var ret string
		return ret
	}
	return *o.NodeRefType.Get()
}

// GetNodeRefTypeOk returns a tuple with the NodeRefType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualImageLocation) GetNodeRefTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NodeRefType.Get(), o.NodeRefType.IsSet()
}

// HasNodeRefType returns a boolean if a field has been set.
func (o *VirtualImageLocation) HasNodeRefType() bool {
	if o != nil && o.NodeRefType.IsSet() {
		return true
	}

	return false
}

// SetNodeRefType gets a reference to the given NullableString and assigns it to the NodeRefType field.
func (o *VirtualImageLocation) SetNodeRefType(v string) {
	o.NodeRefType.Set(&v)
}
// SetNodeRefTypeNil sets the value for NodeRefType to be an explicit nil
func (o *VirtualImageLocation) SetNodeRefTypeNil() {
	o.NodeRefType.Set(nil)
}

// UnsetNodeRefType ensures that no value is present for NodeRefType, not even an explicit nil
func (o *VirtualImageLocation) UnsetNodeRefType() {
	o.NodeRefType.Unset()
}

// GetNodeRefId returns the NodeRefId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualImageLocation) GetNodeRefId() string {
	if o == nil || o.NodeRefId.Get() == nil {
		var ret string
		return ret
	}
	return *o.NodeRefId.Get()
}

// GetNodeRefIdOk returns a tuple with the NodeRefId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualImageLocation) GetNodeRefIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NodeRefId.Get(), o.NodeRefId.IsSet()
}

// HasNodeRefId returns a boolean if a field has been set.
func (o *VirtualImageLocation) HasNodeRefId() bool {
	if o != nil && o.NodeRefId.IsSet() {
		return true
	}

	return false
}

// SetNodeRefId gets a reference to the given NullableString and assigns it to the NodeRefId field.
func (o *VirtualImageLocation) SetNodeRefId(v string) {
	o.NodeRefId.Set(&v)
}
// SetNodeRefIdNil sets the value for NodeRefId to be an explicit nil
func (o *VirtualImageLocation) SetNodeRefIdNil() {
	o.NodeRefId.Set(nil)
}

// UnsetNodeRefId ensures that no value is present for NodeRefId, not even an explicit nil
func (o *VirtualImageLocation) UnsetNodeRefId() {
	o.NodeRefId.Unset()
}

// GetSubRefType returns the SubRefType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualImageLocation) GetSubRefType() string {
	if o == nil || o.SubRefType.Get() == nil {
		var ret string
		return ret
	}
	return *o.SubRefType.Get()
}

// GetSubRefTypeOk returns a tuple with the SubRefType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualImageLocation) GetSubRefTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SubRefType.Get(), o.SubRefType.IsSet()
}

// HasSubRefType returns a boolean if a field has been set.
func (o *VirtualImageLocation) HasSubRefType() bool {
	if o != nil && o.SubRefType.IsSet() {
		return true
	}

	return false
}

// SetSubRefType gets a reference to the given NullableString and assigns it to the SubRefType field.
func (o *VirtualImageLocation) SetSubRefType(v string) {
	o.SubRefType.Set(&v)
}
// SetSubRefTypeNil sets the value for SubRefType to be an explicit nil
func (o *VirtualImageLocation) SetSubRefTypeNil() {
	o.SubRefType.Set(nil)
}

// UnsetSubRefType ensures that no value is present for SubRefType, not even an explicit nil
func (o *VirtualImageLocation) UnsetSubRefType() {
	o.SubRefType.Unset()
}

// GetSubRefId returns the SubRefId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualImageLocation) GetSubRefId() string {
	if o == nil || o.SubRefId.Get() == nil {
		var ret string
		return ret
	}
	return *o.SubRefId.Get()
}

// GetSubRefIdOk returns a tuple with the SubRefId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualImageLocation) GetSubRefIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SubRefId.Get(), o.SubRefId.IsSet()
}

// HasSubRefId returns a boolean if a field has been set.
func (o *VirtualImageLocation) HasSubRefId() bool {
	if o != nil && o.SubRefId.IsSet() {
		return true
	}

	return false
}

// SetSubRefId gets a reference to the given NullableString and assigns it to the SubRefId field.
func (o *VirtualImageLocation) SetSubRefId(v string) {
	o.SubRefId.Set(&v)
}
// SetSubRefIdNil sets the value for SubRefId to be an explicit nil
func (o *VirtualImageLocation) SetSubRefIdNil() {
	o.SubRefId.Set(nil)
}

// UnsetSubRefId ensures that no value is present for SubRefId, not even an explicit nil
func (o *VirtualImageLocation) UnsetSubRefId() {
	o.SubRefId.Unset()
}

// GetIsPublic returns the IsPublic field value if set, zero value otherwise.
func (o *VirtualImageLocation) GetIsPublic() bool {
	if o == nil || o.IsPublic == nil {
		var ret bool
		return ret
	}
	return *o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualImageLocation) GetIsPublicOk() (*bool, bool) {
	if o == nil || o.IsPublic == nil {
		return nil, false
	}
	return o.IsPublic, true
}

// HasIsPublic returns a boolean if a field has been set.
func (o *VirtualImageLocation) HasIsPublic() bool {
	if o != nil && o.IsPublic != nil {
		return true
	}

	return false
}

// SetIsPublic gets a reference to the given bool and assigns it to the IsPublic field.
func (o *VirtualImageLocation) SetIsPublic(v bool) {
	o.IsPublic = &v
}

// GetSystemImage returns the SystemImage field value if set, zero value otherwise.
func (o *VirtualImageLocation) GetSystemImage() bool {
	if o == nil || o.SystemImage == nil {
		var ret bool
		return ret
	}
	return *o.SystemImage
}

// GetSystemImageOk returns a tuple with the SystemImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualImageLocation) GetSystemImageOk() (*bool, bool) {
	if o == nil || o.SystemImage == nil {
		return nil, false
	}
	return o.SystemImage, true
}

// HasSystemImage returns a boolean if a field has been set.
func (o *VirtualImageLocation) HasSystemImage() bool {
	if o != nil && o.SystemImage != nil {
		return true
	}

	return false
}

// SetSystemImage gets a reference to the given bool and assigns it to the SystemImage field.
func (o *VirtualImageLocation) SetSystemImage(v bool) {
	o.SystemImage = &v
}

// GetDiskIndex returns the DiskIndex field value if set, zero value otherwise.
func (o *VirtualImageLocation) GetDiskIndex() int64 {
	if o == nil || o.DiskIndex == nil {
		var ret int64
		return ret
	}
	return *o.DiskIndex
}

// GetDiskIndexOk returns a tuple with the DiskIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualImageLocation) GetDiskIndexOk() (*int64, bool) {
	if o == nil || o.DiskIndex == nil {
		return nil, false
	}
	return o.DiskIndex, true
}

// HasDiskIndex returns a boolean if a field has been set.
func (o *VirtualImageLocation) HasDiskIndex() bool {
	if o != nil && o.DiskIndex != nil {
		return true
	}

	return false
}

// SetDiskIndex gets a reference to the given int64 and assigns it to the DiskIndex field.
func (o *VirtualImageLocation) SetDiskIndex(v int64) {
	o.DiskIndex = &v
}

// GetPricePlan returns the PricePlan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VirtualImageLocation) GetPricePlan() string {
	if o == nil || o.PricePlan.Get() == nil {
		var ret string
		return ret
	}
	return *o.PricePlan.Get()
}

// GetPricePlanOk returns a tuple with the PricePlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VirtualImageLocation) GetPricePlanOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PricePlan.Get(), o.PricePlan.IsSet()
}

// HasPricePlan returns a boolean if a field has been set.
func (o *VirtualImageLocation) HasPricePlan() bool {
	if o != nil && o.PricePlan.IsSet() {
		return true
	}

	return false
}

// SetPricePlan gets a reference to the given NullableString and assigns it to the PricePlan field.
func (o *VirtualImageLocation) SetPricePlan(v string) {
	o.PricePlan.Set(&v)
}
// SetPricePlanNil sets the value for PricePlan to be an explicit nil
func (o *VirtualImageLocation) SetPricePlanNil() {
	o.PricePlan.Set(nil)
}

// UnsetPricePlan ensures that no value is present for PricePlan, not even an explicit nil
func (o *VirtualImageLocation) UnsetPricePlan() {
	o.PricePlan.Unset()
}

// GetVolumes returns the Volumes field value if set, zero value otherwise.
func (o *VirtualImageLocation) GetVolumes() []map[string]interface{} {
	if o == nil || o.Volumes == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Volumes
}

// GetVolumesOk returns a tuple with the Volumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualImageLocation) GetVolumesOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Volumes == nil {
		return nil, false
	}
	return o.Volumes, true
}

// HasVolumes returns a boolean if a field has been set.
func (o *VirtualImageLocation) HasVolumes() bool {
	if o != nil && o.Volumes != nil {
		return true
	}

	return false
}

// SetVolumes gets a reference to the given []map[string]interface{} and assigns it to the Volumes field.
func (o *VirtualImageLocation) SetVolumes(v []map[string]interface{}) {
	o.Volumes = &v
}

// GetStorageControllers returns the StorageControllers field value if set, zero value otherwise.
func (o *VirtualImageLocation) GetStorageControllers() []map[string]interface{} {
	if o == nil || o.StorageControllers == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.StorageControllers
}

// GetStorageControllersOk returns a tuple with the StorageControllers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualImageLocation) GetStorageControllersOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.StorageControllers == nil {
		return nil, false
	}
	return o.StorageControllers, true
}

// HasStorageControllers returns a boolean if a field has been set.
func (o *VirtualImageLocation) HasStorageControllers() bool {
	if o != nil && o.StorageControllers != nil {
		return true
	}

	return false
}

// SetStorageControllers gets a reference to the given []map[string]interface{} and assigns it to the StorageControllers field.
func (o *VirtualImageLocation) SetStorageControllers(v []map[string]interface{}) {
	o.StorageControllers = &v
}

// GetNetworkInterfaces returns the NetworkInterfaces field value if set, zero value otherwise.
func (o *VirtualImageLocation) GetNetworkInterfaces() []map[string]interface{} {
	if o == nil || o.NetworkInterfaces == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.NetworkInterfaces
}

// GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualImageLocation) GetNetworkInterfacesOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.NetworkInterfaces == nil {
		return nil, false
	}
	return o.NetworkInterfaces, true
}

// HasNetworkInterfaces returns a boolean if a field has been set.
func (o *VirtualImageLocation) HasNetworkInterfaces() bool {
	if o != nil && o.NetworkInterfaces != nil {
		return true
	}

	return false
}

// SetNetworkInterfaces gets a reference to the given []map[string]interface{} and assigns it to the NetworkInterfaces field.
func (o *VirtualImageLocation) SetNetworkInterfaces(v []map[string]interface{}) {
	o.NetworkInterfaces = &v
}

// GetVirtualImage returns the VirtualImage field value if set, zero value otherwise.
func (o *VirtualImageLocation) GetVirtualImage() VirtualImageLocationVirtualImage {
	if o == nil || o.VirtualImage == nil {
		var ret VirtualImageLocationVirtualImage
		return ret
	}
	return *o.VirtualImage
}

// GetVirtualImageOk returns a tuple with the VirtualImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualImageLocation) GetVirtualImageOk() (*VirtualImageLocationVirtualImage, bool) {
	if o == nil || o.VirtualImage == nil {
		return nil, false
	}
	return o.VirtualImage, true
}

// HasVirtualImage returns a boolean if a field has been set.
func (o *VirtualImageLocation) HasVirtualImage() bool {
	if o != nil && o.VirtualImage != nil {
		return true
	}

	return false
}

// SetVirtualImage gets a reference to the given VirtualImageLocationVirtualImage and assigns it to the VirtualImage field.
func (o *VirtualImageLocation) SetVirtualImage(v VirtualImageLocationVirtualImage) {
	o.VirtualImage = &v
}

func (o VirtualImageLocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Cloud != nil {
		toSerialize["cloud"] = o.Cloud
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.InternalId.IsSet() {
		toSerialize["internalId"] = o.InternalId.Get()
	}
	if o.ExternalId != nil {
		toSerialize["externalId"] = o.ExternalId
	}
	if o.ExternalDiskId != nil {
		toSerialize["externalDiskId"] = o.ExternalDiskId
	}
	if o.RemotePath.IsSet() {
		toSerialize["remotePath"] = o.RemotePath.Get()
	}
	if o.ImagePath.IsSet() {
		toSerialize["imagePath"] = o.ImagePath.Get()
	}
	if o.ImageName != nil {
		toSerialize["imageName"] = o.ImageName
	}
	if o.ImageRegion != nil {
		toSerialize["imageRegion"] = o.ImageRegion
	}
	if o.ImageFolder.IsSet() {
		toSerialize["imageFolder"] = o.ImageFolder.Get()
	}
	if o.RefType != nil {
		toSerialize["refType"] = o.RefType
	}
	if o.RefId != nil {
		toSerialize["refId"] = o.RefId
	}
	if o.NodeRefType.IsSet() {
		toSerialize["nodeRefType"] = o.NodeRefType.Get()
	}
	if o.NodeRefId.IsSet() {
		toSerialize["nodeRefId"] = o.NodeRefId.Get()
	}
	if o.SubRefType.IsSet() {
		toSerialize["subRefType"] = o.SubRefType.Get()
	}
	if o.SubRefId.IsSet() {
		toSerialize["subRefId"] = o.SubRefId.Get()
	}
	if o.IsPublic != nil {
		toSerialize["isPublic"] = o.IsPublic
	}
	if o.SystemImage != nil {
		toSerialize["systemImage"] = o.SystemImage
	}
	if o.DiskIndex != nil {
		toSerialize["diskIndex"] = o.DiskIndex
	}
	if o.PricePlan.IsSet() {
		toSerialize["pricePlan"] = o.PricePlan.Get()
	}
	if o.Volumes != nil {
		toSerialize["volumes"] = o.Volumes
	}
	if o.StorageControllers != nil {
		toSerialize["storageControllers"] = o.StorageControllers
	}
	if o.NetworkInterfaces != nil {
		toSerialize["networkInterfaces"] = o.NetworkInterfaces
	}
	if o.VirtualImage != nil {
		toSerialize["virtualImage"] = o.VirtualImage
	}
	return json.Marshal(toSerialize)
}

type NullableVirtualImageLocation struct {
	value *VirtualImageLocation
	isSet bool
}

func (v NullableVirtualImageLocation) Get() *VirtualImageLocation {
	return v.value
}

func (v *NullableVirtualImageLocation) Set(val *VirtualImageLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualImageLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualImageLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualImageLocation(val *VirtualImageLocation) *NullableVirtualImageLocation {
	return &NullableVirtualImageLocation{value: val, isSet: true}
}

func (v NullableVirtualImageLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualImageLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

