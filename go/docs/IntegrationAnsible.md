# IntegrationAnsible

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**IntegrationType** | Pointer to [**InlineResponse20079LoadBalancerMonitorLoadBalancerType**](inline_response_200_79_loadBalancerMonitor_loadBalancer_type.md) |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**ServiceKey** | Pointer to [**InlineResponse20040AppDeployInstance**](inline_response_200_40_appDeploy_instance.md) |  | [optional] 
**IsPlugin** | Pointer to **bool** |  | [optional] 
**Config** | Pointer to [**IntegrationAnsibleConfig**](integrationAnsible_config.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusDate** | Pointer to **time.Time** |  | [optional] 
**StatusMessage** | Pointer to **NullableString** |  | [optional] 
**LastSync** | Pointer to **NullableString** |  | [optional] 
**LastSyncDuration** | Pointer to **NullableString** |  | [optional] 
**Credential** | Pointer to [**Creds**](creds.md) |  | [optional] 

## Methods

### NewIntegrationAnsible

`func NewIntegrationAnsible() *IntegrationAnsible`

NewIntegrationAnsible instantiates a new IntegrationAnsible object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationAnsibleWithDefaults

`func NewIntegrationAnsibleWithDefaults() *IntegrationAnsible`

NewIntegrationAnsibleWithDefaults instantiates a new IntegrationAnsible object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IntegrationAnsible) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IntegrationAnsible) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IntegrationAnsible) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *IntegrationAnsible) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IntegrationAnsible) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IntegrationAnsible) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IntegrationAnsible) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IntegrationAnsible) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *IntegrationAnsible) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IntegrationAnsible) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IntegrationAnsible) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *IntegrationAnsible) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetType

`func (o *IntegrationAnsible) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IntegrationAnsible) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IntegrationAnsible) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IntegrationAnsible) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIntegrationType

`func (o *IntegrationAnsible) GetIntegrationType() InlineResponse20079LoadBalancerMonitorLoadBalancerType`

GetIntegrationType returns the IntegrationType field if non-nil, zero value otherwise.

### GetIntegrationTypeOk

`func (o *IntegrationAnsible) GetIntegrationTypeOk() (*InlineResponse20079LoadBalancerMonitorLoadBalancerType, bool)`

GetIntegrationTypeOk returns a tuple with the IntegrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationType

`func (o *IntegrationAnsible) SetIntegrationType(v InlineResponse20079LoadBalancerMonitorLoadBalancerType)`

SetIntegrationType sets IntegrationType field to given value.

### HasIntegrationType

`func (o *IntegrationAnsible) HasIntegrationType() bool`

HasIntegrationType returns a boolean if a field has been set.

### GetUrl

`func (o *IntegrationAnsible) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IntegrationAnsible) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IntegrationAnsible) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *IntegrationAnsible) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetServiceKey

`func (o *IntegrationAnsible) GetServiceKey() InlineResponse20040AppDeployInstance`

GetServiceKey returns the ServiceKey field if non-nil, zero value otherwise.

### GetServiceKeyOk

`func (o *IntegrationAnsible) GetServiceKeyOk() (*InlineResponse20040AppDeployInstance, bool)`

GetServiceKeyOk returns a tuple with the ServiceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceKey

`func (o *IntegrationAnsible) SetServiceKey(v InlineResponse20040AppDeployInstance)`

SetServiceKey sets ServiceKey field to given value.

### HasServiceKey

`func (o *IntegrationAnsible) HasServiceKey() bool`

HasServiceKey returns a boolean if a field has been set.

### GetIsPlugin

`func (o *IntegrationAnsible) GetIsPlugin() bool`

GetIsPlugin returns the IsPlugin field if non-nil, zero value otherwise.

### GetIsPluginOk

`func (o *IntegrationAnsible) GetIsPluginOk() (*bool, bool)`

GetIsPluginOk returns a tuple with the IsPlugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlugin

`func (o *IntegrationAnsible) SetIsPlugin(v bool)`

SetIsPlugin sets IsPlugin field to given value.

### HasIsPlugin

`func (o *IntegrationAnsible) HasIsPlugin() bool`

HasIsPlugin returns a boolean if a field has been set.

### GetConfig

`func (o *IntegrationAnsible) GetConfig() IntegrationAnsibleConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *IntegrationAnsible) GetConfigOk() (*IntegrationAnsibleConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *IntegrationAnsible) SetConfig(v IntegrationAnsibleConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *IntegrationAnsible) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetStatus

`func (o *IntegrationAnsible) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IntegrationAnsible) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IntegrationAnsible) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IntegrationAnsible) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDate

`func (o *IntegrationAnsible) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *IntegrationAnsible) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *IntegrationAnsible) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *IntegrationAnsible) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### GetStatusMessage

`func (o *IntegrationAnsible) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *IntegrationAnsible) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *IntegrationAnsible) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *IntegrationAnsible) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *IntegrationAnsible) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *IntegrationAnsible) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetLastSync

`func (o *IntegrationAnsible) GetLastSync() string`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *IntegrationAnsible) GetLastSyncOk() (*string, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *IntegrationAnsible) SetLastSync(v string)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *IntegrationAnsible) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.

### SetLastSyncNil

`func (o *IntegrationAnsible) SetLastSyncNil(b bool)`

 SetLastSyncNil sets the value for LastSync to be an explicit nil

### UnsetLastSync
`func (o *IntegrationAnsible) UnsetLastSync()`

UnsetLastSync ensures that no value is present for LastSync, not even an explicit nil
### GetLastSyncDuration

`func (o *IntegrationAnsible) GetLastSyncDuration() string`

GetLastSyncDuration returns the LastSyncDuration field if non-nil, zero value otherwise.

### GetLastSyncDurationOk

`func (o *IntegrationAnsible) GetLastSyncDurationOk() (*string, bool)`

GetLastSyncDurationOk returns a tuple with the LastSyncDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncDuration

`func (o *IntegrationAnsible) SetLastSyncDuration(v string)`

SetLastSyncDuration sets LastSyncDuration field to given value.

### HasLastSyncDuration

`func (o *IntegrationAnsible) HasLastSyncDuration() bool`

HasLastSyncDuration returns a boolean if a field has been set.

### SetLastSyncDurationNil

`func (o *IntegrationAnsible) SetLastSyncDurationNil(b bool)`

 SetLastSyncDurationNil sets the value for LastSyncDuration to be an explicit nil

### UnsetLastSyncDuration
`func (o *IntegrationAnsible) UnsetLastSyncDuration()`

UnsetLastSyncDuration ensures that no value is present for LastSyncDuration, not even an explicit nil
### GetCredential

`func (o *IntegrationAnsible) GetCredential() Creds`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *IntegrationAnsible) GetCredentialOk() (*Creds, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *IntegrationAnsible) SetCredential(v Creds)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *IntegrationAnsible) HasCredential() bool`

HasCredential returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


