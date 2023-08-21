# InlineResponse20079LoadBalancerMonitor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**LoadBalancer** | Pointer to [**InlineResponse20079LoadBalancerMonitorLoadBalancer**](inline_response_200_79_loadBalancerMonitor_loadBalancer.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **NullableString** |  | [optional] 
**Category** | Pointer to **NullableString** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**MonitorType** | Pointer to **string** |  | [optional] 
**MonitorInterval** | Pointer to **int64** |  | [optional] 
**MonitorTimeout** | Pointer to **int64** |  | [optional] 
**SendData** | Pointer to **NullableString** |  | [optional] 
**SendVersion** | Pointer to **string** |  | [optional] 
**SendType** | Pointer to **string** |  | [optional] 
**ReceiveData** | Pointer to **NullableString** |  | [optional] 
**ReceiveCode** | Pointer to **string** |  | [optional] 
**DisabledData** | Pointer to **NullableString** |  | [optional] 
**MonitorUsername** | Pointer to **NullableString** |  | [optional] 
**MonitorPassword** | Pointer to **NullableString** |  | [optional] 
**MonitorDestination** | Pointer to **string** |  | [optional] 
**MonitorReverse** | Pointer to **bool** |  | [optional] 
**MonitorTransparent** | Pointer to **bool** |  | [optional] 
**MonitorAdaptive** | Pointer to **bool** |  | [optional] 
**AliasAddress** | Pointer to **NullableString** |  | [optional] 
**AliasPort** | Pointer to **int64** |  | [optional] 
**InternalId** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**MonitorSource** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusMessage** | Pointer to **NullableString** |  | [optional] 
**StatusDate** | Pointer to **NullableTime** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**MaxRetry** | Pointer to **int64** |  | [optional] 
**FallCount** | Pointer to **int64** |  | [optional] 
**RiseCount** | Pointer to **int64** |  | [optional] 
**DataLength** | Pointer to **NullableString** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewInlineResponse20079LoadBalancerMonitor

`func NewInlineResponse20079LoadBalancerMonitor() *InlineResponse20079LoadBalancerMonitor`

NewInlineResponse20079LoadBalancerMonitor instantiates a new InlineResponse20079LoadBalancerMonitor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20079LoadBalancerMonitorWithDefaults

`func NewInlineResponse20079LoadBalancerMonitorWithDefaults() *InlineResponse20079LoadBalancerMonitor`

NewInlineResponse20079LoadBalancerMonitorWithDefaults instantiates a new InlineResponse20079LoadBalancerMonitor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse20079LoadBalancerMonitor) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse20079LoadBalancerMonitor) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse20079LoadBalancerMonitor) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse20079LoadBalancerMonitor) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLoadBalancer

`func (o *InlineResponse20079LoadBalancerMonitor) GetLoadBalancer() InlineResponse20079LoadBalancerMonitorLoadBalancer`

GetLoadBalancer returns the LoadBalancer field if non-nil, zero value otherwise.

### GetLoadBalancerOk

`func (o *InlineResponse20079LoadBalancerMonitor) GetLoadBalancerOk() (*InlineResponse20079LoadBalancerMonitorLoadBalancer, bool)`

GetLoadBalancerOk returns a tuple with the LoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancer

`func (o *InlineResponse20079LoadBalancerMonitor) SetLoadBalancer(v InlineResponse20079LoadBalancerMonitorLoadBalancer)`

SetLoadBalancer sets LoadBalancer field to given value.

### HasLoadBalancer

`func (o *InlineResponse20079LoadBalancerMonitor) HasLoadBalancer() bool`

HasLoadBalancer returns a boolean if a field has been set.

### GetName

`func (o *InlineResponse20079LoadBalancerMonitor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineResponse20079LoadBalancerMonitor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineResponse20079LoadBalancerMonitor) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineResponse20079LoadBalancerMonitor) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *InlineResponse20079LoadBalancerMonitor) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *InlineResponse20079LoadBalancerMonitor) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *InlineResponse20079LoadBalancerMonitor) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *InlineResponse20079LoadBalancerMonitor) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *InlineResponse20079LoadBalancerMonitor) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *InlineResponse20079LoadBalancerMonitor) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetCategory

`func (o *InlineResponse20079LoadBalancerMonitor) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *InlineResponse20079LoadBalancerMonitor) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *InlineResponse20079LoadBalancerMonitor) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *InlineResponse20079LoadBalancerMonitor) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *InlineResponse20079LoadBalancerMonitor) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *InlineResponse20079LoadBalancerMonitor) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetVisibility

`func (o *InlineResponse20079LoadBalancerMonitor) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *InlineResponse20079LoadBalancerMonitor) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *InlineResponse20079LoadBalancerMonitor) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *InlineResponse20079LoadBalancerMonitor) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetDescription

`func (o *InlineResponse20079LoadBalancerMonitor) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse20079LoadBalancerMonitor) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse20079LoadBalancerMonitor) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse20079LoadBalancerMonitor) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMonitorType

`func (o *InlineResponse20079LoadBalancerMonitor) GetMonitorType() string`

GetMonitorType returns the MonitorType field if non-nil, zero value otherwise.

### GetMonitorTypeOk

`func (o *InlineResponse20079LoadBalancerMonitor) GetMonitorTypeOk() (*string, bool)`

GetMonitorTypeOk returns a tuple with the MonitorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorType

`func (o *InlineResponse20079LoadBalancerMonitor) SetMonitorType(v string)`

SetMonitorType sets MonitorType field to given value.

### HasMonitorType

`func (o *InlineResponse20079LoadBalancerMonitor) HasMonitorType() bool`

HasMonitorType returns a boolean if a field has been set.

### GetMonitorInterval

`func (o *InlineResponse20079LoadBalancerMonitor) GetMonitorInterval() int64`

GetMonitorInterval returns the MonitorInterval field if non-nil, zero value otherwise.

### GetMonitorIntervalOk

`func (o *InlineResponse20079LoadBalancerMonitor) GetMonitorIntervalOk() (*int64, bool)`

GetMonitorIntervalOk returns a tuple with the MonitorInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorInterval

`func (o *InlineResponse20079LoadBalancerMonitor) SetMonitorInterval(v int64)`

SetMonitorInterval sets MonitorInterval field to given value.

### HasMonitorInterval

`func (o *InlineResponse20079LoadBalancerMonitor) HasMonitorInterval() bool`

HasMonitorInterval returns a boolean if a field has been set.

### GetMonitorTimeout

`func (o *InlineResponse20079LoadBalancerMonitor) GetMonitorTimeout() int64`

GetMonitorTimeout returns the MonitorTimeout field if non-nil, zero value otherwise.

### GetMonitorTimeoutOk

`func (o *InlineResponse20079LoadBalancerMonitor) GetMonitorTimeoutOk() (*int64, bool)`

GetMonitorTimeoutOk returns a tuple with the MonitorTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorTimeout

`func (o *InlineResponse20079LoadBalancerMonitor) SetMonitorTimeout(v int64)`

SetMonitorTimeout sets MonitorTimeout field to given value.

### HasMonitorTimeout

`func (o *InlineResponse20079LoadBalancerMonitor) HasMonitorTimeout() bool`

HasMonitorTimeout returns a boolean if a field has been set.

### GetSendData

`func (o *InlineResponse20079LoadBalancerMonitor) GetSendData() string`

GetSendData returns the SendData field if non-nil, zero value otherwise.

### GetSendDataOk

`func (o *InlineResponse20079LoadBalancerMonitor) GetSendDataOk() (*string, bool)`

GetSendDataOk returns a tuple with the SendData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendData

`func (o *InlineResponse20079LoadBalancerMonitor) SetSendData(v string)`

SetSendData sets SendData field to given value.

### HasSendData

`func (o *InlineResponse20079LoadBalancerMonitor) HasSendData() bool`

HasSendData returns a boolean if a field has been set.

### SetSendDataNil

`func (o *InlineResponse20079LoadBalancerMonitor) SetSendDataNil(b bool)`

 SetSendDataNil sets the value for SendData to be an explicit nil

### UnsetSendData
`func (o *InlineResponse20079LoadBalancerMonitor) UnsetSendData()`

UnsetSendData ensures that no value is present for SendData, not even an explicit nil
### GetSendVersion

`func (o *InlineResponse20079LoadBalancerMonitor) GetSendVersion() string`

GetSendVersion returns the SendVersion field if non-nil, zero value otherwise.

### GetSendVersionOk

`func (o *InlineResponse20079LoadBalancerMonitor) GetSendVersionOk() (*string, bool)`

GetSendVersionOk returns a tuple with the SendVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendVersion

`func (o *InlineResponse20079LoadBalancerMonitor) SetSendVersion(v string)`

SetSendVersion sets SendVersion field to given value.

### HasSendVersion

`func (o *InlineResponse20079LoadBalancerMonitor) HasSendVersion() bool`

HasSendVersion returns a boolean if a field has been set.

### GetSendType

`func (o *InlineResponse20079LoadBalancerMonitor) GetSendType() string`

GetSendType returns the SendType field if non-nil, zero value otherwise.

### GetSendTypeOk

`func (o *InlineResponse20079LoadBalancerMonitor) GetSendTypeOk() (*string, bool)`

GetSendTypeOk returns a tuple with the SendType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendType

`func (o *InlineResponse20079LoadBalancerMonitor) SetSendType(v string)`

SetSendType sets SendType field to given value.

### HasSendType

`func (o *InlineResponse20079LoadBalancerMonitor) HasSendType() bool`

HasSendType returns a boolean if a field has been set.

### GetReceiveData

`func (o *InlineResponse20079LoadBalancerMonitor) GetReceiveData() string`

GetReceiveData returns the ReceiveData field if non-nil, zero value otherwise.

### GetReceiveDataOk

`func (o *InlineResponse20079LoadBalancerMonitor) GetReceiveDataOk() (*string, bool)`

GetReceiveDataOk returns a tuple with the ReceiveData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveData

`func (o *InlineResponse20079LoadBalancerMonitor) SetReceiveData(v string)`

SetReceiveData sets ReceiveData field to given value.

### HasReceiveData

`func (o *InlineResponse20079LoadBalancerMonitor) HasReceiveData() bool`

HasReceiveData returns a boolean if a field has been set.

### SetReceiveDataNil

`func (o *InlineResponse20079LoadBalancerMonitor) SetReceiveDataNil(b bool)`

 SetReceiveDataNil sets the value for ReceiveData to be an explicit nil

### UnsetReceiveData
`func (o *InlineResponse20079LoadBalancerMonitor) UnsetReceiveData()`

UnsetReceiveData ensures that no value is present for ReceiveData, not even an explicit nil
### GetReceiveCode

`func (o *InlineResponse20079LoadBalancerMonitor) GetReceiveCode() string`

GetReceiveCode returns the ReceiveCode field if non-nil, zero value otherwise.

### GetReceiveCodeOk

`func (o *InlineResponse20079LoadBalancerMonitor) GetReceiveCodeOk() (*string, bool)`

GetReceiveCodeOk returns a tuple with the ReceiveCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveCode

`func (o *InlineResponse20079LoadBalancerMonitor) SetReceiveCode(v string)`

SetReceiveCode sets ReceiveCode field to given value.

### HasReceiveCode

`func (o *InlineResponse20079LoadBalancerMonitor) HasReceiveCode() bool`

HasReceiveCode returns a boolean if a field has been set.

### GetDisabledData

`func (o *InlineResponse20079LoadBalancerMonitor) GetDisabledData() string`

GetDisabledData returns the DisabledData field if non-nil, zero value otherwise.

### GetDisabledDataOk

`func (o *InlineResponse20079LoadBalancerMonitor) GetDisabledDataOk() (*string, bool)`

GetDisabledDataOk returns a tuple with the DisabledData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledData

`func (o *InlineResponse20079LoadBalancerMonitor) SetDisabledData(v string)`

SetDisabledData sets DisabledData field to given value.

### HasDisabledData

`func (o *InlineResponse20079LoadBalancerMonitor) HasDisabledData() bool`

HasDisabledData returns a boolean if a field has been set.

### SetDisabledDataNil

`func (o *InlineResponse20079LoadBalancerMonitor) SetDisabledDataNil(b bool)`

 SetDisabledDataNil sets the value for DisabledData to be an explicit nil

### UnsetDisabledData
`func (o *InlineResponse20079LoadBalancerMonitor) UnsetDisabledData()`

UnsetDisabledData ensures that no value is present for DisabledData, not even an explicit nil
### GetMonitorUsername

`func (o *InlineResponse20079LoadBalancerMonitor) GetMonitorUsername() string`

GetMonitorUsername returns the MonitorUsername field if non-nil, zero value otherwise.

### GetMonitorUsernameOk

`func (o *InlineResponse20079LoadBalancerMonitor) GetMonitorUsernameOk() (*string, bool)`

GetMonitorUsernameOk returns a tuple with the MonitorUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorUsername

`func (o *InlineResponse20079LoadBalancerMonitor) SetMonitorUsername(v string)`

SetMonitorUsername sets MonitorUsername field to given value.

### HasMonitorUsername

`func (o *InlineResponse20079LoadBalancerMonitor) HasMonitorUsername() bool`

HasMonitorUsername returns a boolean if a field has been set.

### SetMonitorUsernameNil

`func (o *InlineResponse20079LoadBalancerMonitor) SetMonitorUsernameNil(b bool)`

 SetMonitorUsernameNil sets the value for MonitorUsername to be an explicit nil

### UnsetMonitorUsername
`func (o *InlineResponse20079LoadBalancerMonitor) UnsetMonitorUsername()`

UnsetMonitorUsername ensures that no value is present for MonitorUsername, not even an explicit nil
### GetMonitorPassword

`func (o *InlineResponse20079LoadBalancerMonitor) GetMonitorPassword() string`

GetMonitorPassword returns the MonitorPassword field if non-nil, zero value otherwise.

### GetMonitorPasswordOk

`func (o *InlineResponse20079LoadBalancerMonitor) GetMonitorPasswordOk() (*string, bool)`

GetMonitorPasswordOk returns a tuple with the MonitorPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorPassword

`func (o *InlineResponse20079LoadBalancerMonitor) SetMonitorPassword(v string)`

SetMonitorPassword sets MonitorPassword field to given value.

### HasMonitorPassword

`func (o *InlineResponse20079LoadBalancerMonitor) HasMonitorPassword() bool`

HasMonitorPassword returns a boolean if a field has been set.

### SetMonitorPasswordNil

`func (o *InlineResponse20079LoadBalancerMonitor) SetMonitorPasswordNil(b bool)`

 SetMonitorPasswordNil sets the value for MonitorPassword to be an explicit nil

### UnsetMonitorPassword
`func (o *InlineResponse20079LoadBalancerMonitor) UnsetMonitorPassword()`

UnsetMonitorPassword ensures that no value is present for MonitorPassword, not even an explicit nil
### GetMonitorDestination

`func (o *InlineResponse20079LoadBalancerMonitor) GetMonitorDestination() string`

GetMonitorDestination returns the MonitorDestination field if non-nil, zero value otherwise.

### GetMonitorDestinationOk

`func (o *InlineResponse20079LoadBalancerMonitor) GetMonitorDestinationOk() (*string, bool)`

GetMonitorDestinationOk returns a tuple with the MonitorDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorDestination

`func (o *InlineResponse20079LoadBalancerMonitor) SetMonitorDestination(v string)`

SetMonitorDestination sets MonitorDestination field to given value.

### HasMonitorDestination

`func (o *InlineResponse20079LoadBalancerMonitor) HasMonitorDestination() bool`

HasMonitorDestination returns a boolean if a field has been set.

### GetMonitorReverse

`func (o *InlineResponse20079LoadBalancerMonitor) GetMonitorReverse() bool`

GetMonitorReverse returns the MonitorReverse field if non-nil, zero value otherwise.

### GetMonitorReverseOk

`func (o *InlineResponse20079LoadBalancerMonitor) GetMonitorReverseOk() (*bool, bool)`

GetMonitorReverseOk returns a tuple with the MonitorReverse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorReverse

`func (o *InlineResponse20079LoadBalancerMonitor) SetMonitorReverse(v bool)`

SetMonitorReverse sets MonitorReverse field to given value.

### HasMonitorReverse

`func (o *InlineResponse20079LoadBalancerMonitor) HasMonitorReverse() bool`

HasMonitorReverse returns a boolean if a field has been set.

### GetMonitorTransparent

`func (o *InlineResponse20079LoadBalancerMonitor) GetMonitorTransparent() bool`

GetMonitorTransparent returns the MonitorTransparent field if non-nil, zero value otherwise.

### GetMonitorTransparentOk

`func (o *InlineResponse20079LoadBalancerMonitor) GetMonitorTransparentOk() (*bool, bool)`

GetMonitorTransparentOk returns a tuple with the MonitorTransparent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorTransparent

`func (o *InlineResponse20079LoadBalancerMonitor) SetMonitorTransparent(v bool)`

SetMonitorTransparent sets MonitorTransparent field to given value.

### HasMonitorTransparent

`func (o *InlineResponse20079LoadBalancerMonitor) HasMonitorTransparent() bool`

HasMonitorTransparent returns a boolean if a field has been set.

### GetMonitorAdaptive

`func (o *InlineResponse20079LoadBalancerMonitor) GetMonitorAdaptive() bool`

GetMonitorAdaptive returns the MonitorAdaptive field if non-nil, zero value otherwise.

### GetMonitorAdaptiveOk

`func (o *InlineResponse20079LoadBalancerMonitor) GetMonitorAdaptiveOk() (*bool, bool)`

GetMonitorAdaptiveOk returns a tuple with the MonitorAdaptive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorAdaptive

`func (o *InlineResponse20079LoadBalancerMonitor) SetMonitorAdaptive(v bool)`

SetMonitorAdaptive sets MonitorAdaptive field to given value.

### HasMonitorAdaptive

`func (o *InlineResponse20079LoadBalancerMonitor) HasMonitorAdaptive() bool`

HasMonitorAdaptive returns a boolean if a field has been set.

### GetAliasAddress

`func (o *InlineResponse20079LoadBalancerMonitor) GetAliasAddress() string`

GetAliasAddress returns the AliasAddress field if non-nil, zero value otherwise.

### GetAliasAddressOk

`func (o *InlineResponse20079LoadBalancerMonitor) GetAliasAddressOk() (*string, bool)`

GetAliasAddressOk returns a tuple with the AliasAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasAddress

`func (o *InlineResponse20079LoadBalancerMonitor) SetAliasAddress(v string)`

SetAliasAddress sets AliasAddress field to given value.

### HasAliasAddress

`func (o *InlineResponse20079LoadBalancerMonitor) HasAliasAddress() bool`

HasAliasAddress returns a boolean if a field has been set.

### SetAliasAddressNil

`func (o *InlineResponse20079LoadBalancerMonitor) SetAliasAddressNil(b bool)`

 SetAliasAddressNil sets the value for AliasAddress to be an explicit nil

### UnsetAliasAddress
`func (o *InlineResponse20079LoadBalancerMonitor) UnsetAliasAddress()`

UnsetAliasAddress ensures that no value is present for AliasAddress, not even an explicit nil
### GetAliasPort

`func (o *InlineResponse20079LoadBalancerMonitor) GetAliasPort() int64`

GetAliasPort returns the AliasPort field if non-nil, zero value otherwise.

### GetAliasPortOk

`func (o *InlineResponse20079LoadBalancerMonitor) GetAliasPortOk() (*int64, bool)`

GetAliasPortOk returns a tuple with the AliasPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasPort

`func (o *InlineResponse20079LoadBalancerMonitor) SetAliasPort(v int64)`

SetAliasPort sets AliasPort field to given value.

### HasAliasPort

`func (o *InlineResponse20079LoadBalancerMonitor) HasAliasPort() bool`

HasAliasPort returns a boolean if a field has been set.

### GetInternalId

`func (o *InlineResponse20079LoadBalancerMonitor) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *InlineResponse20079LoadBalancerMonitor) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *InlineResponse20079LoadBalancerMonitor) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *InlineResponse20079LoadBalancerMonitor) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetExternalId

`func (o *InlineResponse20079LoadBalancerMonitor) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *InlineResponse20079LoadBalancerMonitor) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *InlineResponse20079LoadBalancerMonitor) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *InlineResponse20079LoadBalancerMonitor) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetMonitorSource

`func (o *InlineResponse20079LoadBalancerMonitor) GetMonitorSource() string`

GetMonitorSource returns the MonitorSource field if non-nil, zero value otherwise.

### GetMonitorSourceOk

`func (o *InlineResponse20079LoadBalancerMonitor) GetMonitorSourceOk() (*string, bool)`

GetMonitorSourceOk returns a tuple with the MonitorSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorSource

`func (o *InlineResponse20079LoadBalancerMonitor) SetMonitorSource(v string)`

SetMonitorSource sets MonitorSource field to given value.

### HasMonitorSource

`func (o *InlineResponse20079LoadBalancerMonitor) HasMonitorSource() bool`

HasMonitorSource returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse20079LoadBalancerMonitor) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse20079LoadBalancerMonitor) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse20079LoadBalancerMonitor) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse20079LoadBalancerMonitor) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *InlineResponse20079LoadBalancerMonitor) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *InlineResponse20079LoadBalancerMonitor) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *InlineResponse20079LoadBalancerMonitor) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *InlineResponse20079LoadBalancerMonitor) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *InlineResponse20079LoadBalancerMonitor) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *InlineResponse20079LoadBalancerMonitor) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetStatusDate

`func (o *InlineResponse20079LoadBalancerMonitor) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *InlineResponse20079LoadBalancerMonitor) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *InlineResponse20079LoadBalancerMonitor) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *InlineResponse20079LoadBalancerMonitor) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### SetStatusDateNil

`func (o *InlineResponse20079LoadBalancerMonitor) SetStatusDateNil(b bool)`

 SetStatusDateNil sets the value for StatusDate to be an explicit nil

### UnsetStatusDate
`func (o *InlineResponse20079LoadBalancerMonitor) UnsetStatusDate()`

UnsetStatusDate ensures that no value is present for StatusDate, not even an explicit nil
### GetEnabled

`func (o *InlineResponse20079LoadBalancerMonitor) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse20079LoadBalancerMonitor) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse20079LoadBalancerMonitor) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse20079LoadBalancerMonitor) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMaxRetry

`func (o *InlineResponse20079LoadBalancerMonitor) GetMaxRetry() int64`

GetMaxRetry returns the MaxRetry field if non-nil, zero value otherwise.

### GetMaxRetryOk

`func (o *InlineResponse20079LoadBalancerMonitor) GetMaxRetryOk() (*int64, bool)`

GetMaxRetryOk returns a tuple with the MaxRetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetry

`func (o *InlineResponse20079LoadBalancerMonitor) SetMaxRetry(v int64)`

SetMaxRetry sets MaxRetry field to given value.

### HasMaxRetry

`func (o *InlineResponse20079LoadBalancerMonitor) HasMaxRetry() bool`

HasMaxRetry returns a boolean if a field has been set.

### GetFallCount

`func (o *InlineResponse20079LoadBalancerMonitor) GetFallCount() int64`

GetFallCount returns the FallCount field if non-nil, zero value otherwise.

### GetFallCountOk

`func (o *InlineResponse20079LoadBalancerMonitor) GetFallCountOk() (*int64, bool)`

GetFallCountOk returns a tuple with the FallCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallCount

`func (o *InlineResponse20079LoadBalancerMonitor) SetFallCount(v int64)`

SetFallCount sets FallCount field to given value.

### HasFallCount

`func (o *InlineResponse20079LoadBalancerMonitor) HasFallCount() bool`

HasFallCount returns a boolean if a field has been set.

### GetRiseCount

`func (o *InlineResponse20079LoadBalancerMonitor) GetRiseCount() int64`

GetRiseCount returns the RiseCount field if non-nil, zero value otherwise.

### GetRiseCountOk

`func (o *InlineResponse20079LoadBalancerMonitor) GetRiseCountOk() (*int64, bool)`

GetRiseCountOk returns a tuple with the RiseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiseCount

`func (o *InlineResponse20079LoadBalancerMonitor) SetRiseCount(v int64)`

SetRiseCount sets RiseCount field to given value.

### HasRiseCount

`func (o *InlineResponse20079LoadBalancerMonitor) HasRiseCount() bool`

HasRiseCount returns a boolean if a field has been set.

### GetDataLength

`func (o *InlineResponse20079LoadBalancerMonitor) GetDataLength() string`

GetDataLength returns the DataLength field if non-nil, zero value otherwise.

### GetDataLengthOk

`func (o *InlineResponse20079LoadBalancerMonitor) GetDataLengthOk() (*string, bool)`

GetDataLengthOk returns a tuple with the DataLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLength

`func (o *InlineResponse20079LoadBalancerMonitor) SetDataLength(v string)`

SetDataLength sets DataLength field to given value.

### HasDataLength

`func (o *InlineResponse20079LoadBalancerMonitor) HasDataLength() bool`

HasDataLength returns a boolean if a field has been set.

### SetDataLengthNil

`func (o *InlineResponse20079LoadBalancerMonitor) SetDataLengthNil(b bool)`

 SetDataLengthNil sets the value for DataLength to be an explicit nil

### UnsetDataLength
`func (o *InlineResponse20079LoadBalancerMonitor) UnsetDataLength()`

UnsetDataLength ensures that no value is present for DataLength, not even an explicit nil
### GetConfig

`func (o *InlineResponse20079LoadBalancerMonitor) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *InlineResponse20079LoadBalancerMonitor) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *InlineResponse20079LoadBalancerMonitor) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *InlineResponse20079LoadBalancerMonitor) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreatedBy

`func (o *InlineResponse20079LoadBalancerMonitor) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *InlineResponse20079LoadBalancerMonitor) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *InlineResponse20079LoadBalancerMonitor) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *InlineResponse20079LoadBalancerMonitor) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *InlineResponse20079LoadBalancerMonitor) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *InlineResponse20079LoadBalancerMonitor) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetDateCreated

`func (o *InlineResponse20079LoadBalancerMonitor) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *InlineResponse20079LoadBalancerMonitor) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *InlineResponse20079LoadBalancerMonitor) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *InlineResponse20079LoadBalancerMonitor) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *InlineResponse20079LoadBalancerMonitor) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *InlineResponse20079LoadBalancerMonitor) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *InlineResponse20079LoadBalancerMonitor) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *InlineResponse20079LoadBalancerMonitor) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


