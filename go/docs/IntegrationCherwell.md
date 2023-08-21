# IntegrationCherwell

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**IntegrationType** | Pointer to [**InlineResponse20079LoadBalancerMonitorLoadBalancerType**](inline_response_200_79_loadBalancerMonitor_loadBalancer_type.md) |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**PasswordHash** | Pointer to **string** |  | [optional] 
**IsPlugin** | Pointer to **bool** |  | [optional] 
**Config** | Pointer to [**IntegrationCherwellConfig**](integrationCherwell_config.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusDate** | Pointer to **NullableTime** |  | [optional] 
**StatusMessage** | Pointer to **NullableString** |  | [optional] 
**LastSync** | Pointer to **NullableString** |  | [optional] 
**LastSyncDuration** | Pointer to **NullableString** |  | [optional] 
**Credential** | Pointer to [**Creds**](creds.md) |  | [optional] 

## Methods

### NewIntegrationCherwell

`func NewIntegrationCherwell() *IntegrationCherwell`

NewIntegrationCherwell instantiates a new IntegrationCherwell object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationCherwellWithDefaults

`func NewIntegrationCherwellWithDefaults() *IntegrationCherwell`

NewIntegrationCherwellWithDefaults instantiates a new IntegrationCherwell object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IntegrationCherwell) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IntegrationCherwell) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IntegrationCherwell) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *IntegrationCherwell) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IntegrationCherwell) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IntegrationCherwell) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IntegrationCherwell) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IntegrationCherwell) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *IntegrationCherwell) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IntegrationCherwell) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IntegrationCherwell) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *IntegrationCherwell) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetType

`func (o *IntegrationCherwell) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IntegrationCherwell) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IntegrationCherwell) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IntegrationCherwell) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIntegrationType

`func (o *IntegrationCherwell) GetIntegrationType() InlineResponse20079LoadBalancerMonitorLoadBalancerType`

GetIntegrationType returns the IntegrationType field if non-nil, zero value otherwise.

### GetIntegrationTypeOk

`func (o *IntegrationCherwell) GetIntegrationTypeOk() (*InlineResponse20079LoadBalancerMonitorLoadBalancerType, bool)`

GetIntegrationTypeOk returns a tuple with the IntegrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationType

`func (o *IntegrationCherwell) SetIntegrationType(v InlineResponse20079LoadBalancerMonitorLoadBalancerType)`

SetIntegrationType sets IntegrationType field to given value.

### HasIntegrationType

`func (o *IntegrationCherwell) HasIntegrationType() bool`

HasIntegrationType returns a boolean if a field has been set.

### GetUrl

`func (o *IntegrationCherwell) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IntegrationCherwell) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IntegrationCherwell) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *IntegrationCherwell) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUsername

`func (o *IntegrationCherwell) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *IntegrationCherwell) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *IntegrationCherwell) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *IntegrationCherwell) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *IntegrationCherwell) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *IntegrationCherwell) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *IntegrationCherwell) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *IntegrationCherwell) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPasswordHash

`func (o *IntegrationCherwell) GetPasswordHash() string`

GetPasswordHash returns the PasswordHash field if non-nil, zero value otherwise.

### GetPasswordHashOk

`func (o *IntegrationCherwell) GetPasswordHashOk() (*string, bool)`

GetPasswordHashOk returns a tuple with the PasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHash

`func (o *IntegrationCherwell) SetPasswordHash(v string)`

SetPasswordHash sets PasswordHash field to given value.

### HasPasswordHash

`func (o *IntegrationCherwell) HasPasswordHash() bool`

HasPasswordHash returns a boolean if a field has been set.

### GetIsPlugin

`func (o *IntegrationCherwell) GetIsPlugin() bool`

GetIsPlugin returns the IsPlugin field if non-nil, zero value otherwise.

### GetIsPluginOk

`func (o *IntegrationCherwell) GetIsPluginOk() (*bool, bool)`

GetIsPluginOk returns a tuple with the IsPlugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlugin

`func (o *IntegrationCherwell) SetIsPlugin(v bool)`

SetIsPlugin sets IsPlugin field to given value.

### HasIsPlugin

`func (o *IntegrationCherwell) HasIsPlugin() bool`

HasIsPlugin returns a boolean if a field has been set.

### GetConfig

`func (o *IntegrationCherwell) GetConfig() IntegrationCherwellConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *IntegrationCherwell) GetConfigOk() (*IntegrationCherwellConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *IntegrationCherwell) SetConfig(v IntegrationCherwellConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *IntegrationCherwell) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetStatus

`func (o *IntegrationCherwell) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IntegrationCherwell) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IntegrationCherwell) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IntegrationCherwell) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDate

`func (o *IntegrationCherwell) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *IntegrationCherwell) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *IntegrationCherwell) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *IntegrationCherwell) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### SetStatusDateNil

`func (o *IntegrationCherwell) SetStatusDateNil(b bool)`

 SetStatusDateNil sets the value for StatusDate to be an explicit nil

### UnsetStatusDate
`func (o *IntegrationCherwell) UnsetStatusDate()`

UnsetStatusDate ensures that no value is present for StatusDate, not even an explicit nil
### GetStatusMessage

`func (o *IntegrationCherwell) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *IntegrationCherwell) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *IntegrationCherwell) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *IntegrationCherwell) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *IntegrationCherwell) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *IntegrationCherwell) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetLastSync

`func (o *IntegrationCherwell) GetLastSync() string`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *IntegrationCherwell) GetLastSyncOk() (*string, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *IntegrationCherwell) SetLastSync(v string)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *IntegrationCherwell) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.

### SetLastSyncNil

`func (o *IntegrationCherwell) SetLastSyncNil(b bool)`

 SetLastSyncNil sets the value for LastSync to be an explicit nil

### UnsetLastSync
`func (o *IntegrationCherwell) UnsetLastSync()`

UnsetLastSync ensures that no value is present for LastSync, not even an explicit nil
### GetLastSyncDuration

`func (o *IntegrationCherwell) GetLastSyncDuration() string`

GetLastSyncDuration returns the LastSyncDuration field if non-nil, zero value otherwise.

### GetLastSyncDurationOk

`func (o *IntegrationCherwell) GetLastSyncDurationOk() (*string, bool)`

GetLastSyncDurationOk returns a tuple with the LastSyncDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncDuration

`func (o *IntegrationCherwell) SetLastSyncDuration(v string)`

SetLastSyncDuration sets LastSyncDuration field to given value.

### HasLastSyncDuration

`func (o *IntegrationCherwell) HasLastSyncDuration() bool`

HasLastSyncDuration returns a boolean if a field has been set.

### SetLastSyncDurationNil

`func (o *IntegrationCherwell) SetLastSyncDurationNil(b bool)`

 SetLastSyncDurationNil sets the value for LastSyncDuration to be an explicit nil

### UnsetLastSyncDuration
`func (o *IntegrationCherwell) UnsetLastSyncDuration()`

UnsetLastSyncDuration ensures that no value is present for LastSyncDuration, not even an explicit nil
### GetCredential

`func (o *IntegrationCherwell) GetCredential() Creds`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *IntegrationCherwell) GetCredentialOk() (*Creds, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *IntegrationCherwell) SetCredential(v Creds)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *IntegrationCherwell) HasCredential() bool`

HasCredential returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


