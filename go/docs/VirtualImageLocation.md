# VirtualImageLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Cloud** | Pointer to [**InlineResponse20079LoadBalancerMonitorLoadBalancerType**](inline_response_200_79_loadBalancerMonitor_loadBalancer_type.md) |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**InternalId** | Pointer to **NullableString** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**ExternalDiskId** | Pointer to **string** |  | [optional] 
**RemotePath** | Pointer to **NullableString** |  | [optional] 
**ImagePath** | Pointer to **NullableString** |  | [optional] 
**ImageName** | Pointer to **string** |  | [optional] 
**ImageRegion** | Pointer to **string** |  | [optional] 
**ImageFolder** | Pointer to **NullableString** |  | [optional] 
**RefType** | Pointer to **string** |  | [optional] 
**RefId** | Pointer to **int64** |  | [optional] 
**NodeRefType** | Pointer to **NullableString** |  | [optional] 
**NodeRefId** | Pointer to **NullableString** |  | [optional] 
**SubRefType** | Pointer to **NullableString** |  | [optional] 
**SubRefId** | Pointer to **NullableString** |  | [optional] 
**IsPublic** | Pointer to **bool** |  | [optional] 
**SystemImage** | Pointer to **bool** |  | [optional] 
**DiskIndex** | Pointer to **int64** |  | [optional] 
**PricePlan** | Pointer to **NullableString** |  | [optional] 
**Volumes** | Pointer to **[]map[string]interface{}** |  | [optional] 
**StorageControllers** | Pointer to **[]map[string]interface{}** |  | [optional] 
**NetworkInterfaces** | Pointer to **[]map[string]interface{}** |  | [optional] 
**VirtualImage** | Pointer to [**VirtualImageLocationVirtualImage**](virtualImageLocation_virtualImage.md) |  | [optional] 

## Methods

### NewVirtualImageLocation

`func NewVirtualImageLocation() *VirtualImageLocation`

NewVirtualImageLocation instantiates a new VirtualImageLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualImageLocationWithDefaults

`func NewVirtualImageLocationWithDefaults() *VirtualImageLocation`

NewVirtualImageLocationWithDefaults instantiates a new VirtualImageLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VirtualImageLocation) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VirtualImageLocation) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VirtualImageLocation) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *VirtualImageLocation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCloud

`func (o *VirtualImageLocation) GetCloud() InlineResponse20079LoadBalancerMonitorLoadBalancerType`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *VirtualImageLocation) GetCloudOk() (*InlineResponse20079LoadBalancerMonitorLoadBalancerType, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *VirtualImageLocation) SetCloud(v InlineResponse20079LoadBalancerMonitorLoadBalancerType)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *VirtualImageLocation) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetCode

`func (o *VirtualImageLocation) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *VirtualImageLocation) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *VirtualImageLocation) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *VirtualImageLocation) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetInternalId

`func (o *VirtualImageLocation) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *VirtualImageLocation) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *VirtualImageLocation) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *VirtualImageLocation) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *VirtualImageLocation) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *VirtualImageLocation) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetExternalId

`func (o *VirtualImageLocation) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *VirtualImageLocation) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *VirtualImageLocation) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *VirtualImageLocation) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetExternalDiskId

`func (o *VirtualImageLocation) GetExternalDiskId() string`

GetExternalDiskId returns the ExternalDiskId field if non-nil, zero value otherwise.

### GetExternalDiskIdOk

`func (o *VirtualImageLocation) GetExternalDiskIdOk() (*string, bool)`

GetExternalDiskIdOk returns a tuple with the ExternalDiskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDiskId

`func (o *VirtualImageLocation) SetExternalDiskId(v string)`

SetExternalDiskId sets ExternalDiskId field to given value.

### HasExternalDiskId

`func (o *VirtualImageLocation) HasExternalDiskId() bool`

HasExternalDiskId returns a boolean if a field has been set.

### GetRemotePath

`func (o *VirtualImageLocation) GetRemotePath() string`

GetRemotePath returns the RemotePath field if non-nil, zero value otherwise.

### GetRemotePathOk

`func (o *VirtualImageLocation) GetRemotePathOk() (*string, bool)`

GetRemotePathOk returns a tuple with the RemotePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePath

`func (o *VirtualImageLocation) SetRemotePath(v string)`

SetRemotePath sets RemotePath field to given value.

### HasRemotePath

`func (o *VirtualImageLocation) HasRemotePath() bool`

HasRemotePath returns a boolean if a field has been set.

### SetRemotePathNil

`func (o *VirtualImageLocation) SetRemotePathNil(b bool)`

 SetRemotePathNil sets the value for RemotePath to be an explicit nil

### UnsetRemotePath
`func (o *VirtualImageLocation) UnsetRemotePath()`

UnsetRemotePath ensures that no value is present for RemotePath, not even an explicit nil
### GetImagePath

`func (o *VirtualImageLocation) GetImagePath() string`

GetImagePath returns the ImagePath field if non-nil, zero value otherwise.

### GetImagePathOk

`func (o *VirtualImageLocation) GetImagePathOk() (*string, bool)`

GetImagePathOk returns a tuple with the ImagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagePath

`func (o *VirtualImageLocation) SetImagePath(v string)`

SetImagePath sets ImagePath field to given value.

### HasImagePath

`func (o *VirtualImageLocation) HasImagePath() bool`

HasImagePath returns a boolean if a field has been set.

### SetImagePathNil

`func (o *VirtualImageLocation) SetImagePathNil(b bool)`

 SetImagePathNil sets the value for ImagePath to be an explicit nil

### UnsetImagePath
`func (o *VirtualImageLocation) UnsetImagePath()`

UnsetImagePath ensures that no value is present for ImagePath, not even an explicit nil
### GetImageName

`func (o *VirtualImageLocation) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *VirtualImageLocation) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *VirtualImageLocation) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *VirtualImageLocation) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### GetImageRegion

`func (o *VirtualImageLocation) GetImageRegion() string`

GetImageRegion returns the ImageRegion field if non-nil, zero value otherwise.

### GetImageRegionOk

`func (o *VirtualImageLocation) GetImageRegionOk() (*string, bool)`

GetImageRegionOk returns a tuple with the ImageRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageRegion

`func (o *VirtualImageLocation) SetImageRegion(v string)`

SetImageRegion sets ImageRegion field to given value.

### HasImageRegion

`func (o *VirtualImageLocation) HasImageRegion() bool`

HasImageRegion returns a boolean if a field has been set.

### GetImageFolder

`func (o *VirtualImageLocation) GetImageFolder() string`

GetImageFolder returns the ImageFolder field if non-nil, zero value otherwise.

### GetImageFolderOk

`func (o *VirtualImageLocation) GetImageFolderOk() (*string, bool)`

GetImageFolderOk returns a tuple with the ImageFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageFolder

`func (o *VirtualImageLocation) SetImageFolder(v string)`

SetImageFolder sets ImageFolder field to given value.

### HasImageFolder

`func (o *VirtualImageLocation) HasImageFolder() bool`

HasImageFolder returns a boolean if a field has been set.

### SetImageFolderNil

`func (o *VirtualImageLocation) SetImageFolderNil(b bool)`

 SetImageFolderNil sets the value for ImageFolder to be an explicit nil

### UnsetImageFolder
`func (o *VirtualImageLocation) UnsetImageFolder()`

UnsetImageFolder ensures that no value is present for ImageFolder, not even an explicit nil
### GetRefType

`func (o *VirtualImageLocation) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *VirtualImageLocation) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *VirtualImageLocation) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *VirtualImageLocation) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *VirtualImageLocation) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *VirtualImageLocation) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *VirtualImageLocation) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *VirtualImageLocation) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetNodeRefType

`func (o *VirtualImageLocation) GetNodeRefType() string`

GetNodeRefType returns the NodeRefType field if non-nil, zero value otherwise.

### GetNodeRefTypeOk

`func (o *VirtualImageLocation) GetNodeRefTypeOk() (*string, bool)`

GetNodeRefTypeOk returns a tuple with the NodeRefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeRefType

`func (o *VirtualImageLocation) SetNodeRefType(v string)`

SetNodeRefType sets NodeRefType field to given value.

### HasNodeRefType

`func (o *VirtualImageLocation) HasNodeRefType() bool`

HasNodeRefType returns a boolean if a field has been set.

### SetNodeRefTypeNil

`func (o *VirtualImageLocation) SetNodeRefTypeNil(b bool)`

 SetNodeRefTypeNil sets the value for NodeRefType to be an explicit nil

### UnsetNodeRefType
`func (o *VirtualImageLocation) UnsetNodeRefType()`

UnsetNodeRefType ensures that no value is present for NodeRefType, not even an explicit nil
### GetNodeRefId

`func (o *VirtualImageLocation) GetNodeRefId() string`

GetNodeRefId returns the NodeRefId field if non-nil, zero value otherwise.

### GetNodeRefIdOk

`func (o *VirtualImageLocation) GetNodeRefIdOk() (*string, bool)`

GetNodeRefIdOk returns a tuple with the NodeRefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeRefId

`func (o *VirtualImageLocation) SetNodeRefId(v string)`

SetNodeRefId sets NodeRefId field to given value.

### HasNodeRefId

`func (o *VirtualImageLocation) HasNodeRefId() bool`

HasNodeRefId returns a boolean if a field has been set.

### SetNodeRefIdNil

`func (o *VirtualImageLocation) SetNodeRefIdNil(b bool)`

 SetNodeRefIdNil sets the value for NodeRefId to be an explicit nil

### UnsetNodeRefId
`func (o *VirtualImageLocation) UnsetNodeRefId()`

UnsetNodeRefId ensures that no value is present for NodeRefId, not even an explicit nil
### GetSubRefType

`func (o *VirtualImageLocation) GetSubRefType() string`

GetSubRefType returns the SubRefType field if non-nil, zero value otherwise.

### GetSubRefTypeOk

`func (o *VirtualImageLocation) GetSubRefTypeOk() (*string, bool)`

GetSubRefTypeOk returns a tuple with the SubRefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubRefType

`func (o *VirtualImageLocation) SetSubRefType(v string)`

SetSubRefType sets SubRefType field to given value.

### HasSubRefType

`func (o *VirtualImageLocation) HasSubRefType() bool`

HasSubRefType returns a boolean if a field has been set.

### SetSubRefTypeNil

`func (o *VirtualImageLocation) SetSubRefTypeNil(b bool)`

 SetSubRefTypeNil sets the value for SubRefType to be an explicit nil

### UnsetSubRefType
`func (o *VirtualImageLocation) UnsetSubRefType()`

UnsetSubRefType ensures that no value is present for SubRefType, not even an explicit nil
### GetSubRefId

`func (o *VirtualImageLocation) GetSubRefId() string`

GetSubRefId returns the SubRefId field if non-nil, zero value otherwise.

### GetSubRefIdOk

`func (o *VirtualImageLocation) GetSubRefIdOk() (*string, bool)`

GetSubRefIdOk returns a tuple with the SubRefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubRefId

`func (o *VirtualImageLocation) SetSubRefId(v string)`

SetSubRefId sets SubRefId field to given value.

### HasSubRefId

`func (o *VirtualImageLocation) HasSubRefId() bool`

HasSubRefId returns a boolean if a field has been set.

### SetSubRefIdNil

`func (o *VirtualImageLocation) SetSubRefIdNil(b bool)`

 SetSubRefIdNil sets the value for SubRefId to be an explicit nil

### UnsetSubRefId
`func (o *VirtualImageLocation) UnsetSubRefId()`

UnsetSubRefId ensures that no value is present for SubRefId, not even an explicit nil
### GetIsPublic

`func (o *VirtualImageLocation) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *VirtualImageLocation) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *VirtualImageLocation) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *VirtualImageLocation) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetSystemImage

`func (o *VirtualImageLocation) GetSystemImage() bool`

GetSystemImage returns the SystemImage field if non-nil, zero value otherwise.

### GetSystemImageOk

`func (o *VirtualImageLocation) GetSystemImageOk() (*bool, bool)`

GetSystemImageOk returns a tuple with the SystemImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemImage

`func (o *VirtualImageLocation) SetSystemImage(v bool)`

SetSystemImage sets SystemImage field to given value.

### HasSystemImage

`func (o *VirtualImageLocation) HasSystemImage() bool`

HasSystemImage returns a boolean if a field has been set.

### GetDiskIndex

`func (o *VirtualImageLocation) GetDiskIndex() int64`

GetDiskIndex returns the DiskIndex field if non-nil, zero value otherwise.

### GetDiskIndexOk

`func (o *VirtualImageLocation) GetDiskIndexOk() (*int64, bool)`

GetDiskIndexOk returns a tuple with the DiskIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskIndex

`func (o *VirtualImageLocation) SetDiskIndex(v int64)`

SetDiskIndex sets DiskIndex field to given value.

### HasDiskIndex

`func (o *VirtualImageLocation) HasDiskIndex() bool`

HasDiskIndex returns a boolean if a field has been set.

### GetPricePlan

`func (o *VirtualImageLocation) GetPricePlan() string`

GetPricePlan returns the PricePlan field if non-nil, zero value otherwise.

### GetPricePlanOk

`func (o *VirtualImageLocation) GetPricePlanOk() (*string, bool)`

GetPricePlanOk returns a tuple with the PricePlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePlan

`func (o *VirtualImageLocation) SetPricePlan(v string)`

SetPricePlan sets PricePlan field to given value.

### HasPricePlan

`func (o *VirtualImageLocation) HasPricePlan() bool`

HasPricePlan returns a boolean if a field has been set.

### SetPricePlanNil

`func (o *VirtualImageLocation) SetPricePlanNil(b bool)`

 SetPricePlanNil sets the value for PricePlan to be an explicit nil

### UnsetPricePlan
`func (o *VirtualImageLocation) UnsetPricePlan()`

UnsetPricePlan ensures that no value is present for PricePlan, not even an explicit nil
### GetVolumes

`func (o *VirtualImageLocation) GetVolumes() []map[string]interface{}`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *VirtualImageLocation) GetVolumesOk() (*[]map[string]interface{}, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *VirtualImageLocation) SetVolumes(v []map[string]interface{})`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *VirtualImageLocation) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetStorageControllers

`func (o *VirtualImageLocation) GetStorageControllers() []map[string]interface{}`

GetStorageControllers returns the StorageControllers field if non-nil, zero value otherwise.

### GetStorageControllersOk

`func (o *VirtualImageLocation) GetStorageControllersOk() (*[]map[string]interface{}, bool)`

GetStorageControllersOk returns a tuple with the StorageControllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageControllers

`func (o *VirtualImageLocation) SetStorageControllers(v []map[string]interface{})`

SetStorageControllers sets StorageControllers field to given value.

### HasStorageControllers

`func (o *VirtualImageLocation) HasStorageControllers() bool`

HasStorageControllers returns a boolean if a field has been set.

### GetNetworkInterfaces

`func (o *VirtualImageLocation) GetNetworkInterfaces() []map[string]interface{}`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *VirtualImageLocation) GetNetworkInterfacesOk() (*[]map[string]interface{}, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *VirtualImageLocation) SetNetworkInterfaces(v []map[string]interface{})`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *VirtualImageLocation) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.

### GetVirtualImage

`func (o *VirtualImageLocation) GetVirtualImage() VirtualImageLocationVirtualImage`

GetVirtualImage returns the VirtualImage field if non-nil, zero value otherwise.

### GetVirtualImageOk

`func (o *VirtualImageLocation) GetVirtualImageOk() (*VirtualImageLocationVirtualImage, bool)`

GetVirtualImageOk returns a tuple with the VirtualImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualImage

`func (o *VirtualImageLocation) SetVirtualImage(v VirtualImageLocationVirtualImage)`

SetVirtualImage sets VirtualImage field to given value.

### HasVirtualImage

`func (o *VirtualImageLocation) HasVirtualImage() bool`

HasVirtualImage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


