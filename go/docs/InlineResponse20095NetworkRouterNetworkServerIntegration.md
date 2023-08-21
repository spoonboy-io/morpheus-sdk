# InlineResponse20095NetworkRouterNetworkServerIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**IntegrationType** | Pointer to [**InlineResponse20079LoadBalancerMonitorLoadBalancerType**](inline_response_200_79_loadBalancerMonitor_loadBalancer_type.md) |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**RefType** | Pointer to **string** |  | [optional] 
**RefId** | Pointer to **string** |  | [optional] 
**IsPlugin** | Pointer to **bool** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusDate** | Pointer to **time.Time** |  | [optional] 
**StatusMessage** | Pointer to **string** |  | [optional] 
**LastSync** | Pointer to **time.Time** |  | [optional] 
**LastSyncDuration** | Pointer to **int64** |  | [optional] 

## Methods

### NewInlineResponse20095NetworkRouterNetworkServerIntegration

`func NewInlineResponse20095NetworkRouterNetworkServerIntegration() *InlineResponse20095NetworkRouterNetworkServerIntegration`

NewInlineResponse20095NetworkRouterNetworkServerIntegration instantiates a new InlineResponse20095NetworkRouterNetworkServerIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20095NetworkRouterNetworkServerIntegrationWithDefaults

`func NewInlineResponse20095NetworkRouterNetworkServerIntegrationWithDefaults() *InlineResponse20095NetworkRouterNetworkServerIntegration`

NewInlineResponse20095NetworkRouterNetworkServerIntegrationWithDefaults instantiates a new InlineResponse20095NetworkRouterNetworkServerIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetType

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIntegrationType

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) GetIntegrationType() InlineResponse20079LoadBalancerMonitorLoadBalancerType`

GetIntegrationType returns the IntegrationType field if non-nil, zero value otherwise.

### GetIntegrationTypeOk

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) GetIntegrationTypeOk() (*InlineResponse20079LoadBalancerMonitorLoadBalancerType, bool)`

GetIntegrationTypeOk returns a tuple with the IntegrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationType

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) SetIntegrationType(v InlineResponse20079LoadBalancerMonitorLoadBalancerType)`

SetIntegrationType sets IntegrationType field to given value.

### HasIntegrationType

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) HasIntegrationType() bool`

HasIntegrationType returns a boolean if a field has been set.

### GetUrl

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetPort

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetUsername

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetRefType

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetIsPlugin

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) GetIsPlugin() bool`

GetIsPlugin returns the IsPlugin field if non-nil, zero value otherwise.

### GetIsPluginOk

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) GetIsPluginOk() (*bool, bool)`

GetIsPluginOk returns a tuple with the IsPlugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlugin

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) SetIsPlugin(v bool)`

SetIsPlugin sets IsPlugin field to given value.

### HasIsPlugin

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) HasIsPlugin() bool`

HasIsPlugin returns a boolean if a field has been set.

### GetConfig

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDate

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### GetStatusMessage

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetLastSync

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) GetLastSync() time.Time`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) GetLastSyncOk() (*time.Time, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) SetLastSync(v time.Time)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.

### GetLastSyncDuration

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) GetLastSyncDuration() int64`

GetLastSyncDuration returns the LastSyncDuration field if non-nil, zero value otherwise.

### GetLastSyncDurationOk

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) GetLastSyncDurationOk() (*int64, bool)`

GetLastSyncDurationOk returns a tuple with the LastSyncDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncDuration

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) SetLastSyncDuration(v int64)`

SetLastSyncDuration sets LastSyncDuration field to given value.

### HasLastSyncDuration

`func (o *InlineResponse20095NetworkRouterNetworkServerIntegration) HasLastSyncDuration() bool`

HasLastSyncDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


