# Zone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Location** | Pointer to **NullableString** |  | [optional] 
**Owner** | Pointer to [**InlineResponse20040AppDeployInstance**](inline_response_200_40_appDeploy_instance.md) |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to [**InlineResponse20040AppDeployInstance**](inline_response_200_40_appDeploy_instance.md) |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusMessage** | Pointer to **NullableString** |  | [optional] 
**StatusDate** | Pointer to **NullableTime** |  | [optional] 
**CostStatus** | Pointer to **NullableString** |  | [optional] 
**CostStatusMessage** | Pointer to **NullableString** |  | [optional] 
**CostStatusDate** | Pointer to **NullableTime** |  | [optional] 
**CostLastSyncDuration** | Pointer to **NullableInt64** |  | [optional] 
**CostLastSync** | Pointer to **NullableTime** |  | [optional] 
**ZoneType** | Pointer to [**InlineResponse20079LoadBalancerMonitorLoadBalancerType**](inline_response_200_79_loadBalancerMonitor_loadBalancer_type.md) |  | [optional] 
**ZoneTypeId** | Pointer to **int64** |  | [optional] 
**GuidanceMode** | Pointer to **NullableString** |  | [optional] 
**StorageMode** | Pointer to **string** |  | [optional] 
**AgentMode** | Pointer to **string** |  | [optional] 
**UserDataLinux** | Pointer to **NullableString** |  | [optional] 
**UserDataWindows** | Pointer to **NullableString** |  | [optional] 
**ConsoleKeymap** | Pointer to **NullableString** |  | [optional] 
**ContainerMode** | Pointer to **string** |  | [optional] 
**CostingMode** | Pointer to **NullableString** |  | [optional] 
**ServiceVersion** | Pointer to **NullableString** |  | [optional] 
**SecurityMode** | Pointer to **string** |  | [optional] 
**InventoryLevel** | Pointer to **string** |  | [optional] 
**Timezone** | Pointer to **NullableString** |  | [optional] 
**ApiProxy** | Pointer to **NullableString** |  | [optional] 
**ProvisioningProxy** | Pointer to **NullableString** |  | [optional] 
**NetworkDomain** | Pointer to [**NullableInlineResponse20082LoadBalancerInstanceSslCert**](inline_response_200_82_loadBalancerInstance_sslCert.md) |  | [optional] 
**DomainName** | Pointer to **string** |  | [optional] 
**RegionCode** | Pointer to **NullableString** |  | [optional] 
**AutoRecoverPowerState** | Pointer to **bool** |  | [optional] 
**ScalePriority** | Pointer to **int64** |  | [optional] 
**Config** | Pointer to [**AnyOfzoneVcenterConfigzoneAwsConfigzoneAzureConfigzoneGcpConfig**](anyOf&lt;zoneVcenterConfig,zoneAwsConfig,zoneAzureConfig,zoneGcpConfig&gt;.md) |  | [optional] 
**Credential** | Pointer to [**AnyOfobjectobject**](anyOf&lt;object,object&gt;.md) |  | [optional] 
**ImagePath** | Pointer to **NullableString** | Logo image URL | [optional] 
**DarkImagePath** | Pointer to **NullableString** | Dark logo image URL | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**LastSync** | Pointer to **NullableTime** |  | [optional] 
**LastSyncDuration** | Pointer to **NullableInt64** |  | [optional] 
**NextRunDate** | Pointer to **NullableTime** |  | [optional] 
**Groups** | Pointer to [**[]ZoneGroups**](ZoneGroups.md) |  | [optional] 
**SecurityServer** | Pointer to [**NullableInlineResponse20082LoadBalancerInstanceSslCert**](inline_response_200_82_loadBalancerInstance_sslCert.md) |  | [optional] 
**NetworkServer** | Pointer to [**NullableInlineResponse20082LoadBalancerInstanceSslCert**](inline_response_200_82_loadBalancerInstance_sslCert.md) |  | [optional] 
**Stats** | Pointer to [**ZoneStats**](zone_stats.md) |  | [optional] 
**ServerCount** | Pointer to **int64** |  | [optional] 

## Methods

### NewZone

`func NewZone() *Zone`

NewZone instantiates a new Zone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneWithDefaults

`func NewZoneWithDefaults() *Zone`

NewZoneWithDefaults instantiates a new Zone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Zone) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Zone) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Zone) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Zone) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *Zone) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Zone) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Zone) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Zone) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetExternalId

`func (o *Zone) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *Zone) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *Zone) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *Zone) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *Zone) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *Zone) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetName

`func (o *Zone) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Zone) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Zone) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Zone) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *Zone) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Zone) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Zone) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Zone) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetLocation

`func (o *Zone) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Zone) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Zone) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Zone) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *Zone) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *Zone) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetOwner

`func (o *Zone) GetOwner() InlineResponse20040AppDeployInstance`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Zone) GetOwnerOk() (*InlineResponse20040AppDeployInstance, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Zone) SetOwner(v InlineResponse20040AppDeployInstance)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Zone) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetAccountId

`func (o *Zone) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Zone) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Zone) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *Zone) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccount

`func (o *Zone) GetAccount() InlineResponse20040AppDeployInstance`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Zone) GetAccountOk() (*InlineResponse20040AppDeployInstance, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Zone) SetAccount(v InlineResponse20040AppDeployInstance)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *Zone) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetVisibility

`func (o *Zone) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *Zone) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *Zone) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *Zone) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetEnabled

`func (o *Zone) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Zone) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Zone) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Zone) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetStatus

`func (o *Zone) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Zone) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Zone) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Zone) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *Zone) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *Zone) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *Zone) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *Zone) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *Zone) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *Zone) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetStatusDate

`func (o *Zone) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *Zone) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *Zone) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *Zone) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### SetStatusDateNil

`func (o *Zone) SetStatusDateNil(b bool)`

 SetStatusDateNil sets the value for StatusDate to be an explicit nil

### UnsetStatusDate
`func (o *Zone) UnsetStatusDate()`

UnsetStatusDate ensures that no value is present for StatusDate, not even an explicit nil
### GetCostStatus

`func (o *Zone) GetCostStatus() string`

GetCostStatus returns the CostStatus field if non-nil, zero value otherwise.

### GetCostStatusOk

`func (o *Zone) GetCostStatusOk() (*string, bool)`

GetCostStatusOk returns a tuple with the CostStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostStatus

`func (o *Zone) SetCostStatus(v string)`

SetCostStatus sets CostStatus field to given value.

### HasCostStatus

`func (o *Zone) HasCostStatus() bool`

HasCostStatus returns a boolean if a field has been set.

### SetCostStatusNil

`func (o *Zone) SetCostStatusNil(b bool)`

 SetCostStatusNil sets the value for CostStatus to be an explicit nil

### UnsetCostStatus
`func (o *Zone) UnsetCostStatus()`

UnsetCostStatus ensures that no value is present for CostStatus, not even an explicit nil
### GetCostStatusMessage

`func (o *Zone) GetCostStatusMessage() string`

GetCostStatusMessage returns the CostStatusMessage field if non-nil, zero value otherwise.

### GetCostStatusMessageOk

`func (o *Zone) GetCostStatusMessageOk() (*string, bool)`

GetCostStatusMessageOk returns a tuple with the CostStatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostStatusMessage

`func (o *Zone) SetCostStatusMessage(v string)`

SetCostStatusMessage sets CostStatusMessage field to given value.

### HasCostStatusMessage

`func (o *Zone) HasCostStatusMessage() bool`

HasCostStatusMessage returns a boolean if a field has been set.

### SetCostStatusMessageNil

`func (o *Zone) SetCostStatusMessageNil(b bool)`

 SetCostStatusMessageNil sets the value for CostStatusMessage to be an explicit nil

### UnsetCostStatusMessage
`func (o *Zone) UnsetCostStatusMessage()`

UnsetCostStatusMessage ensures that no value is present for CostStatusMessage, not even an explicit nil
### GetCostStatusDate

`func (o *Zone) GetCostStatusDate() time.Time`

GetCostStatusDate returns the CostStatusDate field if non-nil, zero value otherwise.

### GetCostStatusDateOk

`func (o *Zone) GetCostStatusDateOk() (*time.Time, bool)`

GetCostStatusDateOk returns a tuple with the CostStatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostStatusDate

`func (o *Zone) SetCostStatusDate(v time.Time)`

SetCostStatusDate sets CostStatusDate field to given value.

### HasCostStatusDate

`func (o *Zone) HasCostStatusDate() bool`

HasCostStatusDate returns a boolean if a field has been set.

### SetCostStatusDateNil

`func (o *Zone) SetCostStatusDateNil(b bool)`

 SetCostStatusDateNil sets the value for CostStatusDate to be an explicit nil

### UnsetCostStatusDate
`func (o *Zone) UnsetCostStatusDate()`

UnsetCostStatusDate ensures that no value is present for CostStatusDate, not even an explicit nil
### GetCostLastSyncDuration

`func (o *Zone) GetCostLastSyncDuration() int64`

GetCostLastSyncDuration returns the CostLastSyncDuration field if non-nil, zero value otherwise.

### GetCostLastSyncDurationOk

`func (o *Zone) GetCostLastSyncDurationOk() (*int64, bool)`

GetCostLastSyncDurationOk returns a tuple with the CostLastSyncDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostLastSyncDuration

`func (o *Zone) SetCostLastSyncDuration(v int64)`

SetCostLastSyncDuration sets CostLastSyncDuration field to given value.

### HasCostLastSyncDuration

`func (o *Zone) HasCostLastSyncDuration() bool`

HasCostLastSyncDuration returns a boolean if a field has been set.

### SetCostLastSyncDurationNil

`func (o *Zone) SetCostLastSyncDurationNil(b bool)`

 SetCostLastSyncDurationNil sets the value for CostLastSyncDuration to be an explicit nil

### UnsetCostLastSyncDuration
`func (o *Zone) UnsetCostLastSyncDuration()`

UnsetCostLastSyncDuration ensures that no value is present for CostLastSyncDuration, not even an explicit nil
### GetCostLastSync

`func (o *Zone) GetCostLastSync() time.Time`

GetCostLastSync returns the CostLastSync field if non-nil, zero value otherwise.

### GetCostLastSyncOk

`func (o *Zone) GetCostLastSyncOk() (*time.Time, bool)`

GetCostLastSyncOk returns a tuple with the CostLastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostLastSync

`func (o *Zone) SetCostLastSync(v time.Time)`

SetCostLastSync sets CostLastSync field to given value.

### HasCostLastSync

`func (o *Zone) HasCostLastSync() bool`

HasCostLastSync returns a boolean if a field has been set.

### SetCostLastSyncNil

`func (o *Zone) SetCostLastSyncNil(b bool)`

 SetCostLastSyncNil sets the value for CostLastSync to be an explicit nil

### UnsetCostLastSync
`func (o *Zone) UnsetCostLastSync()`

UnsetCostLastSync ensures that no value is present for CostLastSync, not even an explicit nil
### GetZoneType

`func (o *Zone) GetZoneType() InlineResponse20079LoadBalancerMonitorLoadBalancerType`

GetZoneType returns the ZoneType field if non-nil, zero value otherwise.

### GetZoneTypeOk

`func (o *Zone) GetZoneTypeOk() (*InlineResponse20079LoadBalancerMonitorLoadBalancerType, bool)`

GetZoneTypeOk returns a tuple with the ZoneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneType

`func (o *Zone) SetZoneType(v InlineResponse20079LoadBalancerMonitorLoadBalancerType)`

SetZoneType sets ZoneType field to given value.

### HasZoneType

`func (o *Zone) HasZoneType() bool`

HasZoneType returns a boolean if a field has been set.

### GetZoneTypeId

`func (o *Zone) GetZoneTypeId() int64`

GetZoneTypeId returns the ZoneTypeId field if non-nil, zero value otherwise.

### GetZoneTypeIdOk

`func (o *Zone) GetZoneTypeIdOk() (*int64, bool)`

GetZoneTypeIdOk returns a tuple with the ZoneTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneTypeId

`func (o *Zone) SetZoneTypeId(v int64)`

SetZoneTypeId sets ZoneTypeId field to given value.

### HasZoneTypeId

`func (o *Zone) HasZoneTypeId() bool`

HasZoneTypeId returns a boolean if a field has been set.

### GetGuidanceMode

`func (o *Zone) GetGuidanceMode() string`

GetGuidanceMode returns the GuidanceMode field if non-nil, zero value otherwise.

### GetGuidanceModeOk

`func (o *Zone) GetGuidanceModeOk() (*string, bool)`

GetGuidanceModeOk returns a tuple with the GuidanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuidanceMode

`func (o *Zone) SetGuidanceMode(v string)`

SetGuidanceMode sets GuidanceMode field to given value.

### HasGuidanceMode

`func (o *Zone) HasGuidanceMode() bool`

HasGuidanceMode returns a boolean if a field has been set.

### SetGuidanceModeNil

`func (o *Zone) SetGuidanceModeNil(b bool)`

 SetGuidanceModeNil sets the value for GuidanceMode to be an explicit nil

### UnsetGuidanceMode
`func (o *Zone) UnsetGuidanceMode()`

UnsetGuidanceMode ensures that no value is present for GuidanceMode, not even an explicit nil
### GetStorageMode

`func (o *Zone) GetStorageMode() string`

GetStorageMode returns the StorageMode field if non-nil, zero value otherwise.

### GetStorageModeOk

`func (o *Zone) GetStorageModeOk() (*string, bool)`

GetStorageModeOk returns a tuple with the StorageMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageMode

`func (o *Zone) SetStorageMode(v string)`

SetStorageMode sets StorageMode field to given value.

### HasStorageMode

`func (o *Zone) HasStorageMode() bool`

HasStorageMode returns a boolean if a field has been set.

### GetAgentMode

`func (o *Zone) GetAgentMode() string`

GetAgentMode returns the AgentMode field if non-nil, zero value otherwise.

### GetAgentModeOk

`func (o *Zone) GetAgentModeOk() (*string, bool)`

GetAgentModeOk returns a tuple with the AgentMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentMode

`func (o *Zone) SetAgentMode(v string)`

SetAgentMode sets AgentMode field to given value.

### HasAgentMode

`func (o *Zone) HasAgentMode() bool`

HasAgentMode returns a boolean if a field has been set.

### GetUserDataLinux

`func (o *Zone) GetUserDataLinux() string`

GetUserDataLinux returns the UserDataLinux field if non-nil, zero value otherwise.

### GetUserDataLinuxOk

`func (o *Zone) GetUserDataLinuxOk() (*string, bool)`

GetUserDataLinuxOk returns a tuple with the UserDataLinux field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDataLinux

`func (o *Zone) SetUserDataLinux(v string)`

SetUserDataLinux sets UserDataLinux field to given value.

### HasUserDataLinux

`func (o *Zone) HasUserDataLinux() bool`

HasUserDataLinux returns a boolean if a field has been set.

### SetUserDataLinuxNil

`func (o *Zone) SetUserDataLinuxNil(b bool)`

 SetUserDataLinuxNil sets the value for UserDataLinux to be an explicit nil

### UnsetUserDataLinux
`func (o *Zone) UnsetUserDataLinux()`

UnsetUserDataLinux ensures that no value is present for UserDataLinux, not even an explicit nil
### GetUserDataWindows

`func (o *Zone) GetUserDataWindows() string`

GetUserDataWindows returns the UserDataWindows field if non-nil, zero value otherwise.

### GetUserDataWindowsOk

`func (o *Zone) GetUserDataWindowsOk() (*string, bool)`

GetUserDataWindowsOk returns a tuple with the UserDataWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDataWindows

`func (o *Zone) SetUserDataWindows(v string)`

SetUserDataWindows sets UserDataWindows field to given value.

### HasUserDataWindows

`func (o *Zone) HasUserDataWindows() bool`

HasUserDataWindows returns a boolean if a field has been set.

### SetUserDataWindowsNil

`func (o *Zone) SetUserDataWindowsNil(b bool)`

 SetUserDataWindowsNil sets the value for UserDataWindows to be an explicit nil

### UnsetUserDataWindows
`func (o *Zone) UnsetUserDataWindows()`

UnsetUserDataWindows ensures that no value is present for UserDataWindows, not even an explicit nil
### GetConsoleKeymap

`func (o *Zone) GetConsoleKeymap() string`

GetConsoleKeymap returns the ConsoleKeymap field if non-nil, zero value otherwise.

### GetConsoleKeymapOk

`func (o *Zone) GetConsoleKeymapOk() (*string, bool)`

GetConsoleKeymapOk returns a tuple with the ConsoleKeymap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleKeymap

`func (o *Zone) SetConsoleKeymap(v string)`

SetConsoleKeymap sets ConsoleKeymap field to given value.

### HasConsoleKeymap

`func (o *Zone) HasConsoleKeymap() bool`

HasConsoleKeymap returns a boolean if a field has been set.

### SetConsoleKeymapNil

`func (o *Zone) SetConsoleKeymapNil(b bool)`

 SetConsoleKeymapNil sets the value for ConsoleKeymap to be an explicit nil

### UnsetConsoleKeymap
`func (o *Zone) UnsetConsoleKeymap()`

UnsetConsoleKeymap ensures that no value is present for ConsoleKeymap, not even an explicit nil
### GetContainerMode

`func (o *Zone) GetContainerMode() string`

GetContainerMode returns the ContainerMode field if non-nil, zero value otherwise.

### GetContainerModeOk

`func (o *Zone) GetContainerModeOk() (*string, bool)`

GetContainerModeOk returns a tuple with the ContainerMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerMode

`func (o *Zone) SetContainerMode(v string)`

SetContainerMode sets ContainerMode field to given value.

### HasContainerMode

`func (o *Zone) HasContainerMode() bool`

HasContainerMode returns a boolean if a field has been set.

### GetCostingMode

`func (o *Zone) GetCostingMode() string`

GetCostingMode returns the CostingMode field if non-nil, zero value otherwise.

### GetCostingModeOk

`func (o *Zone) GetCostingModeOk() (*string, bool)`

GetCostingModeOk returns a tuple with the CostingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingMode

`func (o *Zone) SetCostingMode(v string)`

SetCostingMode sets CostingMode field to given value.

### HasCostingMode

`func (o *Zone) HasCostingMode() bool`

HasCostingMode returns a boolean if a field has been set.

### SetCostingModeNil

`func (o *Zone) SetCostingModeNil(b bool)`

 SetCostingModeNil sets the value for CostingMode to be an explicit nil

### UnsetCostingMode
`func (o *Zone) UnsetCostingMode()`

UnsetCostingMode ensures that no value is present for CostingMode, not even an explicit nil
### GetServiceVersion

`func (o *Zone) GetServiceVersion() string`

GetServiceVersion returns the ServiceVersion field if non-nil, zero value otherwise.

### GetServiceVersionOk

`func (o *Zone) GetServiceVersionOk() (*string, bool)`

GetServiceVersionOk returns a tuple with the ServiceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceVersion

`func (o *Zone) SetServiceVersion(v string)`

SetServiceVersion sets ServiceVersion field to given value.

### HasServiceVersion

`func (o *Zone) HasServiceVersion() bool`

HasServiceVersion returns a boolean if a field has been set.

### SetServiceVersionNil

`func (o *Zone) SetServiceVersionNil(b bool)`

 SetServiceVersionNil sets the value for ServiceVersion to be an explicit nil

### UnsetServiceVersion
`func (o *Zone) UnsetServiceVersion()`

UnsetServiceVersion ensures that no value is present for ServiceVersion, not even an explicit nil
### GetSecurityMode

`func (o *Zone) GetSecurityMode() string`

GetSecurityMode returns the SecurityMode field if non-nil, zero value otherwise.

### GetSecurityModeOk

`func (o *Zone) GetSecurityModeOk() (*string, bool)`

GetSecurityModeOk returns a tuple with the SecurityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityMode

`func (o *Zone) SetSecurityMode(v string)`

SetSecurityMode sets SecurityMode field to given value.

### HasSecurityMode

`func (o *Zone) HasSecurityMode() bool`

HasSecurityMode returns a boolean if a field has been set.

### GetInventoryLevel

`func (o *Zone) GetInventoryLevel() string`

GetInventoryLevel returns the InventoryLevel field if non-nil, zero value otherwise.

### GetInventoryLevelOk

`func (o *Zone) GetInventoryLevelOk() (*string, bool)`

GetInventoryLevelOk returns a tuple with the InventoryLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryLevel

`func (o *Zone) SetInventoryLevel(v string)`

SetInventoryLevel sets InventoryLevel field to given value.

### HasInventoryLevel

`func (o *Zone) HasInventoryLevel() bool`

HasInventoryLevel returns a boolean if a field has been set.

### GetTimezone

`func (o *Zone) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *Zone) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *Zone) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *Zone) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### SetTimezoneNil

`func (o *Zone) SetTimezoneNil(b bool)`

 SetTimezoneNil sets the value for Timezone to be an explicit nil

### UnsetTimezone
`func (o *Zone) UnsetTimezone()`

UnsetTimezone ensures that no value is present for Timezone, not even an explicit nil
### GetApiProxy

`func (o *Zone) GetApiProxy() string`

GetApiProxy returns the ApiProxy field if non-nil, zero value otherwise.

### GetApiProxyOk

`func (o *Zone) GetApiProxyOk() (*string, bool)`

GetApiProxyOk returns a tuple with the ApiProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiProxy

`func (o *Zone) SetApiProxy(v string)`

SetApiProxy sets ApiProxy field to given value.

### HasApiProxy

`func (o *Zone) HasApiProxy() bool`

HasApiProxy returns a boolean if a field has been set.

### SetApiProxyNil

`func (o *Zone) SetApiProxyNil(b bool)`

 SetApiProxyNil sets the value for ApiProxy to be an explicit nil

### UnsetApiProxy
`func (o *Zone) UnsetApiProxy()`

UnsetApiProxy ensures that no value is present for ApiProxy, not even an explicit nil
### GetProvisioningProxy

`func (o *Zone) GetProvisioningProxy() string`

GetProvisioningProxy returns the ProvisioningProxy field if non-nil, zero value otherwise.

### GetProvisioningProxyOk

`func (o *Zone) GetProvisioningProxyOk() (*string, bool)`

GetProvisioningProxyOk returns a tuple with the ProvisioningProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningProxy

`func (o *Zone) SetProvisioningProxy(v string)`

SetProvisioningProxy sets ProvisioningProxy field to given value.

### HasProvisioningProxy

`func (o *Zone) HasProvisioningProxy() bool`

HasProvisioningProxy returns a boolean if a field has been set.

### SetProvisioningProxyNil

`func (o *Zone) SetProvisioningProxyNil(b bool)`

 SetProvisioningProxyNil sets the value for ProvisioningProxy to be an explicit nil

### UnsetProvisioningProxy
`func (o *Zone) UnsetProvisioningProxy()`

UnsetProvisioningProxy ensures that no value is present for ProvisioningProxy, not even an explicit nil
### GetNetworkDomain

`func (o *Zone) GetNetworkDomain() InlineResponse20082LoadBalancerInstanceSslCert`

GetNetworkDomain returns the NetworkDomain field if non-nil, zero value otherwise.

### GetNetworkDomainOk

`func (o *Zone) GetNetworkDomainOk() (*InlineResponse20082LoadBalancerInstanceSslCert, bool)`

GetNetworkDomainOk returns a tuple with the NetworkDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDomain

`func (o *Zone) SetNetworkDomain(v InlineResponse20082LoadBalancerInstanceSslCert)`

SetNetworkDomain sets NetworkDomain field to given value.

### HasNetworkDomain

`func (o *Zone) HasNetworkDomain() bool`

HasNetworkDomain returns a boolean if a field has been set.

### SetNetworkDomainNil

`func (o *Zone) SetNetworkDomainNil(b bool)`

 SetNetworkDomainNil sets the value for NetworkDomain to be an explicit nil

### UnsetNetworkDomain
`func (o *Zone) UnsetNetworkDomain()`

UnsetNetworkDomain ensures that no value is present for NetworkDomain, not even an explicit nil
### GetDomainName

`func (o *Zone) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *Zone) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *Zone) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *Zone) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetRegionCode

`func (o *Zone) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *Zone) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *Zone) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *Zone) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### SetRegionCodeNil

`func (o *Zone) SetRegionCodeNil(b bool)`

 SetRegionCodeNil sets the value for RegionCode to be an explicit nil

### UnsetRegionCode
`func (o *Zone) UnsetRegionCode()`

UnsetRegionCode ensures that no value is present for RegionCode, not even an explicit nil
### GetAutoRecoverPowerState

`func (o *Zone) GetAutoRecoverPowerState() bool`

GetAutoRecoverPowerState returns the AutoRecoverPowerState field if non-nil, zero value otherwise.

### GetAutoRecoverPowerStateOk

`func (o *Zone) GetAutoRecoverPowerStateOk() (*bool, bool)`

GetAutoRecoverPowerStateOk returns a tuple with the AutoRecoverPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRecoverPowerState

`func (o *Zone) SetAutoRecoverPowerState(v bool)`

SetAutoRecoverPowerState sets AutoRecoverPowerState field to given value.

### HasAutoRecoverPowerState

`func (o *Zone) HasAutoRecoverPowerState() bool`

HasAutoRecoverPowerState returns a boolean if a field has been set.

### GetScalePriority

`func (o *Zone) GetScalePriority() int64`

GetScalePriority returns the ScalePriority field if non-nil, zero value otherwise.

### GetScalePriorityOk

`func (o *Zone) GetScalePriorityOk() (*int64, bool)`

GetScalePriorityOk returns a tuple with the ScalePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScalePriority

`func (o *Zone) SetScalePriority(v int64)`

SetScalePriority sets ScalePriority field to given value.

### HasScalePriority

`func (o *Zone) HasScalePriority() bool`

HasScalePriority returns a boolean if a field has been set.

### GetConfig

`func (o *Zone) GetConfig() AnyOfzoneVcenterConfigzoneAwsConfigzoneAzureConfigzoneGcpConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Zone) GetConfigOk() (*AnyOfzoneVcenterConfigzoneAwsConfigzoneAzureConfigzoneGcpConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Zone) SetConfig(v AnyOfzoneVcenterConfigzoneAwsConfigzoneAzureConfigzoneGcpConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *Zone) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCredential

`func (o *Zone) GetCredential() AnyOfobjectobject`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *Zone) GetCredentialOk() (*AnyOfobjectobject, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *Zone) SetCredential(v AnyOfobjectobject)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *Zone) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetImagePath

`func (o *Zone) GetImagePath() string`

GetImagePath returns the ImagePath field if non-nil, zero value otherwise.

### GetImagePathOk

`func (o *Zone) GetImagePathOk() (*string, bool)`

GetImagePathOk returns a tuple with the ImagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagePath

`func (o *Zone) SetImagePath(v string)`

SetImagePath sets ImagePath field to given value.

### HasImagePath

`func (o *Zone) HasImagePath() bool`

HasImagePath returns a boolean if a field has been set.

### SetImagePathNil

`func (o *Zone) SetImagePathNil(b bool)`

 SetImagePathNil sets the value for ImagePath to be an explicit nil

### UnsetImagePath
`func (o *Zone) UnsetImagePath()`

UnsetImagePath ensures that no value is present for ImagePath, not even an explicit nil
### GetDarkImagePath

`func (o *Zone) GetDarkImagePath() string`

GetDarkImagePath returns the DarkImagePath field if non-nil, zero value otherwise.

### GetDarkImagePathOk

`func (o *Zone) GetDarkImagePathOk() (*string, bool)`

GetDarkImagePathOk returns a tuple with the DarkImagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDarkImagePath

`func (o *Zone) SetDarkImagePath(v string)`

SetDarkImagePath sets DarkImagePath field to given value.

### HasDarkImagePath

`func (o *Zone) HasDarkImagePath() bool`

HasDarkImagePath returns a boolean if a field has been set.

### SetDarkImagePathNil

`func (o *Zone) SetDarkImagePathNil(b bool)`

 SetDarkImagePathNil sets the value for DarkImagePath to be an explicit nil

### UnsetDarkImagePath
`func (o *Zone) UnsetDarkImagePath()`

UnsetDarkImagePath ensures that no value is present for DarkImagePath, not even an explicit nil
### GetDateCreated

`func (o *Zone) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *Zone) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *Zone) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *Zone) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Zone) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Zone) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Zone) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Zone) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLastSync

`func (o *Zone) GetLastSync() time.Time`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *Zone) GetLastSyncOk() (*time.Time, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *Zone) SetLastSync(v time.Time)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *Zone) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.

### SetLastSyncNil

`func (o *Zone) SetLastSyncNil(b bool)`

 SetLastSyncNil sets the value for LastSync to be an explicit nil

### UnsetLastSync
`func (o *Zone) UnsetLastSync()`

UnsetLastSync ensures that no value is present for LastSync, not even an explicit nil
### GetLastSyncDuration

`func (o *Zone) GetLastSyncDuration() int64`

GetLastSyncDuration returns the LastSyncDuration field if non-nil, zero value otherwise.

### GetLastSyncDurationOk

`func (o *Zone) GetLastSyncDurationOk() (*int64, bool)`

GetLastSyncDurationOk returns a tuple with the LastSyncDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncDuration

`func (o *Zone) SetLastSyncDuration(v int64)`

SetLastSyncDuration sets LastSyncDuration field to given value.

### HasLastSyncDuration

`func (o *Zone) HasLastSyncDuration() bool`

HasLastSyncDuration returns a boolean if a field has been set.

### SetLastSyncDurationNil

`func (o *Zone) SetLastSyncDurationNil(b bool)`

 SetLastSyncDurationNil sets the value for LastSyncDuration to be an explicit nil

### UnsetLastSyncDuration
`func (o *Zone) UnsetLastSyncDuration()`

UnsetLastSyncDuration ensures that no value is present for LastSyncDuration, not even an explicit nil
### GetNextRunDate

`func (o *Zone) GetNextRunDate() time.Time`

GetNextRunDate returns the NextRunDate field if non-nil, zero value otherwise.

### GetNextRunDateOk

`func (o *Zone) GetNextRunDateOk() (*time.Time, bool)`

GetNextRunDateOk returns a tuple with the NextRunDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRunDate

`func (o *Zone) SetNextRunDate(v time.Time)`

SetNextRunDate sets NextRunDate field to given value.

### HasNextRunDate

`func (o *Zone) HasNextRunDate() bool`

HasNextRunDate returns a boolean if a field has been set.

### SetNextRunDateNil

`func (o *Zone) SetNextRunDateNil(b bool)`

 SetNextRunDateNil sets the value for NextRunDate to be an explicit nil

### UnsetNextRunDate
`func (o *Zone) UnsetNextRunDate()`

UnsetNextRunDate ensures that no value is present for NextRunDate, not even an explicit nil
### GetGroups

`func (o *Zone) GetGroups() []ZoneGroups`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *Zone) GetGroupsOk() (*[]ZoneGroups, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *Zone) SetGroups(v []ZoneGroups)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *Zone) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetSecurityServer

`func (o *Zone) GetSecurityServer() InlineResponse20082LoadBalancerInstanceSslCert`

GetSecurityServer returns the SecurityServer field if non-nil, zero value otherwise.

### GetSecurityServerOk

`func (o *Zone) GetSecurityServerOk() (*InlineResponse20082LoadBalancerInstanceSslCert, bool)`

GetSecurityServerOk returns a tuple with the SecurityServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityServer

`func (o *Zone) SetSecurityServer(v InlineResponse20082LoadBalancerInstanceSslCert)`

SetSecurityServer sets SecurityServer field to given value.

### HasSecurityServer

`func (o *Zone) HasSecurityServer() bool`

HasSecurityServer returns a boolean if a field has been set.

### SetSecurityServerNil

`func (o *Zone) SetSecurityServerNil(b bool)`

 SetSecurityServerNil sets the value for SecurityServer to be an explicit nil

### UnsetSecurityServer
`func (o *Zone) UnsetSecurityServer()`

UnsetSecurityServer ensures that no value is present for SecurityServer, not even an explicit nil
### GetNetworkServer

`func (o *Zone) GetNetworkServer() InlineResponse20082LoadBalancerInstanceSslCert`

GetNetworkServer returns the NetworkServer field if non-nil, zero value otherwise.

### GetNetworkServerOk

`func (o *Zone) GetNetworkServerOk() (*InlineResponse20082LoadBalancerInstanceSslCert, bool)`

GetNetworkServerOk returns a tuple with the NetworkServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkServer

`func (o *Zone) SetNetworkServer(v InlineResponse20082LoadBalancerInstanceSslCert)`

SetNetworkServer sets NetworkServer field to given value.

### HasNetworkServer

`func (o *Zone) HasNetworkServer() bool`

HasNetworkServer returns a boolean if a field has been set.

### SetNetworkServerNil

`func (o *Zone) SetNetworkServerNil(b bool)`

 SetNetworkServerNil sets the value for NetworkServer to be an explicit nil

### UnsetNetworkServer
`func (o *Zone) UnsetNetworkServer()`

UnsetNetworkServer ensures that no value is present for NetworkServer, not even an explicit nil
### GetStats

`func (o *Zone) GetStats() ZoneStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *Zone) GetStatsOk() (*ZoneStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *Zone) SetStats(v ZoneStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *Zone) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetServerCount

`func (o *Zone) GetServerCount() int64`

GetServerCount returns the ServerCount field if non-nil, zero value otherwise.

### GetServerCountOk

`func (o *Zone) GetServerCountOk() (*int64, bool)`

GetServerCountOk returns a tuple with the ServerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCount

`func (o *Zone) SetServerCount(v int64)`

SetServerCount sets ServerCount field to given value.

### HasServerCount

`func (o *Zone) HasServerCount() bool`

HasServerCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


