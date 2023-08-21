# VirtualImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**OwnerId** | Pointer to **int64** |  | [optional] 
**Tenant** | Pointer to [**InlineResponse20040AppDeployInstance**](inline_response_200_40_appDeploy_instance.md) |  | [optional] 
**ImageType** | Pointer to **string** |  | [optional] 
**UserUploaded** | Pointer to **bool** |  | [optional] 
**UserDefined** | Pointer to **bool** |  | [optional] 
**SystemImage** | Pointer to **bool** |  | [optional] 
**IsCloudInit** | Pointer to **bool** |  | [optional] 
**SshUsername** | Pointer to **NullableString** |  | [optional] 
**SshPassword** | Pointer to **NullableString** |  | [optional] 
**SshPasswordHash** | Pointer to **NullableString** |  | [optional] 
**SshKey** | Pointer to **NullableString** |  | [optional] 
**OsType** | Pointer to [**NullableVirtualImageOsType**](virtualImage_osType.md) |  | [optional] 
**MinRam** | Pointer to **NullableInt64** |  | [optional] 
**MinRamGB** | Pointer to **NullableInt64** |  | [optional] 
**MinDisk** | Pointer to **NullableString** |  | [optional] 
**MinDiskGB** | Pointer to **NullableString** |  | [optional] 
**RawSize** | Pointer to **NullableInt64** |  | [optional] 
**RawSizeGB** | Pointer to **NullableFloat32** |  | [optional] 
**TrialVersion** | Pointer to **bool** |  | [optional] 
**VirtioSupported** | Pointer to **bool** |  | [optional] 
**Uefi** | Pointer to **NullableString** |  | [optional] 
**IsAutoJoinDomain** | Pointer to **bool** |  | [optional] 
**VmToolsInstalled** | Pointer to **bool** |  | [optional] 
**InstallAgent** | Pointer to **bool** |  | [optional] 
**IsForceCustomization** | Pointer to **bool** |  | [optional] 
**IsSysprep** | Pointer to **bool** |  | [optional] 
**FipsEnabled** | Pointer to **bool** |  | [optional] 
**UserData** | Pointer to **NullableString** |  | [optional] 
**ConsoleKeymap** | Pointer to **NullableString** |  | [optional] 
**StorageProvider** | Pointer to **NullableString** |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Accounts** | Pointer to [**[]InlineResponse20040AppDeployInstance**](InlineResponse20040AppDeployInstance.md) |  | [optional] 
**Config** | Pointer to [**VirtualImageConfig**](virtualImage_config.md) |  | [optional] 
**Volumes** | Pointer to **[]map[string]interface{}** |  | [optional] 
**StorageControllers** | Pointer to **[]map[string]interface{}** |  | [optional] 
**NetworkInterfaces** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Tags** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Locations** | Pointer to **[]map[string]interface{}** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewVirtualImage

`func NewVirtualImage() *VirtualImage`

NewVirtualImage instantiates a new VirtualImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualImageWithDefaults

`func NewVirtualImageWithDefaults() *VirtualImage`

NewVirtualImageWithDefaults instantiates a new VirtualImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VirtualImage) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VirtualImage) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VirtualImage) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *VirtualImage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *VirtualImage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualImage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualImage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualImage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *VirtualImage) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VirtualImage) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VirtualImage) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VirtualImage) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *VirtualImage) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *VirtualImage) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLabels

`func (o *VirtualImage) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *VirtualImage) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *VirtualImage) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *VirtualImage) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetOwnerId

`func (o *VirtualImage) GetOwnerId() int64`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *VirtualImage) GetOwnerIdOk() (*int64, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *VirtualImage) SetOwnerId(v int64)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *VirtualImage) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetTenant

`func (o *VirtualImage) GetTenant() InlineResponse20040AppDeployInstance`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *VirtualImage) GetTenantOk() (*InlineResponse20040AppDeployInstance, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *VirtualImage) SetTenant(v InlineResponse20040AppDeployInstance)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *VirtualImage) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetImageType

`func (o *VirtualImage) GetImageType() string`

GetImageType returns the ImageType field if non-nil, zero value otherwise.

### GetImageTypeOk

`func (o *VirtualImage) GetImageTypeOk() (*string, bool)`

GetImageTypeOk returns a tuple with the ImageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageType

`func (o *VirtualImage) SetImageType(v string)`

SetImageType sets ImageType field to given value.

### HasImageType

`func (o *VirtualImage) HasImageType() bool`

HasImageType returns a boolean if a field has been set.

### GetUserUploaded

`func (o *VirtualImage) GetUserUploaded() bool`

GetUserUploaded returns the UserUploaded field if non-nil, zero value otherwise.

### GetUserUploadedOk

`func (o *VirtualImage) GetUserUploadedOk() (*bool, bool)`

GetUserUploadedOk returns a tuple with the UserUploaded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUploaded

`func (o *VirtualImage) SetUserUploaded(v bool)`

SetUserUploaded sets UserUploaded field to given value.

### HasUserUploaded

`func (o *VirtualImage) HasUserUploaded() bool`

HasUserUploaded returns a boolean if a field has been set.

### GetUserDefined

`func (o *VirtualImage) GetUserDefined() bool`

GetUserDefined returns the UserDefined field if non-nil, zero value otherwise.

### GetUserDefinedOk

`func (o *VirtualImage) GetUserDefinedOk() (*bool, bool)`

GetUserDefinedOk returns a tuple with the UserDefined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefined

`func (o *VirtualImage) SetUserDefined(v bool)`

SetUserDefined sets UserDefined field to given value.

### HasUserDefined

`func (o *VirtualImage) HasUserDefined() bool`

HasUserDefined returns a boolean if a field has been set.

### GetSystemImage

`func (o *VirtualImage) GetSystemImage() bool`

GetSystemImage returns the SystemImage field if non-nil, zero value otherwise.

### GetSystemImageOk

`func (o *VirtualImage) GetSystemImageOk() (*bool, bool)`

GetSystemImageOk returns a tuple with the SystemImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemImage

`func (o *VirtualImage) SetSystemImage(v bool)`

SetSystemImage sets SystemImage field to given value.

### HasSystemImage

`func (o *VirtualImage) HasSystemImage() bool`

HasSystemImage returns a boolean if a field has been set.

### GetIsCloudInit

`func (o *VirtualImage) GetIsCloudInit() bool`

GetIsCloudInit returns the IsCloudInit field if non-nil, zero value otherwise.

### GetIsCloudInitOk

`func (o *VirtualImage) GetIsCloudInitOk() (*bool, bool)`

GetIsCloudInitOk returns a tuple with the IsCloudInit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCloudInit

`func (o *VirtualImage) SetIsCloudInit(v bool)`

SetIsCloudInit sets IsCloudInit field to given value.

### HasIsCloudInit

`func (o *VirtualImage) HasIsCloudInit() bool`

HasIsCloudInit returns a boolean if a field has been set.

### GetSshUsername

`func (o *VirtualImage) GetSshUsername() string`

GetSshUsername returns the SshUsername field if non-nil, zero value otherwise.

### GetSshUsernameOk

`func (o *VirtualImage) GetSshUsernameOk() (*string, bool)`

GetSshUsernameOk returns a tuple with the SshUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUsername

`func (o *VirtualImage) SetSshUsername(v string)`

SetSshUsername sets SshUsername field to given value.

### HasSshUsername

`func (o *VirtualImage) HasSshUsername() bool`

HasSshUsername returns a boolean if a field has been set.

### SetSshUsernameNil

`func (o *VirtualImage) SetSshUsernameNil(b bool)`

 SetSshUsernameNil sets the value for SshUsername to be an explicit nil

### UnsetSshUsername
`func (o *VirtualImage) UnsetSshUsername()`

UnsetSshUsername ensures that no value is present for SshUsername, not even an explicit nil
### GetSshPassword

`func (o *VirtualImage) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *VirtualImage) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *VirtualImage) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *VirtualImage) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.

### SetSshPasswordNil

`func (o *VirtualImage) SetSshPasswordNil(b bool)`

 SetSshPasswordNil sets the value for SshPassword to be an explicit nil

### UnsetSshPassword
`func (o *VirtualImage) UnsetSshPassword()`

UnsetSshPassword ensures that no value is present for SshPassword, not even an explicit nil
### GetSshPasswordHash

`func (o *VirtualImage) GetSshPasswordHash() string`

GetSshPasswordHash returns the SshPasswordHash field if non-nil, zero value otherwise.

### GetSshPasswordHashOk

`func (o *VirtualImage) GetSshPasswordHashOk() (*string, bool)`

GetSshPasswordHashOk returns a tuple with the SshPasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPasswordHash

`func (o *VirtualImage) SetSshPasswordHash(v string)`

SetSshPasswordHash sets SshPasswordHash field to given value.

### HasSshPasswordHash

`func (o *VirtualImage) HasSshPasswordHash() bool`

HasSshPasswordHash returns a boolean if a field has been set.

### SetSshPasswordHashNil

`func (o *VirtualImage) SetSshPasswordHashNil(b bool)`

 SetSshPasswordHashNil sets the value for SshPasswordHash to be an explicit nil

### UnsetSshPasswordHash
`func (o *VirtualImage) UnsetSshPasswordHash()`

UnsetSshPasswordHash ensures that no value is present for SshPasswordHash, not even an explicit nil
### GetSshKey

`func (o *VirtualImage) GetSshKey() string`

GetSshKey returns the SshKey field if non-nil, zero value otherwise.

### GetSshKeyOk

`func (o *VirtualImage) GetSshKeyOk() (*string, bool)`

GetSshKeyOk returns a tuple with the SshKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKey

`func (o *VirtualImage) SetSshKey(v string)`

SetSshKey sets SshKey field to given value.

### HasSshKey

`func (o *VirtualImage) HasSshKey() bool`

HasSshKey returns a boolean if a field has been set.

### SetSshKeyNil

`func (o *VirtualImage) SetSshKeyNil(b bool)`

 SetSshKeyNil sets the value for SshKey to be an explicit nil

### UnsetSshKey
`func (o *VirtualImage) UnsetSshKey()`

UnsetSshKey ensures that no value is present for SshKey, not even an explicit nil
### GetOsType

`func (o *VirtualImage) GetOsType() VirtualImageOsType`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *VirtualImage) GetOsTypeOk() (*VirtualImageOsType, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *VirtualImage) SetOsType(v VirtualImageOsType)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *VirtualImage) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### SetOsTypeNil

`func (o *VirtualImage) SetOsTypeNil(b bool)`

 SetOsTypeNil sets the value for OsType to be an explicit nil

### UnsetOsType
`func (o *VirtualImage) UnsetOsType()`

UnsetOsType ensures that no value is present for OsType, not even an explicit nil
### GetMinRam

`func (o *VirtualImage) GetMinRam() int64`

GetMinRam returns the MinRam field if non-nil, zero value otherwise.

### GetMinRamOk

`func (o *VirtualImage) GetMinRamOk() (*int64, bool)`

GetMinRamOk returns a tuple with the MinRam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRam

`func (o *VirtualImage) SetMinRam(v int64)`

SetMinRam sets MinRam field to given value.

### HasMinRam

`func (o *VirtualImage) HasMinRam() bool`

HasMinRam returns a boolean if a field has been set.

### SetMinRamNil

`func (o *VirtualImage) SetMinRamNil(b bool)`

 SetMinRamNil sets the value for MinRam to be an explicit nil

### UnsetMinRam
`func (o *VirtualImage) UnsetMinRam()`

UnsetMinRam ensures that no value is present for MinRam, not even an explicit nil
### GetMinRamGB

`func (o *VirtualImage) GetMinRamGB() int64`

GetMinRamGB returns the MinRamGB field if non-nil, zero value otherwise.

### GetMinRamGBOk

`func (o *VirtualImage) GetMinRamGBOk() (*int64, bool)`

GetMinRamGBOk returns a tuple with the MinRamGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRamGB

`func (o *VirtualImage) SetMinRamGB(v int64)`

SetMinRamGB sets MinRamGB field to given value.

### HasMinRamGB

`func (o *VirtualImage) HasMinRamGB() bool`

HasMinRamGB returns a boolean if a field has been set.

### SetMinRamGBNil

`func (o *VirtualImage) SetMinRamGBNil(b bool)`

 SetMinRamGBNil sets the value for MinRamGB to be an explicit nil

### UnsetMinRamGB
`func (o *VirtualImage) UnsetMinRamGB()`

UnsetMinRamGB ensures that no value is present for MinRamGB, not even an explicit nil
### GetMinDisk

`func (o *VirtualImage) GetMinDisk() string`

GetMinDisk returns the MinDisk field if non-nil, zero value otherwise.

### GetMinDiskOk

`func (o *VirtualImage) GetMinDiskOk() (*string, bool)`

GetMinDiskOk returns a tuple with the MinDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDisk

`func (o *VirtualImage) SetMinDisk(v string)`

SetMinDisk sets MinDisk field to given value.

### HasMinDisk

`func (o *VirtualImage) HasMinDisk() bool`

HasMinDisk returns a boolean if a field has been set.

### SetMinDiskNil

`func (o *VirtualImage) SetMinDiskNil(b bool)`

 SetMinDiskNil sets the value for MinDisk to be an explicit nil

### UnsetMinDisk
`func (o *VirtualImage) UnsetMinDisk()`

UnsetMinDisk ensures that no value is present for MinDisk, not even an explicit nil
### GetMinDiskGB

`func (o *VirtualImage) GetMinDiskGB() string`

GetMinDiskGB returns the MinDiskGB field if non-nil, zero value otherwise.

### GetMinDiskGBOk

`func (o *VirtualImage) GetMinDiskGBOk() (*string, bool)`

GetMinDiskGBOk returns a tuple with the MinDiskGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDiskGB

`func (o *VirtualImage) SetMinDiskGB(v string)`

SetMinDiskGB sets MinDiskGB field to given value.

### HasMinDiskGB

`func (o *VirtualImage) HasMinDiskGB() bool`

HasMinDiskGB returns a boolean if a field has been set.

### SetMinDiskGBNil

`func (o *VirtualImage) SetMinDiskGBNil(b bool)`

 SetMinDiskGBNil sets the value for MinDiskGB to be an explicit nil

### UnsetMinDiskGB
`func (o *VirtualImage) UnsetMinDiskGB()`

UnsetMinDiskGB ensures that no value is present for MinDiskGB, not even an explicit nil
### GetRawSize

`func (o *VirtualImage) GetRawSize() int64`

GetRawSize returns the RawSize field if non-nil, zero value otherwise.

### GetRawSizeOk

`func (o *VirtualImage) GetRawSizeOk() (*int64, bool)`

GetRawSizeOk returns a tuple with the RawSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawSize

`func (o *VirtualImage) SetRawSize(v int64)`

SetRawSize sets RawSize field to given value.

### HasRawSize

`func (o *VirtualImage) HasRawSize() bool`

HasRawSize returns a boolean if a field has been set.

### SetRawSizeNil

`func (o *VirtualImage) SetRawSizeNil(b bool)`

 SetRawSizeNil sets the value for RawSize to be an explicit nil

### UnsetRawSize
`func (o *VirtualImage) UnsetRawSize()`

UnsetRawSize ensures that no value is present for RawSize, not even an explicit nil
### GetRawSizeGB

`func (o *VirtualImage) GetRawSizeGB() float32`

GetRawSizeGB returns the RawSizeGB field if non-nil, zero value otherwise.

### GetRawSizeGBOk

`func (o *VirtualImage) GetRawSizeGBOk() (*float32, bool)`

GetRawSizeGBOk returns a tuple with the RawSizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawSizeGB

`func (o *VirtualImage) SetRawSizeGB(v float32)`

SetRawSizeGB sets RawSizeGB field to given value.

### HasRawSizeGB

`func (o *VirtualImage) HasRawSizeGB() bool`

HasRawSizeGB returns a boolean if a field has been set.

### SetRawSizeGBNil

`func (o *VirtualImage) SetRawSizeGBNil(b bool)`

 SetRawSizeGBNil sets the value for RawSizeGB to be an explicit nil

### UnsetRawSizeGB
`func (o *VirtualImage) UnsetRawSizeGB()`

UnsetRawSizeGB ensures that no value is present for RawSizeGB, not even an explicit nil
### GetTrialVersion

`func (o *VirtualImage) GetTrialVersion() bool`

GetTrialVersion returns the TrialVersion field if non-nil, zero value otherwise.

### GetTrialVersionOk

`func (o *VirtualImage) GetTrialVersionOk() (*bool, bool)`

GetTrialVersionOk returns a tuple with the TrialVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialVersion

`func (o *VirtualImage) SetTrialVersion(v bool)`

SetTrialVersion sets TrialVersion field to given value.

### HasTrialVersion

`func (o *VirtualImage) HasTrialVersion() bool`

HasTrialVersion returns a boolean if a field has been set.

### GetVirtioSupported

`func (o *VirtualImage) GetVirtioSupported() bool`

GetVirtioSupported returns the VirtioSupported field if non-nil, zero value otherwise.

### GetVirtioSupportedOk

`func (o *VirtualImage) GetVirtioSupportedOk() (*bool, bool)`

GetVirtioSupportedOk returns a tuple with the VirtioSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtioSupported

`func (o *VirtualImage) SetVirtioSupported(v bool)`

SetVirtioSupported sets VirtioSupported field to given value.

### HasVirtioSupported

`func (o *VirtualImage) HasVirtioSupported() bool`

HasVirtioSupported returns a boolean if a field has been set.

### GetUefi

`func (o *VirtualImage) GetUefi() string`

GetUefi returns the Uefi field if non-nil, zero value otherwise.

### GetUefiOk

`func (o *VirtualImage) GetUefiOk() (*string, bool)`

GetUefiOk returns a tuple with the Uefi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUefi

`func (o *VirtualImage) SetUefi(v string)`

SetUefi sets Uefi field to given value.

### HasUefi

`func (o *VirtualImage) HasUefi() bool`

HasUefi returns a boolean if a field has been set.

### SetUefiNil

`func (o *VirtualImage) SetUefiNil(b bool)`

 SetUefiNil sets the value for Uefi to be an explicit nil

### UnsetUefi
`func (o *VirtualImage) UnsetUefi()`

UnsetUefi ensures that no value is present for Uefi, not even an explicit nil
### GetIsAutoJoinDomain

`func (o *VirtualImage) GetIsAutoJoinDomain() bool`

GetIsAutoJoinDomain returns the IsAutoJoinDomain field if non-nil, zero value otherwise.

### GetIsAutoJoinDomainOk

`func (o *VirtualImage) GetIsAutoJoinDomainOk() (*bool, bool)`

GetIsAutoJoinDomainOk returns a tuple with the IsAutoJoinDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoJoinDomain

`func (o *VirtualImage) SetIsAutoJoinDomain(v bool)`

SetIsAutoJoinDomain sets IsAutoJoinDomain field to given value.

### HasIsAutoJoinDomain

`func (o *VirtualImage) HasIsAutoJoinDomain() bool`

HasIsAutoJoinDomain returns a boolean if a field has been set.

### GetVmToolsInstalled

`func (o *VirtualImage) GetVmToolsInstalled() bool`

GetVmToolsInstalled returns the VmToolsInstalled field if non-nil, zero value otherwise.

### GetVmToolsInstalledOk

`func (o *VirtualImage) GetVmToolsInstalledOk() (*bool, bool)`

GetVmToolsInstalledOk returns a tuple with the VmToolsInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmToolsInstalled

`func (o *VirtualImage) SetVmToolsInstalled(v bool)`

SetVmToolsInstalled sets VmToolsInstalled field to given value.

### HasVmToolsInstalled

`func (o *VirtualImage) HasVmToolsInstalled() bool`

HasVmToolsInstalled returns a boolean if a field has been set.

### GetInstallAgent

`func (o *VirtualImage) GetInstallAgent() bool`

GetInstallAgent returns the InstallAgent field if non-nil, zero value otherwise.

### GetInstallAgentOk

`func (o *VirtualImage) GetInstallAgentOk() (*bool, bool)`

GetInstallAgentOk returns a tuple with the InstallAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallAgent

`func (o *VirtualImage) SetInstallAgent(v bool)`

SetInstallAgent sets InstallAgent field to given value.

### HasInstallAgent

`func (o *VirtualImage) HasInstallAgent() bool`

HasInstallAgent returns a boolean if a field has been set.

### GetIsForceCustomization

`func (o *VirtualImage) GetIsForceCustomization() bool`

GetIsForceCustomization returns the IsForceCustomization field if non-nil, zero value otherwise.

### GetIsForceCustomizationOk

`func (o *VirtualImage) GetIsForceCustomizationOk() (*bool, bool)`

GetIsForceCustomizationOk returns a tuple with the IsForceCustomization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsForceCustomization

`func (o *VirtualImage) SetIsForceCustomization(v bool)`

SetIsForceCustomization sets IsForceCustomization field to given value.

### HasIsForceCustomization

`func (o *VirtualImage) HasIsForceCustomization() bool`

HasIsForceCustomization returns a boolean if a field has been set.

### GetIsSysprep

`func (o *VirtualImage) GetIsSysprep() bool`

GetIsSysprep returns the IsSysprep field if non-nil, zero value otherwise.

### GetIsSysprepOk

`func (o *VirtualImage) GetIsSysprepOk() (*bool, bool)`

GetIsSysprepOk returns a tuple with the IsSysprep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSysprep

`func (o *VirtualImage) SetIsSysprep(v bool)`

SetIsSysprep sets IsSysprep field to given value.

### HasIsSysprep

`func (o *VirtualImage) HasIsSysprep() bool`

HasIsSysprep returns a boolean if a field has been set.

### GetFipsEnabled

`func (o *VirtualImage) GetFipsEnabled() bool`

GetFipsEnabled returns the FipsEnabled field if non-nil, zero value otherwise.

### GetFipsEnabledOk

`func (o *VirtualImage) GetFipsEnabledOk() (*bool, bool)`

GetFipsEnabledOk returns a tuple with the FipsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFipsEnabled

`func (o *VirtualImage) SetFipsEnabled(v bool)`

SetFipsEnabled sets FipsEnabled field to given value.

### HasFipsEnabled

`func (o *VirtualImage) HasFipsEnabled() bool`

HasFipsEnabled returns a boolean if a field has been set.

### GetUserData

`func (o *VirtualImage) GetUserData() string`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *VirtualImage) GetUserDataOk() (*string, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *VirtualImage) SetUserData(v string)`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *VirtualImage) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### SetUserDataNil

`func (o *VirtualImage) SetUserDataNil(b bool)`

 SetUserDataNil sets the value for UserData to be an explicit nil

### UnsetUserData
`func (o *VirtualImage) UnsetUserData()`

UnsetUserData ensures that no value is present for UserData, not even an explicit nil
### GetConsoleKeymap

`func (o *VirtualImage) GetConsoleKeymap() string`

GetConsoleKeymap returns the ConsoleKeymap field if non-nil, zero value otherwise.

### GetConsoleKeymapOk

`func (o *VirtualImage) GetConsoleKeymapOk() (*string, bool)`

GetConsoleKeymapOk returns a tuple with the ConsoleKeymap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleKeymap

`func (o *VirtualImage) SetConsoleKeymap(v string)`

SetConsoleKeymap sets ConsoleKeymap field to given value.

### HasConsoleKeymap

`func (o *VirtualImage) HasConsoleKeymap() bool`

HasConsoleKeymap returns a boolean if a field has been set.

### SetConsoleKeymapNil

`func (o *VirtualImage) SetConsoleKeymapNil(b bool)`

 SetConsoleKeymapNil sets the value for ConsoleKeymap to be an explicit nil

### UnsetConsoleKeymap
`func (o *VirtualImage) UnsetConsoleKeymap()`

UnsetConsoleKeymap ensures that no value is present for ConsoleKeymap, not even an explicit nil
### GetStorageProvider

`func (o *VirtualImage) GetStorageProvider() string`

GetStorageProvider returns the StorageProvider field if non-nil, zero value otherwise.

### GetStorageProviderOk

`func (o *VirtualImage) GetStorageProviderOk() (*string, bool)`

GetStorageProviderOk returns a tuple with the StorageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProvider

`func (o *VirtualImage) SetStorageProvider(v string)`

SetStorageProvider sets StorageProvider field to given value.

### HasStorageProvider

`func (o *VirtualImage) HasStorageProvider() bool`

HasStorageProvider returns a boolean if a field has been set.

### SetStorageProviderNil

`func (o *VirtualImage) SetStorageProviderNil(b bool)`

 SetStorageProviderNil sets the value for StorageProvider to be an explicit nil

### UnsetStorageProvider
`func (o *VirtualImage) UnsetStorageProvider()`

UnsetStorageProvider ensures that no value is present for StorageProvider, not even an explicit nil
### GetExternalId

`func (o *VirtualImage) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *VirtualImage) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *VirtualImage) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *VirtualImage) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *VirtualImage) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *VirtualImage) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetVisibility

`func (o *VirtualImage) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *VirtualImage) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *VirtualImage) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *VirtualImage) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetAccounts

`func (o *VirtualImage) GetAccounts() []InlineResponse20040AppDeployInstance`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *VirtualImage) GetAccountsOk() (*[]InlineResponse20040AppDeployInstance, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *VirtualImage) SetAccounts(v []InlineResponse20040AppDeployInstance)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *VirtualImage) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetConfig

`func (o *VirtualImage) GetConfig() VirtualImageConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *VirtualImage) GetConfigOk() (*VirtualImageConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *VirtualImage) SetConfig(v VirtualImageConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *VirtualImage) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetVolumes

`func (o *VirtualImage) GetVolumes() []map[string]interface{}`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *VirtualImage) GetVolumesOk() (*[]map[string]interface{}, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *VirtualImage) SetVolumes(v []map[string]interface{})`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *VirtualImage) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetStorageControllers

`func (o *VirtualImage) GetStorageControllers() []map[string]interface{}`

GetStorageControllers returns the StorageControllers field if non-nil, zero value otherwise.

### GetStorageControllersOk

`func (o *VirtualImage) GetStorageControllersOk() (*[]map[string]interface{}, bool)`

GetStorageControllersOk returns a tuple with the StorageControllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageControllers

`func (o *VirtualImage) SetStorageControllers(v []map[string]interface{})`

SetStorageControllers sets StorageControllers field to given value.

### HasStorageControllers

`func (o *VirtualImage) HasStorageControllers() bool`

HasStorageControllers returns a boolean if a field has been set.

### GetNetworkInterfaces

`func (o *VirtualImage) GetNetworkInterfaces() []map[string]interface{}`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *VirtualImage) GetNetworkInterfacesOk() (*[]map[string]interface{}, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *VirtualImage) SetNetworkInterfaces(v []map[string]interface{})`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *VirtualImage) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.

### GetTags

`func (o *VirtualImage) GetTags() []map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VirtualImage) GetTagsOk() (*[]map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VirtualImage) SetTags(v []map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *VirtualImage) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetLocations

`func (o *VirtualImage) GetLocations() []map[string]interface{}`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *VirtualImage) GetLocationsOk() (*[]map[string]interface{}, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *VirtualImage) SetLocations(v []map[string]interface{})`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *VirtualImage) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetDateCreated

`func (o *VirtualImage) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *VirtualImage) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *VirtualImage) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *VirtualImage) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *VirtualImage) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *VirtualImage) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *VirtualImage) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *VirtualImage) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetStatus

`func (o *VirtualImage) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VirtualImage) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VirtualImage) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VirtualImage) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


