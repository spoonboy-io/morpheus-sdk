# ImageBuild

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to [**InlineResponse20040AppDeployInstance**](inline_response_200_40_appDeploy_instance.md) |  | [optional] 
**Type** | Pointer to [**InlineResponse20079LoadBalancerMonitorLoadBalancerType**](inline_response_200_79_loadBalancerMonitor_loadBalancer_type.md) |  | [optional] 
**Site** | Pointer to [**InlineResponse20040AppDeployInstance**](inline_response_200_40_appDeploy_instance.md) |  | [optional] 
**Zone** | Pointer to [**InlineResponse20040AppDeployInstance**](inline_response_200_40_appDeploy_instance.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**BootScript** | Pointer to [**ImageBuildsBootScript**](imageBuilds_bootScript.md) |  | [optional] 
**BootCommand** | Pointer to **NullableString** |  | [optional] 
**PreseedScript** | Pointer to [**ImageBuildsBootScript**](imageBuilds_bootScript.md) |  | [optional] 
**Scripts** | Pointer to [**[]ImageBuildsScripts**](ImageBuildsScripts.md) |  | [optional] 
**SshUsername** | Pointer to **string** |  | [optional] 
**SshPassword** | Pointer to **string** |  | [optional] 
**StorageProvider** | Pointer to **NullableString** |  | [optional] 
**BuildOutputName** | Pointer to **NullableString** |  | [optional] 
**ConversionFormats** | Pointer to **NullableString** |  | [optional] 
**IsCloudInit** | Pointer to **bool** |  | [optional] 
**VmToolsInstalled** | Pointer to **bool** |  | [optional] 
**KeepResults** | Pointer to **NullableInt64** |  | [optional] 
**Config** | Pointer to [**ImageBuildConfig**](imageBuild_config.md) |  | [optional] 
**LastResult** | Pointer to [**NullableImageBuildLastResult**](imageBuild_lastResult.md) |  | [optional] 
**ExecutionCount** | Pointer to **int64** |  | [optional] 

## Methods

### NewImageBuild

`func NewImageBuild() *ImageBuild`

NewImageBuild instantiates a new ImageBuild object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageBuildWithDefaults

`func NewImageBuildWithDefaults() *ImageBuild`

NewImageBuildWithDefaults instantiates a new ImageBuild object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ImageBuild) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImageBuild) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImageBuild) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ImageBuild) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *ImageBuild) GetAccount() InlineResponse20040AppDeployInstance`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ImageBuild) GetAccountOk() (*InlineResponse20040AppDeployInstance, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ImageBuild) SetAccount(v InlineResponse20040AppDeployInstance)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ImageBuild) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetType

`func (o *ImageBuild) GetType() InlineResponse20079LoadBalancerMonitorLoadBalancerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ImageBuild) GetTypeOk() (*InlineResponse20079LoadBalancerMonitorLoadBalancerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ImageBuild) SetType(v InlineResponse20079LoadBalancerMonitorLoadBalancerType)`

SetType sets Type field to given value.

### HasType

`func (o *ImageBuild) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSite

`func (o *ImageBuild) GetSite() InlineResponse20040AppDeployInstance`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *ImageBuild) GetSiteOk() (*InlineResponse20040AppDeployInstance, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *ImageBuild) SetSite(v InlineResponse20040AppDeployInstance)`

SetSite sets Site field to given value.

### HasSite

`func (o *ImageBuild) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetZone

`func (o *ImageBuild) GetZone() InlineResponse20040AppDeployInstance`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *ImageBuild) GetZoneOk() (*InlineResponse20040AppDeployInstance, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *ImageBuild) SetZone(v InlineResponse20040AppDeployInstance)`

SetZone sets Zone field to given value.

### HasZone

`func (o *ImageBuild) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetName

`func (o *ImageBuild) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImageBuild) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImageBuild) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ImageBuild) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ImageBuild) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ImageBuild) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ImageBuild) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ImageBuild) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ImageBuild) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ImageBuild) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetBootScript

`func (o *ImageBuild) GetBootScript() ImageBuildsBootScript`

GetBootScript returns the BootScript field if non-nil, zero value otherwise.

### GetBootScriptOk

`func (o *ImageBuild) GetBootScriptOk() (*ImageBuildsBootScript, bool)`

GetBootScriptOk returns a tuple with the BootScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootScript

`func (o *ImageBuild) SetBootScript(v ImageBuildsBootScript)`

SetBootScript sets BootScript field to given value.

### HasBootScript

`func (o *ImageBuild) HasBootScript() bool`

HasBootScript returns a boolean if a field has been set.

### GetBootCommand

`func (o *ImageBuild) GetBootCommand() string`

GetBootCommand returns the BootCommand field if non-nil, zero value otherwise.

### GetBootCommandOk

`func (o *ImageBuild) GetBootCommandOk() (*string, bool)`

GetBootCommandOk returns a tuple with the BootCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootCommand

`func (o *ImageBuild) SetBootCommand(v string)`

SetBootCommand sets BootCommand field to given value.

### HasBootCommand

`func (o *ImageBuild) HasBootCommand() bool`

HasBootCommand returns a boolean if a field has been set.

### SetBootCommandNil

`func (o *ImageBuild) SetBootCommandNil(b bool)`

 SetBootCommandNil sets the value for BootCommand to be an explicit nil

### UnsetBootCommand
`func (o *ImageBuild) UnsetBootCommand()`

UnsetBootCommand ensures that no value is present for BootCommand, not even an explicit nil
### GetPreseedScript

`func (o *ImageBuild) GetPreseedScript() ImageBuildsBootScript`

GetPreseedScript returns the PreseedScript field if non-nil, zero value otherwise.

### GetPreseedScriptOk

`func (o *ImageBuild) GetPreseedScriptOk() (*ImageBuildsBootScript, bool)`

GetPreseedScriptOk returns a tuple with the PreseedScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreseedScript

`func (o *ImageBuild) SetPreseedScript(v ImageBuildsBootScript)`

SetPreseedScript sets PreseedScript field to given value.

### HasPreseedScript

`func (o *ImageBuild) HasPreseedScript() bool`

HasPreseedScript returns a boolean if a field has been set.

### GetScripts

`func (o *ImageBuild) GetScripts() []ImageBuildsScripts`

GetScripts returns the Scripts field if non-nil, zero value otherwise.

### GetScriptsOk

`func (o *ImageBuild) GetScriptsOk() (*[]ImageBuildsScripts, bool)`

GetScriptsOk returns a tuple with the Scripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScripts

`func (o *ImageBuild) SetScripts(v []ImageBuildsScripts)`

SetScripts sets Scripts field to given value.

### HasScripts

`func (o *ImageBuild) HasScripts() bool`

HasScripts returns a boolean if a field has been set.

### GetSshUsername

`func (o *ImageBuild) GetSshUsername() string`

GetSshUsername returns the SshUsername field if non-nil, zero value otherwise.

### GetSshUsernameOk

`func (o *ImageBuild) GetSshUsernameOk() (*string, bool)`

GetSshUsernameOk returns a tuple with the SshUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUsername

`func (o *ImageBuild) SetSshUsername(v string)`

SetSshUsername sets SshUsername field to given value.

### HasSshUsername

`func (o *ImageBuild) HasSshUsername() bool`

HasSshUsername returns a boolean if a field has been set.

### GetSshPassword

`func (o *ImageBuild) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *ImageBuild) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *ImageBuild) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *ImageBuild) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.

### GetStorageProvider

`func (o *ImageBuild) GetStorageProvider() string`

GetStorageProvider returns the StorageProvider field if non-nil, zero value otherwise.

### GetStorageProviderOk

`func (o *ImageBuild) GetStorageProviderOk() (*string, bool)`

GetStorageProviderOk returns a tuple with the StorageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProvider

`func (o *ImageBuild) SetStorageProvider(v string)`

SetStorageProvider sets StorageProvider field to given value.

### HasStorageProvider

`func (o *ImageBuild) HasStorageProvider() bool`

HasStorageProvider returns a boolean if a field has been set.

### SetStorageProviderNil

`func (o *ImageBuild) SetStorageProviderNil(b bool)`

 SetStorageProviderNil sets the value for StorageProvider to be an explicit nil

### UnsetStorageProvider
`func (o *ImageBuild) UnsetStorageProvider()`

UnsetStorageProvider ensures that no value is present for StorageProvider, not even an explicit nil
### GetBuildOutputName

`func (o *ImageBuild) GetBuildOutputName() string`

GetBuildOutputName returns the BuildOutputName field if non-nil, zero value otherwise.

### GetBuildOutputNameOk

`func (o *ImageBuild) GetBuildOutputNameOk() (*string, bool)`

GetBuildOutputNameOk returns a tuple with the BuildOutputName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildOutputName

`func (o *ImageBuild) SetBuildOutputName(v string)`

SetBuildOutputName sets BuildOutputName field to given value.

### HasBuildOutputName

`func (o *ImageBuild) HasBuildOutputName() bool`

HasBuildOutputName returns a boolean if a field has been set.

### SetBuildOutputNameNil

`func (o *ImageBuild) SetBuildOutputNameNil(b bool)`

 SetBuildOutputNameNil sets the value for BuildOutputName to be an explicit nil

### UnsetBuildOutputName
`func (o *ImageBuild) UnsetBuildOutputName()`

UnsetBuildOutputName ensures that no value is present for BuildOutputName, not even an explicit nil
### GetConversionFormats

`func (o *ImageBuild) GetConversionFormats() string`

GetConversionFormats returns the ConversionFormats field if non-nil, zero value otherwise.

### GetConversionFormatsOk

`func (o *ImageBuild) GetConversionFormatsOk() (*string, bool)`

GetConversionFormatsOk returns a tuple with the ConversionFormats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionFormats

`func (o *ImageBuild) SetConversionFormats(v string)`

SetConversionFormats sets ConversionFormats field to given value.

### HasConversionFormats

`func (o *ImageBuild) HasConversionFormats() bool`

HasConversionFormats returns a boolean if a field has been set.

### SetConversionFormatsNil

`func (o *ImageBuild) SetConversionFormatsNil(b bool)`

 SetConversionFormatsNil sets the value for ConversionFormats to be an explicit nil

### UnsetConversionFormats
`func (o *ImageBuild) UnsetConversionFormats()`

UnsetConversionFormats ensures that no value is present for ConversionFormats, not even an explicit nil
### GetIsCloudInit

`func (o *ImageBuild) GetIsCloudInit() bool`

GetIsCloudInit returns the IsCloudInit field if non-nil, zero value otherwise.

### GetIsCloudInitOk

`func (o *ImageBuild) GetIsCloudInitOk() (*bool, bool)`

GetIsCloudInitOk returns a tuple with the IsCloudInit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCloudInit

`func (o *ImageBuild) SetIsCloudInit(v bool)`

SetIsCloudInit sets IsCloudInit field to given value.

### HasIsCloudInit

`func (o *ImageBuild) HasIsCloudInit() bool`

HasIsCloudInit returns a boolean if a field has been set.

### GetVmToolsInstalled

`func (o *ImageBuild) GetVmToolsInstalled() bool`

GetVmToolsInstalled returns the VmToolsInstalled field if non-nil, zero value otherwise.

### GetVmToolsInstalledOk

`func (o *ImageBuild) GetVmToolsInstalledOk() (*bool, bool)`

GetVmToolsInstalledOk returns a tuple with the VmToolsInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmToolsInstalled

`func (o *ImageBuild) SetVmToolsInstalled(v bool)`

SetVmToolsInstalled sets VmToolsInstalled field to given value.

### HasVmToolsInstalled

`func (o *ImageBuild) HasVmToolsInstalled() bool`

HasVmToolsInstalled returns a boolean if a field has been set.

### GetKeepResults

`func (o *ImageBuild) GetKeepResults() int64`

GetKeepResults returns the KeepResults field if non-nil, zero value otherwise.

### GetKeepResultsOk

`func (o *ImageBuild) GetKeepResultsOk() (*int64, bool)`

GetKeepResultsOk returns a tuple with the KeepResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepResults

`func (o *ImageBuild) SetKeepResults(v int64)`

SetKeepResults sets KeepResults field to given value.

### HasKeepResults

`func (o *ImageBuild) HasKeepResults() bool`

HasKeepResults returns a boolean if a field has been set.

### SetKeepResultsNil

`func (o *ImageBuild) SetKeepResultsNil(b bool)`

 SetKeepResultsNil sets the value for KeepResults to be an explicit nil

### UnsetKeepResults
`func (o *ImageBuild) UnsetKeepResults()`

UnsetKeepResults ensures that no value is present for KeepResults, not even an explicit nil
### GetConfig

`func (o *ImageBuild) GetConfig() ImageBuildConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ImageBuild) GetConfigOk() (*ImageBuildConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ImageBuild) SetConfig(v ImageBuildConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ImageBuild) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetLastResult

`func (o *ImageBuild) GetLastResult() ImageBuildLastResult`

GetLastResult returns the LastResult field if non-nil, zero value otherwise.

### GetLastResultOk

`func (o *ImageBuild) GetLastResultOk() (*ImageBuildLastResult, bool)`

GetLastResultOk returns a tuple with the LastResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastResult

`func (o *ImageBuild) SetLastResult(v ImageBuildLastResult)`

SetLastResult sets LastResult field to given value.

### HasLastResult

`func (o *ImageBuild) HasLastResult() bool`

HasLastResult returns a boolean if a field has been set.

### SetLastResultNil

`func (o *ImageBuild) SetLastResultNil(b bool)`

 SetLastResultNil sets the value for LastResult to be an explicit nil

### UnsetLastResult
`func (o *ImageBuild) UnsetLastResult()`

UnsetLastResult ensures that no value is present for LastResult, not even an explicit nil
### GetExecutionCount

`func (o *ImageBuild) GetExecutionCount() int64`

GetExecutionCount returns the ExecutionCount field if non-nil, zero value otherwise.

### GetExecutionCountOk

`func (o *ImageBuild) GetExecutionCountOk() (*int64, bool)`

GetExecutionCountOk returns a tuple with the ExecutionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionCount

`func (o *ImageBuild) SetExecutionCount(v int64)`

SetExecutionCount sets ExecutionCount field to given value.

### HasExecutionCount

`func (o *ImageBuild) HasExecutionCount() bool`

HasExecutionCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


